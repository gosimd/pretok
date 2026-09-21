# Changelog

## Unreleased

- Expand the HF 0.23.2 Pipeline oracle from 4,475 to 5,062 cases, preserving
  existing fixtures and adding component boundaries, long inputs, nested
  sequences and composed offset ranges for all nine core pre-tokenizers.
- Breaking: make Pipeline the public pre-tokenization API. Move `Regexp`,
  `Compile`, `Cursor`, their tests, examples and microbenchmarks to
  `internal/regexp`; move regex measurement infrastructure to
  `internal/regexp/bench`. Existing public Pipeline tests and benchmarks remain
  in the root package. No new test or benchmark cases are added by this move.
  Use a Pipeline Split stage for registered regex patterns. Public `Span` and
  error sentinels remain part of the Pipeline API; `ErrUninitialized` now reports
  an uninitialized pipeline. The v0.1.0 tag retains its historical API.
- Add `CompileJSON`, immutable `Pipeline`, `Walk`, `WalkText`, `WalkTextUnsafe`,
  and callback-scoped `Piece.Text` / `Piece.OriginalRange`. Pipelines support the nine documented
  built-ins plus exact-registry/literal `Split` and `ByteLevel`; see
  [the pipeline contract](docs/Pipeline.md). Text-only walking skips output
  alignment construction while preserving stage context. `WalkText` copies strings
  for safe retention; `WalkTextUnsafe` avoids defensive copies but requires
  `strings.Clone` inside the callback before retaining or asynchronously using text.
- Add HF Tokenizers 0.23.2 pipeline fixtures, built-in Unicode tables, offset
  regressions, concurrent/reentrant execution tests and allocation benchmarks.
  These are separate from the existing regex profile pinned to HF 0.22.2.
- P01 `Split` validates the complete input, then scans 1,024-rune windows decoded
  by `github.com/gosimd/unicode v0.1.2`. Supported `GOEXPERIMENT=simd` builds use
  SIMD UTF-8 kernels; other configurations retain byte scanning.
- Rune scratch is local to each call. Short inputs keep byte scanning;
  internal cursor-copy semantics, byte boundaries and error atomicity remain
  unchanged by the buffering optimization.

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
