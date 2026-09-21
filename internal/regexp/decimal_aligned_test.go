package regexp_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	regex "github.com/gosimd/pretok/internal/regexp"
)

const patternP25 = `\p{Nd}{1,510}(?=(?>\p{Nd}{510})*(?:\P{Nd}|$))|\G\p{Nd}{510}`

var p25Runs = regexp.MustCompile(`[` + referenceRangeBody("../scan/decimal_tables.go", "decimalRanges") + `]+`)

// Independent RE2 run selection plus right-to-left grouping of rune boundaries.
// The production scanner counts forwards and carries a cached byte endpoint.
func p25ReferencePairs(src string) [][]int {
	var pairs [][]int
	for _, run := range p25Runs.FindAllStringIndex(src, -1) {
		var bounds []int
		for i := range src[run[0]:run[1]] {
			bounds = append(bounds, run[0]+i)
		}
		bounds = append(bounds, run[1])
		var reverse [][]int
		for end := len(bounds) - 1; end > 0; {
			start := max(0, end-510)
			reverse = append(reverse, []int{bounds[start], bounds[end]})
			end = start
		}
		slices.Reverse(reverse)
		pairs = append(pairs, reverse...)
	}
	return pairs
}
func assertP25Reference(t testing.TB, re *regex.Regexp, src string) {
	t.Helper()
	assertReferencePairs(t, re, src, p25ReferencePairs(src))
}

// Group lengths and byte endpoints captured with HF 0.22.2 before implementation.
func TestP25HFOffsets(t *testing.T) {
	re := compiledPattern(t, patternP25)
	for _, tc := range []struct {
		n      int
		groups []int
	}{
		{0, nil}, {1, []int{1}}, {2, []int{2}}, {509, []int{509}}, {510, []int{510}},
		{511, []int{1, 510}}, {512, []int{2, 510}}, {1019, []int{509, 510}},
		{1020, []int{510, 510}}, {1021, []int{1, 510, 510}}, {1022, []int{2, 510, 510}},
		{1530, []int{510, 510, 510}}, {1531, []int{1, 510, 510, 510}},
	} {
		var pairs [][]int
		end := 0
		for _, n := range tc.groups {
			pairs = append(pairs, []int{end, end + n})
			end += n
		}
		assertReferencePairs(t, re, strings.Repeat("1", tc.n), pairs)
	}
	assertReferencePairs(t, re, "x123y45", [][]int{{1, 4}, {5, 7}})
	assertReferencePairs(t, re, "12²34", [][]int{{0, 2}, {4, 6}})
	assertReferencePairs(t, re, "١２𝟛4٥", [][]int{{0, 12}})
	assertReferencePairs(t, re, "123\U00011de0456", [][]int{{0, 3}, {7, 10}})
	for _, near := range []string{`\p{Nd}{1,510}`, `\p{Nd}{1,510}(?=(?:\p{Nd}{510})*(?:\P{Nd}|$))|\G\p{Nd}{510}`, strings.ReplaceAll(patternP25, "510", "511"), strings.ReplaceAll(patternP25, "Nd", "N"), patternP25 + " ", "(?:" + patternP25 + ")"} {
		if r, err := regex.Compile(near); r != nil || !errors.Is(err, regex.ErrUnsupportedPattern) {
			t.Fatalf("accepted %q", near)
		}
	}
}

func TestP25Runs(t *testing.T) {
	re := compiledPattern(t, patternP25)
	for n := 1; n <= 1022; n++ {
		for _, digits := range []string{"1", "١", "１", "𝟙"} {
			src := "x" + strings.Repeat(digits, n) + "²123"
			assertP25Reference(t, re, src)
		}
	}
	for _, n := range []int{1529, 1530, 1531, 1532, 4095, 4096, 4097, 65535, 65536, 65537, 1 << 20} {
		assertP25Reference(t, re, strings.Repeat("1", n)+"\r\n"+strings.Repeat("𝟙٢３4", 256)+"X")
	}
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if r >= 0xD800 && r <= 0xDFFF {
			continue
		}
		assertP25Reference(t, re, "12"+string(r)+"34")
	}
}

func drain(t testing.TB, c *regex.Cursor) []regex.Span {
	t.Helper()
	var spans []regex.Span
	for s, ok := c.Next(); ok; s, ok = c.Next() {
		spans = append(spans, s)
	}
	if s, ok := c.Next(); ok || s != (regex.Span{}) {
		t.Fatal("cursor not exhausted")
	}
	return spans
}
func TestP25CursorCacheOwnership(t *testing.T) {
	re := compiledPattern(t, patternP25)
	src := "x" + strings.Repeat("𝟙٢３4", 511) + "²" + strings.Repeat("١", 511)
	want := collect(t, re, src)
	for cut := 0; cut <= len(want); cut++ {
		c, _ := re.Cursor(src)
		for i := 0; i < cut; i++ {
			c.Next()
		}
		copy := c
		if got := drain(t, &c); !slices.Equal(got, want[cut:]) {
			t.Fatal("original progress differs")
		}
		if got := drain(t, &copy); !slices.Equal(got, want[cut:]) {
			t.Fatal("copy shares cached progress")
		}
	}
	// Independent cursors on a shared plan, interleaved across different run ends.
	a, _ := re.Cursor(src)
	b, _ := re.Cursor(strings.Repeat("1", 512))
	a.Next()
	b.Next()
	if got := drain(t, &b); !slices.Equal(got, []regex.Span{{Start: 2, End: 512}}) {
		t.Fatal("run cache leaked between subjects")
	}
	if got := drain(t, &a); !slices.Equal(got, want[1:]) {
		t.Fatal("second cursor changed first")
	}
	for _, text := range []string{src, "123x456", strings.Repeat("1", 1022)} {
		testConcurrentReuse(t, re, text)
	}
	wantSplit, _ := re.Split(src, nil)
	dst := make([]regex.Span, 0, len(wantSplit)+1)
	dst = append(dst, regex.Span{Start: 91, End: 92})
	got, err := re.Split(src, dst)
	if err != nil || &got[0] != &dst[0] || !slices.Equal(got[1:], wantSplit) {
		t.Fatal("cached split append/reuse failed")
	}
	if allocs := testing.AllocsPerRun(100, func() { re.Split(src, got[:0]) }); allocs != 0 {
		t.Fatalf("reused Split allocates: %v", allocs)
	}
	if allocs := testing.AllocsPerRun(100, func() {
		c, _ := re.Cursor(src)
		for _, ok := c.Next(); ok; _, ok = c.Next() {
		}
	}); allocs != 0 {
		t.Fatalf("Cursor allocates: %v", allocs)
	}
}

func TestP25Profile(t *testing.T) {
	data, err := os.ReadFile("../../testdata/patterns/P25/manifest.json")
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
	sum := sha256.Sum256([]byte(patternP25))
	if m.Pattern != patternP25 || hex.EncodeToString(sum[:]) != m.PatternSHA {
		t.Fatal("pattern differs from provenance")
	}
	if m.Profile != "../P10/manifest.json" {
		t.Fatal("unexpected shared profile")
	}
	data, err = os.ReadFile("../../testdata/patterns/P25/" + m.Profile)
	if err != nil {
		t.Fatal(err)
	}
	sum = sha256.Sum256(data)
	if hex.EncodeToString(sum[:]) != m.ProfileSHA {
		t.Fatal("shared Unicode profile changed")
	}
}
func FuzzP25Partition(f *testing.F) {
	for _, n := range []int{509, 510, 511, 512, 1019, 1020, 1021, 1022, 1531} {
		f.Add("x" + strings.Repeat("1", n) + "²" + strings.Repeat("١", 511))
	}
	fuzzPartition(f, patternP25)
}
func BenchmarkP25Split(b *testing.B)      { benchmarkSplit(b, patternP25, false) }
func BenchmarkP25SplitFresh(b *testing.B) { benchmarkSplit(b, patternP25, true) }
