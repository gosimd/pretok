package pretok_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

// Each execution input describes a valid configuration, rather than relying on
// mutated JSON to survive parsing. FuzzCompileJSON separately covers rejection.
// Kind indices: Bert, delimiter, digits, metaspace, punctuation, scripts,
// whitespace, whitespace-split, literal Split, regex Split, ByteLevel.
func pipelineFuzzConfig(recipe []byte, patterns []string) ([]byte, bool) {
	pos := 0
	take := func() byte {
		if pos == len(recipe) {
			return 0
		}
		b := recipe[pos]
		pos++
		return b
	}
	behaviors := []string{"Removed", "Isolated", "Contiguous", "MergedWithPrevious", "MergedWithNext"}
	transformed := false
	leaf := func() map[string]any {
		kind, options := take()%11, take()
		a, b, c := take(), take(), take()
		common := []rune{' ', 'é', '▁', '😀', 0, '\n', 'a', '中', '_'}
		r := common[int(a)%len(common)]
		if a&128 != 0 {
			r = rune((uint32(a)<<16 | uint32(b)<<8 | uint32(c)) % (utf8.MaxRune + 1))
			if !utf8.ValidRune(r) {
				r = utf8.RuneError
			}
		}
		switch kind {
		case 0:
			return map[string]any{"type": "BertPreTokenizer"}
		case 1:
			return map[string]any{"type": "CharDelimiterSplit", "delimiter": string(r)}
		case 2:
			return map[string]any{"type": "Digits", "individual_digits": options&1 != 0}
		case 3:
			transformed = true
			return map[string]any{"type": "Metaspace", "replacement": string(r), "prepend_scheme": []string{"always", "first", "never"}[options%3], "split": options&4 != 0}
		case 4:
			return map[string]any{"type": "Punctuation", "behavior": behaviors[options%5]}
		case 5:
			return map[string]any{"type": "UnicodeScripts"}
		case 6:
			return map[string]any{"type": "Whitespace"}
		case 7:
			return map[string]any{"type": "WhitespaceSplit"}
		case 8, 9:
			pattern := map[string]string{"String": strings.Repeat(string(r), 1+int(options>>5)%3)}
			if kind == 9 {
				pattern = map[string]string{"Regex": patterns[int(options)%len(patterns)]}
			}
			return map[string]any{"type": "Split", "pattern": pattern, "behavior": behaviors[options%5], "invert": options&8 != 0}
		default:
			transformed = true
			return map[string]any{"type": "ByteLevel", "add_prefix_space": options&1 != 0, "trim_offsets": options&2 != 0, "use_regex": options&4 != 0}
		}
	}
	sequence := func(children []any) map[string]any {
		return map[string]any{"type": "Sequence", "pretokenizers": children}
	}
	mode := take() % 4
	var config any
	switch mode {
	case 0:
		config = leaf()
	case 1, 2:
		count := int(take() % 5)
		children := make([]any, 0, count)
		for range count {
			var child any = leaf()
			if mode == 2 {
				child = sequence([]any{sequence([]any{}), child})
			}
			children = append(children, child)
		}
		config = sequence(children)
		if mode == 2 {
			for n := int(take() % 3); n > 0; n-- {
				config = sequence([]any{config})
			}
		}
	case 3:
		config = nil
	}
	data, err := json.Marshal(config)
	if err != nil {
		panic(err) // All generated fields are JSON primitives.
	}
	return data, transformed
}

func FuzzPipelineExecution(f *testing.F) {
	data, err := os.ReadFile("testdata/corpus_patterns.json")
	if err != nil {
		f.Fatal(err)
	}
	var registry []struct{ Pattern string }
	if err := json.Unmarshal(data, &registry); err != nil || len(registry) == 0 {
		f.Fatal("invalid pattern corpus", err)
	}
	patterns := make([]string, len(registry))
	for i, entry := range registry {
		patterns[i] = entry.Pattern
	}
	// Unregistered expressions exercise Go fallback, including empty matches.
	patterns = append(patterns, `[0-9]+`, `\b`, `(?m)^.+$`, `a*`, ``, `é|中`, `\d`, `a|ab`)
	text := "  a_é\u0301中ーБ²Ⅳ٣!😀▁\x00\r\n 123 "
	for kind := byte(0); kind < 11; kind++ {
		// Cover both booleans and each behavior/prefix option before mutation.
		for options := byte(0); options < 16; options++ {
			f.Add([]byte{0, kind, options, options % 9, 0, 0}, text)
		}
		for _, mode := range []byte{1, 2} {
			f.Add([]byte{mode, 3, kind, 0, 1, 0, 0, (kind + 3) % 11, 5, 2, 0, 0, (kind + 7) % 11, 9, 3, 0, 0, 2}, text)
		}
	}
	for i := range patterns {
		f.Add([]byte{0, 9, byte(i), 0, 0, 0}, text)
	}
	for _, input := range []string{"", " ", "e\u0301", "漢あアー", strings.Repeat("é 中12!\n", 256)} {
		f.Add([]byte{2, 3, 7, 0, 0, 0, 0, 3, 5, 2, 0, 0, 10, 0, 0, 0, 0, 2}, input)
	}
	f.Add([]byte{1, 0}, text) // Empty sequence.
	f.Add([]byte{3}, text)    // JSON null identity.
	f.Fuzz(func(t *testing.T, recipe []byte, input string) {
		// Bounds constrain repeated transformation growth. Invalid UTF-8 belongs
		// to the compact facade tests, not a component-by-method fuzz matrix.
		if len(recipe) > 64 || len(input) > 4096 {
			t.Skip()
		}
		// Keep byte mutations productive for stage semantics: construct valid
		// test text here, before calling the API. Production input is untouched.
		input = strings.ToValidUTF8(input, "\uFFFD")
		if len(input) > 4096 {
			t.Skip()
		}
		config, transformed := pipelineFuzzConfig(recipe, patterns)
		t.Cleanup(func() {
			if t.Failed() {
				t.Logf("config=%s input=%q", config, input)
			}
		})
		p := compilePipeline(t, string(config))
		var pieces []string
		previous := pretok.Span{}
		err := p.Walk(input, func(piece pretok.Piece) error {
			text := piece.Text()
			if text == "" || !utf8.ValidString(text) {
				t.Fatal("empty or invalid output")
			}
			checkRange := func(start, end int) pretok.Span {
				span, err := piece.OriginalRange(start, end)
				if err != nil || span.Start < 0 || span.End > len(input) || span.Start >= span.End || !utf8.ValidString(input[span.Start:span.End]) {
					t.Fatalf("invalid mapping [%d,%d): %v, %v", start, end, span, err)
				}
				return span
			}
			whole := checkRange(0, len(text))
			// Mappings may overlap after expansion, but their endpoints cannot
			// move backwards. Removed separators may leave gaps.
			if whole.Start < previous.Start || whole.End < previous.End {
				t.Fatalf("out-of-order mapping: %v after %v", whole, previous)
			}
			previous = whole
			if !transformed && input[whole.Start:whole.End] != text {
				t.Fatal("boundary-only stages changed source bytes")
			}
			for _, i := range []int{0, len(text) / 2, len(text) - 1} {
				part := checkRange(i, i+1)
				if part.Start < whole.Start || part.End > whole.End {
					t.Fatal("byte mapping outside whole-piece mapping")
				}
				prefix, suffix := checkRange(0, i+1), checkRange(i, len(text))
				if prefix.Start != whole.Start || prefix.End != part.End || suffix.Start != part.Start || suffix.End != whole.End {
					t.Fatal("inconsistent byte/prefix/suffix mappings")
				}
			}
			pieces = append(pieces, strings.Clone(text))
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		i := 0
		err = p.WalkTextUnsafe(input, func(text string) error {
			if i >= len(pieces) || text != pieces[i] {
				t.Fatalf("text-only execution differs at piece %d", i)
			}
			i++
			return nil
		})
		if err != nil || i != len(pieces) {
			t.Fatal("text-only execution differs", err, i, len(pieces))
		}
	})
}
