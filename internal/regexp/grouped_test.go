package regexp_test

import (
	"regexp"
	"slices"
	"strings"
	"testing"

	regex "github.com/gosimd/pretok/internal/regexp"
)

const patternP12 = `(?:[\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF\u3400-\u4DBF\uf900-\uFAFF\uFF65-\uFF9F\u2F00-\u2FDF]+|[\u0E00-\u0E7F]+|[\u0E80-\u0EFF]+|[\u1780-\u17FF]+|[\u1000-\u109F\uAA60-\uAA7F\uA9E0-\uA9FF]+|[\uAC00-\uD7AF\u1100-\u11FF]+)`
const patternP15 = `\p{Hangul}+|[！…“”‘’—：；，、-〿︰-﹏]+|[ㄅ-ㄯ]+|[一-龥぀-ゟ゠-ヿ]+`

// Independent expansion of the HF 0.22.2 Hangul Script property. The explicit
// extraction in benchmark/prepare/grouped.py verifies production table data.
const referenceHangul = `[\x{1100}-\x{11FF}\x{302E}-\x{302F}\x{3131}-\x{318E}\x{3200}-\x{321E}\x{3260}-\x{327E}\x{A960}-\x{A97C}\x{AC00}-\x{D7A3}\x{D7B0}-\x{D7C6}\x{D7CB}-\x{D7FB}\x{FFA0}-\x{FFBE}\x{FFC2}-\x{FFC7}\x{FFCA}-\x{FFCF}\x{FFD2}-\x{FFD7}\x{FFDA}-\x{FFDC}]`

func groupedReferencePattern(p string) string {
	p = regexp.MustCompile(`\\u([0-9a-fA-F]{4})`).ReplaceAllString(p, `\x{$1}`)
	return strings.ReplaceAll(p, `\p{Hangul}`, referenceHangul)
}

// Pinned HF examples include order reversal at the overlapping tone marks.
func TestGroupedHFExamples(t *testing.T) {
	for _, tc := range []struct {
		p, src string
		want   []regex.Span
	}{
		{patternP12, "中ก", []regex.Span{{0, 3}, {3, 6}}},
		{patternP12, "中あカ가ᄀกກកက", []regex.Span{{0, 9}, {9, 15}, {15, 18}, {18, 21}, {21, 24}, {24, 27}}},
		{patternP12, "ꥠힰ", nil},
		{patternP15, "中ก", []regex.Span{{0, 3}}},
		{patternP15, "〮、", []regex.Span{{0, 3}, {3, 6}}},
		{patternP15, "、〮", []regex.Span{{0, 6}}},
		{patternP15, "가〮、", []regex.Span{{0, 6}, {6, 9}}},
		{patternP15, "、〮가", []regex.Span{{0, 6}, {6, 9}}},
		{patternP15, "〯〮、", []regex.Span{{0, 6}, {6, 9}}},
		{patternP15, "〮〯가", []regex.Span{{0, 9}}},
		{patternP15, "ㄅ中", []regex.Span{{0, 3}, {3, 6}}},
		{patternP15, "A〮、X", []regex.Span{{1, 4}, {4, 7}}},
		{patternP15, "ﾠﾡ", []regex.Span{{0, 6}}},
		{patternP15, "ꥠힰ", []regex.Span{{0, 6}}},
	} {
		re := compiledPattern(t, tc.p)
		if got := collect(t, re, tc.src); !slices.Equal(got, tc.want) {
			t.Fatalf("%q: %v want %v", tc.src, got, tc.want)
		}
		assertReference(t, re, tc.src, setReferences[tc.p])
	}
}
func TestGroupedTransitions(t *testing.T) {
	alphabet := []rune{'中', 'あ', '가', 'ᄀ', 'ก', 'ກ', 'ក', 'က', '〮', '〯', '、', 'ㄅ', 'A', ' '}
	for _, p := range []string{patternP12, patternP15} {
		re := compiledPattern(t, p)
		var visit func(string, int)
		visit = func(src string, left int) {
			assertReference(t, re, src, setReferences[p])
			if left > 0 {
				for _, r := range alphabet {
					visit(src+string(r), left-1)
				}
			}
		}
		visit("", 4)
		for _, n := range []int{1, 2, 3, 15, 16, 17, 31, 32, 33, 63, 64, 65, 4096} {
			for _, body := range []string{"中あ", "ก", "가〮〯", "、〮〯", "ㄅ"} {
				assertReference(t, re, "A"+strings.Repeat(body, n)+"、가X", setReferences[p])
			}
		}
	}
}
func FuzzP12Partition(f *testing.F)       { fuzzPartition(f, patternP12) }
func FuzzP15Partition(f *testing.F)       { fuzzPartition(f, patternP15) }
func BenchmarkP12Split(b *testing.B)      { benchmarkSplit(b, patternP12, false) }
func BenchmarkP12SplitFresh(b *testing.B) { benchmarkSplit(b, patternP12, true) }
func BenchmarkP15Split(b *testing.B)      { benchmarkSplit(b, patternP15, false) }
func BenchmarkP15SplitFresh(b *testing.B) { benchmarkSplit(b, patternP15, true) }
