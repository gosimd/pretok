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
	"slices"
	"unicode/utf8"

	"github.com/gosimd/pretok"
)

type Manifest struct {
	Schema     int          `json:"schema"`
	Pattern    string       `json:"pattern"`
	PatternSHA string       `json:"pattern_sha256"`
	FixtureSHA string       `json:"fixture_sha256"`
	Corpora    []CorpusSpec `json:"corpora"`
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
	Spans    int    `json:"spans"`
}
type Request struct {
	Text     string
	Expected []pretok.Span
}
type Batch struct {
	ID, Kind                    string
	Size, Bytes, Scalars, Spans int
	Requests                    []Request
}

func Digest(data []byte) string { sum := sha256.Sum256(data); return hex.EncodeToString(sum[:]) }
func file(root, name string) (string, error) {
	if name == "" || !filepath.IsLocal(name) {
		return "", fmt.Errorf("invalid corpus path %q", name)
	}
	return filepath.Join(root, name), nil
}
func ReadGzip(root, name, sha string) ([]byte, error) {
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
	r, err := gzip.NewReader(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
func Load(root string) (Manifest, []Batch, error) {
	var manifest Manifest
	raw, err := os.ReadFile(filepath.Join(root, "manifest.json"))
	if err != nil {
		return manifest, nil, fmt.Errorf("read corpus manifest: %w; run benchmark/prepare/corpus.py first", err)
	}
	if err = json.Unmarshal(raw, &manifest); err != nil {
		return manifest, nil, err
	}
	if manifest.Schema != 1 || Digest([]byte(manifest.Pattern)) != manifest.PatternSHA {
		return manifest, nil, fmt.Errorf("invalid manifest schema or pattern hash")
	}
	re, err := pretok.Compile(manifest.Pattern)
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
			raw, err := ReadGzip(root, spec.File, spec.SHA)
			if err != nil {
				return manifest, nil, err
			}
			var cases []struct {
				Start int      `json:"start"`
				End   int      `json:"end"`
				Spans [][2]int `json:"spans"`
			}
			if err = json.Unmarshal(raw, &cases); err != nil {
				return manifest, nil, err
			}
			if len(cases) != spec.Requests {
				return manifest, nil, fmt.Errorf("request count mismatch")
			}
			b := Batch{ID: c.ID, Kind: c.Kind, Size: spec.Size}
			previous := 0
			for i, tc := range cases {
				if tc.Start != previous || tc.End < tc.Start || tc.End > len(text) || !utf8.ValidString(text[tc.Start:tc.End]) {
					return manifest, nil, fmt.Errorf("invalid %s/%d request %d", c.ID, spec.Size, i)
				}
				previous = tc.End
				src := text[tc.Start:tc.End]
				want := make([]pretok.Span, len(tc.Spans))
				end := 0
				for j, s := range tc.Spans {
					if s[0] != end || s[1] <= s[0] || s[1] > len(src) || !utf8.ValidString(src[s[0]:s[1]]) {
						return manifest, nil, fmt.Errorf("invalid reference span %s/%d/%d/%d", c.ID, spec.Size, i, j)
					}
					want[j] = pretok.Span{Start: s[0], End: s[1]}
					end = s[1]
				}
				if end != len(src) {
					return manifest, nil, fmt.Errorf("incomplete reference partition")
				}
				got, err := re.Split(src, nil)
				if err != nil {
					return manifest, nil, err
				}
				if !slices.Equal(got, want) {
					return manifest, nil, mismatch("pretok-split", c.ID, spec.Size, i, src, got, want)
				}
				cursor, err := re.Cursor(src)
				if err != nil {
					return manifest, nil, err
				}
				got = got[:0]
				for s, ok := cursor.Next(); ok; s, ok = cursor.Next() {
					got = append(got, s)
				}
				if !slices.Equal(got, want) {
					return manifest, nil, mismatch("pretok-cursor", c.ID, spec.Size, i, src, got, want)
				}
				b.Requests = append(b.Requests, Request{Text: src, Expected: want})
				b.Bytes += len(src)
				b.Scalars += utf8.RuneCountInString(src)
				b.Spans += len(want)
			}
			if previous != len(text) || b.Spans != spec.Spans {
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
func mismatch(engine, id string, size, index int, src string, got, want []pretok.Span) error {
	at := 0
	for at < len(got) && at < len(want) && got[at] == want[at] {
		at++
	}
	start := 0
	if at < len(want) {
		start = want[at].Start
	}
	lo := max(0, start-30)
	hi := min(len(src), start+90)
	return fmt.Errorf("%s %s/%d request %d: first mismatch span %d, got %v expected %v, nearby bytes %q", engine, id, size, index, at, got[max(0, at-1):min(len(got), at+2)], want[max(0, at-1):min(len(want), at+2)], src[lo:hi])
}
