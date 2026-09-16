package scan

import (
	"unicode/utf8"

	simdutf8 "github.com/gosimd/unicode/utf8"
)

// RuneBufferSize is the per-Split P01 scratch capacity, in Unicode scalars.
const RuneBufferSize = 1024

// RuneSplitMinBytes keeps short inputs on the byte scanner without a rune buffer.
const RuneSplitMinBytes = 256

// QwenMarksCursor scans P01 using a caller-owned rune window. It is internal to
// Split: unlike the public Cursor, its copies must not share the mutable buffer.
// Input must already be valid UTF-8. Selected matches cover the complete input.
type QwenMarksCursor struct {
	src                string
	buf                []rune
	i, n, decoded, pos int
}

func (q *QwenMarksCursor) peek(k int) rune {
	if i := q.i + k; i < q.n {
		return q.buf[i]
	}
	return q.more(k)
}

// Keep refilling out of peek's inline budget: an ordinary rune access must be
// an indexed load, not a call. The buffer has room for contraction lookahead.
//
//go:noinline
func (q *QwenMarksCursor) more(k int) rune {
	q.fill(k)
	if q.i+k >= q.n {
		return -1
	}
	return q.buf[q.i+k]
}
func (q *QwenMarksCursor) fill(k int) {
	if q.decoded == len(q.src) {
		return
	}
	keep := copy(q.buf, q.buf[q.i:q.n])
	q.i = 0
	q.n = keep
	for q.n <= k && q.decoded < len(q.src) {
		n, b := simdutf8.DecodeValidInto(q.buf[q.n:], q.src[q.decoded:])
		q.n += n
		q.decoded += b
	}
}
func (q *QwenMarksCursor) take() { r := q.buf[q.i]; q.i++; q.pos += utf8.RuneLen(r) }

// Rewind within the decoded window when possible. Discarding the entire window
// for the whitespace branch's one-rune backtrack would repeatedly decode text.
func (q *QwenMarksCursor) reset(pos int) {
	for q.i > 0 && q.pos > pos {
		q.i--
		q.pos -= utf8.RuneLen(q.buf[q.i])
	}
	if q.pos == pos {
		return
	}
	q.pos = pos
	q.decoded = pos
	q.i = 0
	q.n = 0
}
func marksSymbol(r rune) bool {
	return r >= 0 && !space(r) && !lettersAndMarks.contains(r) && !number(r)
}
func runeContraction(q *QwenMarksCursor) int {
	if q.peek(0) != '\'' {
		return 0
	}
	r := foldContraction(q.peek(1))
	if singleContraction(r) {
		return 2
	}
	if suffix := contractionSuffix(r); suffix != 0 && foldContraction(q.peek(2)) == suffix {
		return 3
	}
	return 0
}
func (q *QwenMarksCursor) match() int {
	if n := runeContraction(q); n > 0 {
		for i := 0; i < n; i++ {
			q.take()
		}
		return q.pos
	}
	r := q.peek(0)
	if lettersAndMarks.contains(r) {
		q.take()
		for lettersAndMarks.contains(q.peek(0)) {
			q.take()
		}
		return q.pos
	}
	if r != '\r' && r != '\n' && !number(r) && lettersAndMarks.contains(q.peek(1)) {
		q.take()
		q.take()
		for lettersAndMarks.contains(q.peek(0)) {
			q.take()
		}
		return q.pos
	}
	if number(r) {
		q.take()
		return q.pos
	}
	if marksSymbol(r) || (r == ' ' && marksSymbol(q.peek(1))) {
		q.take()
		for marksSymbol(q.peek(0)) {
			q.take()
		}
		for q.peek(0) == '\r' || q.peek(0) == '\n' {
			q.take()
		}
		return q.pos
	}
	start := q.pos
	run := whitespaceObservation{lastRuneStart: -1, lastCRLFEnd: -1}
	for r = q.peek(0); r >= 0 && space(r); r = q.peek(0) {
		run.lastRuneStart = q.pos
		q.take()
		if r == '\r' || r == '\n' {
			run.lastCRLFEnd = q.pos
		}
	}
	run.end = q.pos
	end := run.matchEnd(start, len(q.src))
	if end < q.pos {
		q.reset(end)
	}
	return end
}

// NewQwenMarksCursor borrows src and scratch until scanning finishes. Scratch
// must have room for the three-rune contraction lookahead.
func NewQwenMarksCursor(src string, scratch []rune) QwenMarksCursor {
	if len(scratch) < 3 {
		panic("pretok: rune scratch too small")
	}
	return QwenMarksCursor{src: src, buf: scratch}
}

// Next returns the selected match's byte boundaries; no window splits a token.
func (q *QwenMarksCursor) Next() (start, end int, ok bool) {
	if q.pos == len(q.src) {
		return 0, 0, false
	}
	start = q.pos
	return start, q.match(), true
}
