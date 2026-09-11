package pretok_test

import (
	"errors"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

const patternP22 = "[!\"#$%&'()*+,\\-./:;<=>?@\\[\\\\\\]^_`{|}~][A-Za-z]+|[^\r\n\\p{L}\\p{P}\\p{S}]?[\\p{L}\\p{M}]+| ?[\\p{P}\\p{S}]+[\r\n]*|\\s*[\r\n]+|\\s+(?!\\S)|\\s+"

// Pinned class data with independent RE2 matching and explicit lookahead
// evaluation. The full native HF gate independently checks these tables/choices.
func referenceRangeBody(path, name string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	section := strings.Split(strings.Split(string(data), "var "+name)[1], "}\n")[0]
	var b strings.Builder
	for _, m := range regexp.MustCompile(`\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}`).FindAllStringSubmatch(section, -1) {
		b.WriteString(`\x{` + m[1] + `}-\x{` + m[2] + `}`)
	}
	return b.String()
}

var p22Head = func() *regexp.Regexp {
	head, _, _ := strings.Cut(patternP22, `|\s*`)
	head = strings.ReplaceAll(head, `[\p{L}\p{M}]`, `[`+referenceRangeBody("internal/scan/word_tables.go", "wordMarkRanges")+`]`)
	head = strings.ReplaceAll(head, `\p{L}`, referenceRangeBody("internal/scan/unicode_tables.go", "letterRanges"))
	head = strings.ReplaceAll(head, `\p{P}\p{S}`, referenceRangeBody("internal/scan/punctuation_tables.go", "punctuationSymbolRanges"))
	return regexp.MustCompile(`\A(?:` + head + `)`)
}()
var p22Newline = regexp.MustCompile(`\A[` + referenceSpace + `]*[\r\n]+`)
var p22Whitespace = regexp.MustCompile(`\A[` + referenceSpace + `]+`)

func assertP22Reference(t testing.TB, re *pretok.Regexp, src string) {
	t.Helper()
	var want, split []pretok.Span
	previous := 0
	for pos := 0; pos < len(src); {
		tail := src[pos:]
		end := 0
		if m := p22Head.FindStringIndex(tail); m != nil {
			end = m[1]
		} else if m := p22Newline.FindStringIndex(tail); m != nil {
			end = m[1]
		} else if m := p22Whitespace.FindStringIndex(tail); m != nil {
			end = m[1]
			// With a following non-space, (?!\S) can succeed only by giving
			// back the last whitespace scalar. A singleton uses the final \s+.
			if end < len(tail) {
				_, width := utf8.DecodeLastRuneInString(tail[:end])
				if end > width {
					end -= width
				}
			}
		}
		if end == 0 {
			_, width := utf8.DecodeRuneInString(tail)
			pos += width
			continue
		}
		span := pretok.Span{Start: pos, End: pos + end}
		want = append(want, span)
		if pos > previous {
			split = append(split, pretok.Span{Start: previous, End: pos})
		}
		split = append(split, span)
		previous = span.End
		pos = span.End
	}
	if previous < len(src) {
		split = append(split, pretok.Span{Start: previous, End: len(src)})
	}
	if got := collect(t, re, src); !slices.Equal(got, want) {
		t.Fatalf("%q matches %v want %v", src, got, want)
	}
	if got, err := re.Split(src, nil); err != nil || !slices.Equal(got, split) {
		t.Fatalf("%q split %v want %v: %v", src, got, split, err)
	}
}
func TestP22HFExamples(t *testing.T) {
	re := compiledPattern(t, patternP22)
	for _, tc := range []struct {
		src  string
		want []pretok.Span
	}{
		{"!abcé", []pretok.Span{{0, 4}, {4, 6}}},
		{"123abc", []pretok.Span{{2, 6}}},
		{" 123abc", []pretok.Span{{0, 1}, {3, 7}}},
		{"!a\u0301", []pretok.Span{{0, 2}, {2, 4}}},
		{"1\u0301a", []pretok.Span{{0, 4}}},
		{"\x00abc", []pretok.Span{{0, 4}}},
		{"\u200dabc", []pretok.Span{{0, 6}}},
		{"😀abc", []pretok.Span{{0, 4}, {4, 7}}},
		{"!😀abc", []pretok.Span{{0, 5}, {5, 8}}},
		{"!!abc", []pretok.Span{{0, 2}, {2, 5}}},
		{"\u2003\u2003X", []pretok.Span{{0, 3}, {3, 7}}},
		{"\u0301", []pretok.Span{{0, 2}}},
		{"12\u0301", []pretok.Span{{1, 4}}},
		{"١٢abc", []pretok.Span{{2, 7}}},
	} {
		if got := collect(t, re, tc.src); !slices.Equal(got, tc.want) {
			t.Fatalf("%q: %v want %v", tc.src, got, tc.want)
		}
		assertP22Reference(t, re, tc.src)
	}
}
func TestP22Transitions(t *testing.T) {
	re := compiledPattern(t, patternP22)
	for c := rune(0); c < 128; c++ {
		for _, tail := range []string{"abcé", "a\u0301", "é", "中", "!a", "1abc", "\r\n"} {
			assertP22Reference(t, re, string(c)+tail)
		}
	}
	alphabet := []rune{'A', 'é', '中', '1', '١', '!', '😀', '\u0301', ' ', '\r', '\n', '\u200d'}
	var visit func(string, int)
	visit = func(s string, left int) {
		assertP22Reference(t, re, s)
		if left > 0 {
			for _, r := range alphabet {
				visit(s+string(r), left-1)
			}
		}
	}
	visit("", 4)
	for _, n := range []int{1, 2, 3, 15, 16, 17, 31, 32, 33, 63, 64, 65, 4096} {
		for _, body := range []string{"1", "١", "a", "é", "中", "!", "😀", "\u0301", " ", "\u2003"} {
			assertP22Reference(t, re, "!"+strings.Repeat(body, n)+"abcé\r\nX")
		}
	}
	for _, near := range []string{patternP22 + " ", "(?:" + patternP22 + ")", strings.ReplaceAll(patternP22, "\r", `\r`)} {
		if _, err := pretok.Compile(near); !errors.Is(err, pretok.ErrUnsupportedPattern) {
			t.Fatal("accepted near miss", near)
		}
	}
}
func TestP22AllScalars(t *testing.T) {
	re := compiledPattern(t, patternP22)
	var b strings.Builder
	check := func() { assertP22Reference(t, re, b.String()); b.Reset() }
	for r := rune(0); r <= 0x10ffff; r++ {
		if r >= 0xd800 && r <= 0xdfff {
			continue
		}
		b.WriteRune(r)
		b.WriteString("A ")
		if b.Len() >= 16384 {
			check()
		}
	}
	check()
}
func FuzzP22Partition(f *testing.F)       { fuzzPartition(f, patternP22) }
func BenchmarkP22Split(b *testing.B)      { benchmarkSplit(b, patternP22, false) }
func BenchmarkP22SplitFresh(b *testing.B) { benchmarkSplit(b, patternP22, true) }
