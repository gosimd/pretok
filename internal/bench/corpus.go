// Package bench contains offline benchmark loading, measurement, and reporting.
// It is not part of pretok's runtime or public API.
package bench

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

type PatternSpec struct {
	FullCoverage   bool   `json:"full_coverage"`
	FixturePattern string `json:"fixture_pattern"`
	Pattern        string `json:"pattern"`
	PatternSHA     string `json:"pattern_sha256"`
	FixtureSHA     string `json:"fixture_sha256"`
}
type Manifest struct {
	Schema   int                    `json:"schema"`
	Patterns map[string]PatternSpec `json:"patterns"`
	Corpora  []CorpusSpec           `json:"corpora"`
	// Selected profile for the caller; all inputs come from the common manifest.
	PatternSpec `json:"-"`
}
type CorpusSpec struct {
	ID            string      `json:"id"`
	Kind          string      `json:"kind"`
	File          string      `json:"file"`
	SHA           string      `json:"sha256"`
	CompressedSHA string      `json:"compressed_sha256"`
	Bytes         int         `json:"bytes"`
	Scalars       int         `json:"scalars"`
	SourceFile    string      `json:"source_file"`
	SourceSHA     string      `json:"source_sha256"`
	Batches       []BatchSpec `json:"batches"`
}
type BatchSpec struct {
	Size     int    `json:"size"`
	File     string `json:"file"`
	SHA      string `json:"sha256"`
	Requests int    `json:"requests"`
}
type Request struct {
	Text       string
	Start, End int
}
type Batch struct {
	ID, Kind                    string
	Size, Bytes, Scalars, Spans int
	Requests                    []Request
	InputSHA                    string
	Matches, Capacity           int
	Verified                    bool
}

func Digest(data []byte) string { sum := sha256.Sum256(data); return hex.EncodeToString(sum[:]) }
func file(root, name string) (string, error) {
	if name == "" || !filepath.IsLocal(name) {
		return "", fmt.Errorf("invalid corpus path %q", name)
	}
	return filepath.Join(root, name), nil
}
func readChecked(root, name, sha string) ([]byte, error) {
	path, err := file(root, name)
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w; prepare the corpus first", path, err)
	}
	if Digest(raw) != sha {
		return nil, fmt.Errorf("hash mismatch: %s", path)
	}
	return raw, nil
}
func ReadGzip(root, name, sha string) ([]byte, error) {
	raw, err := readChecked(root, name, sha)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

// Load reads only pinned input bytes and request ranges. Reference generation
// and verification are a separate required step before measurement.
func Load(root, patternID string) (Manifest, []Batch, error) {
	var manifest Manifest
	path, err := file(root, "manifest.json")
	if err != nil {
		return manifest, nil, err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return manifest, nil, fmt.Errorf("read corpus manifest: %w; run benchmark/prepare/corpus.py first", err)
	}
	if err = json.Unmarshal(raw, &manifest); err != nil {
		return manifest, nil, err
	}
	if manifest.Schema != 3 {
		return manifest, nil, fmt.Errorf("unsupported corpus schema %d", manifest.Schema)
	}
	selected, ok := manifest.Patterns[patternID]
	if !ok {
		return manifest, nil, fmt.Errorf("missing pattern %q in common corpus manifest", patternID)
	}
	manifest.PatternSpec = selected
	if Digest([]byte(manifest.Pattern)) != manifest.PatternSHA {
		return manifest, nil, fmt.Errorf("invalid manifest schema or pattern hash")
	}
	_, err = pretok.Compile(manifest.Pattern)
	if err != nil {
		return manifest, nil, err
	}
	var batches []Batch
	seen := map[string]bool{}
	for _, c := range manifest.Corpora {
		if seen[c.ID] {
			return manifest, nil, fmt.Errorf("duplicate corpus %q", c.ID)
		}
		seen[c.ID] = true
		data, err := ReadGzip(root, c.File, c.CompressedSHA)
		if err != nil {
			return manifest, nil, err
		}
		if len(data) != c.Bytes || Digest(data) != c.SHA || !utf8.Valid(data) || utf8.RuneCount(data) != c.Scalars {
			return manifest, nil, fmt.Errorf("invalid corpus %s", c.ID)
		}
		if c.SourceFile != "" {
			p, e := file(root, c.SourceFile)
			if e != nil {
				return manifest, nil, e
			}
			source, e := os.ReadFile(p)
			if e != nil {
				return manifest, nil, e
			}
			if Digest(source) != c.SourceSHA {
				return manifest, nil, fmt.Errorf("source manifest mismatch for %s", c.ID)
			}
		}
		text := string(data)
		sizes := map[int]bool{}
		for _, spec := range c.Batches {
			if sizes[spec.Size] {
				return manifest, nil, fmt.Errorf("duplicate size for %s", c.ID)
			}
			sizes[spec.Size] = true
			raw, err := readChecked(root, spec.File, spec.SHA)
			if err != nil {
				return manifest, nil, err
			}
			var cases []struct {
				Start int `json:"start"`
				End   int `json:"end"`
			}
			if err = json.Unmarshal(raw, &cases); err != nil {
				return manifest, nil, err
			}
			if len(cases) != spec.Requests {
				return manifest, nil, fmt.Errorf("request count mismatch")
			}
			b := Batch{ID: c.ID, Kind: c.Kind, Size: spec.Size, InputSHA: c.SHA}
			previous := 0
			for i, tc := range cases {
				if tc.Start != previous || tc.End < tc.Start || tc.End > len(text) || !utf8.ValidString(text[tc.Start:tc.End]) {
					return manifest, nil, fmt.Errorf("invalid %s/%d request %d", c.ID, spec.Size, i)
				}
				previous = tc.End
				src := text[tc.Start:tc.End]
				b.Requests = append(b.Requests, Request{Text: src, Start: tc.Start, End: tc.End})
				b.Bytes += len(src)
				b.Scalars += utf8.RuneCountInString(src)
			}
			if previous != len(text) {
				return manifest, nil, fmt.Errorf("batch totals mismatch")
			}
			batches = append(batches, b)
		}
	}
	if len(batches) == 0 {
		return manifest, nil, fmt.Errorf("empty benchmark corpus")
	}
	return manifest, batches, nil
}
