package scan

// exactSet describes only a registered literal class. ASCII membership is
// resolved at plan construction; non-ASCII ranges keep their exact endpoints.
type exactSet struct {
	ascii                 [128]bool
	ranges                []runeRange
	negated, excludeSpace bool
}

func newExactSet(ranges []runeRange, negated, excludeSpace bool) *exactSet {
	s := &exactSet{ranges: ranges, negated: negated, excludeSpace: excludeSpace}
	for r := rune(0); r < 128; r++ {
		s.ascii[r] = s.nonASCII(r)
	}
	return s
}
func (s *exactSet) nonASCII(r rune) bool {
	inside := inRanges(r, s.ranges) || s.excludeSpace && space(r)
	return inside != s.negated
}
func (s *exactSet) contains(r rune) bool {
	if r < 128 {
		return s.ascii[r]
	}
	return s.nonASCII(r)
}

type setPrefix uint8

const (
	noSetPrefix setPrefix = iota
	spaceSetPrefix
	whitespaceSetPrefix
)

type exactRun struct {
	set         *exactSet
	prefix      setPrefix
	singleASCII bool
}

func (p exactRun) isPrefix(r rune) bool {
	return p.prefix == spaceSetPrefix && r == ' ' || p.prefix == whitespaceSetPrefix && space(r)
}
func (p exactRun) find(src string, pos int) (int, int) {
	// P11 has a one-byte body and no prefix. UTF-8 continuation bytes cannot
	// equal CR/LF, so unmatched multibyte text needs no decoding here.
	if p.singleASCII {
		for i := pos; i < len(src); i++ {
			if src[i] < 128 && p.set.ascii[src[i]] {
				return i, i + 1
			}
		}
		return len(src), len(src)
	}
	prefixStart := -1
	for pos < len(src) {
		r, size := runeAt(src, pos)
		if p.set.contains(r) {
			start := pos
			if prefixStart >= 0 {
				start = prefixStart
			}
			pos = p.set.runEnd(src, pos+size)
			// A prefix that is also a body member (P18 U+3000) is part
			// of this same maximal run. Even a singleton succeeds: this
			// is the endpoint after the regex backs out its optional prefix.
			return start, pos
		}
		prefixStart = -1
		if p.isPrefix(r) {
			prefixStart = pos
		}
		pos += size
	}
	return len(src), len(src)
}
func exactPlan(ranges []runeRange, prefix setPrefix) Plan {
	p := exactRun{set: newExactSet(ranges, false, false), prefix: prefix}
	return Plan{find: p.find}
}
func Newline() Plan {
	p := exactRun{set: newExactSet([]runeRange{{'\n', '\n'}, {'\r', '\r'}}, false, false), singleASCII: true}
	return Plan{find: p.find}
}
func CJK() Plan              { return exactPlan(cjkRanges[:], noSetPrefix) }
func BroadRanges() Plan      { return exactPlan(broadScriptRanges[:], noSetPrefix) }
func PrefixedSymbols() Plan  { return exactPlan(prefixedSymbolsRanges[:], whitespaceSetPrefix) }
func PrefixedAlphabet() Plan { return exactPlan(prefixedAlphabetRanges[:], whitespaceSetPrefix) }
func Complement() Plan {
	// Oniguruma treats P23's inner bracket expression as a nested union.
	// Parentheses and pipe are literals; square brackets are not excluded.
	ranges := []runeRange{{'!', '!'}, {'(', ')'}, {',', ','}, {'.', '.'}, {'?', '?'}, {'|', '|'}, {0x60C, 0x60C}, {0x6D4, 0x6D4}, {0x964, 0x964}, {0x2026, 0x2026}, {0x3001, 0x3002}, {0xFF0C, 0xFF0C}}
	p := exactRun{set: newExactSet(ranges, true, true), prefix: spaceSetPrefix}
	return Plan{find: p.find}
}

// runEnd scans a possibly empty suffix within one fixed set. Plain, grouped,
// and P22 runs share this loop; classification never changes mid-run.
func (s *exactSet) runEnd(src string, pos int) int {
	for pos < len(src) {
		r, size := runeAt(src, pos)
		if !s.contains(r) {
			break
		}
		pos += size
	}
	return pos
}
