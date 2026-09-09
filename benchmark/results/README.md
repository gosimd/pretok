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
| P05 (Qwen letter-only words, single Unicode numbers) | Scalar | Apple M5 | [2026-09-09](P05/2026-09-09-apple-m5-scalar/REPORT.md) |

The P05 run contains 1,710 samples: ten samples of at least 500 ms per measurement,
using Go 1.27.0, Rust 1.90.0, HF Tokenizers 0.22.2, and onig 6.5.1. Its original
environment and raw samples are preserved without rewriting historical metadata.
The other 24 regex strings remain unsupported and have no pretok measurements.

Before publishing a local run, check successful correctness gates, `complete:
true`, the intended full matrix and sampling settings, and environment metadata.
Copy the complete three-file result directory into its pattern directory and add
an entry above. Do not publish partial runs or merely rename them as completed.

For future patterns, reuse the same source text pools and request boundaries;
generate separate pattern-specific oracle spans and manifests. Keep results per
pattern, language, size, engine, and backend. Cross-pattern comparisons also need
matching corpus hashes and reference/split settings. The current runner and oracle
generator still target P05; this directory convention does not add pattern support.
