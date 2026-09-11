package pretok_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

const patternP14 = `\A\p{Nd}{3}|\G\p{Nd}{3}`

var p14Reference = regexp.MustCompile(`[` + referenceRangeBody("internal/scan/decimal_tables.go", "decimalRanges") + `]{3}`)

// RE2 finds triples independently of the scanner. Only the uninterrupted prefix
// belongs to P14; a later match beyond any gap must not restart the sequence.
func assertP14Reference(t testing.TB, re *pretok.Regexp, src string) {
	t.Helper()
	pairs := p14Reference.FindAllStringIndex(src, -1)
	end, n := 0, 0
	for _, pair := range pairs {
		if pair[0] != end {
			break
		}
		end = pair[1]
		n++
	}
	assertReferencePairs(t, re, src, pairs[:n])
}

// Byte offsets captured from the pinned HF 0.22.2 oracle before implementation.
func TestP14HFOffsets(t *testing.T) {
	re := compiledPattern(t, patternP14)
	for _, tc := range []struct {
		src   string
		pairs [][]int
	}{
		{"", nil}, {"1", nil}, {"12", nil}, {"123", [][]int{{0, 3}}},
		{"1234", [][]int{{0, 3}}}, {"123456", [][]int{{0, 3}, {3, 6}}},
		{"1234567", [][]int{{0, 3}, {3, 6}}}, {"x123456", nil},
		{"123x456", [][]int{{0, 3}}}, {"1234x567", [][]int{{0, 3}}},
		{"123\n456", [][]int{{0, 3}}}, {"１２٣٤𝟝67", [][]int{{0, 8}, {8, 15}}},
		{"𝟙٢３٤𝟝6", [][]int{{0, 9}, {9, 16}}}, {"123²456", [][]int{{0, 3}}},
		{"123\U00011de0456", [][]int{{0, 3}}}, {"123\U00010d40456", [][]int{{0, 3}, {3, 9}}},
	} {
		assertReferencePairs(t, re, tc.src, tc.pairs)
		assertP14Reference(t, re, tc.src)
	}
	// A separately supplied slice establishes a new subject with its own byte zero.
	assertReferencePairs(t, re, ("x１２٣٤𝟝67")[1:], [][]int{{0, 8}, {8, 15}})
	for _, near := range []string{`\G\p{Nd}{3}`, `\A\p{Nd}{3}`, `\p{Nd}{3}`, `\A\p{N}{3}|\G\p{N}{3}`, patternP14 + " ", "(?:" + patternP14 + ")"} {
		if r, err := pretok.Compile(near); r != nil || !errors.Is(err, pretok.ErrUnsupportedPattern) {
			t.Fatalf("accepted %q", near)
		}
	}
	// Copy at a partial-triple failure: both copies stay exhausted past later digits.
	c, _ := re.Cursor("1234x567")
	c.Next()
	copy := c
	for range 3 {
		for _, cursor := range []*pretok.Cursor{&c, &copy} {
			if span, ok := cursor.Next(); ok || span != (pretok.Span{}) {
				t.Fatal("search resumed after gap")
			}
		}
	}
}

func TestP14DecimalRuns(t *testing.T) {
	re := compiledPattern(t, patternP14)
	for n := 0; n <= 129; n++ {
		for _, digit := range []string{"1", "١", "１", "𝟙"} {
			src := strings.Repeat(digit, n)
			for _, text := range []string{src, src + "x123456", src + "\n١٢٣", "x" + src, "1x" + src, "12x" + src} {
				assertP14Reference(t, re, text)
			}
		}
	}
	for _, n := range []int{509, 510, 511, 512, 1020, 1021, 1022, 4095, 4096, 4097, 65535, 65536, 65537} {
		for _, digit := range []string{"1", "١", "𝟙"} {
			assertP14Reference(t, re, strings.Repeat(digit, n)+"x123")
		}
	}
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if r >= 0xD800 && r <= 0xDFFF {
			continue
		}
		assertP14Reference(t, re, "12"+string(r)+"456")
	}
}

func TestP14Profile(t *testing.T) {
	data, err := os.ReadFile("testdata/patterns/P14/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var m struct {
		Pattern    string `json:"pattern"`
		PatternSHA string `json:"pattern_sha256"`
		Profile    string `json:"shared_unicode_profile"`
		ProfileSHA string `json:"shared_unicode_profile_sha256"`
	}
	if err = json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(patternP14))
	if m.Pattern != patternP14 || hex.EncodeToString(sum[:]) != m.PatternSHA {
		t.Fatal("pattern differs from provenance")
	}
	if m.Profile != "../P10/manifest.json" {
		t.Fatal("unexpected shared Unicode profile")
	}
	data, err = os.ReadFile("testdata/patterns/P14/" + m.Profile)
	if err != nil {
		t.Fatal(err)
	}
	sum = sha256.Sum256(data)
	if hex.EncodeToString(sum[:]) != m.ProfileSHA {
		t.Fatal("shared Unicode profile changed without P14 validation")
	}
}

func FuzzP14Partition(f *testing.F)       { fuzzPartition(f, patternP14) }
func BenchmarkP14Split(b *testing.B)      { benchmarkSplit(b, patternP14, false) }
func BenchmarkP14SplitFresh(b *testing.B) { benchmarkSplit(b, patternP14, true) }
