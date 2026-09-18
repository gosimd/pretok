package pipeline

import (
	"strings"
	"unicode/utf8"

	"github.com/gosimd/pretok/internal/scan"
)

// segments streams selected matches and maximal gaps, before delimiter policy.
// Regex receives the entire previous-stage piece (anchors and lookaround matter).
func (s stage) segments(src string, emit func(int, int, bool) error) error {
	if s.kind == whitespaceStage {
		for pos := 0; pos < len(src); {
			start := pos
			c, w := utf8.DecodeRuneInString(src[pos:])
			mask := classify(c)
			pos += w
			for pos < len(src) {
				c, w = utf8.DecodeRuneInString(src[pos:])
				next := classify(c)
				if next&spaceMask != mask&spaceMask || mask&spaceMask == 0 && next&wordMask != mask&wordMask {
					break
				}
				pos += w
			}
			if err := emit(start, pos, mask&spaceMask != 0); err != nil {
				return err
			}
		}
		return nil
	}
	pos, gap := 0, 0
	state := scan.State{}
	for pos < len(src) {
		start, end := pos, pos
		switch {
		case s.regex != nil:
			if s.regex.HasSearch() {
				start, end, state = s.regex.Find(src, pos, state)
			} else {
				for start < len(src) {
					end = s.regex.Match(src, start)
					if end > start {
						break
					}
					_, w := utf8.DecodeRuneInString(src[start:])
					start += w
				}
				if start == len(src) {
					end = start
				}
			}
		case s.literal != "":
			i := strings.Index(src[pos:], s.literal)
			if i < 0 {
				start = len(src)
				end = start
			} else {
				start = pos + i
				end = start + len(s.literal)
			}
		default:
			for start < len(src) {
				c, w := utf8.DecodeRuneInString(src[start:])
				mask := classify(c)
				match := s.kind == digitsStage && mask&numberMask != 0 || s.kind == punctuationStage && mask&punctMask != 0 || s.kind == whitespaceSplitStage && mask&spaceMask != 0
				if match {
					end = start + w
					break
				}
				start += w
			}
			if start == len(src) {
				end = start
			}
		}
		if end <= start {
			break
		}
		if start > gap {
			if err := emit(gap, start, s.invert); err != nil {
				return err
			}
		}
		if err := emit(start, end, !s.invert); err != nil {
			return err
		}
		gap = end
		pos = end
	}
	if gap < len(src) {
		return emit(gap, len(src), s.invert)
	}
	return nil
}

func splitSegments(s stage, src string, emit func(int, int) error) error {
	mode := s.mode
	if s.kind == whitespaceStage || s.kind == whitespaceSplitStage {
		mode = removed
	}
	if s.kind == digitsStage && !s.individual {
		mode = contiguous
	}
	start, end := -1, 0
	previous := false
	err := s.segments(src, func(a, b int, match bool) error {
		if mode == removed {
			if !match {
				return emit(a, b)
			}
			return nil
		}
		if mode == isolated {
			return emit(a, b)
		}
		merge := start >= 0 && (mode == contiguous && previous == match || mode == mergedPrevious && match && !previous || mode == mergedNext && previous && !match)
		if merge {
			end = b
		} else {
			if start >= 0 {
				if err := emit(start, end); err != nil {
					return err
				}
			}
			start, end = a, b
		}
		previous = match
		return nil
	})
	if err != nil {
		return err
	}
	if start >= 0 {
		return emit(start, end)
	}
	return nil
}

func bertPieces(src string, emit func(int, int) error) error {
	start := 0
	for i, c := range src {
		mask := classify(c)
		if mask&(spaceMask|punctMask) == 0 {
			continue
		}
		if start < i {
			if err := emit(start, i); err != nil {
				return err
			}
		}
		end := i + utf8.RuneLen(c)
		if mask&spaceMask == 0 {
			if err := emit(i, end); err != nil {
				return err
			}
		}
		start = end
	}
	if start < len(src) {
		return emit(start, len(src))
	}
	return nil
}

func scriptPieces(src string, emit func(int, int) error) error {
	// HF starts its first range at the first non-Any character. Leading Any
	// characters (including spaces), or an all-Any subject, produce no piece.
	start, last := -1, uint16(0)
	for i, c := range src {
		script := scriptOf(c)
		if script == 0 {
			continue
		}
		if start < 0 {
			start = i
		} else if script != last {
			if err := emit(start, i); err != nil {
				return err
			}
			start = i
		}
		last = script
	}
	if start >= 0 {
		return emit(start, len(src))
	}
	return nil
}

const (
	wordMask uint8 = 1 << iota
	numberMask
	spaceMask
	punctMask
)

type classRange struct {
	lo, hi rune
	mask   uint8
}
type scriptRange struct {
	lo, hi rune
	script uint16
}

func classify(c rune) uint8 {
	if c < 128 {
		return asciiClasses[c]
	}
	lo, hi := 0, len(classRanges)
	for lo < hi {
		mid := lo + (hi-lo)/2
		r := classRanges[mid]
		if c < r.lo {
			hi = mid
		} else if c > r.hi {
			lo = mid + 1
		} else {
			return r.mask
		}
	}
	return 0
}

func scriptOf(c rune) uint16 {
	lo, hi := 0, len(scriptRanges)
	for lo < hi {
		mid := lo + (hi-lo)/2
		r := scriptRanges[mid]
		if c < r.lo {
			hi = mid
		} else if c > r.hi {
			lo = mid + 1
		} else {
			return r.script
		}
	}
	return 0
}
