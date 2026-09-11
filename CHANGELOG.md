# Changelog

## v0.1.0 — Unreleased

First planned release of `github.com/gosimd/pretok`. Requires Go 1.27.0 or newer.
Publication and its date will be recorded when the release tag is published.

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
