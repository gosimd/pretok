# Published benchmark results

Store only selected, completed measurements here. Local verification, smoke,
interrupted, and exploratory runs belong in ignored `benchmark/runs/`.

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

Each full P01/P05 run contains 1,710 samples: ten samples of at least 500 ms per measurement,
using Go 1.27.0, Rust 1.90.0, HF Tokenizers 0.22.2, and onig 6.5.1. Its original
environment and raw samples are preserved without rewriting historical metadata.
P01 is also implemented; the other 23 regex strings remain unsupported.

The [P05 refactor regression check](P05/2026-09-10-apple-m5-scalar-regression/REPORT.md)
is a separate, completed before/after Go microbenchmark with identical harnesses.
It does not replace the full cross-engine P05 report.

Before publishing a local run, check successful correctness gates, `complete:
true`, the intended full matrix and sampling settings, and environment metadata.
Copy the complete three-file result directory into its pattern directory and add
an entry above. Do not publish partial runs or merely rename them as completed.

For future patterns, reuse the same source text pools and request boundaries;
store their oracle spans in the shared request files and their profiles in the
common corpus manifest. Keep results per
pattern, language, size, engine, and backend. Cross-pattern comparisons also need
matching corpus hashes and reference/split settings. The current runner and oracle
generator support P01 and P05. Generation refreshes both together; `-pattern`
selects one profile for measurement.

The recorded September 9–10 full runs predate corpus schema 2. Their original
metadata and hashes are retained: those corpus-manifest hashes identify the old
serialization. The migration compared all 42 pattern/batch combinations and
preserved input bytes, request ranges, reference spans, and statistics exactly.
Only storage/loading changed; these historical reports are not new measurements
of the current runner.
