# pretok

`pretok` is a Go library for fast pre-tokenization for LLM tokenizers, compiling
common tokenizer regex patterns into specialized scanners, with SIMD acceleration
where it helps.

Module path: `github.com/gosimd/pretok`.

**Status:** scalar implementation for two exact regex strings: P05 (Qwen
letter-only words) and P01 (letters and marks). Both use single Unicode numbers
and share one parameterized scanner template. `Compile`, `Cursor`, and `Split`
are checked against pinned Hugging Face fixtures. The other 23 corpus strings
remain unsupported. No SIMD backend is implemented.

## Scope

`pretok` specializes in regex patterns used for pre-tokenization, including those
found in Hugging Face `tokenizer.json` files. It compiles supported patterns once
and splits text repeatedly without a regex engine in the runtime path.

It is not a general-purpose regex engine. A valid regex outside the supported
subset must produce an explicit unsupported-pattern error. Compilation must never
silently approximate a pattern or fall back to a general regex engine.

The library finds pre-token boundaries. Text normalization, special-token handling,
ByteLevel encoding, BPE merges, vocabulary lookup, token IDs, and post-processing
belong to the surrounding tokenizer. Compiling a regex alone does not establish
compatibility with an entire tokenizer pipeline.

## How it works

The first implementation uses an explicit registry of known regex strings. The
current [corpus](research/pretokenize_regex/all_patterns.json) contains 25 unique
strings; this is the initial research target, not a claim of implemented support.

```text
pre-tokenization regex string
        |
        v
exact lookup in the supported-pattern registry
        |
        v
select a prewritten scanner plan + immutable parameters
        |
        v
scan text using specialized scalar / SIMD code
```

Lookup compares the exact regex string after JSON decoding by the caller. It does
not parse an AST, normalize syntax, trim whitespace, or unescape the regex again.
Unknown strings return an unsupported-pattern error, including malformed strings
and equivalent spellings that have not been registered explicitly.

Several known strings may select the same scanner when matching behavior has been
verified to be equivalent. Supporting 25 strings does not require 25 independent
scanner implementations. Register each entry only after its scalar plan has been
validated against the pinned reference. No general regex engine runs as a fallback.

Shared work such as whitespace scanning is designed into the prewritten plans.
Alternative priority, greediness, Unicode properties, assertions, and consumed
boundaries must still match the reference. Exact lookup removes the need for a
parser; it does not remove the need to establish each pattern's semantics.

A future compiler may add parsing and structural recognition behind the same
public API. That is deferred, not a prerequisite for the first implementation.

## Public API

See [Regexp](docs/Regexp.md) for the API contract, ownership rules, and the
handling of unsupported patterns, including general-purpose regexes.

`Regexp` names the compiled expression: `pretok.Regexp`. Scanner kernels remain
implementation details. The library handles one regex at a time; it does not
provide a full `PreTokenizer` or a `CompileSequence` pipeline.

```go
type Span struct {
    Start int
    End   int
}

func Compile(pattern string) (*Regexp, error)
func (r *Regexp) Cursor(src string) (Cursor, error)
func (c *Cursor) Next() (Span, bool)
func (r *Regexp) Split(src string, dst []Span) ([]Span, error)
```

Usage with the supported expression:

```go
re, err := pretok.Compile(pattern)
if err != nil {
    return err
}

spans, err := re.Split(input, dst[:0])
if err != nil {
    return err
}
for _, span := range spans {
    piece := input[span.Start:span.End]
    _ = piece
}
```

The runtime contract is:

- `Compile` returns an immutable `Regexp`, reusable across concurrent calls.
  The first implementation accepts only registered strings and reports all other
  strings as `ErrUnsupportedPattern`, without separate syntax diagnostics.
- `Cursor` iterates non-empty regex matches only, skipping unmatched gaps.
  It borrows its input and owns mutable iteration state. A zero or exhausted
  cursor returns `Span{}, false` from `Next`.
- `Split` isolates each non-empty match and preserves each maximal unmatched gap.
  It appends spans to caller-owned `dst`; pass `dst[:0]` for a fresh result.
  Empty input adds no spans. The resulting pieces cover the supplied input.
- Spans are half-open byte offsets `[Start, End)` into the exact supplied string,
  on UTF-8 character boundaries. They are not rune indices.
- Runtime entry points validate UTF-8 before scanning. Invalid UTF-8 returns
  `ErrInvalidUTF8`; `Split` leaves `dst` unchanged. Input bytes are never replaced.
- Patterns that can produce zero-length matches are initially unsupported.
- `Regexp` and `Split` retain neither input nor destination storage. A cursor
  borrows input for its lifetime. Execution scratch and CPU dispatch stay internal.

`Split` corresponds to isolated splitting with gaps retained. A caller needing
matches only can consume `Cursor`. Other HF split behaviors are not implied by
regex support. Composition of regex stages belongs to the surrounding tokenizer,
which passes each prior-stage piece separately and preserves its boundaries.

## Compatibility driven by a tokenizer corpus

Support will grow from observed tokenizer patterns rather than an attempt to cover
all regex syntax:

1. Collect pre-tokenizer regex patterns and their surrounding configuration from
   a documented selection of tokenizer files, pinned to model revisions.
2. Preserve exact decoded strings and provenance. Manually group patterns that
   can share scanner code; do not normalize expressions during compilation.
3. Build reference outputs with a pinned Hugging Face tokenizer implementation and
   the relevant split configuration.
4. Implement scalar plans and compare exact pieces and byte offsets with the
   reference, including adversarial inputs and end-of-input cases.
5. Add SIMD paths where profiling and repeatable benchmarks justify them.

A future support table should list verified pattern families, source revisions,
split settings, and known limits. The number of families and their coverage must
come from the collected corpus, not an assumed estimate. Unsupported syntax or
semantics remain an explicit compile-time error.

The first [research survey](research/pretokenize_regex/README.md) collects explicit
`pretokenize_regex` fields from Hugging Face `tokenizer_config.json` and
`pre_tokenizer` regexes from `tokenizer.json`, with
[results and pinned sources](research/pretokenize_regex/TOKENIZER_REPORT.md).
It preserves surrounding split settings. Discovery alone does not establish
support. The initial P05 implementation has
its own [oracle fixtures](testdata/qwen/manifest.json); P01 has [separate fixtures](testdata/patterns/P01/manifest.json) and a [maintained regeneration script](testdata/patterns/README.md). Full tokenizer compatibility
is outside this package.

The [structural and semantic analysis](research/regex_analysis/ANALYSIS.md) groups
the observed alternatives into parameterized scanner operations and records
pinned oracle experiments. Its AST/IR compiler proposal is deferred; the first
implementation selects prewritten plans by exact string lookup.

## Layout

```text
regexp.go              public compilation and scanning facade
internal/compile/      exact-string registry and scanner-plan selection
internal/scan/         scanner plans, scalar constructions, pinned Unicode tables
testdata/              pinned patterns, provenance, and reference outputs
```

The root API stays focused on compilation and splitting. The registry, execution
plans, scratch buffers, and CPU dispatch are implementation details. Parser and IR
packages are deferred until broader recognition is needed.

## Performance and validation

[Benchmark methodology](docs/Benchmark.md) defines the implemented Oniguruma and
native HF comparisons, multilingual corpus, timing boundaries, and report metrics.
See [benchmark/README.md](benchmark/README.md) to build and run the offline suite.
The [P01 Apple M5 scalar report](benchmark/results/P01/2026-09-10-apple-m5-scalar/REPORT.md)
and [original P05 report](benchmark/results/P05/2026-09-09-apple-m5-scalar/REPORT.md)
include all languages and sizes, raw samples, and environment fingerprints.

The [Regexp contract](docs/Regexp.md) lists the exact supported string and plan
components. [Oracle regeneration](research/qwen/README.md) is separate from ordinary
offline tests:

```sh
go test ./...
go test -race ./...
go test -run '^$' -fuzz FuzzQwenPartition -fuzztime 10s .
go test -run '^$' -fuzz FuzzP01Partition -fuzztime 10s .
```

The checked-in fixtures contain 7,738 P05 cases and 26,820 P01 cases. Unicode class membership and the
contraction case folds are generated from every Unicode scalar using the pinned
native oracle, rather than Go's evolving Unicode tables. These checks establish
bounded differential evidence, not proof of all input combinations.

Correctness comes first: reference-versus-scalar tests establish compatibility,
and scalar-versus-SIMD tests establish backend parity. Tests should cover Unicode,
contractions, number boundaries, punctuation, mixed whitespace, CR/LF, and EOF,
as well as explicitly registered equivalent spellings and rejected unknown strings.

Benchmarks separate compilation cost from repeated scanning and report
throughput, latency, and allocations. Both short prompts and long multilingual
inputs matter. Performance claims must name the machine, Go version, backend,
corpus, and comparison baseline, and include regressions.

## License

[MIT](LICENSE).
