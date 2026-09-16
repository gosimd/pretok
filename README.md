# pretok

`github.com/gosimd/pretok` compiles common LLM pre-tokenization regex strings into
specialized Go scanners. Compile once, then find byte boundaries without a regex
engine in the scanning path.

**Status:** all 25 exact strings in the [supported-pattern registry](docs/SupportedPatterns.md)
are implemented and checked against pinned Hugging Face references. On supported
SIMD builds, P01 `Split` uses a bounded rune buffer for UTF-8 conversion;
token classification and boundary decisions remain scalar. The
[v0.1.0 public contract](docs/Regexp.md#v010-release-contract) is unchanged.

## Installation

Requires **Go 1.27.0 or newer** and `github.com/gosimd/unicode v0.1.2` (resolved by
Go modules). Using the library requires neither Rust, Python nor a native regex
library. The dependency is pinned in `go.mod` and verified through `go.sum`.

Install the v0.1.0 version with:

```sh
go get github.com/gosimd/pretok@v0.1.0
```

The P01 buffering optimization described here is in the current checkout; it is
not part of the v0.1.0 tag.

The examples and tests for the current checkout can be run with:

```sh
GOWORK=off go test -mod=readonly ./...
```

Build your application with `GOEXPERIMENT=simd` to enable gosimd/unicode's supported
ARM64/AMD64 SIMD paths. Without this setting, or on unsupported configurations,
P01 keeps its portable byte scanner with the same boundaries.

On eligible builds, P01 `Split` validates the entire input, then decodes into a local 1,024-rune
(4 KiB) buffer. Tokens can cross buffers. Inputs shorter than 256 bytes use the
existing byte scanner. Public `Cursor` keeps its independent, copyable byte-scanning
state; it does not borrow mutable rune scratch. Other patterns keep their existing
scanners.

## Quick start

This is the exact P07 pattern: groups of one to three Unicode number scalars.
The example is also checked by [Go example tests](example_test.go).

```go
package main

import (
    "fmt"

    "github.com/gosimd/pretok"
)

func main() {
    re, err := pretok.Compile(`\p{N}{1,3}`)
    if err != nil {
        panic(err)
    }

    input := "x1234y"
    spans, err := re.Split(input, nil)
    if err != nil {
        panic(err)
    }
    for _, span := range spans {
        fmt.Printf("%q [%d,%d)\n", input[span.Start:span.End], span.Start, span.End)
    }
}
// Output:
// "x" [0,1)
// "123" [1,4)
// "4" [4,5)
// "y" [5,6)
```

For repeated calls, keep the compiled expression and reuse caller-owned capacity:

```go
spans, err = re.Split(nextInput, spans[:0])
```

Check `err` before consuming the result. `Split` appends to its destination;
passing `spans` instead of `spans[:0]` retains the existing entries as a prefix.

## API and behavior

```go
func Compile(pattern string) (*Regexp, error)
func (r *Regexp) Cursor(src string) (Cursor, error)
func (c *Cursor) Next() (Span, bool)
func (r *Regexp) Split(src string, dst []Span) ([]Span, error)
```

- `Compile` accepts the exact string after JSON decoding by the caller. It does
  not trim, unescape again, parse or normalize. Any unregistered string, including
  malformed syntax and equivalent spellings, returns `ErrUnsupportedPattern`.
- `Cursor` returns non-empty matches and skips unmatched gaps. For `x1234y`
  above it returns only `123` and `4`.
- `Split` isolates each match and each maximal unmatched gap. Its appended pieces
  cover the input; empty input appends nothing.
- `Span{Start, End}` uses half-open UTF-8 byte offsets into the supplied string.
  The bytes are preserved. All scanning entry points reject malformed UTF-8 with
  `ErrInvalidUTF8`; `Split` leaves the destination unchanged on error.
- A compiled `Regexp` is immutable and supports concurrent scans with independent
  cursors and buffers. A cursor borrows its input and owns its iteration state.
  `Regexp` and `Split` retain neither input nor destination buffers.
- Scanning with a nil or zero-value `Regexp` returns `ErrUninitialized`. Use
  `errors.Is` for error handling; see the [executable examples](example_test.go).

The [Regexp contract](docs/Regexp.md) specifies error precedence, cursor copies,
ownership and v0.1.x compatibility. The [pattern registry](docs/SupportedPatterns.md)
lists exact expressions, stable P01–P25 IDs and pinned tokenizer examples.

## Integration boundary

The library handles regex boundaries only. It does not implement BPE, vocabulary
lookup, token IDs, special tokens, normalization, ByteLevel mapping or a complete
Hugging Face tokenizer pipeline. There is no generic regex fallback.

`Split` implements isolated splitting with unmatched gaps retained. Other HF
split behaviors and `invert` settings are not inferred from a pattern. When
composing stages, the caller passes each prior-stage piece separately and adds
its base offset where needed. Use `Cursor` only when dropping gaps is intended.
Regex support alone does not establish complete tokenizer compatibility.

## Validation and performance

After the pinned Go dependency has been downloaded, ordinary tests run offline:

```sh
GOWORK=off go test -mod=readonly ./...
GOWORK=off go test -mod=readonly -race ./...
GOWORK=off go vet -mod=readonly ./...
GOWORK=off GOEXPERIMENT=simd go test -mod=readonly ./...
GOWORK=off GOEXPERIMENT=simd go test -mod=readonly -race ./...
```

[Reproducibility](docs/Reproducibility.md) describes the pinned Unicode tables,
fixture generators and live HF correctness checks. Benchmark reference spans
are generated in memory; they are not stored in the shared input corpus.

See the [benchmark instructions](benchmark/README.md), [methodology](docs/Benchmark.md)
and [completed reports](benchmark/results/README.md). Reports distinguish matches
from isolated pieces, fresh from reused destinations, and natural text from
synthetic cases. Their machine, toolchain, source and corpus fingerprints scope
the measurements; historical results are not measurements of every later commit.

The [P01 integration report](benchmark/results/P01/2026-09-16-apple-m5-rune-split/REPORT.md)
measures the buffered public Split against its previous implementation, including
ordinary-build controls, short inputs, fresh/reused output and known regressions.

## Repository guide

| Path | Purpose |
| --- | --- |
| [regexp.go](regexp.go) | Public API |
| [internal/compile](internal/compile) | Exact-string registry and plan selection |
| [internal/scan](internal/scan) | Scanner plans, P01 rune buffering and pinned Unicode tables |
| [testdata](testdata) | Pattern provenance, generators and regression fixtures |
| [benchmark](benchmark) | Shared corpus, native reference runner and reports |
| [CHANGELOG.md](CHANGELOG.md) | Release notes |

## License

Library code is [MIT licensed](LICENSE). Benchmark text has separate
[licenses and attribution](benchmark/corpus/LICENSE.md); retain its source
manifests and notices when redistributing the corpus.
