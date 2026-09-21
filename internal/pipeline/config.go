// Package pipeline implements compiled HF pre-tokenizer stages. Public API and
// input validation belong to the root facade; execution state is per call.
package pipeline

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/gosimd/pretok/internal/compile"
	"github.com/gosimd/pretok/internal/scan"
)

var (
	ErrInvalidConfig      = errors.New("pretok: invalid pre-tokenizer configuration")
	ErrUnsupportedType    = errors.New("pretok: unsupported pre-tokenizer")
	ErrUnsupportedPattern = errors.New("unsupported regex pattern")
	ErrInvalidRange       = errors.New("pretok: invalid piece range")
)

type kind uint8

const (
	splitStage kind = iota
	byteLevelStage
	metaspaceStage
	whitespaceStage
	whitespaceSplitStage
	digitsStage
	punctuationStage
	bertStage
	scriptsStage
)

type behavior uint8

const (
	isolated behavior = iota
	removed
	contiguous
	mergedPrevious
	mergedNext
)

type stage struct {
	kind                                kind
	mode                                behavior
	invert                              bool
	literal                             string
	regex                               *scan.Plan
	fallback                            *regexp.Regexp
	prefix, useRegex, individual, split bool
	prepend                             string
}

// Program contains immutable stages; no input, scratch, or callbacks are stored.
type Program struct{ stages []stage }

// Compile accepts one JSON pre_tokenizer value. Limits bound recursive compile
// and execution stacks; unknown fields and duplicate keys are rejected.
func Compile(config []byte) (*Program, error) {
	if !utf8.Valid(config) {
		return nil, ErrInvalidConfig
	}
	if err := checkJSON(config); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}
	p := new(Program)
	if err := p.add(config, 0, "pre_tokenizer"); err != nil {
		return nil, err
	}
	return p, nil
}

func checkJSON(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	var value func(int) error
	value = func(depth int) error {
		if depth > 80 {
			return errors.New("JSON nesting limit exceeded")
		}
		t, err := d.Token()
		if err != nil {
			return err
		}
		if delim, ok := t.(json.Delim); ok {
			switch delim {
			case '{':
				seen := map[string]bool{}
				for d.More() {
					key, err := d.Token()
					if err != nil {
						return err
					}
					s, ok := key.(string)
					if !ok {
						return errors.New("object key required")
					}
					if seen[s] {
						return fmt.Errorf("duplicate key %q", s)
					}
					seen[s] = true
					if err := value(depth + 1); err != nil {
						return err
					}
				}
			case '[':
				for d.More() {
					if err := value(depth + 1); err != nil {
						return err
					}
				}
			default:
				return errors.New("unexpected delimiter")
			}
			_, err = d.Token()
			return err
		}
		return nil
	}
	if err := value(0); err != nil {
		return err
	}
	if _, err := d.Token(); err != io.EOF {
		return errors.New("trailing JSON data")
	}
	return checkSurrogates(data)
}

// encoding/json replaces unpaired UTF-16 escapes with U+FFFD. Configuration
// strings must never be silently changed, even when the result would compile.
func checkSurrogates(data []byte) error {
	quoted := false
	for i := 0; i < len(data); i++ {
		if data[i] == '"' {
			quoted = !quoted
			continue
		}
		if !quoted || data[i] != '\\' {
			continue
		}
		i++ // Syntax has already been checked by the decoder.
		if data[i] != 'u' {
			continue
		}
		n, _ := strconv.ParseUint(string(data[i+1:i+5]), 16, 16)
		i += 4
		if n >= 0xD800 && n <= 0xDBFF {
			if i+6 >= len(data) || data[i+1] != '\\' || data[i+2] != 'u' {
				return errors.New("unpaired surrogate")
			}
			low, err := strconv.ParseUint(string(data[i+3:i+7]), 16, 16)
			if err != nil || low < 0xDC00 || low > 0xDFFF {
				return errors.New("unpaired surrogate")
			}
			i += 6
		} else if n >= 0xDC00 && n <= 0xDFFF {
			return errors.New("unpaired surrogate")
		}
	}
	return nil
}

func (p *Program) add(raw []byte, depth int, path string) error {
	bad := func(msg string) error { return fmt.Errorf("%w at %s: %s", ErrInvalidConfig, path, msg) }
	if depth > 32 {
		return bad("pipeline exceeds 32 nesting levels or 256 stages")
	}
	if bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
		if depth == 0 {
			return nil
		}
		return bad("null sequence member")
	}
	var obj map[string]json.RawMessage
	if json.Unmarshal(raw, &obj) != nil || obj == nil {
		return bad("expected component object")
	}
	var typ string
	if json.Unmarshal(obj["type"], &typ) != nil || typ == "" {
		return bad("type must be a string")
	}
	delete(obj, "type")
	str := func(key, def string) (string, error) {
		v, ok := obj[key]
		delete(obj, key)
		if !ok {
			return def, nil
		}
		var s string
		if bytes.Equal(v, []byte("null")) || json.Unmarshal(v, &s) != nil {
			return "", bad(key + " must be a string")
		}
		return s, nil
	}
	boolean := func(key string, def bool) (bool, error) {
		v, ok := obj[key]
		delete(obj, key)
		if !ok {
			return def, nil
		}
		var b bool
		if bytes.Equal(v, []byte("null")) || json.Unmarshal(v, &b) != nil {
			return false, bad(key + " must be a boolean")
		}
		return b, nil
	}
	s := stage{kind: splitStage}
	var err error
	switch typ {
	case "Sequence":
		v, ok := obj["pretokenizers"]
		delete(obj, "pretokenizers")
		var children []json.RawMessage
		if !ok || bytes.Equal(v, []byte("null")) || json.Unmarshal(v, &children) != nil {
			return bad("pretokenizers must be an array")
		}
		if len(obj) != 0 {
			return bad("unknown Sequence field")
		}
		for i, c := range children {
			if err := p.add(c, depth+1, fmt.Sprintf("%s.pretokenizers[%d]", path, i)); err != nil {
				return err
			}
		}
		return nil
	case "Split":
		var pattern map[string]string
		v := obj["pattern"]
		delete(obj, "pattern")
		if json.Unmarshal(v, &pattern) != nil || len(pattern) != 1 {
			return bad("pattern must contain exactly one String or Regex")
		}
		if literal, ok := pattern["String"]; ok {
			if literal == "" {
				return bad("empty patterns are unsupported")
			}
			s.literal = literal
		} else if expr, ok := pattern["Regex"]; ok {
			var found bool
			s.regex, found = compile.Lookup(expr)
			if !found {
				// The shared registry miss is internal/regexp.Compile's
				// ErrUnsupportedPattern. Keep registered plans on their original path.
				s.fallback, err = regexp.Compile(expr)
				if err != nil {
					return fmt.Errorf("%w at %s.pattern.Regex: %w", ErrUnsupportedPattern, path, err)
				}
			}
		} else {
			return bad("unknown pattern kind")
		}
		mode, e := str("behavior", "Isolated")
		if e != nil {
			return e
		}
		s.mode, err = parseBehavior(mode)
		if err != nil {
			return bad(err.Error())
		}
		s.invert, err = boolean("invert", false)
	case "ByteLevel":
		s.kind = byteLevelStage
		s.prefix, err = boolean("add_prefix_space", true)
		if err != nil {
			return err
		}
		_, err = boolean("trim_offsets", true)
		if err != nil {
			return err
		}
		s.useRegex, err = boolean("use_regex", true)
	case "Metaspace":
		s.kind = metaspaceStage
		s.literal, err = str("replacement", "▁")
		if err != nil {
			return err
		}
		if utf8.RuneCountInString(s.literal) != 1 {
			return bad("replacement must be one Unicode scalar")
		}
		s.prepend, err = str("prepend_scheme", "always")
		if err != nil {
			return err
		}
		if s.prepend != "always" && s.prepend != "first" && s.prepend != "never" {
			return bad("invalid prepend_scheme")
		}
		s.split, err = boolean("split", true)
	case "Whitespace":
		s.kind = whitespaceStage
	case "WhitespaceSplit":
		s.kind = whitespaceSplitStage
	case "Digits":
		s.kind = digitsStage
		s.individual, err = boolean("individual_digits", false)
	case "Punctuation":
		s.kind = punctuationStage
		mode, e := str("behavior", "Isolated")
		if e != nil {
			return e
		}
		s.mode, err = parseBehavior(mode)
	case "BertPreTokenizer":
		s.kind = bertStage
	case "UnicodeScripts":
		s.kind = scriptsStage
	case "CharDelimiterSplit":
		s.literal, err = str("delimiter", "")
		if err == nil && utf8.RuneCountInString(s.literal) != 1 {
			return bad("delimiter must be one Unicode scalar")
		}
		s.mode = removed
	default:
		return fmt.Errorf("%w at %s: %q", ErrUnsupportedType, path, typ)
	}
	if err != nil {
		return bad(err.Error())
	}
	for key := range obj {
		return bad("unknown field " + key)
	}
	if len(p.stages) >= 256 {
		return bad("pipeline exceeds 256 stages")
	}
	p.stages = append(p.stages, s)
	return nil
}

func parseBehavior(s string) (behavior, error) {
	switch s {
	case "Isolated":
		return isolated, nil
	case "Removed":
		return removed, nil
	case "Contiguous":
		return contiguous, nil
	case "MergedWithPrevious":
		return mergedPrevious, nil
	case "MergedWithNext":
		return mergedNext, nil
	}
	return 0, fmt.Errorf("invalid split behavior %q", s)
}
