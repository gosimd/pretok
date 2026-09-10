package bench

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gosimd/pretok"
	"github.com/gosimd/pretok/internal/compile"
)

func fixture(t *testing.T, spans [][2]int) (string, Manifest) {
	return fixtureText(t, "a b", map[string][][2]int{"P05": spans})
}
func fixtureText(t *testing.T, text string, spans map[string][][2]int) (string, Manifest) {
	t.Helper()
	dir := t.TempDir()
	write := func(name string, v []byte) string {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		if _, err := w.Write(v); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, name), b.Bytes(), 0644); err != nil {
			t.Fatal(err)
		}
		return Digest(b.Bytes())
	}
	request, _ := json.Marshal([]any{map[string]any{"start": 0, "end": len(text), "spans": spans}})
	totals := map[string]int{}
	for id, s := range spans {
		totals[id] = len(s)
	}
	profiles := map[string]PatternSpec{}
	for id, p := range map[string]string{"P01": compile.QwenMarksPattern, "P05": compile.QwenPattern} {
		profiles[id] = PatternSpec{Pattern: p, PatternSHA: Digest([]byte(p))}
	}
	manifest := Manifest{Schema: 2, Patterns: profiles, Corpora: []CorpusSpec{{ID: "test", Kind: "natural", File: "text.gz", SHA: Digest([]byte(text)), CompressedSHA: write("text.gz", []byte(text)), Bytes: len(text), Scalars: utf8.RuneCountInString(text), Batches: []BatchSpec{{Size: 256, File: "requests.gz", SHA: write("requests.gz", request), Requests: 1, Spans: totals}}}}}
	saveManifest(t, dir, manifest)
	return dir, manifest
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
	dir, _ := fixture(t, [][2]int{{0, 1}, {1, 3}})
	_, batches, err := Load(dir, "P05")
	if err != nil {
		t.Fatal(err)
	}
	if len(batches) != 1 || batches[0].Bytes != 3 || batches[0].Spans != 2 {
		t.Fatal("wrong totals")
	}
	re, _ := pretok.Compile(compile.QwenPattern)
	for _, engine := range []string{"pretok-cursor", "pretok-reused", "pretok-fresh"} {
		run, err := Operation(re, &batches[0], engine)
		if err != nil {
			t.Fatal(err)
		}
		got := run()
		if again := run(); again != got {
			t.Fatal("operation not repeatable")
		}
	}
}
func TestRejectCorruptOrIncorrectReference(t *testing.T) {
	// This partitions the input, but disagrees with the regex: counts are not enough.
	dir, _ := fixture(t, [][2]int{{0, 2}, {2, 3}})
	if _, _, err := Load(dir, "P05"); err == nil || !strings.Contains(err.Error(), "first mismatch") {
		t.Fatalf("want mismatch, got %v", err)
	}
	dir, m := fixture(t, [][2]int{{0, 1}, {1, 3}})
	m.Corpora[0].CompressedSHA = "wrong"
	saveManifest(t, dir, m)
	if _, _, err := Load(dir, "P05"); err == nil || !strings.Contains(err.Error(), "hash mismatch") {
		t.Fatalf("want hash error, got %v", err)
	}
	if _, _, err := Load(t.TempDir(), "P05"); err == nil {
		t.Fatal("missing corpus silently accepted")
	}
}
func TestRejectTraversalAndDuplicateBatches(t *testing.T) {
	if _, err := ReadGzip(t.TempDir(), "../outside.gz", ""); err == nil {
		t.Fatal("accepted path traversal")
	}
	dir, m := fixture(t, [][2]int{{0, 1}, {1, 3}})
	m.Corpora[0].Batches = append(m.Corpora[0].Batches, m.Corpora[0].Batches[0])
	saveManifest(t, dir, m)
	if _, _, err := Load(dir, "P05"); err == nil {
		t.Fatal("duplicate batch accepted")
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
	dir, _ := fixtureText(t, "a\u0301b", map[string][][2]int{"P01": {{0, 4}}, "P05": {{0, 1}, {1, 4}}})
	for id, n := range map[string]int{"P01": 1, "P05": 2} {
		got, batches, err := Load(dir, id)
		if err != nil {
			t.Fatal(err)
		}
		if got.Pattern != got.Patterns[id].Pattern || len(batches) != 1 || batches[0].Spans != n || batches[0].Requests[0].Text != "a\u0301b" {
			t.Fatal("wrong reference selected")
		}
	}
	if _, _, err := Load(dir, "P99"); err == nil {
		t.Fatal("accepted unknown profile")
	}
}

func TestMissingAndSwappedPatternReferences(t *testing.T) {
	dir, _ := fixture(t, [][2]int{{0, 1}, {1, 3}})
	if _, _, err := Load(dir, "P01"); err == nil || !strings.Contains(err.Error(), "missing P01 reference") {
		t.Fatalf("want missing reference, got %v", err)
	}
	dir, _ = fixtureText(t, "a\u0301b", map[string][][2]int{"P01": {{0, 1}, {1, 4}}, "P05": {{0, 4}}})
	for _, id := range []string{"P01", "P05"} {
		if _, _, err := Load(dir, id); err == nil || !strings.Contains(err.Error(), "first mismatch") {
			t.Fatalf("%s: swapped references accepted: %v", id, err)
		}
	}
	dir, m := fixtureText(t, "a\u0301b", map[string][][2]int{"P01": {{0, 4}}})
	delete(m.Corpora[0].Batches[0].Spans, "P01")
	saveManifest(t, dir, m)
	if _, _, err := Load(dir, "P01"); err == nil {
		t.Fatal("missing per-pattern total accepted")
	}
}
