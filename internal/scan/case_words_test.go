package scan

import (
	"testing"
	"unicode/utf8"
)

// Explicit greedy quantifier backtracking, independent of the scanner's
// last-overlap metadata. Character membership is separately validated by HF.
func referenceCaseWord(src string) int {
	if src == "" {
		return 0
	}
	r, w := runeAt(src, 0)
	starts := []int{0}
	if r != '\r' && r != '\n' && !letters.contains(r) && !number(r) {
		starts = []int{w, 0}
	}
	for _, requireLower := range []bool{true, false} {
		for _, start := range starts {
			positions := []int{start}
			for end := start; end < len(src); {
				r, w := runeAt(src, end)
				if caseKind(r)&caseUpper == 0 {
					break
				}
				end += w
				positions = append(positions, end)
			}
			minUpper := 0
			if !requireLower {
				minUpper = 1
			}
			for i := len(positions) - 1; i >= minUpper; i-- {
				end := positions[i]
				for end < len(src) {
					r, w := runeAt(src, end)
					if caseKind(r)&caseLower == 0 {
						break
					}
					end += w
				}
				if !requireLower || end > positions[i] {
					return end
				}
			}
		}
	}
	return 0
}

func checkCaseWord(t testing.TB, src string) {
	t.Helper()
	if src == "" {
		return
	}
	got := (prefixedCaseWords{}).match(src, 0)
	if want := referenceCaseWord(src); got != want {
		t.Fatalf("%q: endpoint %d want %d", src, got, want)
	}
}

func TestCaseWordBranchPriority(t *testing.T) {
	alphabet := []rune{'A', 'a', 'ǅ', '中', 'ʰ', '\u0301', ' ', '\n', '1', '/'}
	var visit func(string, int)
	visit = func(src string, n int) {
		checkCaseWord(t, src)
		if n == 0 {
			return
		}
		for _, r := range alphabet {
			visit(src+string(r), n-1)
		}
	}
	visit("", 4)
}

func FuzzCaseWordCandidates(f *testing.F) {
	for _, s := range []string{"AA中BB", "\u0301ABC", "fooBar", "\u0301\u0301ABC", "!\u0301A", "ǅaǅA", "\xff"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, src string) {
		if len(src) > 256 || !utf8.ValidString(src) {
			return
		}
		checkCaseWord(t, src)
	})
}
