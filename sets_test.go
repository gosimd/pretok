package pretok_test

import (
	"errors"
	"github.com/gosimd/pretok"
	"regexp"
	"slices"
	"strings"
	"testing"
)

// Exact independent literals; these are not copied from the production registry.
const patternP11 = "[\r\n]"
const patternP21 = `[一-龥぀-ゟ゠-ヿ]+`
const patternP17 = `[一-龥ࠀ-一가-퟿]+`
const patternP18 = `\s?[!-/:-~！-／：-～‘-‟　-。]+`
const patternP08 = `\s?[A-Za-zµÀ-ÖØ-öø-ƺƼ-ƿǄ-ʓʕ-ʯͰ-ͳͶͷͻ-ͽͿΆΈ-ΊΌΎ-ΡΣ-ϵϷ-ҁҊ-ԯԱ-ՖႠ-ჅᎠ-Ᏽᏸ-ᏽᲐ-ᲺᲽ-Ჿᴀ-ᴫᵫ-ᵷᵹ-ᶚḀ-ἕἘ-Ἕἠ-ὅὈ-Ὅὐ-ὗὙὛὝὟ-ώᾀ-ᾴᾶ-ᾼιῂ-ῄῆ-ῌῐ-ΐῖ-Ίῠ-Ῥῲ-ῴῶ-ῼℂℇℊ-ℓℕℙ-ℝℤΩℨK-ℭℯ-ℴℹℼ-ℿⅅ-ⅉⅎↃↄⰀ-ⱻⱾ-ⳤⳫ-ⳮⳲⳳꙀ-ꙭꚀ-ꚛꜢ-ꝯꝱ-ꞇꞋ-ꞎꭰ-ꮿﬀ-ﬆﬓ-ﬗＡ-Ｚａ-ｚ𐐀-𐑏𐒰-𐓓𐓘-𐓻𐲀-𐲲𐳀-𐳲𑢠-𑣟𞤀-𞥃]+`
const patternP23 = ` ?[^(\s|[.,!?…。，、।۔،])]+`

// HF 0.22.2 whitespace profile expanded for Go RE2 (whose \s is ASCII).
const referenceSpace = `\x{9}-\x{D}\x{20}\x{85}\x{A0}\x{1680}\x{2000}-\x{200A}\x{2028}\x{2029}\x{202F}\x{205F}\x{3000}`

var setPatterns = map[string]string{"P11": patternP11, "P21": patternP21, "P17": patternP17, "P18": patternP18, "P08": patternP08, "P23": patternP23}
var setReferences = func() map[string]*regexp.Regexp {
	out := make(map[string]*regexp.Regexp)
	for _, p := range setPatterns {
		translated := strings.ReplaceAll(p, `\s`, `[`+referenceSpace+`]`)
		if p == patternP23 {
			translated = ` ?[^` + referenceSpace + `().,!?.|…。，、।۔،]+`
		}
		out[p] = regexp.MustCompile(translated)
	}
	return out
}()

func TestExactSets(t *testing.T) {
	for id, p := range setPatterns {
		t.Run(id, func(t *testing.T) {
			re := compiledPattern(t, p)
			for _, src := range []string{"", "\r\n", "\u3000", "\u3000\u3000", "\u2003\u2003X", "[abc] (abc) a|b .,!?:;_-", " abc  def", "A\u0301中𐐀"} {
				assertReference(t, re, src, setReferences[p])
			}
			for _, near := range []string{p + " ", "(?:" + p + ")", strings.TrimSuffix(p, "+") + "{1,}"} {
				if near == p {
					continue
				}
				if _, err := pretok.Compile(near); !errors.Is(err, pretok.ErrUnsupportedPattern) {
					t.Fatalf("near miss accepted: %q", near)
				}
			}
		})
	}
}

// Captured from pinned HF, converted from scalar offsets to UTF-8 byte offsets.
func TestExactSetsHFExamples(t *testing.T) {
	for _, tc := range []struct {
		p, src string
		want   []pretok.Span
	}{
		{patternP11, "\r\n", []pretok.Span{{0, 1}, {1, 2}}},
		{patternP21, "A\u0301中𐐀", []pretok.Span{{3, 6}}},
		{patternP17, "\u2003\u2003X", []pretok.Span{{0, 6}}},
		{patternP18, "\u3000", []pretok.Span{{0, 3}}},
		{patternP18, "\u3000\u3000", []pretok.Span{{0, 6}}},
		{patternP18, "\u2003\u2003X", []pretok.Span{{3, 7}}},
		{patternP08, "A\u0301中𐐀", []pretok.Span{{0, 1}, {6, 10}}},
		{patternP23, "[abc] (abc) a|b .,!?:;_", []pretok.Span{{0, 5}, {7, 10}, {11, 13}, {14, 15}, {20, 23}}},
	} {
		got := collect(t, compiledPattern(t, tc.p), tc.src)
		if !slices.Equal(got, tc.want) {
			t.Fatalf("%q: %v want %v", tc.src, got, tc.want)
		}
	}
}

// Check every valid scalar against an independent regex class, in chunks with
// separators. This covers all holes, endpoints and supplementary-plane ranges.
func TestExactSetsAllScalars(t *testing.T) {
	var b strings.Builder
	check := func() {
		src := b.String()
		for _, p := range setPatterns {
			assertReference(t, compiledPattern(t, p), src, setReferences[p])
		}
		b.Reset()
	}
	for r := rune(0); r <= 0x10FFFF; r++ {
		if r >= 0xD800 && r <= 0xDFFF {
			continue
		}
		b.WriteRune(r)
		b.WriteByte(' ')
		if b.Len() >= 16384 {
			check()
		}
	}
	check()
}
func FuzzP11Partition(f *testing.F)       { fuzzPartition(f, patternP11) }
func BenchmarkP11Split(b *testing.B)      { benchmarkSplit(b, patternP11, false) }
func BenchmarkP11SplitFresh(b *testing.B) { benchmarkSplit(b, patternP11, true) }
func FuzzP21Partition(f *testing.F)       { fuzzPartition(f, patternP21) }
func BenchmarkP21Split(b *testing.B)      { benchmarkSplit(b, patternP21, false) }
func BenchmarkP21SplitFresh(b *testing.B) { benchmarkSplit(b, patternP21, true) }
func FuzzP17Partition(f *testing.F)       { fuzzPartition(f, patternP17) }
func BenchmarkP17Split(b *testing.B)      { benchmarkSplit(b, patternP17, false) }
func BenchmarkP17SplitFresh(b *testing.B) { benchmarkSplit(b, patternP17, true) }
func FuzzP18Partition(f *testing.F)       { fuzzPartition(f, patternP18) }
func BenchmarkP18Split(b *testing.B)      { benchmarkSplit(b, patternP18, false) }
func BenchmarkP18SplitFresh(b *testing.B) { benchmarkSplit(b, patternP18, true) }
func FuzzP08Partition(f *testing.F)       { fuzzPartition(f, patternP08) }
func BenchmarkP08Split(b *testing.B)      { benchmarkSplit(b, patternP08, false) }
func BenchmarkP08SplitFresh(b *testing.B) { benchmarkSplit(b, patternP08, true) }
func FuzzP23Partition(f *testing.F)       { fuzzPartition(f, patternP23) }
func BenchmarkP23Split(b *testing.B)      { benchmarkSplit(b, patternP23, false) }
func BenchmarkP23SplitFresh(b *testing.B) { benchmarkSplit(b, patternP23, true) }
