package pretok_test

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

const pipelineExample = `{"type":"Sequence","pretokenizers":[{"type":"Split","pattern":{"Regex":"[0-9]"},"behavior":"Isolated","invert":false},{"type":"ByteLevel","add_prefix_space":false,"trim_offsets":false,"use_regex":false}]}`

func compilePipeline(t testing.TB, config string) *pretok.Pipeline {
	t.Helper()
	p, err := pretok.CompileJSON([]byte(config))
	if err != nil {
		t.Fatal(err)
	}
	return p
}

type pipelineCase struct {
	Config json.RawMessage `json:"config"`
	Input  string          `json:"input"`
	Pieces []struct {
		Text   string   `json:"text"`
		Span   [2]int   `json:"span"`
		Ranges [][4]int `json:"ranges"`
	} `json:"pieces"`
}

func TestPipelineHFOracle(t *testing.T) {
	f, err := os.Open("testdata/pipeline/cases.json.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	z, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer z.Close()
	data, err := io.ReadAll(z)
	if err != nil {
		t.Fatal(err)
	}
	manifest, err := os.ReadFile("testdata/pipeline/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var meta struct {
		Version string `json:"version"`
		Count   int    `json:"case_count"`
		Hash    string `json:"cases_sha256"`
		Tables  string `json:"tables_sha256"`
	}
	if err := json.Unmarshal(manifest, &meta); err != nil {
		t.Fatal(err)
	}
	if meta.Version != "0.23.2" || fmt.Sprintf("%x", sha256.Sum256(data)) != meta.Hash {
		t.Fatal("oracle provenance mismatch")
	}
	tables, err := os.ReadFile("internal/pipeline/tables.go")
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%x", sha256.Sum256(tables)) != meta.Tables {
		t.Fatal("built-in table hash mismatch")
	}
	var cases []pipelineCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) != meta.Count {
		t.Fatal("oracle count mismatch")
	}
	for index, c := range cases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			p := compilePipeline(t, string(c.Config))
			i := 0
			err := p.Walk(c.Input, func(piece pretok.Piece) error {
				if i >= len(c.Pieces) {
					t.Fatalf("extra %q config=%s input=%q", piece.Text(), c.Config, c.Input)
				}
				want := c.Pieces[i]
				i++
				if piece.Text() != want.Text {
					t.Fatalf("text=%q want=%q config=%s input=%q", piece.Text(), want.Text, c.Config, c.Input)
				}
				span, err := piece.OriginalRange(0, len(piece.Text()))
				if err != nil || span != (pretok.Span{Start: want.Span[0], End: want.Span[1]}) {
					t.Fatalf("whole span=%v %v want=%v config=%s input=%q", span, err, want.Span, c.Config, c.Input)
				}
				for _, r := range want.Ranges {
					got, err := piece.OriginalRange(r[0], r[1])
					if err != nil || got != (pretok.Span{Start: r[2], End: r[3]}) {
						t.Fatalf("%q range=%v got=%v %v config=%s input=%q", piece.Text(), r, got, err, c.Config, c.Input)
					}
				}
				return nil
			})
			if err != nil || i != len(c.Pieces) {
				t.Fatalf("Walk: %d/%d pieces, %v", i, len(c.Pieces), err)
			}
			i = 0
			var retained []string
			err = p.WalkText(c.Input, func(s string) error {
				if i >= len(c.Pieces) || s != c.Pieces[i].Text {
					t.Fatalf("WalkText differs at %d: %q config=%s input=%q", i, s, c.Config, c.Input)
				}
				i++
				retained = append(retained, s)
				return nil
			})
			if err != nil || i != len(c.Pieces) {
				t.Fatalf("WalkText: %d/%d pieces, %v", i, len(c.Pieces), err)
			}
			i = 0
			err = p.WalkTextUnsafe(c.Input, func(s string) error {
				if i >= len(c.Pieces) || s != c.Pieces[i].Text {
					t.Fatalf("WalkTextUnsafe differs at %d: %q config=%s input=%q", i, s, c.Config, c.Input)
				}
				i++
				return nil
			})
			if err != nil || i != len(c.Pieces) {
				t.Fatalf("WalkTextUnsafe: %d/%d pieces, %v", i, len(c.Pieces), err)
			}
			for i, s := range retained {
				if s != c.Pieces[i].Text {
					t.Fatalf("retained WalkText string changed at %d: %q", i, s)
				}
			}
		})
	}
}

func TestCompilePipelineErrors(t *testing.T) {
	for _, config := range []string{`{"type":"CharDelimiterSplit","delimiter":"\ud800"}`, `{"type":"CharDelimiterSplit","delimiter":"\udc00"}`, `{"type":"CharDelimiterSplit","delimiter":"\ud800a"}`} {
		if _, err := pretok.CompileJSON([]byte(config)); !errors.Is(err, pretok.ErrInvalidConfig) {
			t.Errorf("unpaired surrogate accepted: %s %v", config, err)
		}
	}
	if _, err := pretok.CompileJSON([]byte(`{"type":"CharDelimiterSplit","delimiter":"\ud83d\ude00"}`)); err != nil {
		t.Fatal("valid surrogate pair", err)
	}
	if _, err := pretok.CompileJSON([]byte("{\"type\":\"Whitespace\",\"\xff\":0}")); !errors.Is(err, pretok.ErrInvalidConfig) {
		t.Fatal("invalid configuration UTF-8", err)
	}
	for _, config := range []string{"", `{}`, `[]`, `true`, `{"pre_tokenizer":null}`, `{"type":"Digits","typo":true}`, `{"type":"Digits","individual_digits":null}`, `{"type":"Digits","type":"Whitespace"}`, `{"type":"Digits"} {}`, `{"type":"Sequence"}`, `{"type":"Sequence","pretokenizers":[null]}`, `{"type":"Split","pattern":{"String":""}}`, `{"type":"Split","pattern":{"String":"a","Regex":"[0-9]"}}`, `{"type":"CharDelimiterSplit","delimiter":"ab"}`, `{"type":"Metaspace","replacement":""}`, `{"type":"Metaspace","prepend_scheme":"sometimes"}`, `{"type":"Punctuation","behavior":"isolated"}`} {
		p, err := pretok.CompileJSON([]byte(config))
		if p != nil || !errors.Is(err, pretok.ErrInvalidConfig) {
			t.Errorf("%s => %v %v", config, p, err)
		}
	}
	for _, expr := range []string{`[0-9]+`, `[`, ` [0-9]`, `\d`} {
		config, _ := json.Marshal(map[string]any{"type": "Split", "pattern": map[string]string{"Regex": expr}})
		if _, err := pretok.CompileJSON(config); !errors.Is(err, pretok.ErrUnsupportedPattern) {
			t.Errorf("%s: %v", config, err)
		}
	}
	if _, err := pretok.CompileJSON([]byte(`{"type":"Unknown"}`)); !errors.Is(err, pretok.ErrUnsupportedPreTokenizer) {
		t.Fatal(err)
	}
	nested := `{"type":"Digits"}`
	for range 34 {
		nested = `{"type":"Sequence","pretokenizers":[` + nested + `]}`
	}
	if _, err := pretok.CompileJSON([]byte(nested)); !errors.Is(err, pretok.ErrInvalidConfig) {
		t.Fatal("missing nesting bound", err)
	}
	stages := strings.TrimSuffix(strings.Repeat(`{"type":"Digits"},`, 257), ",")
	if _, err := pretok.CompileJSON([]byte(`{"type":"Sequence","pretokenizers":[` + stages + `]}`)); !errors.Is(err, pretok.ErrInvalidConfig) {
		t.Fatal("missing stage bound", err)
	}
}

func TestPipelineJSONOwnershipAndTrim(t *testing.T) {
	config := []byte(`{"type":"ByteLevel","add_prefix_space":true,"trim_offsets":true,"use_regex":false}`)
	p, err := pretok.CompileJSON(config)
	if err != nil {
		t.Fatal(err)
	}
	for i := range config {
		config[i] = 0
	}
	err = p.Walk(" é ", func(piece pretok.Piece) error {
		got, err := piece.OriginalRange(0, len(piece.Text()))
		if err != nil || got != (pretok.Span{Start: 0, End: 4}) {
			t.Fatalf("trim_offsets incorrectly applied: %v %v", got, err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPipelineErrorsAndCancellation(t *testing.T) {
	for _, p := range []*pretok.Pipeline{nil, new(pretok.Pipeline)} {
		if err := p.Walk("", nil); err != pretok.ErrUninitialized {
			t.Fatal(err)
		}
		if err := p.WalkText("", nil); err != pretok.ErrUninitialized {
			t.Fatal(err)
		}
		if err := p.WalkTextUnsafe("", nil); err != pretok.ErrUninitialized {
			t.Fatal(err)
		}
	}
	p := compilePipeline(t, pipelineExample)
	if p.Walk("", nil) != pretok.ErrNilCallback || p.WalkText("", nil) != pretok.ErrNilCallback || p.WalkTextUnsafe("", nil) != pretok.ErrNilCallback {
		t.Fatal("nil callback")
	}
	for _, input := range []string{"ok\xff", "\xc0\xaf", strings.Repeat("é1", 2000) + "\xed\xa0\x80"} {
		if err := p.Walk(input, func(pretok.Piece) error { t.Fatal("callback on invalid UTF-8"); return nil }); err != pretok.ErrInvalidUTF8 {
			t.Fatal(err)
		}
		if err := p.WalkText(input, func(string) error { t.Fatal("callback on invalid UTF-8"); return nil }); err != pretok.ErrInvalidUTF8 {
			t.Fatal(err)
		}
		if err := p.WalkTextUnsafe(input, func(string) error { t.Fatal("callback on invalid UTF-8"); return nil }); err != pretok.ErrInvalidUTF8 {
			t.Fatal(err)
		}
	}
	stop := errors.New("consumer failed")
	for _, mode := range []string{"Walk", "WalkText", "WalkTextUnsafe"} {
		count := 0
		var err error
		if mode == "WalkTextUnsafe" {
			err = p.WalkTextUnsafe("a123é", func(string) error { count++; return stop })
		} else if mode == "WalkText" {
			err = p.WalkText("a123é", func(string) error { count++; return stop })
		} else {
			err = p.Walk("a123é", func(pretok.Piece) error { count++; return stop })
		}
		if err != stop || count != 1 {
			t.Fatalf("cancellation: %v %d", err, count)
		}
	}
	err := p.Walk("é", func(piece pretok.Piece) error {
		for _, bounds := range [][2]int{{-1, 1}, {0, 0}, {2, 1}, {0, len(piece.Text()) + 1}} {
			if _, err := piece.OriginalRange(bounds[0], bounds[1]); err != pretok.ErrInvalidRange {
				t.Fatal(bounds, err)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := (pretok.Piece{}).OriginalRange(0, 1); err != pretok.ErrInvalidRange {
		t.Fatal(err)
	}
}

func TestPipelineConcurrencyReentrancyAndPanic(t *testing.T) {
	p := compilePipeline(t, pipelineExample)
	want := []string{"ab", "1", "2", "Ã©"}
	check := func() error {
		var got []string
		err := p.Walk("ab12é", func(piece pretok.Piece) error {
			before := strings.Clone(piece.Text())
			if err := p.WalkText("中7", func(string) error { return nil }); err != nil {
				return err
			}
			if err := p.WalkTextUnsafe("中7", func(string) error { return nil }); err != nil {
				return err
			}
			if piece.Text() != before {
				return errors.New("nested call overwrote outer piece")
			}
			got = append(got, before)
			return nil
		})
		if err != nil {
			return err
		}
		if !reflect.DeepEqual(got, want) {
			return fmt.Errorf("got %q", got)
		}
		return nil
	}
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 25 {
				if err := check(); err != nil {
					t.Error(err)
					return
				}
			}
		}()
	}
	wg.Wait()
	for _, walk := range []func(string, func(string) error) error{p.WalkText, p.WalkTextUnsafe} {
		func() {
			defer func() {
				if recover() != "stop" {
					t.Error("panic did not propagate")
				}
			}()
			_ = walk("é", func(string) error { panic("stop") })
		}()
	}
	if err := check(); err != nil {
		t.Fatal(err)
	}
}

func FuzzPipelineWalkParity(f *testing.F) {
	for _, s := range []string{"", "ab12é", "é😀中", "a\xff", " a\r\n", "▁"} {
		f.Add(s)
	}
	configs := []string{pipelineExample, `{"type":"ByteLevel"}`, `{"type":"Sequence","pretokenizers":[{"type":"Metaspace","prepend_scheme":"first"},{"type":"ByteLevel","use_regex":false},{"type":"Whitespace"}]}`}
	ps := make([]*pretok.Pipeline, len(configs))
	for i, c := range configs {
		ps[i] = compilePipeline(f, c)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 16384 {
			t.Skip()
		}
		for _, p := range ps {
			var got []string
			err := p.Walk(s, func(piece pretok.Piece) error {
				if piece.Text() == "" || !utf8.ValidString(piece.Text()) {
					t.Fatal("invalid output")
				}
				span, e := piece.OriginalRange(0, len(piece.Text()))
				if e != nil || span.Start < 0 || span.End > len(s) || span.Start >= span.End || !utf8.ValidString(s[span.Start:span.End]) {
					t.Fatal("invalid source range", span, e)
				}
				got = append(got, strings.Clone(piece.Text()))
				return nil
			})
			if !utf8.ValidString(s) {
				if err != pretok.ErrInvalidUTF8 || len(got) != 0 {
					t.Fatal(err)
				}
				continue
			}
			if err != nil {
				t.Fatal(err)
			}
			for _, walk := range []func(string, func(string) error) error{p.WalkText, p.WalkTextUnsafe} {
				i := 0
				err = walk(s, func(piece string) error {
					if i >= len(got) || piece != got[i] {
						t.Fatal("WalkText mismatch")
					}
					i++
					return nil
				})
				if err != nil || i != len(got) {
					t.Fatal(err, i, len(got))
				}
			}
		}
	})
}

// A consumer may retain safe strings or explicitly clone unsafe strings. Readers
// begin asynchronously after callback return, while later execution reuses its
// transformation buffers. This must also be race-free.
func TestPipelineTextOwnership(t *testing.T) {
	p := compilePipeline(t, pipelineExample)
	want := []string{"ab", "1", "2", "Ã©"}
	for _, cloneUnsafe := range []bool{false, true} {
		walk := p.WalkText
		if cloneUnsafe {
			walk = p.WalkTextUnsafe
		}
		var retained []string
		start := make(chan struct{})
		var wg sync.WaitGroup
		err := walk("ab12é", func(s string) error {
			if cloneUnsafe {
				s = strings.Clone(s)
			}
			i := len(retained)
			retained = append(retained, s)
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-start
				for range 100 {
					if s != want[i] {
						t.Errorf("retained string %d changed: %q", i, s)
						return
					}
				}
			}()
			return nil
		})
		close(start)
		if err != nil {
			wg.Wait()
			t.Fatal(err)
		}
		for range 20 {
			if err := p.WalkTextUnsafe("zz98à", func(string) error { return nil }); err != nil {
				t.Error(err)
			}
		}
		wg.Wait()
		if !reflect.DeepEqual(retained, want) {
			t.Fatalf("retained strings: %q", retained)
		}
		// Consumer errors must not invalidate text already delivered safely.
		stop := errors.New("stop")
		var first string
		if err := walk("ab12é", func(s string) error {
			if cloneUnsafe {
				s = strings.Clone(s)
			}
			first = s
			return stop
		}); err != stop {
			t.Fatal(err)
		}
		if err := p.WalkTextUnsafe("zz98à", func(string) error { return nil }); err != nil {
			t.Fatal(err)
		}
		if first != "ab" {
			t.Fatalf("string invalidated after callback error: %q", first)
		}
	}
}

func FuzzCompileJSON(f *testing.F) {
	for _, config := range []string{pipelineExample, `null`, `{"type":"Metaspace"}`, `{"type":"CharDelimiterSplit","delimiter":"\ud83d\ude00"}`, `{"type":"CharDelimiterSplit","delimiter":"\ud800"}`, `{"type":"Whitespace","type":"Digits"}`, `[`, "\xff"} {
		f.Add(config)
	}
	f.Fuzz(func(t *testing.T, config string) {
		if len(config) > 32768 {
			t.Skip()
		}
		p, err := pretok.CompileJSON([]byte(config))
		if err != nil {
			if p != nil {
				t.Fatal("non-nil pipeline on error")
			}
			return
		}
		if p == nil {
			t.Fatal("nil pipeline on success")
		}
	})
}
