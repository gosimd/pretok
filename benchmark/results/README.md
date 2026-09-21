# Published benchmark results

Store only selected, completed measurements here. Local verification, smoke,
interrupted, and exploratory runs belong in ignored `benchmark/runs/`.

The [integrated P01 rune-splitting comparison](P01/2026-09-16-apple-m5-rune-split/REPORT.md)
uses the published gosimd/unicode v0.1.2 dependency and the regex API that was
public at the measured revisions (now [internal/regexp](../../internal/regexp)), including
per-call stack scratch. It compares ordinary and SIMD builds against the same
baseline with 600 interleaved samples, and includes regressions. Shared reproduction tools live in
[benchmark/compare](../compare). At those revisions, Cursor keeps byte scanning;
its long P01 inputs use SIMD
validation on eligible builds.

Organize published results by exact pattern ID, then date, machine, and backend:

```text
results/<pattern-id>/<date>-<machine>-<backend>/
    REPORT.md
    environment.json
    samples.jsonl
```

Pattern IDs come from the fixed corpus registry; do not renumber them when adding
support. The exact decoded regex SHA-256 in `environment.json` is authoritative.
A model name is a descriptive alias, not a result key: models can share a regex,
and different revisions of one model can use different regexes. Even patterns
sharing one scanner plan need separately validated measurements.

| Pattern | Implementation | Machine | Published run |
|---|---|---|---|
| P01 (Qwen words with marks, single Unicode numbers) | Scalar | Apple M5 | [2026-09-10](P01/2026-09-10-apple-m5-scalar/REPORT.md) |
| P05 (Qwen letter-only words, single Unicode numbers) | Scalar | Apple M5 | [2026-09-09](P05/2026-09-09-apple-m5-scalar/REPORT.md) |
| P09 (P05 with literal CR/LF) | Scalar, P05 plan | Apple M5 | [2026-09-10](P09/2026-09-10-apple-m5-scalar/REPORT.md) |
| P13 (possessive spelling, P05-equivalent boundaries) | Scalar, P05 plan | Apple M5 | [2026-09-10](P13/2026-09-10-apple-m5-scalar/REPORT.md) |
| P16 (one ASCII digit) | Scalar | Apple M5 | [2026-09-10](P16/2026-09-10-apple-m5-scalar/REPORT.md) |
| P07 (1–3 Unicode N scalars) | Scalar | Apple M5 | [2026-09-10](P07/2026-09-10-apple-m5-scalar/REPORT.md) |
| P24 (Qwen with N{1,3}) | Scalar | Apple M5 | [2026-09-10](P24/2026-09-10-apple-m5-scalar/REPORT.md) |
| P20 (Qwen with N+) | Scalar | Apple M5 | [2026-09-10](P20/2026-09-10-apple-m5-scalar/REPORT.md) |
| P19 (optional word contractions, single N) | Scalar, shared case-word plan | Apple M5 | [2026-09-10](P19/2026-09-10-apple-m5-scalar/REPORT.md) |
| P06 (optional word contractions, N{1,3}) | Scalar, shared case-word plan | Apple M5 | [2026-09-10](P06/2026-09-10-apple-m5-scalar/REPORT.md) |

| P11 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P11/2026-09-10-apple-m5-scalar/REPORT.md) |
| P21 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P21/2026-09-10-apple-m5-scalar/REPORT.md) |
| P17 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P17/2026-09-10-apple-m5-scalar/REPORT.md) |
| P18 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P18/2026-09-10-apple-m5-scalar/REPORT.md) |
| P08 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P08/2026-09-10-apple-m5-scalar/REPORT.md) |
| P23 | Scalar, shared exact-set search | Apple M5 | [2026-09-10](P23/2026-09-10-apple-m5-scalar/REPORT.md) |

| P12 (ordered group runs) | Scalar, shared grouped-set search | Apple M5 | [2026-09-11](P12/2026-09-11-apple-m5-scalar/REPORT.md) |
| P15 (ordered group runs) | Scalar, shared grouped-set search | Apple M5 | [2026-09-11](P15/2026-09-11-apple-m5-scalar/REPORT.md) |

| P22 (ASCII-prefixed words and P\|S runs) | Scalar, shared classes and whitespace | Apple M5 | [2026-09-11](P22/2026-09-11-apple-m5-scalar/REPORT.md) |
| P10 (anchored Nd remainder) | Scalar, decimal run metadata | Apple M5 | [2026-09-11](P10/2026-09-11-apple-m5-scalar/REPORT.md) |
| P14 (consecutive Nd triples) | Scalar, shared bounded decimal scanner | Apple M5 | [2026-09-11](P14/2026-09-11-apple-m5-scalar/REPORT.md) |
| P25 (right-aligned Nd groups) | Scalar, cursor-owned run cache | Apple M5 | [2026-09-11](P25/2026-09-11-apple-m5-scalar/REPORT.md) |

Each full P01/P05/P09/P13/P24/P20/P19/P06 run contains 1,710 samples: ten samples of at least 500 ms per measurement,
using Go 1.27.0, Rust 1.90.0, HF Tokenizers 0.22.2, and onig 6.5.1. Its original
environment and raw samples are preserved without rewriting historical metadata.
P16/P07/P08/P11/P12/P15/P17/P18/P21/P22/P23 runs contain 1,500 samples each: the match-at diagnostic is excluded
because these patterns can leave gaps. Their Cursor and Split outputs are
verified independently before timing.

All 25 corpus strings are implemented. P02/P03/P04 have no published full runs.

The [P05 refactor regression check](P05/2026-09-10-apple-m5-scalar-regression/REPORT.md)
is a separate, completed before/after Go microbenchmark with identical harnesses.
It does not replace the full cross-engine P05 report.

Before publishing a local run, check successful correctness gates, `complete:
true`, the intended full matrix and sampling settings, and environment metadata.
Copy the complete three-file result directory into its pattern directory and add
an entry above. Do not publish partial runs or merely rename them as completed.

For future patterns, reuse the same source text pools and request boundaries;
store only inputs and their profiles in the common corpus manifest. Generate
oracle spans in memory during the correctness stage of each run. Keep results per
pattern, language, size, engine, and backend. Cross-pattern comparisons also need
matching corpus hashes and reference/split settings. The current runner and oracle
generator support P01–P09, P11–P13, and P15–P24. Corpus preparation refreshes shared inputs; `-pattern`
selects one profile for live reference generation and measurement.

The original P01/P05 September 9–10 full runs predate corpus schema 2. Their original
metadata and hashes are retained: those corpus-manifest hashes identify the old
serialization. The migration compared all 42 pattern/batch combinations and
preserved input bytes, request ranges, reference spans, and statistics exactly.
Only storage/loading changed; these historical reports are not new measurements
of the current runner.

Schema 3 removes persisted corpus outputs entirely. Expected matches and splits
are generated by HF before each run; old result metadata still describes its
original runner and corpus format. Input bytes and request ranges are preserved.

The numeric-family P16/P07/P24/P20 runs share identical source, binary, and corpus
fingerprints. Their common synthetic pool adds numeric-boundary scenarios, so
its corpus hash differs from earlier P09/P13 runs. Historical results are retained
with their original inputs and metadata.

The P19/P06 runs share identical source, binary and corpus fingerprints. Their
common corpus adds case-word suffix scenarios to the synthetic pool and contains
68,400 requests across all batches. Historical runs retain their original corpus hashes
and metadata.

The P11/P21/P17/P18/P08/P23 runs share identical source, Go/native binaries and
corpus fingerprints. Each contains 1,500 samples; together they contain 9,000.
The shared corpus contains 94,775 requests over 21 batches, including literal-set
boundaries, prefix overlap and nested-class diagnostics. All 19 supported patterns
passed exact live HF/Oniguruma comparisons on these inputs. Historical reports
retain their original corpus and source fingerprints.

The P12/P15 runs share identical source, binaries, corpus and settings. Each contains 1,500 samples (3,000 total), ten repetitions of at least 500 ms across all 21 batches plus compile measurements. Their common corpus contains 118,090 requests, including ordered group transitions and overlaps. Both passed exact live HF/Oniguruma validation before timing. Earlier results retain their original source and corpus hashes.

The P22 run contains 1,500 samples: ten repetitions of at least 500 ms across all
21 batches, plus separate compilation measurements. Its corpus contains 141,415
requests, including ASCII prefix priority, positive P|S boundaries and numeric
gaps. All 22 supported patterns passed exact live HF/Oniguruma comparison on
these inputs; P22 was checked again immediately before timing. Historical runs
retain their original corpus, implementation and binary fingerprints.

The P10 run contains 1,570 samples across the expanded 22 shared batches and
separate compilation measurements. All 162,794 requests were verified against
live HF/Oniguruma references. Its shared decimal workload measures successful
matches separately from natural-text rejection. The report retains the
Russian/Arabic regressions against Oniguruma and all source/corpus fingerprints.

The completed P14 run contains 1,570 samples over the same 22 shared batches,
with 32 additional continuation-failure cases in the synthetic pool. All 162,826
requests passed live reference comparison for all 24 supported patterns. Its
report separates successful decimal grouping from natural-text rejection and
preserves raw samples, variation, allocations and source/corpus fingerprints.

The P25 run contains 1,570 samples over 22 shared batches, with 240 additional
right-alignment/cache-transition subjects. All 25 patterns passed the expanded
163,066-request gate. The report includes the scoped before/after cost of the
additional cursor metadata; historical reports keep their original fingerprints.
