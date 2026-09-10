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

func (c *wordClass) contains(r rune) bool {
	if r < utf8.RuneSelf {
		return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
	}
	return inRanges(r, c.ranges)
}

func number(r rune) bool {
	if r < utf8.RuneSelf {
		return r >= '0' && r <= '9'
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
