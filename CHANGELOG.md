# Changelog

## v0.2.1

- Add standard-library `regexp` fallback to Pipeline `Split` stages when an
  expression is absent from the specialized registry. Compile once and reuse
  the result. Registered expressions retain their existing scanner and HF
  semantics; the internal regex API remains registry-only.
- Accept arbitrary expressions supported by Go `regexp`, including unregistered
  equivalent spellings. Fallback uses Go syntax and matching rules, with no
  general HF/Oniguruma compatibility guarantee. Unsupported Go syntax, including
  lookaround and backreferences, still fails unless the exact expression has a
  registered scanner. Compilation errors wrap both `ErrUnsupportedPattern` and
  the underlying `*regexp/syntax.Error`, and include the configuration path.
- Apply all five Split behaviors and `invert` to matches on the complete
  previous-stage piece. Empty matches create boundaries but never empty output
  pieces; Go's adjacent-empty-match suppression applies. Preserve byte alignment,
  transformed sequences and callback ownership contracts. Fallback materializes
  match indices per input piece and makes no specialized performance claim.
- Add fallback regression tests, differential fuzzing against Go `regexp`,
  concurrent reuse checks and execution/allocation benchmarks for all three
  walking methods. See [fallback semantics](docs/Pipeline.md#standard-library-regex-fallback).

## v0.2.0

Requires Go 1.27.0 or newer and `github.com/gosimd/unicode v0.1.2`.

### Breaking API change and migration

- Replace the root `Regexp`, `Compile`, `Cursor`, and `Regexp.Split` API with
  `CompileJSON` and `Pipeline`. The module path remains
  `github.com/gosimd/pretok`; the historical API remains available at v0.1.0.
  Regex implementation, tests, examples and microbenchmarks now live in
  `internal/regexp`, which is not importable by external consumers. Regex
  measurement infrastructure lives in `internal/regexp/bench`.
- Migrate `Regexp.Split` to a Pipeline `Split` stage with `behavior: "Isolated"`
  and `invert: false`. Migrate match-only cursor iteration with
  `behavior: "Removed"` and `invert: true`. Consume results through callbacks;
  use `Walk` and `Piece.OriginalRange` to collect source spans. See the
  [migration guide and complete example](README.md#migrating-from-v010).
- Keep public `Span`, `ErrUnsupportedPattern`, `ErrInvalidUTF8`, and
  `ErrUninitialized`. The latter now describes a nil or zero-value Pipeline.
  Use `errors.Is` for sentinel checks. Add `ErrInvalidConfig`,
  `ErrUnsupportedPreTokenizer`, `ErrNilCallback`, and `ErrInvalidRange`.

### Pipeline features and contracts

- Compile the JSON value of `pre_tokenizer` into an immutable, reusable Pipeline.
  Support `BertPreTokenizer`, `CharDelimiterSplit`, `Digits`, `Metaspace`,
  `Punctuation`, nested `Sequence`, `UnicodeScripts`, `Whitespace`,
  `WhitespaceSplit`, literal/exact-registry `Split`, and `ByteLevel`.
  The [25-string regex registry](docs/SupportedPatterns.md) is unchanged;
  there is no general regex fallback. See the [configuration contract](docs/Pipeline.md)
  for supported settings, defaults and limits.
- `Walk` emits final transformed pieces with original UTF-8 byte alignment.
  `Piece`, its text and alignment are callback-scoped: copy text and resolve
  offsets before returning. Original ranges can overlap or have gaps.
- `WalkText` defensively copies emitted strings for safe retention and
  asynchronous use. `WalkTextUnsafe` avoids those copies, but later execution
  may overwrite the backing memory. Clone with `strings.Clone` inside the
  callback before retaining or asynchronously using its text. Both text-only
  methods skip output alignment construction; neither promises zero allocations.
- Support concurrent and reentrant execution with private execution state.
  Validate the complete input before callbacks, reject malformed UTF-8, and
  emit nothing for empty input. Callback errors stop execution and are returned
  unchanged; earlier callback side effects are not rolled back.
- Scope support to pre-tokenization. Normalization, special-token handling, BPE,
  vocabulary lookup, token IDs and Encoding postprocessing remain outside the
  library. ByteLevel `trim_offsets` does not perform Encoding postprocessing.

### Compatibility evidence and performance scope

- Add 5,062 pinned HF Tokenizers 0.23.2 Pipeline oracle cases and built-in Unicode
  tables. Cases cover all supported types, including long runs, component
  boundaries, nested sequences and composed offsets. The existing regex profile
  remains pinned to HF 0.22.2.
- Add bounded fuzzing of valid components, settings and nested compositions,
  checking text/offset invariants and aligned/text-only parity. Keep shared UTF-8,
  cancellation, ownership and concurrent/reentrant contracts in focused tests.
  Local parity checks complement independent HF fixtures.
- Add Pipeline execution/allocation benchmarks for Whitespace, Metaspace,
  ByteLevel and a Split/ByteLevel sequence, plus a compilation benchmark.
  Coverage is recorded in the [scenario matrix](docs/Pipeline.md#scenario-coverage);
  not every supported component has a benchmark yet. Historical regex benchmark
  reports do not establish Pipeline performance.
- Optimize the internal regex P01 `Split` path with 1,024-rune windows decoded
  by `github.com/gosimd/unicode v0.1.2` after complete input validation.
  Supported `GOEXPERIMENT=simd` builds use SIMD UTF-8 kernels; other configurations
  and short inputs retain byte scanning. Scratch is local to each call; byte
  boundaries, cursor-copy semantics and error atomicity are preserved. This
  buffering optimization does not apply to Pipeline execution.

## v0.1.0

First release of `github.com/gosimd/pretok`. Requires Go 1.27.0 or newer.

### Included

- Scalar scanners for all 25 exact strings in the [pattern registry](docs/SupportedPatterns.md),
  selected through exact-string compilation without a runtime regex engine.
- `Regexp`, `Compile`, `Cursor`, `Span`, and `Split`: match iteration or isolated
  splitting with gaps retained, UTF-8 byte offsets, caller-owned buffer reuse,
  and immutable compiled plans.
- Sentinel errors for unsupported patterns, malformed UTF-8 input and
  uninitialized expressions. See the [release contract](docs/Regexp.md#v010-release-contract).
- Pinned Unicode/reference profiles, offline tests, executable examples and
  reproducible generators. Live HF/Oniguruma comparison uses the shared corpus
  and creates expected spans in memory.
- [Completed scalar benchmark reports](benchmark/results/README.md), with
  source and environment fingerprints and separate natural/synthetic workloads.

### Scope and measured limits

Only registered exact strings are supported. Equivalent unregistered spellings
and malformed expressions both return `ErrUnsupportedPattern`. SIMD, general
regex parsing, normalization, ByteLevel mapping, BPE and full tokenizer pipelines
are outside this release.

Performance depends on pattern and workload. In the [P25 report](benchmark/results/P25/2026-09-11-apple-m5-scalar/REPORT.md),
adding cursor state raised short P14 rejection from 10.22 to 11.78 ns on the
measured Apple M5 setup. Historical benchmark reports describe their recorded
revisions, not a new measurement of the release candidate.
