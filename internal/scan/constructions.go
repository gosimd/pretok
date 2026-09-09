package scan

// contractionSet implements (?i:'s|'t|'re|'ve|'m|'ll|'d). There is no
// word-boundary assertion. Case folds are generated from the pinned oracle.
type contractionSet struct{}

func (contractionSet) match(src string, pos int) int {
	if src[pos] != '\'' || pos+1 == len(src) {
		return pos
	}
	r, size := runeAt(src, pos+1)
	r = foldContraction(r)
	end := pos + 1 + size
	switch r {
	case 's', 't', 'm', 'd':
		return end
	case 'r', 'v', 'l':
		if end < len(src) {
			next, width := runeAt(src, end)
			next = foldContraction(next)
			if (r == 'r' || r == 'v') && next == 'e' || r == 'l' && next == 'l' {
				return end + width
			}
		}
	}
	return pos
}

// prefixedLetters implements [^\r\n\p{L}\p{N}]?\p{L}+.
// Prefix and body classes are disjoint in this pattern; marks belong only to
// the optional prefix, so no prefix retry or mark-body extension is needed.
type prefixedLetters struct{}

func (prefixedLetters) match(src string, pos int) int {
	r, size := runeAt(src, pos)
	end := pos
	if !letter(r) {
		if r == '\r' || r == '\n' || number(r) {
			return pos
		}
		end += size
	}
	body := end
	for end < len(src) {
		r, size = runeAt(src, end)
		if !letter(r) {
			break
		}
		end += size
	}
	if end == body {
		return pos
	}
	return end
}

// singleNumber matches one Unicode N scalar, including Nd, Nl, and No.
type singleNumber struct{}

func (singleNumber) match(src string, pos int) int {
	r, size := runeAt(src, pos)
	if number(r) {
		return pos + size
	}
	return pos
}

// symbolRun implements an optional ASCII space, a non-empty run outside
// whitespace/L/N, and a greedy suffix of CR and LF. Apostrophes inside an
// already-started symbol run do not restart the higher-priority contraction rule.
type symbolRun struct{}

func (symbolRun) match(src string, pos int) int {
	end := pos
	if src[end] == ' ' {
		end++
	}
	body := end
	for end < len(src) {
		r, size := runeAt(src, end)
		if !symbol(r) {
			break
		}
		end += size
	}
	if end == body {
		return pos
	}
	for end < len(src) && (src[end] == '\r' || src[end] == '\n') {
		end++
	}
	return end
}

// whitespaceRun shares one observation across the final three alternatives:
// \s*[\r\n]+ | \s+(?!\S) | \s+.
// It is consulted only after the higher-priority word and symbol rules.
type whitespaceRun struct{}

type whitespaceObservation struct {
	end           int
	lastRuneStart int
	lastCRLFEnd   int
}

func (whitespaceRun) observe(src string, pos int) whitespaceObservation {
	run := whitespaceObservation{end: pos, lastRuneStart: -1, lastCRLFEnd: -1}
	for run.end < len(src) {
		r, size := runeAt(src, run.end)
		if !space(r) {
			break
		}
		run.lastRuneStart = run.end
		run.end += size
		if r == '\r' || r == '\n' {
			run.lastCRLFEnd = run.end
		}
	}
	return run
}

func (w whitespaceRun) match(src string, pos int) int {
	run := w.observe(src, pos)
	if run.lastCRLFEnd >= 0 {
		return run.lastCRLFEnd
	}
	if run.end == len(src) {
		return run.end
	}
	if run.lastRuneStart > pos {
		return run.lastRuneStart
	}
	return run.end
}
