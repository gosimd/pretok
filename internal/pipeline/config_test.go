package pipeline

import (
	"encoding/json"
	"testing"

	"github.com/gosimd/pretok/internal/compile"
)

func TestRegexBackendSelection(t *testing.T) {
	for _, tc := range []struct {
		name, pattern string
		registered    bool
	}{
		{"registered_go_syntax", compile.ASCIIDigitPattern, true},
		{"registered_hf_syntax", compile.QwenPattern, true},
		{"unregistered_equivalent", `(?:[0-9])`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config, err := json.Marshal(map[string]any{"type": "Split", "pattern": map[string]string{"Regex": tc.pattern}})
			if err != nil {
				t.Fatal(err)
			}
			p, err := Compile(config)
			if err != nil {
				t.Fatal(err)
			}
			s := p.stages[0]
			if (s.regex != nil) != tc.registered || (s.fallback != nil) == tc.registered {
				t.Fatalf("incorrect backend for %q", tc.pattern)
			}
		})
	}
}
