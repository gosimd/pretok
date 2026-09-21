// Package regexp provides internal match iteration and isolated splitting for
// registered pre-tokenization regex strings. Compilation uses exact-string
// lookup, without parsing or a general regex fallback.
package regexp

import (
	"errors"
	"unicode/utf8"

	"github.com/gosimd/pretok/internal/compile"
	"github.com/gosimd/pretok/internal/scan"
	simdutf8 "github.com/gosimd/unicode/utf8"
)

var (
	// ErrUnsupportedPattern means that the exact regex string is not registered.
	// It also covers malformed expressions; Compile does not validate syntax.
	ErrUnsupportedPattern = errors.New("pretok: unsupported pattern")
	// ErrInvalidUTF8 means the supplied input is not valid UTF-8.
	ErrInvalidUTF8 = errors.New("pretok: invalid UTF-8")
	// ErrUninitialized means a nil or zero-value Regexp was used.
	ErrUninitialized = errors.New("pretok: uninitialized regexp")
)

// Span is a half-open byte range [Start, End) in the supplied input string.
// Successful scans produce non-empty spans on UTF-8 character boundaries.
type Span struct {
	Start int
	End   int
}

// Regexp represents one compiled pre-tokenization regular expression.
// It is immutable and safe for concurrent use with separate cursors and output
// buffers. Construct it with Compile; its zero value is not usable for scanning.
// It defines boundaries only; text transformations belong to the caller.
type Regexp struct {
	plan *scan.Plan
}

// Compile selects a prewritten plan by exact regex string. Unregistered strings,
// including malformed expressions and unregistered equivalent spellings, return
// nil and ErrUnsupportedPattern. No parsing, normalization, or fallback occurs.
// See docs/SupportedPatterns.md for the supported exact expressions and
// docs/Regexp.md for the internal scanner contract.
func Compile(pattern string) (*Regexp, error) {
	plan, ok := compile.Lookup(pattern)
	if !ok {
		return nil, ErrUnsupportedPattern
	}
	return &Regexp{plan: plan}, nil
}

// Cursor holds per-input match iteration state. It borrows its input until
// exhaustion or disposal and must not be used concurrently. Its zero value is
// exhausted. Copies have independent progress and may be iterated separately.
type Cursor struct {
	plan  *scan.Plan
	src   string
	pos   int
	state scan.State
}

// Cursor creates an iterator over non-empty, non-overlapping matches, skipping
// unmatched gaps. It validates all of src before iteration begins. A nil or
// zero-value Regexp returns ErrUninitialized; otherwise malformed input returns
// ErrInvalidUTF8. An error is accompanied by an exhausted cursor.
func (r *Regexp) Cursor(src string) (Cursor, error) {
	if r == nil || r.plan == nil {
		return Cursor{}, ErrUninitialized
	}
	var valid bool
	if r.plan.HasRuneSplit() && len(src) >= scan.RuneSplitMinBytes {
		valid = simdutf8.ValidString(src)
	} else {
		valid = utf8.ValidString(src)
	}
	if !valid {
		return Cursor{}, ErrInvalidUTF8
	}
	return Cursor{plan: r.plan, src: src}, nil
}

// Next returns the next match and true, or a zero Span and false after exhaustion.
// Repeated calls on an exhausted cursor keep returning false.
func (c *Cursor) Next() (Span, bool) {
	if c.pos < len(c.src) && c.plan.HasSearch() {
		start, end, state := c.plan.Find(c.src, c.pos, c.state)
		c.state = state
		if end > start {
			return c.matched(start, end)
		}
	} else {
		for c.pos < len(c.src) {
			start := c.pos
			end := c.plan.Match(c.src, start)
			if end > start {
				return c.matched(start, end)
			}
			_, size := utf8.DecodeRuneInString(c.src[start:])
			c.pos += size
		}
	}
	c.src = ""
	return Span{}, false
}

func (c *Cursor) matched(start, end int) (Span, bool) {
	c.pos = end
	if end == len(c.src) {
		c.src = ""
	}
	return Span{Start: start, End: end}, true
}

// Split appends pre-token spans for src to dst. Pass dst[:0] to reuse storage.
// Each non-empty match and each maximal unmatched gap is a separate piece.
// Appended spans are ordered, non-overlapping, and cover src without changing
// bytes. Empty input adds no spans. All offsets are relative to src.
//
// Split returns dst unchanged on error, validating the entire input before
// appending anything. It retains neither src nor dst; execution state is local.
func (r *Regexp) Split(src string, dst []Span) ([]Span, error) {
	cursor, err := r.Cursor(src)
	if err != nil {
		return dst, err
	}
	if r.plan.HasRuneSplit() && len(src) >= scan.RuneSplitMinBytes {
		return splitRuneWords(src, dst), nil
	}
	previous := 0
	for span, ok := cursor.Next(); ok; span, ok = cursor.Next() {
		if span.Start > previous {
			dst = append(dst, Span{Start: previous, End: span.Start})
		}
		dst = append(dst, span)
		previous = span.End
	}
	if previous < len(src) {
		dst = append(dst, Span{Start: previous, End: len(src)})
	}
	return dst, nil
}

// Keep the scratch in a separate call so other patterns and short inputs do not
// acquire a large stack frame. Validation has finished before any destination
// writes; neither the rune window nor input escapes the call.
func splitRuneWords(src string, dst []Span) []Span {
	var buffer [scan.RuneBufferSize]rune
	cursor := scan.NewQwenMarksCursor(src, buffer[:])
	for start, end, ok := cursor.Next(); ok; start, end, ok = cursor.Next() {
		dst = append(dst, Span{Start: start, End: end})
	}
	return dst
}
