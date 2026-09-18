package pipeline

import "unicode/utf8"

type Range struct{ Start, End int }

// Piece uses identity offsets until a transform actually needs an alignment.
// zeroEnd is the prefix originating in the first original scalar. It preserves
// Metaspace First context even when text-only walking builds no alignment map.
type Piece struct {
	text    string
	base    int
	align   []Range
	zeroEnd int
}

func (p Piece) Text() string { return p.text }

func (p Piece) OriginalRange(start, end int) (Range, error) {
	if start < 0 || end <= start || end > len(p.text) {
		return Range{}, ErrInvalidRange
	}
	return p.original(start, end), nil
}

func (p Piece) original(start, end int) Range {
	if p.align != nil {
		return Range{p.align[start].Start, p.align[end-1].End}
	}
	for start > 0 && !utf8.RuneStart(p.text[start]) {
		start--
	}
	for end < len(p.text) && !utf8.RuneStart(p.text[end]) {
		end++
	}
	return Range{p.base + start, p.base + end}
}

func (p Piece) slice(start, end int) Piece {
	q := Piece{text: p.text[start:end], base: p.base + start, zeroEnd: max(0, min(end, p.zeroEnd)-start)}
	if p.align != nil {
		q.align = p.align[start:end]
	}
	return q
}
