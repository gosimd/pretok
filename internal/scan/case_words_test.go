package scan

import (
	"regexp"
	"testing"
	"unicode/utf8"
)

// This independent regex has the same folds for the contraction alphabet
// (including long s); it does not call the production contraction matcher.
var referenceSuffix = regexp.MustCompile(`\A(?i:'s|'t|'re|'ve|'m|'ll|'d)`)

func checkCaseWordSuffix(t testing.TB, src string) {
	t.Helper()
	if src == "" {
		return
	}
	wordEnd := referenceCaseWord(src)
	for _, max := range []int{1, 3} {
		p := newCaseWordPlan(max)
		want := wordEnd
		if wordEnd > 0 {
			if suffix := referenceSuffix.FindStringIndex(src[wordEnd:]); suffix != nil {
				want += suffix[1]
			}
		} else {
			want = p.tail.match(src, 0)
		}
		if got := p.matchWithSuffix(src, 0); got != want {
			t.Fatalf("%q max=%d: %d want %d", src, max, got, want)
		}
	}
}

func TestCaseWordSuffixPriority(t *testing.T) {
	var visit func(string, int)
	visit = func(word string, depth int) {
		for _, suffix := range []string{"'s", "'RE", "'ſ", "'ll", "'r", "'x", "'s's", "’s", ""} {
			checkCaseWordSuffix(t, word+suffix+"ABC")
		}
		if depth == 0 {
			return
		}
		for _, r := range []rune{'A', 'a', '中', '\u0301', ' ', '\n', '1', '!', '\''} {
			visit(word+string(r), depth-1)
		}
	}
	visit("", 4)
}

func FuzzCaseWordSuffix(f *testing.F) {
	for _, s := range []string{"we're", "AA中BB's", "\u0301ABC's", "A'S'S", "!'s", "a'ſ", "\xff"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, src string) {
		if len(src) > 256 || !utf8.ValidString(src) {
			return
		}
		checkCaseWordSuffix(t, src)
	})
}

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
