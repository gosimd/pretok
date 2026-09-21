package pretok_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"regexp/syntax"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

func fallbackConfig(pattern, behavior string, invert bool) string {
	data, err := json.Marshal(map[string]any{
		"type": "Split", "pattern": map[string]string{"Regex": pattern},
		"behavior": behavior, "invert": invert,
	})
	if err != nil {
		panic(err)
	}
	return string(data)
}

func checkFallbackSpans(t testing.TB, p *pretok.Pipeline, input string, want []pretok.Span) {
	t.Helper()
	var got []pretok.Span
	err := p.Walk(input, func(piece pretok.Piece) error {
		span, err := piece.OriginalRange(0, len(piece.Text()))
		if err != nil {
			return err
		}
		if span.Start < 0 || span.End > len(input) || span.Start >= span.End || piece.Text() != input[span.Start:span.End] {
			t.Errorf("invalid source mapping: %q %v in %q", piece.Text(), span, input)
		}
		got = append(got, span)
		return nil
	})
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("Walk(%q): %v, %v; want %v", input, got, err, want)
	}
	for _, walk := range []func(string, func(string) error) error{p.WalkText, p.WalkTextUnsafe} {
		i := 0
		err = walk(input, func(text string) error {
			if i >= len(want) || text != input[want[i].Start:want[i].End] {
				t.Fatalf("text %d = %q, want spans %v", i, text, want)
			}
			i++
			return nil
		})
		if err != nil || i != len(want) {
			t.Fatalf("text walk: count=%d, %v", i, err)
		}
	}
}

func TestPipelineRegexpFallbackModes(t *testing.T) {
	for _, tc := range []struct {
		mode   string
		invert bool
		want   []pretok.Span
	}{
		{"Removed", false, []pretok.Span{{0, 1}, {3, 4}, {6, 7}}},
		{"Removed", true, []pretok.Span{{1, 3}, {4, 6}}},
		{"Isolated", false, []pretok.Span{{0, 1}, {1, 3}, {3, 4}, {4, 6}, {6, 7}}},
		{"Isolated", true, []pretok.Span{{0, 1}, {1, 3}, {3, 4}, {4, 6}, {6, 7}}},
		{"Contiguous", false, []pretok.Span{{0, 1}, {1, 3}, {3, 4}, {4, 6}, {6, 7}}},
		{"Contiguous", true, []pretok.Span{{0, 1}, {1, 3}, {3, 4}, {4, 6}, {6, 7}}},
		{"MergedWithPrevious", false, []pretok.Span{{0, 3}, {3, 6}, {6, 7}}},
		{"MergedWithPrevious", true, []pretok.Span{{0, 1}, {1, 4}, {4, 7}}},
		{"MergedWithNext", false, []pretok.Span{{0, 1}, {1, 4}, {4, 7}}},
		{"MergedWithNext", true, []pretok.Span{{0, 3}, {3, 6}, {6, 7}}},
	} {
		t.Run(fmt.Sprintf("%s/invert=%t", tc.mode, tc.invert), func(t *testing.T) {
			p := compilePipeline(t, fallbackConfig(`,+|;+`, tc.mode, tc.invert))
			checkFallbackSpans(t, p, "a,,b;;c", tc.want)
		})
	}
	p := compilePipeline(t, fallbackConfig(`[,;]`, "Contiguous", false))
	checkFallbackSpans(t, p, "a,;b", []pretok.Span{{0, 1}, {1, 3}, {3, 4}})
}

func TestPipelineRegexpFallbackBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name, pattern, input string
		want                 []pretok.Span
	}{
		{"anchor_is_not_restarted", `^a`, "aaa", []pretok.Span{{0, 1}, {1, 3}}},
		{"word_boundary_context", `a|\bX`, "aX", []pretok.Span{{0, 1}, {1, 2}}},
		{"priority", `ab|a`, "ababa", []pretok.Span{{0, 2}, {2, 4}, {4, 5}}},
		{"empty_regex", ``, "é中", []pretok.Span{{0, 2}, {2, 5}}},
		{"empty_after_match", `a*`, "aé中", []pretok.Span{{0, 1}, {1, 3}, {3, 6}}},
		{"word_boundaries", `\b`, "a b", []pretok.Span{{0, 1}, {1, 2}, {2, 3}}},
		{"eof", `$`, "é中", []pretok.Span{{0, 5}}},
		{"no_match", `xyz`, "é中", []pretok.Span{{0, 5}}},
		{"empty_input", `a*`, "", nil},
		{"exact_spaces", ` [0-9]`, "1 2", []pretok.Span{{0, 1}, {1, 3}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := compilePipeline(t, fallbackConfig(tc.pattern, "Isolated", false))
			checkFallbackSpans(t, p, tc.input, tc.want)
		})
	}
	// A suffix-based search would incorrectly emit X as a second match.
	p := compilePipeline(t, fallbackConfig(`a|\bX`, "Removed", true))
	checkFallbackSpans(t, p, "aX", []pretok.Span{{0, 1}})
}

func TestPipelineRegexpFallbackErrors(t *testing.T) {
	for _, pattern := range []string{`[`, `(?=a)`, `(a)\1`} {
		config := `{"type":"Sequence","pretokenizers":[` + fallbackConfig(pattern, "Isolated", false) + `]}`
		p, err := pretok.CompileJSON([]byte(config))
		var syntaxErr *syntax.Error
		if p != nil || !errors.Is(err, pretok.ErrUnsupportedPattern) || !errors.As(err, &syntaxErr) || !strings.Contains(err.Error(), "pre_tokenizer.pretokenizers[0].pattern.Regex") {
			t.Fatalf("%q: pipeline=%v error=%v", pattern, p, err)
		}
	}
}

func TestPipelineRegexpFallbackSequence(t *testing.T) {
	// Anchors are relative to the complete output piece of the previous stage.
	p := compilePipeline(t, `{"type":"Sequence","pretokenizers":[{"type":"WhitespaceSplit"},`+fallbackConfig(`^é`, "Removed", true)+`,{"type":"ByteLevel","add_prefix_space":false,"use_regex":false}]}`)
	var spans []pretok.Span
	err := p.Walk("éx éy", func(piece pretok.Piece) error {
		if piece.Text() != "Ã©" {
			t.Fatalf("ByteLevel text: %q", piece.Text())
		}
		span, err := piece.OriginalRange(0, len(piece.Text()))
		spans = append(spans, span)
		return err
	})
	if err != nil || !reflect.DeepEqual(spans, []pretok.Span{{0, 2}, {4, 6}}) {
		t.Fatalf("sequence: %v %v", spans, err)
	}
	// Fallback also searches transformed text and preserves its source mapping.
	p = compilePipeline(t, `{"type":"Sequence","pretokenizers":[{"type":"Metaspace","split":false},`+fallbackConfig(`▁[^▁]+`, "Removed", true)+`]}`)
	var texts []string
	spans = nil
	err = p.Walk("é 中", func(piece pretok.Piece) error {
		texts = append(texts, strings.Clone(piece.Text()))
		span, err := piece.OriginalRange(0, len(piece.Text()))
		spans = append(spans, span)
		return err
	})
	if err != nil || !reflect.DeepEqual(texts, []string{"▁é", "▁中"}) || !reflect.DeepEqual(spans, []pretok.Span{{0, 2}, {2, 6}}) {
		t.Fatalf("transformed fallback: %q %v %v", texts, spans, err)
	}
}

func TestPipelineRegexpFallbackReuse(t *testing.T) {
	p := compilePipeline(t, fallbackConfig(`[0-9]+`, "Removed", true))
	stop := errors.New("stop")
	calls := 0
	err := p.WalkTextUnsafe("a12b34", func(string) error { calls++; return stop })
	if err != stop || calls != 1 {
		t.Fatalf("cancellation: %v calls=%d", err, calls)
	}
	var wg sync.WaitGroup
	errorsCh := make(chan error, 8)
	for range 8 {
		wg.Go(func() {
			for range 10 {
				var got []string
				err := p.WalkText("a12b34", func(text string) error {
					got = append(got, text)
					return nil
				})
				if err != nil || !reflect.DeepEqual(got, []string{"12", "34"}) {
					errorsCh <- fmt.Errorf("concurrent fallback: %q %v", got, err)
					return
				}
			}
		})
	}
	wg.Wait()
	close(errorsCh)
	for err := range errorsCh {
		t.Error(err)
	}
}

func FuzzPipelineRegexpFallback(f *testing.F) {
	for _, pattern := range []string{`[0-9]+`, `a|\bX`, `(?m)^.+$`, `a*`, ``, `(?i:é)|中`, `\d`, `a|ab`, `(?U).+`} {
		f.Add(pattern, "aX é中12\r\n ab")
	}
	f.Fuzz(func(t *testing.T, pattern, input string) {
		if len(pattern) > 256 || len(input) > 4096 || !utf8.ValidString(pattern) || !utf8.ValidString(input) {
			t.Skip()
		}
		// The wrapper guarantees an unregistered spelling, even after mutation.
		pattern = "(?:" + pattern + ")"
		re, err := regexp.Compile(pattern)
		if err != nil {
			t.Skip()
		}
		var matches, isolated []pretok.Span
		gap := 0
		for _, m := range re.FindAllStringIndex(input, -1) {
			if m[0] > gap {
				isolated = append(isolated, pretok.Span{Start: gap, End: m[0]})
			}
			if m[1] > m[0] {
				span := pretok.Span{Start: m[0], End: m[1]}
				matches = append(matches, span)
				isolated = append(isolated, span)
			}
			gap = m[1]
		}
		if gap < len(input) {
			isolated = append(isolated, pretok.Span{Start: gap, End: len(input)})
		}
		p := compilePipeline(t, fallbackConfig(pattern, "Removed", true))
		checkFallbackSpans(t, p, input, matches)
		p = compilePipeline(t, fallbackConfig(pattern, "Isolated", false))
		checkFallbackSpans(t, p, input, isolated)
	})
}
