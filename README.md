# pretok

`github.com/gosimd/pretok` compiles Hugging Face pre-tokenizer configurations into
reusable Go pipelines. Compile once, then split and transform text through
callbacks, with optional mapping back to original UTF-8 byte offsets.

The public API is [`Pipeline`](docs/Pipeline.md): `CompileJSON`, `Walk`,
`WalkText`, and `WalkTextUnsafe`. Regex compilation and scanning are internal
implementation details. `Split` stages accept the 25 exact strings in the
[supported-pattern registry](docs/SupportedPatterns.md), with specialized
scanners and no general regex fallback.

**v0.2.0:** this release replaces
the public regex API with `Pipeline`. `Regexp`, `Compile`, and `Cursor` are no
longer exported from the root package. See the [migration guide](#migrating-from-v010)
and [release notes](CHANGELOG.md).
The published v0.1.0 tag retains its [historical regex API](docs/Regexp.md#v010-release-contract);
it does not provide the Pipeline API described here.

## Requirements

Requires **Go 1.27.0 or newer** and `github.com/gosimd/unicode v0.1.2` (resolved by
Go modules). Using the library requires neither Rust, Python nor a native regex
library. The dependency is pinned in `go.mod` and verified through `go.sum`.

Run the examples and tests from this checkout with:

```sh
GOWORK=off go test -mod=readonly ./...
```

## Quick start

Pass the JSON value of `pre_tokenizer`, not the entire `tokenizer.json`.
This example is also checked by [Go example tests](pipeline_example_test.go).

```go
package main

import (
    "fmt"

    "github.com/gosimd/pretok"
)

func main() {
    p, err := pretok.CompileJSON([]byte(`{"type":"Metaspace"}`))
    if err != nil {
        panic(err)
    }
    var pieces []string
    err = p.WalkText("Hello world", func(text string) error {
        pieces = append(pieces, text)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, text := range pieces {
        fmt.Println(text)
    }
}
// Output:
// ▁Hello
// ▁world
```

## API and behavior

```go
func CompileJSON(config []byte) (*Pipeline, error)
func (p *Pipeline) Walk(text string, emit func(Piece) error) error
func (p *Pipeline) WalkText(text string, emit func(string) error) error
func (p *Pipeline) WalkTextUnsafe(text string, emit func(string) error) error
func (p Piece) Text() string
func (p Piece) OriginalRange(start, end int) (Span, error)
```

- Supported components: `BertPreTokenizer`, `CharDelimiterSplit`, `Digits`,
  `Metaspace`, `Punctuation`, `Sequence`, `UnicodeScripts`, `Whitespace`,
  `WhitespaceSplit`, `Split`, and `ByteLevel`. See the [configuration contract](docs/Pipeline.md)
  for supported settings, defaults and compatibility limits.
- `Walk` supplies transformed text and original byte-offset mapping. A `Piece`,
  its text and alignment are valid only during the callback.
- `WalkText` copies each emitted string; callers may retain or process it
  asynchronously after the callback returns.
- `WalkTextUnsafe` avoids defensive copies. Strings are valid only during the
  callback; subsequent execution may overwrite their memory. Use `strings.Clone`
  inside the callback before retaining or asynchronously processing them.
- Compiled pipelines are immutable and support concurrent and reentrant calls.
  The complete input is checked for valid UTF-8 before any callback. Empty input
  emits nothing; callback errors stop execution and are returned unchanged.
- Regex strings in a `Split` configuration are matched exactly after JSON
  decoding, without trimming, additional unescaping, normalization or fallback.

## Migrating from v0.1.0

The module path remains `github.com/gosimd/pretok`, but v0.2.0 is a **breaking
API change**. Replace calls to `pretok.Compile`, `Regexp.Split`, and
`Regexp.Cursor` with `CompileJSON` and Pipeline walking. There is no public regex
compatibility wrapper; Go's `internal` boundary prevents consumers from importing
`internal/regexp`.

| v0.1.0 operation | v0.2.0 replacement |
| --- | --- |
| `pretok.Compile(pattern)` | `CompileJSON` with a `Split` stage whose `pattern` is `{"Regex": pattern}` |
| `re.Split(input, dst)` — matches and unmatched gaps | `Split` with `behavior: "Isolated", invert: false`; append `OriginalRange` results in `Walk` |
| `re.Cursor(input)` / `cursor.Next()` — matches only | `Split` with `behavior: "Removed", invert: true`; consume callbacks instead of pulling a cursor |
| Source byte spans | `Piece.OriginalRange(0, len(piece.Text()))`, resolved inside the callback |
| Retained output strings | `WalkText`; emitted strings remain valid after the callback |
| Synchronous text consumption | `WalkTextUnsafe`; clone inside the callback if retaining text |

For example, the former ``pretok.Compile(`\p{N}{1,3}`)`` followed by
`re.Split(input, dst)` becomes:

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/gosimd/pretok"
)

func main() {
    // Marshal the original Go string: do not manually escape or normalize it.
    config, err := json.Marshal(map[string]any{
        "type": "Split",
        "pattern": map[string]string{"Regex": `\p{N}{1,3}`},
        "behavior": "Isolated",
        "invert": false,
    })
    if err != nil {
        panic(err)
    }
    p, err := pretok.CompileJSON(config) // Compile once, reuse for each input.
    if err != nil {
        panic(err)
    }
    input := "x1234y"
    spans := make([]pretok.Span, 0, 8)
    err = p.Walk(input, func(piece pretok.Piece) error {
        span, err := piece.OriginalRange(0, len(piece.Text()))
        if err != nil {
            return err
        }
        spans = append(spans, span)
        return nil
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(spans)
}
// Output:
// [{0 1} {1 4} {4 5} {5 6}]
```

The caller owns the output slice: keep its length to append, or use `spans[:0]`
before the next input to reuse capacity. Pipeline itself does not materialize an
output list. It validates the complete input before any callback, so invalid
UTF-8 cannot append partial results. A callback error stops walking and is returned
unchanged; earlier callback side effects are not rolled back.

A standalone `Split` preserves source bytes. With transformations such as
Metaspace or ByteLevel, `Piece.Text()` contains transformed text and
`OriginalRange` maps its byte ranges back to the supplied input. Source ranges
can overlap or have gaps. Resolve ranges and copy borrowed text before the
callback returns; storing a `Piece` does not extend its lifetime. Both text-only
methods skip output alignment construction. Avoiding defensive copies with
`WalkTextUnsafe` does not guarantee allocation-free execution.

The 25 accepted regex strings and `ErrUnsupportedPattern` contract are unchanged.
`Span`, `ErrInvalidUTF8`, and `ErrUninitialized` remain public;
`ErrUninitialized` now describes a nil or zero-value Pipeline. Use `errors.Is`
to check sentinels, including errors wrapped by `CompileJSON`, rather than matching
error messages. Configuration and callback errors are documented in the
[Pipeline contract](docs/Pipeline.md).

## Integration boundary

The library implements HF-compatible pre-tokenization for the supported
configurations and pinned reference profiles. Compatibility claims are scoped
to those settings and fixtures. It does not implement BPE, vocabulary lookup,
token IDs, special-token matching, Unicode normalization, Encoding postprocessing,
or a complete Hugging Face tokenizer.

## Validation and performance

After the pinned Go dependency has been downloaded, ordinary tests run offline:

```sh
GOWORK=off go test -mod=readonly ./...
GOWORK=off go test -mod=readonly -race ./...
GOWORK=off go vet -mod=readonly ./...
GOWORK=off GOEXPERIMENT=simd go test -mod=readonly ./...
```

Public Pipeline tests, examples and benchmarks live in the root package.
Internal regex tests and microbenchmarks live alongside their API in
`internal/regexp`; scanner-specific tests remain in `internal/scan`.

```sh
# Public pre-tokenization API, including transformations and sequences.
GOWORK=off go test -mod=readonly -run '^$' -bench 'Benchmark(Pipeline|CompileJSON)' -benchmem .
# Internal regex compilation and boundary scanning.
GOWORK=off go test -mod=readonly -run '^$' -bench . -benchmem ./internal/regexp
```

Existing Pipeline benchmarks cover Whitespace, Metaspace, ByteLevel and a
Split/ByteLevel sequence; benchmark coverage does not yet include every supported
component. All supported types have pinned HF oracle cases. The
[scenario coverage matrix](docs/Pipeline.md#scenario-coverage) maps component
settings to tests and benchmarks; [fuzzing and shared contracts](docs/Pipeline.md#fuzzing-and-shared-contracts)
describe the additional checks.

[Reproducibility](docs/Reproducibility.md) describes pinned Unicode tables,
fixture generators and live HF checks. The [regex benchmark instructions](benchmark/README.md),
[methodology](docs/Benchmark.md) and [completed reports](benchmark/results/README.md)
describe the internal regex subsystem. Historical results retain their measured
source revisions and do not establish current Pipeline performance.

## Repository guide

| Path | Purpose |
| --- | --- |
| [pretok.go](pretok.go) | Public Pipeline API, pieces, source spans and errors |
| [pipeline_test.go](pipeline_test.go), [pipeline_bench_test.go](pipeline_bench_test.go) | Public API compatibility tests and benchmarks |
| [internal/pipeline](internal/pipeline) | Stage execution, private buffers, alignment and HF built-in classes |
| [internal/regexp](internal/regexp) | Internal regex API with its tests, examples and microbenchmarks |
| [internal/regexp/bench](internal/regexp/bench) | Regex corpus validation, measurements and reports |
| [internal/compile](internal/compile) | Exact-string registry and plan selection shared by regex and Pipeline |
| [internal/scan](internal/scan) | Scanner plans, P01 rune buffering and pinned Unicode tables |
| [testdata](testdata) | Pipeline and regex provenance, generators and regression fixtures |
| [benchmark](benchmark) | Shared corpus, native regex reference runner and historical reports |
| [CHANGELOG.md](CHANGELOG.md) | Release notes |

## License

Library code is [MIT licensed](LICENSE). Benchmark text has separate
[licenses and attribution](benchmark/corpus/LICENSE.md); retain its source
manifests and notices when redistributing the corpus.
