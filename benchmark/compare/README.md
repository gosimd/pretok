# Before/after comparison tools

This runner currently implements the P01 public-API comparison: two Git revisions,
each built with and without `GOEXPERIMENT=simd`. It reuses `internal/bench` for
corpus loading, reference verification, operations, timing and allocation counts.
The measurement server in `main.go` is copied into each snapshot so both revisions
execute the same harness. It runs only when explicitly invoked, not during tests.

`reproduce.py --result <published-directory> --work <new-directory>` reads
`baseline_commit` and `implementation_commit` from the result's `environment.json`.
It reconstructs both revisions, runs live HF/Oniguruma correctness gates, then
uses `measure.py` for 600 sequential interleaved samples (five repetitions of at
least 150 ms per cell). The protocol is currently fixed to P01 and the five shared
natural-language corpora. Source-only preparation is available with `--prepare-only`.
Use `--native <executable>` to reuse a built pinned reference runner; otherwise
Cargo builds the locked runner in the implementation snapshot.

`summarize.py <result-directory> --output <generated.csv>` accepts either a new
run with `complete.json` or a published result with completion metadata embedded
in `environment.json`. It checks the 600-sample/120-cell matrix and computes
medians, ranges, allocations and before/after ratios from `samples.jsonl`.

Keep generated results under ignored `benchmark/runs/` or a temporary directory.
Publish selected completed runs as `REPORT.md`, `environment.json`, `samples.jsonl`.
See the [P01 integration report](../results/P01/2026-09-16-apple-m5-rune-split/REPORT.md)
for concrete commands, requirements and the historical measurement method.
