package pretok_test

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/gosimd/pretok"
)

// Selected pieces captured from HF Tokenizers 0.22.2 before implementation.
func TestCaseWordSuffixHF(t *testing.T) {
	for id, pattern := range map[string]string{"P19": patternCaseWordsSuffix, "P06": patternCaseWordsSuffixTriples} {
		t.Run(id, func(t *testing.T) {
			re := compiledPattern(t, pattern)
			for _, tc := range []struct {
				src    string
				pieces []string
			}{
				{"we're", []string{"we're"}},
				{"AA中BB's", []string{"AA中", "BB's"}},
				{"\u0301ABC's", []string{"\u0301", "ABC's"}},
				{"A'S'S", []string{"A'S", "'S"}},
				{"1's", []string{"1", "'s"}},
				{"!'s", []string{"!'", "s"}},
				{"'s", []string{"'s"}},
				{"a'ſ", []string{"a'ſ"}},
			} {
				var want []pretok.Span
				pos := 0
				for _, piece := range tc.pieces {
					want = append(want, pretok.Span{Start: pos, End: pos + len(piece)})
					pos += len(piece)
				}
				if strings.Join(tc.pieces, "") != tc.src {
					t.Fatal("invalid fixture")
				}
				if got := collect(t, re, tc.src); !slices.Equal(got, want) {
					t.Fatalf("%q: %v want %v", tc.src, got, want)
				}
				if got, err := re.Split(tc.src, nil); err != nil || !slices.Equal(got, want) {
					t.Fatalf("%q split %v want %v: %v", tc.src, got, want, err)
				}
			}
			for _, near := range []string{"(?:" + pattern + ")", pattern + " ", strings.ReplaceAll(pattern, "(?i:", "(?:"), strings.Replace(pattern, `[\r\n/]*`, `[\r\n]*`, 1)} {
				if re, err := pretok.Compile(near); re != nil || !errors.Is(err, pretok.ErrUnsupportedPattern) {
					t.Fatalf("accepted %q", near)
				}
			}
		})
	}
	src := "A１²Ⅳ١𐒠3B"
	for pattern, want := range map[string][]pretok.Span{
		patternCaseWordsSuffix:        {{0, 1}, {1, 4}, {4, 6}, {6, 9}, {9, 11}, {11, 15}, {15, 16}, {16, 17}},
		patternCaseWordsSuffixTriples: {{0, 1}, {1, 9}, {9, 16}, {16, 17}},
	} {
		if got := collect(t, compiledPattern(t, pattern), src); !slices.Equal(got, want) {
			t.Fatalf("%q: %v want %v", pattern, got, want)
		}
	}
}
