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
	if singleContraction(r) {
		return end
	}
	if suffix := contractionSuffix(r); suffix != 0 && end < len(src) {
		next, width := runeAt(src, end)
		if foldContraction(next) == suffix {
			return end + width
		}
	}
	return pos
}

// Both byte and rune readers use these predicates on already-folded runes.
func singleContraction(r rune) bool { return r == 's' || r == 't' || r == 'm' || r == 'd' }

func contractionSuffix(r rune) rune {
	switch r {
	case 'r', 'v':
		return 'e'
	case 'l':
		return 'l'
	}
	return 0
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

// singleNumber is the shared fast path for one Unicode N scalar.
type singleNumber struct{}

func (singleNumber) match(src string, pos int) int {
	r, size := runeAt(src, pos)
	if number(r) {
		return pos + size
	}
	return pos
}

// numberRun consumes greedy scalar counts, not bytes. max == -1 is unbounded.
// All registered plans require a positive min; failed partial runs consume nothing.
// ASCII and Unicode N are distinct classes; N includes Nd, Nl and No.
type numberRun struct {
	ascii    bool
	min, max int
}

// find skips unmatched input without re-decoding it in the generic cursor.
// In valid UTF-8 an ASCII digit cannot occur inside a multibyte encoding, so
// ASCII search may inspect bytes; returned positions remain scalar boundaries.
func (p numberRun) find(src string, pos int) (int, int) {
	if p.ascii {
		for pos < len(src) {
			if asciiDigit(src[pos]) {
				if end := p.matchASCII(src, pos); end > pos {
					return pos, end
				}
			}
			pos++
		}
	} else {
		for pos < len(src) {
			r, width := runeAt(src, pos)
			if number(r) {
				end := pos + width
				if p.max != 1 {
					end = p.rest(src, pos, end)
				}
				if end > pos {
					return pos, end
				}
			}
			pos += width
		}
	}
	return len(src), len(src)
}

func (p numberRun) matchASCII(src string, pos int) int {
	end := pos
	for end < len(src) && (p.max < 0 || end-pos < p.max) && asciiDigit(src[end]) {
		end++
	}
	if end-pos < p.min {
		return pos
	}
	return end
}

func (p numberRun) rest(src string, start, end int) int {
	count := 1
	for end < len(src) && (p.max < 0 || count < p.max) {
		r, size := runeAt(src, end)
		if !number(r) {
			break
		}
		end += size
		count++
	}
	if count < p.min {
		return start
	}
	return end
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
	return run.matchEnd(pos, len(src))
}

// matchEnd preserves the priority of the three shared whitespace alternatives.
func (run whitespaceObservation) matchEnd(pos, inputLen int) int {
	if run.lastCRLFEnd >= 0 {
		return run.lastCRLFEnd
	}
	if run.end == inputLen {
		return run.end
	}
	if run.lastRuneStart > pos {
		return run.lastRuneStart
	}
	return run.end
}
