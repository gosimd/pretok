// Package pretok compiles HF pre-tokenizer configurations into reusable
// pipelines. Pipelines split and transform text, with optional mapping back to
// original byte offsets. Regex stages prefer registered specialized scanners
// and otherwise use the standard library's regexp engine.
package pretok

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gosimd/pretok/internal/pipeline"
)

var (
	// ErrUnsupportedPattern means that a Split regex is absent from the
	// specialized registry and cannot be compiled by the standard regexp package.
	// The error also wraps the standard library's syntax error.
	ErrUnsupportedPattern = errors.New("pretok: unsupported pattern")
	// ErrInvalidUTF8 means the supplied input is not valid UTF-8.
	ErrInvalidUTF8 = errors.New("pretok: invalid UTF-8")
	// ErrUninitialized means a nil or zero-value Pipeline was used.
	ErrUninitialized = errors.New("pretok: uninitialized pipeline")
	// ErrInvalidConfig indicates an invalid pre_tokenizer JSON configuration.
	ErrInvalidConfig = pipeline.ErrInvalidConfig
	// ErrUnsupportedPreTokenizer indicates an unsupported component type.
	ErrUnsupportedPreTokenizer = pipeline.ErrUnsupportedType
	// ErrNilCallback indicates a nil Walk, WalkText, or WalkTextUnsafe callback.
	ErrNilCallback = errors.New("pretok: nil callback")
	// ErrInvalidRange indicates an empty, reversed, or out-of-bounds byte range.
	ErrInvalidRange = pipeline.ErrInvalidRange
)

// Span is a half-open UTF-8 byte range [Start, End) in the original input.
// OriginalRange returns non-empty spans on source character boundaries.
type Span struct {
	Start int
	End   int
}

// Pipeline is an immutable compiled pre-tokenizer configuration. It is safe for
// concurrent and reentrant use. Construct it with CompileJSON.
type Pipeline struct{ program *pipeline.Program }

// CompileJSON compiles the contents of tokenizer.json's pre_tokenizer field,
// not the enclosing tokenizer document. A null configuration is an identity
// pipeline. Split regex strings select an exact registered scanner when possible,
// otherwise they are compiled with the standard library's regexp package.
// See docs/Pipeline.md for supported components, defaults, and limits.
func CompileJSON(config []byte) (*Pipeline, error) {
	p, err := pipeline.Compile(config)
	if errors.Is(err, pipeline.ErrUnsupportedPattern) {
		return nil, fmt.Errorf("%w: %w", ErrUnsupportedPattern, err)
	}
	if err != nil {
		return nil, err
	}
	return &Pipeline{program: p}, nil
}

// Piece is one final pre-token after every pipeline stage. The value, its Text,
// and its alignment are borrowed and valid only during the callback. Copy text
// with strings.Clone and resolve offsets before returning from the callback.
// Different pieces may refer to overlapping ranges in the original input.
type Piece struct{ value pipeline.Piece }

// Text returns the transformed UTF-8 text to pass to the model.
func (p Piece) Text() string { return p.value.Text() }

// OriginalRange maps a non-empty half-open byte range in Text to the smallest
// covering UTF-8 byte range in the original Walk input. Endpoints inside a UTF-8
// encoding are allowed and expand to the source character's complete range.
// Prefix characters added by a stage map to that piece's first source character.
// Empty, reversed, or out-of-bounds ranges return ErrInvalidRange.
func (p Piece) OriginalRange(start, end int) (Span, error) {
	s, err := p.value.OriginalRange(start, end)
	return Span{Start: s.Start, End: s.End}, err
}

// Walk emits final pieces in input order without materializing an output list.
// It validates all input before invoking emit. Empty input emits nothing.
// Callback errors are returned unchanged and stop execution immediately; earlier
// callback side effects are not rolled back. Panics propagate to the caller.
// Error precedence is uninitialized pipeline, nil callback, then invalid UTF-8.
func (p *Pipeline) Walk(text string, emit func(Piece) error) error {
	if p == nil || p.program == nil {
		return ErrUninitialized
	}
	if emit == nil {
		return ErrNilCallback
	}
	if !utf8.ValidString(text) {
		return ErrInvalidUTF8
	}
	return p.program.Walk(text, func(v pipeline.Piece) error { return emit(Piece{value: v}) })
}

// WalkText emits exactly the same transformed pieces as Walk, without building
// the output alignment map. Stage-dependent context is still preserved. Each
// string is defensively copied and may be retained or processed asynchronously
// after the callback returns. Error, validation, concurrency, and cancellation
// behavior is the same as Walk. Use WalkTextUnsafe to avoid defensive copies.
func (p *Pipeline) WalkText(text string, emit func(string) error) error {
	if emit == nil {
		return p.WalkTextUnsafe(text, nil)
	}
	return p.WalkTextUnsafe(text, func(piece string) error {
		return emit(strings.Clone(piece))
	})
}

// WalkTextUnsafe emits the same text as WalkText without defensive copying or
// building the output alignment map. Stage-dependent context is preserved.
// Each string is valid only until its callback returns: subsequent execution
// may overwrite its backing memory. Do not retain it or pass it to asynchronous
// processing without first copying it with strings.Clone inside the callback.
// Error, validation, concurrency, and cancellation behavior is the same as Walk.
func (p *Pipeline) WalkTextUnsafe(text string, emit func(string) error) error {
	if p == nil || p.program == nil {
		return ErrUninitialized
	}
	if emit == nil {
		return ErrNilCallback
	}
	if !utf8.ValidString(text) {
		return ErrInvalidUTF8
	}
	return p.program.WalkTextUnsafe(text, emit)
}
