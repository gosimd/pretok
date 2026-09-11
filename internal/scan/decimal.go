package scan

// decimalRun describes an Nd run in UTF-8 bytes and Unicode scalars. A bounded
// scan can serve consecutive groups; an unbounded scan supplies alignment data.
type decimalRun struct {
	End, Count int
}

func decimal(r rune) bool {
	if r < 128 {
		return asciiDigit(byte(r))
	}
	return inRanges(r, decimalRanges[:])
}

// scanDecimals consumes at most limit scalars; zero means an unbounded run.
// Nd is deliberately distinct from the existing N and ASCII digit classes.
func scanDecimals(src string, pos, limit int) decimalRun {
	run := decimalRun{End: pos}
	for run.End < len(src) && (limit == 0 || run.Count < limit) {
		r, width := runeAt(src, run.End)
		if !decimal(r) {
			break
		}
		run.End += width
		run.Count++
	}
	return run
}

// InitialDecimalRemainder builds P10 under the pinned HF/Oniguruma syntax:
// {3}+ repeats exact triples, rather than making {3} possessive. Both anchors
// refer to the whole supplied subject, never a suffix starting at searchStart.
func InitialDecimalRemainder() Plan {
	return Plan{find: initialDecimalRemainder}
}

func initialDecimalRemainder(src string, searchStart int) (int, int) {
	if searchStart == 0 {
		run := scanDecimals(src, 0, 0)
		if run.End == len(src) && run.Count >= 4 && run.Count%3 != 0 {
			// The lookahead consumes no input. Only the leading remainder is
			// returned; the following triples remain one unmatched Split gap.
			prefix := scanDecimals(src, 0, run.Count%3)
			return 0, prefix.End
		}
	}
	return len(src), len(src)
}

// ConsecutiveDecimalTriples builds P14. Initially both anchors select byte zero;
// after a match only \G selects the current search start. A failed triple must
// end iteration, never search for another digit run beyond a gap.
func ConsecutiveDecimalTriples() Plan {
	return Plan{find: consecutiveDecimalTriples}
}

func consecutiveDecimalTriples(src string, searchStart int) (int, int) {
	run := scanDecimals(src, searchStart, 3)
	if run.Count == 3 {
		return searchStart, run.End
	}
	return len(src), len(src)
}

// RightAlignedDecimals builds P25. The first alternative always matches at an
// Nd scalar: its longest acceptable prefix leaves a multiple of 510 digits.
// The later \G alternative therefore cannot change the selected match.
func RightAlignedDecimals() Plan {
	return Plan{findState: rightAlignedDecimals}
}

const decimalGroupSize = 510

func rightAlignedDecimals(src string, searchStart int, state State) (int, int, State) {
	if searchStart < state.decimalEnd {
		// After the first group the cached remainder is a multiple of 510.
		group := scanDecimals(src, searchStart, decimalGroupSize)
		return searchStart, group.End, state
	}
	state = State{}
	for pos := searchStart; pos < len(src); {
		r, width := runeAt(src, pos)
		if !decimal(r) {
			pos += width
			continue
		}
		run := scanDecimals(src, pos, 0)
		state.decimalEnd = run.End
		if run.Count <= decimalGroupSize {
			return pos, run.End, state
		}
		first := (run.Count-1)%decimalGroupSize + 1
		group := scanDecimals(src, pos, first)
		return pos, group.End, state
	}
	return len(src), len(src), state
}
