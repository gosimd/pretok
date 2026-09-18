package pipeline

import (
	"slices"
	"strings"
	"sync"
	"unicode/utf8"
	"unsafe"

	"github.com/gosimd/pretok/internal/scan"
)

// Transform buffers are indexed by stage so child stages never overwrite a
// parent's live input. One active call owns each workspace, including recursion.
type frame struct {
	prefix, output       []byte
	prefixMap, outputMap []Range
}
type workspace struct{ frames []frame }

var workspaces = sync.Pool{New: func() any { return new(workspace) }}

func (p *Program) Walk(text string, emit func(Piece) error) error {
	return p.run(text, true, emit, nil)
}

// WalkTextUnsafe lends each text view only for the duration of emit.
func (p *Program) WalkTextUnsafe(text string, emit func(string) error) error {
	return p.run(text, false, nil, emit)
}

func (p *Program) run(text string, aligned bool, emit func(Piece) error, emitText func(string) error) error {
	if text == "" {
		return nil
	}
	w := workspaces.Get().(*workspace)
	if cap(w.frames) < len(p.stages) {
		w.frames = make([]frame, len(p.stages))
	} else {
		w.frames = w.frames[:len(p.stages)]
	}
	defer func() {
		// Bound pooled retention. No input strings, pieces, or callbacks enter
		// the pool. A panic also returns or discards the exclusive workspace.
		bytes := 0
		frames := w.frames[:cap(w.frames)]
		for i := range frames {
			f := &frames[i]
			bytes += cap(f.prefix) + cap(f.output) + int(unsafe.Sizeof(Range{}))*(cap(f.prefixMap)+cap(f.outputMap))
		}
		if bytes <= 1<<20 && cap(w.frames) <= 256 {
			workspaces.Put(w)
		}
	}()
	_, width := utf8.DecodeRuneInString(text)
	r := runner{program: p, work: w, aligned: aligned, emit: emit, emitText: emitText}
	return r.walk(0, Piece{text: text, zeroEnd: width})
}

type runner struct {
	program  *Program
	work     *workspace
	aligned  bool
	emit     func(Piece) error
	emitText func(string) error
}

func (r *runner) walk(index int, piece Piece) error {
	if piece.text == "" {
		return nil
	}
	if index == len(r.program.stages) {
		if r.aligned {
			return r.emit(piece)
		}
		return r.emitText(piece.text)
	}
	s := r.program.stages[index]
	next := func(a, b int) error { return r.walk(index+1, piece.slice(a, b)) }
	switch s.kind {
	case byteLevelStage:
		f := &r.work.frames[index]
		if s.prefix && !strings.HasPrefix(piece.text, " ") {
			piece = prepend(piece, " ", f, r.aligned)
		}
		transform := func(a, b int) error {
			part := piece.slice(a, b)
			out := byteLevel(part, f, r.aligned)
			return r.walk(index+1, out)
		}
		if s.useRegex {
			return splitSegments(stage{regex: &byteLevelPlan}, piece.text, transform)
		}
		return transform(0, len(piece.text))
	case metaspaceStage:
		piece = metaspace(piece, s, &r.work.frames[index], r.aligned)
		if !s.split {
			return r.walk(index+1, piece)
		}
		return splitSegments(stage{literal: s.literal, mode: mergedNext}, piece.text, func(a, b int) error { return r.walk(index+1, piece.slice(a, b)) })
	case scriptsStage:
		return scriptPieces(piece.text, next)
	case bertStage:
		return bertPieces(piece.text, next)
	default:
		return splitSegments(s, piece.text, next)
	}
}

var byteLevelPlan = scan.ByteLevel()

// borrowedString is confined to per-call transform buffers. A buffer is never
// mutated until all descendant stages and the final callback have returned.
// WalkText clones these views before exposing them; Walk and WalkTextUnsafe
// forbid retaining them without cloning.
func borrowedString(b []byte) string { return unsafe.String(unsafe.SliceData(b), len(b)) }

func appendMapped(dst []byte, mapping []Range, text string, origin Range, aligned bool) ([]byte, []Range) {
	dst = append(dst, text...)
	if aligned {
		for range len(text) {
			mapping = append(mapping, origin)
		}
	}
	return dst, mapping
}

func prepend(p Piece, prefix string, f *frame, aligned bool) Piece {
	f.prefix, f.prefixMap = prepare(f.prefix, f.prefixMap, len(prefix)+len(p.text), aligned)
	_, width := utf8.DecodeRuneInString(p.text)
	origin := Range{}
	if aligned {
		origin = p.original(0, width)
	}
	f.prefix, f.prefixMap = appendMapped(f.prefix, f.prefixMap, prefix, origin, aligned)
	for i, c := range p.text {
		width := utf8.RuneLen(c)
		if aligned {
			origin = p.original(i, i+width)
		}
		f.prefix, f.prefixMap = appendMapped(f.prefix, f.prefixMap, p.text[i:i+width], origin, aligned)
	}
	zero := p.zeroEnd
	if zero > 0 {
		zero += len(prefix)
	}
	q := Piece{text: borrowedString(f.prefix), zeroEnd: zero}
	if aligned {
		q.align = f.prefixMap
	}
	return q
}

var byteChars = func() [256]string {
	var out [256]string
	n := rune(256)
	for b := 0; b < 256; b++ {
		c := rune(b)
		if !(b >= 33 && b <= 126 || b >= 161 && b <= 172 || b >= 174) {
			c = n
			n++
		}
		out[b] = string(c)
	}
	return out
}()

func byteLevel(p Piece, f *frame, aligned bool) Piece {
	size := len(p.text)
	for i := range len(p.text) {
		if p.text[i] < 33 || p.text[i] > 126 {
			size++
		}
	}
	f.output, f.outputMap = prepare(f.output, f.outputMap, size, aligned)
	zero := 0
	for i, c := range p.text {
		width := utf8.RuneLen(c)
		origin := Range{}
		if aligned {
			origin = p.original(i, i+width)
		}
		for j := i; j < i+width; j++ {
			f.output, f.outputMap = appendMapped(f.output, f.outputMap, byteChars[p.text[j]], origin, aligned)
		}
		if i < p.zeroEnd {
			zero = len(f.output)
		}
	}
	q := Piece{text: borrowedString(f.output), zeroEnd: zero}
	if aligned {
		q.align = f.outputMap
	}
	return q
}

func metaspace(p Piece, s stage, f *frame, aligned bool) Piece {
	zero := 0
	add := s.prepend == "always" || s.prepend == "first" && p.zeroEnd > 0
	add = add && !strings.HasPrefix(p.text, " ") && !strings.HasPrefix(p.text, s.literal)
	size := len(p.text) + strings.Count(p.text, " ")*(len(s.literal)-1)
	if add {
		size += len(s.literal)
	}
	f.output, f.outputMap = prepare(f.output, f.outputMap, size, aligned)
	if add {
		_, width := utf8.DecodeRuneInString(p.text)
		origin := Range{}
		if aligned {
			origin = p.original(0, width)
		}
		f.output, f.outputMap = appendMapped(f.output, f.outputMap, s.literal, origin, aligned)
	}
	for i, c := range p.text {
		width := utf8.RuneLen(c)
		origin := Range{}
		if aligned {
			origin = p.original(i, i+width)
		}
		text := p.text[i : i+width]
		if c == ' ' {
			text = s.literal
		}
		f.output, f.outputMap = appendMapped(f.output, f.outputMap, text, origin, aligned)
		if i < p.zeroEnd {
			zero = len(f.output)
		}
	}
	q := Piece{text: borrowedString(f.output), zeroEnd: zero}
	if aligned {
		q.align = f.outputMap
	}
	return q
}

func prepare(text []byte, mapping []Range, size int, aligned bool) ([]byte, []Range) {
	text = slices.Grow(text[:0], size)
	mapping = mapping[:0]
	if aligned {
		mapping = slices.Grow(mapping, size)
	}
	return text, mapping
}
