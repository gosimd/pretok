package pretok_test

import (
	"strings"
	"testing"

	"github.com/gosimd/pretok"
)

var pipelineSink int

func BenchmarkPipeline(b *testing.B) {
	configs := []struct{ name, config string }{
		{pipelineExampleName, pipelineExample},
		{"Whitespace/default", `{"type":"Whitespace"}`},
		{"ByteLevel/prefix=true,trim=true,regex=true", `{"type":"ByteLevel"}`},
		{"Metaspace/prepend=always,split=true,replacement=U+2581", `{"type":"Metaspace"}`},
		{"Split/GoRegexpFallback", `{"type":"Split","pattern":{"Regex":"\\p{L}+|[0-9]+"},"behavior":"Isolated"}`},
	}
	inputs := []struct{ name, text string }{
		{"short_ascii", "Hello, world! 123"},
		{"short_mixed", "Привет é世界 ١٢٣!"},
		{"long_ascii", strings.Repeat("Hello, world! 123\n", 4096)},
		{"long_mixed", strings.Repeat("Привет é世界 ١٢٣!\n", 2048)},
	}
	for _, c := range configs {
		for _, input := range inputs {
			b.Run(c.name+"/"+input.name, func(b *testing.B) {
				p := compilePipeline(b, c.config)
				// Untimed contract checks on the exact benchmark input. Oracle
				// compatibility is tested separately by TestPipelineHFOracle.
				var pieces []string
				if err := p.Walk(input.text, func(piece pretok.Piece) error {
					span, err := piece.OriginalRange(0, len(piece.Text()))
					if err != nil {
						return err
					}
					if span.Start < 0 || span.End > len(input.text) || span.Start >= span.End {
						b.Fatal("invalid benchmark span", span)
					}
					pieces = append(pieces, strings.Clone(piece.Text()))
					return nil
				}); err != nil {
					b.Fatal(err)
				}
				for _, walkText := range []func(string, func(string) error) error{p.WalkText, p.WalkTextUnsafe} {
					i := 0
					if err := walkText(input.text, func(text string) error {
						if i >= len(pieces) || pieces[i] != text {
							b.Fatal("benchmark WalkText mismatch")
						}
						i++
						return nil
					}); err != nil {
						b.Fatal(err)
					}
					if i != len(pieces) {
						b.Fatal("benchmark piece count mismatch")
					}
				}
				for _, name := range []string{"Walk", "WalkText", "WalkTextUnsafe"} {
					textOnly := name != "Walk"
					walkText := p.WalkText
					if name == "WalkTextUnsafe" {
						walkText = p.WalkTextUnsafe
					}
					b.Run(name, func(b *testing.B) {
						b.ReportAllocs()
						b.SetBytes(int64(len(input.text)))
						consume := func(piece pretok.Piece) error {
							span, err := piece.OriginalRange(0, len(piece.Text()))
							pipelineSink += span.End - span.Start
							return err
						}
						consumeText := func(s string) error { pipelineSink += len(s); return nil }
						// Warm reusable buffers before measuring steady-state execution.
						if textOnly {
							if err := walkText(input.text, consumeText); err != nil {
								b.Fatal(err)
							}
						} else {
							if err := p.Walk(input.text, consume); err != nil {
								b.Fatal(err)
							}
						}
						b.ResetTimer()
						for b.Loop() {
							var err error
							if textOnly {
								err = walkText(input.text, consumeText)
							} else {
								err = p.Walk(input.text, consume)
							}
							if err != nil {
								b.Fatal(err)
							}
						}
					})
				}
			})
		}
	}
}

func BenchmarkCompileJSON(b *testing.B) {
	for _, tc := range []struct{ name, config string }{
		{pipelineExampleName, pipelineExample},
		{"Split/GoRegexpFallback", `{"type":"Split","pattern":{"Regex":"\\p{L}+|[0-9]+"},"behavior":"Isolated"}`},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				p, err := pretok.CompileJSON([]byte(tc.config))
				if err != nil {
					b.Fatal(err)
				}
				_ = p
			}
		})
	}
}
