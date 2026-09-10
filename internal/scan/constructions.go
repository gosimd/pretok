package scan

// asciiTriplet matches exactly three ASCII digits. The length and each byte
// are checked before consumption; a failed partial match consumes nothing.
type asciiTriplet struct{}

func (asciiTriplet) match(src string, pos int) int {
	if len(src)-pos >= 3 && asciiDigit(src[pos]) && asciiDigit(src[pos+1]) && asciiDigit(src[pos+2]) {
		return pos + 3
	}
	return pos
}

func asciiDigit(b byte) bool { return b >= '0' && b <= '9' }

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

// prefixedLetters implements the optional non-CR/LF/L/N prefix followed by
// L+ (P05) or (L|M)+ (P01). A leading mark belongs to both the P01 prefix
// and body. Consuming it as body selects the same endpoint as greedy prefix
// consumption followed by a retry without the prefix when the body is empty.
type prefixedLetters struct{ class *wordClass }

func (p prefixedLetters) match(src string, pos int) int {
	end := pos
	haveBody := false
	for end < len(src) {
		r, size := runeAt(src, end)
		if p.class.contains(r) {
			haveBody = true
			end += size
			continue
		}
		if end == pos && r != '\r' && r != '\n' && !number(r) {
			end += size // optional prefix, still requiring a non-empty body
			continue
		}
		break
	}
	if !haveBody {
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
// whitespace/N/word-class, and a greedy suffix of CR and LF. Apostrophes inside an
// already-started symbol run do not restart the higher-priority contraction rule.
// P04 permits slash in the suffix as well.
type symbolRun struct {
	class       *wordClass
	slashSuffix bool // P04 extends the CR/LF suffix to [\r\n/]*.
}

func (p symbolRun) match(src string, pos int) int {
	end := pos
	if src[end] == ' ' {
		end++
	}
	body := end
	for end < len(src) {
		r, size := runeAt(src, end)
		if space(r) || p.class.contains(r) || number(r) {
			break
		}
		end += size
	}
	if end == body {
		return pos
	}
	for end < len(src) && (src[end] == '\r' || src[end] == '\n' || p.slashSuffix && src[end] == '/') {
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
	lastLFStart   int
}

func (whitespaceRun) observe(src string, pos int) whitespaceObservation {
	run := whitespaceObservation{end: pos, lastRuneStart: -1, lastCRLFEnd: -1, lastLFStart: -1}
	for run.end < len(src) {
		r, size := runeAt(src, run.end)
		if !space(r) {
			break
		}
		run.lastRuneStart = run.end
		if r == '\n' {
			run.lastLFStart = run.end
		}
		run.end += size
		if r == '\r' || r == '\n' {
			run.lastCRLFEnd = run.end
		}
	}
	return run
}

// lineEndWhitespace implements P03's \s+$ under the pinned HF/Oniguruma
// profile: $ asserts EOF or the position before any LF, without consuming LF.
// A greedy whitespace run at EOF is consumed whole. Otherwise the last LF
// supplies the rightmost possible endpoint, provided \s+ consumed something.
type lineEndWhitespace struct{}

func (lineEndWhitespace) find(src string, pos int) (int, int) {
	for pos < len(src) {
		r, width := runeAt(src, pos)
		if !space(r) {
			pos += width
			continue
		}
		run := (whitespaceRun{}).observe(src, pos)
		if run.end == len(src) {
			return pos, run.end
		}
		if run.lastLFStart > pos {
			return pos, run.lastLFStart
		}
		// No suffix starting inside this run can satisfy the anchor either.
		pos = run.end
	}
	return len(src), len(src)
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
