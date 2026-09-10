package pretok_test

import (
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/gosimd/pretok"
)

// Use the pinned N set with an independent regex engine; Go's Unicode version
// need not agree with the reference profile used by the scanners.
func pinnedNumberClass() string {
	data, err := os.ReadFile("internal/scan/unicode_tables.go")
	if err != nil {
		panic(err)
	}
	section := strings.Split(strings.Split(string(data), "var numberRanges")[1], "}\n")[0]
	var b strings.Builder
	b.WriteByte('[')
	for _, m := range regexp.MustCompile(`\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}`).FindAllStringSubmatch(section, -1) {
		b.WriteString(`\x{` + m[1] + `}-\x{` + m[2] + `}`)
	}
	b.WriteByte(']')
	return b.String()
}

var p16Reference = regexp.MustCompile(`[0-9]`)
var p07Reference = regexp.MustCompile(pinnedNumberClass() + `{1,3}`)
var oneNumber = regexp.MustCompile(`\A` + pinnedNumberClass() + `\z`)

// P05 provides all non-numeric decisions. An independent regex class selects
// adjacent single-number matches to merge; the live HF gate checks the whole plan.
func assertQwenNumbers(t testing.TB, re *pretok.Regexp, src string, max int) {
	t.Helper()
	var want []pretok.Span
	count := 0
	for _, span := range collect(t, compiled(t), src) {
		if oneNumber.MatchString(src[span.Start:span.End]) {
			if count > 0 && (max < 0 || count < max) {
				want[len(want)-1].End = span.End
				count++
				continue
			}
			count = 1
		} else {
			count = 0
		}
		want = append(want, span)
	}
	if got := collect(t, re, src); !slices.Equal(got, want) {
		t.Fatalf("%q: %v want %v", src, got, want)
	}
	if got, err := re.Split(src, nil); err != nil || !slices.Equal(got, want) {
		t.Fatalf("%q split: %v want %v: %v", src, got, want, err)
	}
}

func TestNumericPatterns(t *testing.T) {
	for id, p := range map[string]string{"P16": patternDigit, "P07": patternNumberTriples, "P24": patternQwenTriples, "P20": patternQwenRuns} {
		t.Run(id, func(t *testing.T) {
			re := compiledPattern(t, p)
			for _, input := range []string{"", "x1234567y", "A１²Ⅳ١𐒠3B", "we're １２345\r\nX", "12\u030134", "\u2003\u2003X"} {
				checkNumeric(t, re, p, input)
			}
			for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65, 4096} {
				for _, text := range []string{"0", "１", "²", "Ⅳ", "١", "𐒠", "1²Ⅳ𐒠"} {
					checkNumeric(t, re, p, "A"+strings.Repeat(text, n)+"B")
				}
			}
			for _, near := range []string{"(?:" + p + ")", p + " "} {
				if re, err := pretok.Compile(near); re != nil || err == nil {
					t.Fatalf("accepted %q", near)
				}
			}
		})
	}
}

func checkNumeric(t testing.TB, re *pretok.Regexp, p, src string) {
	t.Helper()
	switch p {
	case patternDigit:
		assertReference(t, re, src, p16Reference)
	case patternNumberTriples:
		assertReference(t, re, src, p07Reference)
	case patternQwenTriples:
		assertQwenNumbers(t, re, src, 3)
	case patternQwenRuns:
		assertQwenNumbers(t, re, src, -1)
	}
}

// Byte offsets captured from pinned HF 0.22.2 before implementation.
func TestNumericHFOffsets(t *testing.T) {
	src := "A１²Ⅳ١𐒠3B"
	for p, want := range map[string][]pretok.Span{
		patternDigit:         {{15, 16}},
		patternNumberTriples: {{1, 9}, {9, 16}},
		patternQwenTriples:   {{0, 1}, {1, 9}, {9, 16}, {16, 17}},
		patternQwenRuns:      {{0, 1}, {1, 16}, {16, 17}},
	} {
		if got := collect(t, compiledPattern(t, p), src); !slices.Equal(got, want) {
			t.Fatalf("%q: %v want %v", p, got, want)
		}
	}
}
