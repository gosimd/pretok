package scan

import "unicode/utf8"

type runeRange struct{ lo, hi rune }

func inRanges(r rune, ranges []runeRange) bool {
	lo, hi := 0, len(ranges)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if r < ranges[mid].lo {
			hi = mid
		} else if r > ranges[mid].hi {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

// wordClass is chosen once by the plan. Classification has one table lookup,
// with the same ASCII fast path for both L and L|M (marks are non-ASCII).
type wordClass struct{ ranges []runeRange }

var letters = wordClass{ranges: letterRanges[:]}
var lettersAndMarks = wordClass{ranges: wordMarkRanges[:]}

const (
	caseUpper uint8 = 1 // Lu | Lt | Lm | Lo | M
	caseLower uint8 = 2 // Ll | Lm | Lo | M
)

type caseRange struct {
	lo, hi rune
	kind   uint8
}

// caseKind classifies both overlapping P04 word sets with one lookup.
func caseKind(r rune) uint8 {
	if r < utf8.RuneSelf {
		if r >= 'A' && r <= 'Z' {
			return caseUpper
		}
		if r >= 'a' && r <= 'z' {
			return caseLower
		}
		return 0
	}
	lo, hi := 0, len(caseWordRanges)
	for lo < hi {
		mid := lo + (hi-lo)/2
		x := caseWordRanges[mid]
		if r < x.lo {
			hi = mid
		} else if r > x.hi {
			lo = mid + 1
		} else {
			return x.kind
		}
	}
	return 0
}

func (c *wordClass) contains(r rune) bool {
	if r < utf8.RuneSelf {
		return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
	}
	return inRanges(r, c.ranges)
}

func number(r rune) bool {
	if r < utf8.RuneSelf {
		return asciiDigit(byte(r))
	}
	return inRanges(r, numberRanges[:])
}

func space(r rune) bool {
	if r < utf8.RuneSelf {
		return r == ' ' || r >= '\t' && r <= '\r'
	}
	return inRanges(r, spaceRanges[:])
}

func runeAt(src string, pos int) (rune, int) {
	if src[pos] < utf8.RuneSelf {
		return rune(src[pos]), 1
	}
	return utf8.DecodeRuneInString(src[pos:])
}
