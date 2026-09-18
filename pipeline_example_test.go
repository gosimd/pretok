package pretok_test

import (
	"fmt"
	"strings"

	"github.com/gosimd/pretok"
)

func ExamplePipeline_Walk() {
	p, err := pretok.CompileJSON([]byte(`{"type":"Sequence","pretokenizers":[
		{"type":"Split","pattern":{"Regex":"[0-9]"},"behavior":"Isolated","invert":false},
		{"type":"ByteLevel","add_prefix_space":false,"trim_offsets":false,"use_regex":false}
	]}`))
	if err != nil {
		panic(err)
	}
	err = p.Walk("ab12é", func(piece pretok.Piece) error {
		span, err := piece.OriginalRange(0, len(piece.Text()))
		if err != nil {
			return err
		}
		fmt.Printf("%q [%d,%d)\n", piece.Text(), span.Start, span.End)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Output:
	// "ab" [0,2)
	// "1" [2,3)
	// "2" [3,4)
	// "Ã©" [4,6)
}

func ExamplePipeline_WalkText() {
	p, err := pretok.CompileJSON([]byte(`{"type":"Metaspace"}`))
	if err != nil {
		panic(err)
	}
	var pieces []string
	err = p.WalkText("Hello world", func(text string) error {
		pieces = append(pieces, text) // Safe to retain without copying.
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, text := range pieces {
		fmt.Println(text)
	}
	// Output:
	// ▁Hello
	// ▁world
}

func ExamplePipeline_WalkTextUnsafe() {
	p, err := pretok.CompileJSON([]byte(`{"type":"Metaspace"}`))
	if err != nil {
		panic(err)
	}
	var pieces []string
	err = p.WalkTextUnsafe("Hello world", func(text string) error {
		// Synchronous consumers can use text directly. Retention requires a copy.
		pieces = append(pieces, strings.Clone(text))
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, text := range pieces {
		fmt.Println(text)
	}
	// Output:
	// ▁Hello
	// ▁world
}
