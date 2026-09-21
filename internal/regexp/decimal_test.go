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

const patternP10 = `\A\p{Nd}{1,2}(?=\p{Nd}{3}+\z)`

// RE2 checks the entire subject and captures the selected prefix independently
// of the scanner's counting algorithm. HF's {3}+ means repeated triples here.
var p10Reference = func() *regexp.Regexp {
	nd := `[` + referenceRangeBody("../scan/decimal_tables.go", "decimalRanges") + `]`
	return regexp.MustCompile(`\A(` + nd + `{1,2})(?:` + nd + `{3})+\z`)
}()

func assertP10Reference(t testing.TB, re *regex.Regexp, src string) {
	t.Helper()
	var want, split []regex.Span
	if m := p10Reference.FindStringSubmatchIndex(src); m != nil {
		want = []regex.Span{{Start: m[2], End: m[3]}}
		split = append(split, want...)
		split = append(split, regex.Span{Start: m[3], End: len(src)})
	} else if src != "" {
		split = []regex.Span{{End: len(src)}}
	}
	if got := collect(t, re, src); !slices.Equal(got, want) {
		t.Fatalf("%q: %v want %v", src, got, want)
	}
	if got, err := re.Split(src, nil); err != nil || !slices.Equal(got, split) {
		t.Fatalf("%q split: %v want %v (%v)", src, got, split, err)
	}
}

// Byte boundaries captured from tokenizers 0.22.2 before implementation.
func TestP10HFOffsets(t *testing.T) {
	re := compiledPattern(t, patternP10)
	for _, tc := range []struct {
		src string
		end int
	}{
		{"", 0}, {"1", 0}, {"12", 0}, {"123", 0}, {"1234", 1}, {"12345", 2},
		{"123456", 0}, {"1234567", 1}, {"12345678", 2}, {"123456789", 0},
		{"1234\n", 0}, {"\n1234", 0}, {"1234x", 0}, {"x1234", 0}, {"12\x0034", 0},
		{"١٢٣٤", 2}, {"１٢𝟛4٥", 5}, {"𝟙𝟚٣４567", 4},
		{"²123", 0}, {"Ⅳ123", 0}, {"123\U00011de0", 0}, {"123\U00010d40", 1},
	} {
		var want []regex.Span
		if tc.end != 0 {
			want = []regex.Span{{End: tc.end}}
		}
		if got := collect(t, re, tc.src); !slices.Equal(got, want) {
			t.Fatalf("%q: %v want %v", tc.src, got, want)
		}
		assertP10Reference(t, re, tc.src)
	}
	// Anchors refer to the supplied subject, including an independently supplied slice.
	assertP10Reference(t, re, ("x１٢𝟛4٥")[1:])
	for _, near := range []string{`\A\p{Nd}{1,2}(?=(?:\p{Nd}{3})+\z)`, `\A\p{Nd}{1,2}(?=\p{Nd}{3}\z)`, `\A\p{N}{1,2}(?=\p{N}{3}+\z)`, patternP10 + " ", "(?:" + patternP10 + ")"} {
		if r, err := regex.Compile(near); r != nil || !errors.Is(err, regex.ErrUnsupportedPattern) {
			t.Fatalf("accepted %q", near)
		}
	}
}

func TestP10DecimalRuns(t *testing.T) {
	re := compiledPattern(t, patternP10)
	for n := 0; n <= 1025; n++ {
		for _, digit := range []string{"1", "١", "１", "𝟙"} {
			src := strings.Repeat(digit, n)
			for _, text := range []string{src, src + "\n", "x" + src, src + "²"} {
				assertP10Reference(t, re, text)
			}
		}
	}
	for _, n := range []int{4095, 4096, 4097, 65535, 65536, 65537, 1 << 20} {
		src := strings.Repeat("1", n)
		assertP10Reference(t, re, src)
		assertP10Reference(t, re, src+"x")
	}
	// Covers every scalar in a context where Nd membership changes the result.
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if r >= 0xD800 && r <= 0xDFFF {
			continue
		}
		assertP10Reference(t, re, "123"+string(r))
	}
}

func TestP10Profile(t *testing.T) {
	data, err := os.ReadFile("../../testdata/patterns/P10/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var m struct {
		Pattern      string `json:"pattern"`
		PatternSHA   string `json:"pattern_sha256"`
		TablesSHA    string `json:"tables_sha256"`
		GeneratorSHA string `json:"generator_sha256"`
	}
	if err = json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(patternP10))
	if m.Pattern != patternP10 || hex.EncodeToString(sum[:]) != m.PatternSHA {
		t.Fatal("pattern differs from provenance")
	}
	for path, want := range map[string]string{"../scan/decimal_tables.go": m.TablesSHA, "../../testdata/patterns/generate_decimal_tables.py": m.GeneratorSHA} {
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		sum := sha256.Sum256(b)
		if hex.EncodeToString(sum[:]) != want {
			t.Fatalf("%s differs from pinned profile", path)
		}
	}
}

func FuzzP10Partition(f *testing.F)       { fuzzPartition(f, patternP10) }
func BenchmarkP10Split(b *testing.B)      { benchmarkSplit(b, patternP10, false) }
func BenchmarkP10SplitFresh(b *testing.B) { benchmarkSplit(b, patternP10, true) }
