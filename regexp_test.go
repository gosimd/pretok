package pretok_test

import (
	"bufio"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

const patternTriples = `[0-9][0-9][0-9]`

const patternMarks = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?[\p{L}\p{M}]+|\p{N}| ?[^\s\p{L}\p{M}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

const pattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

func compiled(t testing.TB) *pretok.Regexp { return compiledPattern(t, pattern) }

func compiledPattern(t testing.TB, pattern string) *pretok.Regexp {
	t.Helper()
	re, err := pretok.Compile(pattern)
	if err != nil {
		t.Fatal(err)
	}
	return re
}

func collect(t testing.TB, re *pretok.Regexp, src string) []pretok.Span {
	t.Helper()
	c, err := re.Cursor(src)
	if err != nil {
		t.Fatal(err)
	}
	var spans []pretok.Span
	for span, ok := c.Next(); ok; span, ok = c.Next() {
		spans = append(spans, span)
	}
	if span, ok := c.Next(); ok || span != (pretok.Span{}) {
		t.Fatal("cursor not exhausted")
	}
	return spans
}

func TestOracle(t *testing.T) {
	testOracle(t, "testdata/qwen", pattern, "internal/scan/unicode_tables.go")
}

func TestP01Oracle(t *testing.T) {
	testOracle(t, "testdata/patterns/P01", patternMarks, "internal/scan/word_tables.go")
}

func testOracle(t *testing.T, directory, pattern, table string) {
	re, err := pretok.Compile(pattern)
	if err != nil {
		t.Fatal(err)
	}
	manifestData, err := os.ReadFile(directory + "/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		Pattern   string `json:"pattern"`
		Count     int    `json:"case_count"`
		CasesSHA  string `json:"cases_sha256"`
		TablesSHA string `json:"tables_sha256"`
	}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		t.Fatal(err)
	}
	if manifest.Pattern != pattern {
		t.Fatal("fixture pattern differs from compiled expression")
	}
	for path, want := range map[string]string{directory + "/cases.jsonl.gz": manifest.CasesSHA, table: manifest.TablesSHA} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		sum := sha256.Sum256(data)
		if hex.EncodeToString(sum[:]) != want {
			t.Fatalf("%s differs from oracle manifest", path)
		}
	}
	f, err := os.Open(directory + "/cases.jsonl.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	decoded, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer decoded.Close()
	lines := bufio.NewScanner(decoded)
	lines.Buffer(make([]byte, 4096), 4<<20)
	count := 0
	var dst []pretok.Span
	for lines.Scan() {
		count++
		var tc struct {
			Input          string
			Matches, Split []pretok.Span
		}
		if err := json.Unmarshal(lines.Bytes(), &tc); err != nil {
			t.Fatal(err)
		}
		if got := collect(t, re, tc.Input); !slices.Equal(got, tc.Matches) {
			t.Fatalf("case %d input %q: matches %v want %v", count, tc.Input, got, tc.Matches)
		}
		dst, err = re.Split(tc.Input, dst[:0])
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(dst, tc.Split) {
			t.Fatalf("case %d input %q: split %v want %v", count, tc.Input, dst, tc.Split)
		}
		assertPartition(t, tc.Input, dst)
	}
	if err := lines.Err(); err != nil {
		t.Fatal(err)
	}
	if count != manifest.Count {
		t.Fatalf("read %d cases, want %d", count, manifest.Count)
	}
}

func assertPartition(t testing.TB, src string, spans []pretok.Span) {
	t.Helper()
	end := 0
	for _, s := range spans {
		if s.Start != end || s.End <= s.Start || s.End > len(src) {
			t.Fatalf("invalid partition: %v", spans)
		}
		if !utf8.ValidString(src[s.Start:s.End]) {
			t.Fatalf("span cuts UTF-8: %v", s)
		}
		end = s.End
	}
	if end != len(src) {
		t.Fatalf("partition ends at %d, want %d", end, len(src))
	}
}

func TestSupportedCorpusStrings(t *testing.T) {
	data, err := os.ReadFile("testdata/corpus_patterns.json")
	if err != nil {
		t.Fatal(err)
	}
	var rows []struct{ Pattern string }
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, row := range rows {
		re, err := pretok.Compile(row.Pattern)
		if row.Pattern == pattern || row.Pattern == patternMarks || row.Pattern == patternTriples {
			if err != nil || re == nil {
				t.Fatal("registered string rejected")
			}
			count++
		} else if !errors.Is(err, pretok.ErrUnsupportedPattern) || re != nil {
			t.Fatalf("unexpectedly accepted %q", row.Pattern)
		}
	}
	if count != 3 {
		t.Fatalf("registered corpus entries: %d", count)
	}
	for _, p := range []string{"", "[", ".*", `[0-9]{3}`, `[0-9][0-9]`, `[0-9][0-9][0-9] `, `\d{3}`, "(?:" + pattern + ")", pattern + " ", "(?:" + patternMarks + ")", patternMarks + " ", strings.ReplaceAll(pattern, `\r`, "\r")} {
		if re, err := pretok.Compile(p); re != nil || !errors.Is(err, pretok.ErrUnsupportedPattern) {
			t.Fatalf("accepted unknown %q", p)
		}
	}
}

func TestOwnershipAndErrors(t *testing.T) {
	for id, p := range map[string]string{"P01": patternMarks, "P02": patternTriples, "P05": pattern} {
		t.Run(id, func(t *testing.T) { testOwnershipAndErrors(t, compiledPattern(t, p), p) })
	}
}
func testOwnershipAndErrors(t *testing.T, re *pretok.Regexp, pattern string) {
	storage := make([]pretok.Span, 32)
	for i := range storage {
		storage[i] = pretok.Span{Start: 91, End: 92}
	}
	before := slices.Clone(storage)
	for _, src := range []string{"\xff", "valid\xf0\x9f", "\xed\xa0\x80", "\xc0\xaf"} {
		got, err := re.Split(src, storage[:1])
		if !errors.Is(err, pretok.ErrInvalidUTF8) || len(got) != 1 || &got[0] != &storage[0] || !slices.Equal(before, storage) {
			t.Fatal("invalid UTF-8 mutated destination")
		}
		c, err := re.Cursor(src)
		if !errors.Is(err, pretok.ErrInvalidUTF8) {
			t.Fatal(err)
		}
		if _, ok := c.Next(); ok {
			t.Fatal("active cursor after error")
		}
	}
	for _, r := range []*pretok.Regexp{nil, new(pretok.Regexp)} {
		if _, err := r.Split("a", storage[:1]); !errors.Is(err, pretok.ErrUninitialized) {
			t.Fatal(err)
		}
		if _, err := r.Cursor("a"); !errors.Is(err, pretok.ErrUninitialized) {
			t.Fatal(err)
		}
	}
	got, err := re.Split("a  b", storage[:1])
	if err != nil || &got[0] != &storage[0] || got[0] != before[0] {
		t.Fatal("append/reuse contract")
	}
	if want, err := re.Split("a  b", nil); err != nil || !slices.Equal(got[1:], want) {
		t.Fatal("appended spans differ")
	}
	empty, err := re.Split("", got)
	if err != nil || len(empty) != len(got) || &empty[0] != &got[0] {
		t.Fatal("empty input appended spans")
	}
	var zero pretok.Cursor
	if span, ok := zero.Next(); ok || span != (pretok.Span{}) {
		t.Fatal("zero cursor not exhausted")
	}
	c, _ := re.Cursor("123456")
	c.Next()
	copy := c
	if a, ok := c.Next(); !ok {
		t.Fatal("missing match")
	} else if b, ok := copy.Next(); !ok || a != b {
		t.Fatal("cursor copy shares progress")
	}
	piece := strings.TrimPrefix("prefix\u2003\u2003X", "prefix")
	want := []pretok.Span{{0, 3}, {3, 7}}
	if pattern == patternTriples {
		want = nil
	}
	if !slices.Equal(collect(t, re, piece), want) {
		t.Fatal("offsets not relative to subject")
	}

}

func TestConcurrentReuse(t *testing.T) {
	for id, p := range map[string]string{"P01": patternMarks, "P02": patternTriples, "P05": pattern} {
		t.Run(id, func(t *testing.T) { testConcurrentReuse(t, compiledPattern(t, p)) })
	}
}
func testConcurrentReuse(t *testing.T, re *pretok.Regexp) {
	text := "We're 1234567 １２3\r\n\u2003\u2003世界!"
	want, err := re.Split(text, nil)
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for range 16 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var dst []pretok.Span
			for range 50 {
				var err error
				dst, err = re.Split(text, dst[:0])
				if err != nil || !slices.Equal(dst, want) {
					t.Error("concurrent split differs")
					return
				}
			}
		}()
	}
	wg.Wait()
}

func FuzzP02Partition(f *testing.F)  { fuzzPartition(f, patternTriples) }
func FuzzP01Partition(f *testing.F)  { fuzzPartition(f, patternMarks) }
func FuzzQwenPartition(f *testing.F) { fuzzPartition(f, pattern) }
func fuzzPartition(f *testing.F, pattern string) {
	for _, s := range []string{"", "we're", "x123456y12", "１２３123١٢٣", "\u2003\u2003X", "a\u0301b", "²Ⅳ１２3", "\r\n\tX", "\xff"} {
		f.Add(s)
	}
	re := compiledPattern(f, pattern)
	f.Fuzz(func(t *testing.T, src string) {
		got, err := re.Split(src, nil)
		if !utf8.ValidString(src) {
			if !errors.Is(err, pretok.ErrInvalidUTF8) || got != nil {
				t.Fatal("invalid UTF-8 accepted")
			}
			return
		}
		if err != nil {
			t.Fatal(err)
		}
		assertPartition(t, src, got)
		// P01/P05 cover every scalar. P02 also needs independent gap checks.
		if pattern == patternTriples {
			assertP02Reference(t, re, src)
		} else if !slices.Equal(got, collect(t, re, src)) {
			t.Fatal("cursor/split mismatch")
		}
	})
}

func TestQwenExamples(t *testing.T) {
	re := compiled(t)
	for _, tc := range []struct {
		input  string
		pieces []string
	}{
		{"we're I'LL 'ſ", []string{"we", "'re", " I", "'LL", " '", "ſ"}},
		{"'ſ", []string{"'ſ"}},
		{"fooBar HTTPServer", []string{"fooBar", " HTTPServer"}},
		{"a\u0301b", []string{"a", "\u0301b"}},
		{"１２²Ⅳ", []string{"１", "２", "²", "Ⅳ"}},
		{"\u2003\u2003X", []string{"\u2003", "\u2003X"}},
		{" \r\n\tX", []string{" \r\n", "\tX"}},
		{"!\r\n", []string{"!\r\n"}},
		{"!!'s", []string{"!!'", "s"}},
	} {
		spans, err := re.Split(tc.input, nil)
		if err != nil {
			t.Fatal(err)
		}
		var pieces []string
		for _, span := range spans {
			pieces = append(pieces, tc.input[span.Start:span.End])
		}
		if !slices.Equal(pieces, tc.pieces) {
			t.Errorf("%q: got %q want %q", tc.input, pieces, tc.pieces)
		}
	}
}

func BenchmarkCompile(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if re, err := pretok.Compile(pattern); err != nil || re == nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkQwenSplit(b *testing.B)      { benchmarkSplit(b, pattern, false) }
func BenchmarkQwenSplitFresh(b *testing.B) { benchmarkSplit(b, pattern, true) }
func BenchmarkP01Split(b *testing.B)       { benchmarkSplit(b, patternMarks, false) }
func BenchmarkP02Split(b *testing.B)       { benchmarkSplit(b, patternTriples, false) }
func BenchmarkP02SplitFresh(b *testing.B)  { benchmarkSplit(b, patternTriples, true) }
func BenchmarkP01SplitFresh(b *testing.B)  { benchmarkSplit(b, patternMarks, true) }

func benchmarkSplit(b *testing.B, pattern string, fresh bool) {
	re := compiledPattern(b, pattern)
	for name, src := range map[string]string{"short": "Hello, we're here! 123", "multilingual": strings.Repeat("世界 Привет! we're １２3\r\n\u2003\u2003X ", 64)} {
		b.Run(name, func(b *testing.B) {
			dst, err := re.Split(src, nil)
			if err != nil {
				b.Fatal(err)
			}
			b.SetBytes(int64(len(src)))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				if fresh {
					dst = nil
				}
				dst, err = re.Split(src, dst[:0])
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestP01Marks(t *testing.T) {
	re := compiledPattern(t, patternMarks)
	for _, tc := range []struct {
		text   string
		pieces []string
	}{
		{"a\u0301b", []string{"a\u0301b"}},
		{"\u0301", []string{"\u0301"}},
		{"\u0301\u0301", []string{"\u0301\u0301"}},
		{"!\u0301a", []string{"!\u0301a"}},
		{"!!\u0301a", []string{"!!", "\u0301a"}},
		{" \u0301", []string{" \u0301"}},
		{"\u0301!", []string{"\u0301", "!"}},
		{"'s\u0301", []string{"'s", "\u0301"}},
		{"1\u0301", []string{"1", "\u0301"}},
	} {
		spans, err := re.Split(tc.text, nil)
		if err != nil {
			t.Fatal(err)
		}
		var pieces []string
		for _, s := range spans {
			pieces = append(pieces, tc.text[s.Start:s.End])
		}
		if !slices.Equal(pieces, tc.pieces) {
			t.Errorf("%q: got %q want %q", tc.text, pieces, tc.pieces)
		}
	}
}

// RE2 supports this ASCII-only expression; HF remains the external compatibility
// oracle in the massive-testing pipeline. This comparator keeps fuzzing offline.
var p02Reference = regexp.MustCompile(`[0-9]{3}`)

func assertP02Reference(t testing.TB, re *pretok.Regexp, src string) {
	t.Helper()
	var matches, split []pretok.Span
	end := 0
	for _, pair := range p02Reference.FindAllStringIndex(src, -1) {
		if pair[0] > end {
			split = append(split, pretok.Span{Start: end, End: pair[0]})
		}
		span := pretok.Span{Start: pair[0], End: pair[1]}
		matches = append(matches, span)
		split = append(split, span)
		end = pair[1]
	}
	if end < len(src) {
		split = append(split, pretok.Span{Start: end, End: len(src)})
	}
	if got := collect(t, re, src); !slices.Equal(got, matches) {
		t.Fatalf("%q: matches %v want %v", src, got, matches)
	}
	if got, err := re.Split(src, nil); err != nil || !slices.Equal(got, split) {
		t.Fatalf("%q: split %v want %v: %v", src, got, split, err)
	}

}

func TestP02Gaps(t *testing.T) {
	re := compiledPattern(t, patternTriples)
	// Pieces observed with HF Tokenizers 0.22.2, Isolated / invert=false.
	for _, tc := range []struct {
		text   string
		pieces []string
	}{
		{"", nil},
		{"12a34", []string{"12a34"}},
		{"1234567", []string{"123", "456", "7"}},
		{"x123y456z12", []string{"x", "123", "y", "456", "z12"}},
		{"１２３123١٢٣", []string{"１２３", "123", "١٢٣"}},
		{"123\r\n45", []string{"123", "\r\n45"}},
		{"a123456b", []string{"a", "123", "456", "b"}},
	} {
		spans, err := re.Split(tc.text, nil)
		if err != nil {
			t.Fatal(err)
		}
		var pieces []string
		for _, s := range spans {
			pieces = append(pieces, tc.text[s.Start:s.End])
		}
		if !slices.Equal(pieces, tc.pieces) {
			t.Errorf("%q: pieces %q want %q", tc.text, pieces, tc.pieces)
		}
		assertPartition(t, tc.text, spans)
		assertP02Reference(t, re, tc.text)
	}
	// Every residue and UTF-8 transition around common vector/tail boundaries.
	for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 15, 16, 17, 31, 32, 33, 63, 64, 65, 4096} {
		for _, sep := range []string{"", "a", "１", "١", "²", "😀", "\r\n", "\u0301", "/:"} {
			src := sep + strings.Repeat("1", n) + sep + "12" + sep + "345" + sep
			assertP02Reference(t, re, src)
			spans, err := re.Split(src, nil)
			if err != nil {
				t.Fatal(err)
			}
			assertPartition(t, src, spans)
		}
	}
}
