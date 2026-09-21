# pretok

`github.com/gosimd/pretok` compiles Hugging Face pre-tokenizer configurations into
reusable Go pipelines. Compile once, then split and transform text through
callbacks, with optional mapping back to original UTF-8 byte offsets.

The public API is [`Pipeline`](docs/Pipeline.md): `CompileJSON`, `Walk`,
`WalkText`, and `WalkTextUnsafe`. Regex compilation and scanning are internal
implementation details. `Split` stages accept the 25 exact strings in the
[supported-pattern registry](docs/SupportedPatterns.md), with specialized
scanners and no general regex fallback.

**Status:** this is an unreleased breaking API change on `pretok-api`.
`Regexp`, `Compile`, and `Cursor` are no longer exported from the root package.
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

To migrate isolated regex splitting, compile
`{"type":"Split","pattern":{"Regex":"..."},"behavior":"Isolated","invert":false}`
and consume final pieces through `Walk` or either text-only method. Build the JSON
with `encoding/json` when the pattern comes from a Go string. `Walk` maps source
offsets through any subsequent transformations.

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
Split/ByteLevel sequence. The other supported components are also benchmark
targets; this API reorganization does not add new cases or claim complete
benchmark coverage. The existing pipeline oracle tests cover all supported types.

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
