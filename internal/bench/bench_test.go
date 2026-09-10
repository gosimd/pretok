package bench

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
	"github.com/gosimd/pretok/internal/compile"
)

func fixture(t *testing.T, text string) (string, Manifest) {
	t.Helper()
	dir := t.TempDir()
	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	if _, err := w.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "text.gz"), compressed.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}
	request, _ := json.Marshal([]any{map[string]any{"start": 0, "end": len(text)}})
	if err := os.WriteFile(filepath.Join(dir, "requests.json"), request, 0644); err != nil {
		t.Fatal(err)
	}
	profiles := map[string]PatternSpec{}
	for id, p := range map[string]string{"P01": compile.QwenMarksPattern, "P02": compile.ASCIITriplesPattern, "P05": compile.QwenPattern} {
		profiles[id] = PatternSpec{Pattern: p, PatternSHA: Digest([]byte(p))}
	}
	m := Manifest{Schema: 3, Patterns: profiles, Corpora: []CorpusSpec{{ID: "test", Kind: "natural", File: "text.gz", SHA: Digest([]byte(text)), CompressedSHA: Digest(compressed.Bytes()), Bytes: len(text), Scalars: utf8.RuneCountInString(text), Batches: []BatchSpec{{Size: 256, File: "requests.json", SHA: Digest(request), Requests: 1}}}}}
	saveManifest(t, dir, m)
	return dir, m
}
func reference(b Batch, p string, matches, split [][2]int) Reference {
	return Reference{Corpus: b.ID, Size: b.Size, PatternSHA: Digest([]byte(p)), InputSHA: b.InputSHA, Requests: []ReferenceRequest{{Start: 0, End: len(b.Requests[0].Text), Matches: matches, Split: split}}}
}
func saveManifest(t *testing.T, dir string, m Manifest) {
	t.Helper()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(dir, "manifest.json"), b, 0644); err != nil {
		t.Fatal(err)
	}
}
func TestLoadAndGate(t *testing.T) {
	dir, _ := fixture(t, "a b")
	_, batches, err := Load(dir, "P05")
	if err != nil {
		t.Fatal(err)
	}
	b := &batches[0]
	re, _ := pretok.Compile(compile.QwenPattern)
	if b.Verified || b.Spans != 0 {
		t.Fatal("load constructed references")
	}
	if _, err := Operation(re, b, "pretok-reused"); err == nil {
		t.Fatal("unverified input measured")
	}
	spans := [][2]int{{0, 1}, {1, 3}}
	if err := Verify(re, Digest([]byte(compile.QwenPattern)), b, reference(*b, compile.QwenPattern, spans, spans)); err != nil {
		t.Fatal(err)
	}
	if b.Bytes != 3 || b.Spans != 2 || b.Matches != 2 || b.Capacity != 2 {
		t.Fatal("wrong verified totals")
	}
	for _, engine := range []string{"pretok-cursor", "pretok-reused", "pretok-fresh"} {
		op, err := Operation(re, b, engine)
		if err != nil {
			t.Fatal(err)
		}
		if op() != op() {
			t.Fatal("operation not repeatable")
		}
	}
}
func TestRejectCorruptInputs(t *testing.T) {
	dir, m := fixture(t, "a b")
	m.Corpora[0].CompressedSHA = "wrong"
	saveManifest(t, dir, m)
	if _, _, err := Load(dir, "P05"); err == nil || !strings.Contains(err.Error(), "hash mismatch") {
		t.Fatalf("want hash error: %v", err)
	}
	if _, _, err := Load(t.TempDir(), "P05"); err == nil {
		t.Fatal("missing inputs accepted")
	}
	if _, err := ReadGzip(t.TempDir(), "../outside.gz", ""); err == nil {
		t.Fatal("path traversal accepted")
	}
	dir, m = fixture(t, "a b")
	m.Corpora[0].Batches = append(m.Corpora[0].Batches, m.Corpora[0].Batches[0])
	saveManifest(t, dir, m)
	if _, _, err := Load(dir, "P05"); err == nil {
		t.Fatal("duplicate batch accepted")
	}
}
func TestVerifyRejectsWrongOrIncompleteOracle(t *testing.T) {
	dir, _ := fixture(t, "a b")
	m, batches, err := Load(dir, "P05")
	if err != nil {
		t.Fatal(err)
	}
	re, _ := pretok.Compile(m.Pattern)
	for _, kind := range []string{"split", "matches", "pattern", "input", "range", "count", "corpus"} {
		t.Run(kind, func(t *testing.T) {
			b := batches[0]
			spans := [][2]int{{0, 1}, {1, 3}}
			ref := reference(b, m.Pattern, spans, spans)
			switch kind {
			case "split":
				ref.Requests[0].Split = [][2]int{{0, 2}, {2, 3}}
			case "matches":
				ref.Requests[0].Matches = [][2]int{{0, 3}}
			case "pattern":
				ref.PatternSHA = "wrong"
			case "input":
				ref.InputSHA = "wrong"
			case "range":
				ref.Requests[0].End--
			case "count":
				ref.Requests = nil
			case "corpus":
				ref.Corpus = "other"
			}
			if err := Verify(re, m.PatternSHA, &b, ref); err == nil {
				t.Fatal("incorrect oracle accepted")
			}
			if b.Verified {
				t.Fatal("failed batch marked verified")
			}
			if _, err := Operation(re, &b, "pretok-reused"); err == nil {
				t.Fatal("failed batch can be measured")
			}
		})
	}
}
func TestReferenceSpanValidation(t *testing.T) {
	for _, raw := range [][][2]int{{{0, 1}}, {{-1, 3}}, {{0, 4}}, {{1, 2}}, {{0, 3}, {2, 3}}, {{0, 0}}} {
		if _, err := checkedSpans("界", raw, false); err == nil {
			t.Fatalf("accepted %v", raw)
		}
	}
	if _, err := checkedSpans("a b", [][2]int{{2, 3}}, false); err != nil {
		t.Fatal("match gaps must be allowed", err)
	}
	if _, err := checkedSpans("a b", [][2]int{{2, 3}}, true); err == nil {
		t.Fatal("split gaps must be rejected")
	}
}
func TestReportWeightedAggregate(t *testing.T) {
	var samples []Sample
	for _, tc := range []struct {
		id                string
		reference, pretok uint64
	}{{"a", 100, 100}, {"b", 900, 300}} {
		for engine, ns := range map[string]uint64{"onig-iter": tc.reference, "pretok-cursor": tc.pretok} {
			samples = append(samples, Sample{Engine: engine, Corpus: tc.id, Kind: "natural", Size: 256, Iterations: 1, ElapsedNS: ns, Bytes: 100, Requests: 2, Scalars: 80, Spans: 10})
		}
	}
	var out bytes.Buffer
	if err := WriteReport(&out, samples, "environment.json", "P01"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "Combined selected natural corpora | 256 | 2.500x") {
		t.Fatal("aggregate is not ratio of summed times")
	}
	if !strings.HasPrefix(out.String(), "# P01 benchmark") {
		t.Fatal("wrong pattern label")
	}
	if !strings.Contains(out.String(), "N/A") {
		t.Fatal("missing allocations presented as zero")
	}
	if got := quantile([]float64{100, 1, 3, 2}, .5); got != 2.5 {
		t.Fatal(got)
	}
}

func TestSharedCorpusManifest(t *testing.T) {
	dir, _ := fixture(t, "a\u0301b")
	for id, spans := range map[string][][2]int{"P01": {{0, 4}}, "P05": {{0, 1}, {1, 4}}} {
		m, batches, err := Load(dir, id)
		if err != nil {
			t.Fatal(err)
		}
		re, _ := pretok.Compile(m.Pattern)
		if err := Verify(re, m.PatternSHA, &batches[0], reference(batches[0], m.Pattern, spans, spans)); err != nil {
			t.Fatal(err)
		}
		if batches[0].Spans != len(spans) {
			t.Fatal("wrong profile verified")
		}
	}
	if _, _, err := Load(dir, "P99"); err == nil {
		t.Fatal("unknown profile accepted")
	}
}

func TestP02LiveGateWithGaps(t *testing.T) {
	for _, tc := range []struct {
		text           string
		matches, split [][2]int
	}{
		{"a123b12", [][2]int{{1, 4}}, [][2]int{{0, 1}, {1, 4}, {4, 7}}},
		{"１２3", nil, [][2]int{{0, 7}}},
		{"", nil, nil},
	} {
		dir, _ := fixture(t, tc.text)
		m, batches, err := Load(dir, "P02")
		if err != nil {
			t.Fatal(err)
		}
		b := &batches[0]
		re, err := pretok.Compile(m.Pattern)
		if err != nil {
			t.Fatal(err)
		}
		if err = Verify(re, m.PatternSHA, b, reference(*b, m.Pattern, tc.matches, tc.split)); err != nil {
			t.Fatal(err)
		}
		if b.OutputCount("pretok-cursor") != len(tc.matches) || b.OutputCount("onig-iter") != len(tc.matches) || b.OutputCount("pretok-reused") != len(tc.split) || b.OutputCount("hf-split") != len(tc.split) {
			t.Fatal("matches and isolated split counts conflated")
		}
		for _, engine := range []string{"pretok-cursor", "pretok-reused", "pretok-fresh"} {
			op, err := Operation(re, b, engine)
			if err != nil {
				t.Fatal(err)
			}
			if op() != op() {
				t.Fatal("operation not repeatable")
			}
		}
	}
	if slices.Contains(Engines(false), "onig-at") || !slices.Contains(Engines(true), "onig-at") {
		t.Fatal("match-at diagnostic scope")
	}
}
