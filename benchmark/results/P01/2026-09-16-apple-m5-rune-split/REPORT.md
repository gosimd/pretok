# P01 buffered Split: integrated implementation

2026-09-16, Apple M5, darwin/arm64, Go 1.27.0. This is a completed comparison
of the public API before and after integration, not the earlier prepared-worker
prototype. Baseline commit and all source, binary, corpus and dependency identities
are in [environment.json](environment.json).

## Implementation and scope

P01 `Split` on eligible SIMD builds validates **all** input first, then calls
`gosimd/unicode/utf8.DecodeValidInto` with a local **1,024-rune / 4 KiB** buffer.
The dependency is the published `github.com/gosimd/unicode v0.1.2`, without a
sibling checkout or `replace`. No pool, mutable plan state, staged spans or public
API changes were introduced. Rune classification remains scalar and uses the
same pinned tables. Byte and rune scanners share contraction predicates and the
whitespace endpoint decision.

The preceding prototype compared 1K, 2K, 4K and 8K rune windows: reused-output
median throughput differed by at most 1.04% within each workload. We chose 1K
to bound scratch to 4 KiB. Full-input validation preserves the public guarantee
that invalid UTF-8 leaves the destination unchanged, without staging output.
Prototype measurements and profiling artifacts are local development material;
the integrated measurements below are the published performance evidence.

The rune path is selected for P01 inputs of at least 256 bytes in
`GOEXPERIMENT=simd` builds on ARM64 or AMD64 with AVX2. The Unicode module chooses
its NEON/AVX2/AVX-512 kernel. Other configurations and short inputs keep byte
scanning. An initial local comparison found materializing runes without SIMD
slightly slower on Russian/Arabic; the final implementation avoids that path.
The portable-build eligibility condition is a compile-time false constant.

`Cursor.Next` remains a byte scanner with independent value-copy progress.
For eligible long P01 inputs its constructor benefits from SIMD validation.
Validation still precedes all output writes, including when malformed UTF-8 is
at the end of a long input. Tokens and whitespace rollback can cross rune windows.

## Method

Four persistent Go processes use the identical measurement harness:
`old-go`, `new-go`, `old-simd`, `new-simd`. The `-simd` pair is built with
`GOEXPERIMENT=simd`; the `-go` pair uses ordinary builds. Samples are executed
sequentially in a seeded shuffled order, with `GOMAXPROCS=1` per process and no
core affinity. No correctness tests, fuzzing or other benchmark runs overlap the
final timing run.

The shared corpus supplies five languages and request sizes 256 B, 4 KiB, 64 KiB
and 1 MiB. Reused `Split` is timed at all sizes; fresh `Split` and match-only
`Cursor` are also timed at 4 KiB. **600 samples / 120 cells**, five repetitions
of at least 150 ms each. The final request in a batch can be shorter. Medians
below use decimal MB/s and include full-input validation, local scratch setup,
decoding, scanning, byte-offset reconstruction and output writes. Caller capacity
is prepared for reused-output rows; there is no prepared rune worker. Allocations
are measured separately using ten warmed batch iterations.

HF/Oniguruma are correctness references in this run, not throughput baselines.
Historical cross-engine reports retain their original numbers. Five samples do
not establish significance for small differences; raw ranges are preserved.

## Reused Split results

4 KiB requests, medians; comparisons use matching compiler settings:

| Corpus | Before SIMD-build MB/s | After SIMD-build MB/s | Speedup | Before ordinary MB/s | After ordinary MB/s |
| --- | ---: | ---: | ---: | ---: | ---: |
| English | 222.1 | 270.4 | 1.218× | 221.8 | 222.4 |
| Russian | 178.4 | 219.3 | 1.229× | 178.2 | 180.4 |
| Japanese | 200.4 | 219.7 | 1.096× | 201.0 | 201.5 |
| Arabic | 179.8 | 229.0 | 1.274× | 179.6 | 182.3 |
| Chinese | 256.6 | 291.0 | 1.134× | 257.2 | 257.5 |

SIMD-build speedup by request size:

| Corpus | 256 B | 4 KiB | 64 KiB | 1 MiB |
| --- | ---: | ---: | ---: | ---: |
| English | 1.134× | 1.218× | 1.228× | 1.242× |
| Russian | 1.073× | 1.229× | 1.232× | 1.230× |
| Japanese | 0.986× | 1.096× | 1.093× | 1.094× |
| Arabic | 1.077× | 1.274× | 1.286× | 1.282× |
| Chinese | 0.996× | 1.134× | 1.120× | 1.120× |

Japanese 256 B requests are a small regression, **0.986×** (about 1.4% lower
throughput). The largest SIMD speedup is about **1.286×** on Arabic 64 KiB.
We keep the 256 B threshold because other measured languages benefit there;
this is a workload tradeoff, not an across-the-board improvement claim.
Ordinary-build reused-output controls range from 0.985× to 1.072×; they still
execute the byte scanner, so these differences are not rune/SIMD speedups.

All reused Split and Cursor cells report **0 B/request and 0 allocations/request**.
This includes local rune scratch setup in Split, but a new goroutine can still
need runtime stack growth. Fresh-output allocation and throughput figures, Cursor
measurements, per-request latency and min/max sample ranges can be recalculated
from all 600 raw [samples](samples.jsonl) using the summarizer below.

## Verification and generated code

Each of the four builds passed the live native gate: the native runner checked
26,820 pinned P01 fixtures, then generated matches and isolated splits for all
**163,066 requests / 22 batches** in the common corpus. Public Cursor and Split
were both checked against those byte boundaries before timing. This includes the
synthetic/decimal inputs even though their throughput is not measured here.
Gate results are embedded in [environment.json](environment.json). No reference outputs were saved in the corpus.

Additional tests compare the rune reader directly with the retained byte scanner
at scratch capacities 3, 4, 7, 16, 1K, 2K, 4K and 8K. They cover long tokens,
marks, contractions, CR/LF, multibyte whitespace, rollback, EOF and scratch guards.
Public tests cover the 256-byte threshold, malformed suffixes, unchanged complete
destination backing storage, append/reuse, concurrent Split and independent
public cursor copies. Standard and SIMD test/race suites and vet checks passed.
A SIMD fuzz run completed **392,207 executions** without a mismatch. Linux/AMD64
SIMD cross-build passed; it verifies buildability, not runtime performance or
runtime parity on AMD64. Source-only reconstruction also passed offline tests
in both ordinary and SIMD builds; see the reproduction commands below.

Compiler diagnostics showed `peek` inlined at cost 79 and no heap escape of the
rune buffer. Its refill helper is deliberately non-inline. Generated assembly
showed a 4,432-byte frame for the buffered helper, including the 4,096-byte rune
array; short/other-pattern calls avoid this helper. Rune scanning used indexed
32-bit loads with no `runeAt` or `DecodeRune` calls. The published Unicode decoder
used NEON widening and table-lookup instructions. Diagnostic dumps are local
artifacts and are not included in this result. No claim is made that classification
itself is vectorized.

## Reproduce

The shared [comparison tools](../../../compare/README.md) reconstruct the baseline
and implementation commits recorded in `environment.json`, build all four
executables with the same measurement server, run live correctness gates and
repeat the interleaved measurement. No saved patch or prototype is required.
Run from the repository root:

```sh
python3 benchmark/compare/reproduce.py \
  --result benchmark/results/P01/2026-09-16-apple-m5-rune-split \
  --work /private/tmp/pretok-p01-comparison \
  --native /path/to/pretok/benchmark/native/target/release/pretok-reference
```

Use a new `--work` path. Requirements: Python 3.9+, Go 1.27 with the SIMD
experiment, both Git commits and the pinned Go dependency in cache or available
through the proxy. Omitting `--native` builds the locked Rust runner and may
require Cargo dependencies/network access. `--prepare-only` reconstructs sources
without compiling or measuring. Neither the working checkout nor the sibling
Unicode repository is modified. Rebuilt executable hashes can differ due to
source paths and harness formatting; historical hashes remain in the metadata.

To recalculate the historical statistics without rerunning timings:

```sh
python3 benchmark/compare/summarize.py \
  benchmark/results/P01/2026-09-16-apple-m5-rune-split \
  --output benchmark/runs/P01/rune-split-summary.csv
```

To check reconstructed sources after `--prepare-only`, run `GOWORK=off go test
-mod=readonly ./...` and `GOWORK=off GOEXPERIMENT=simd go test -mod=readonly ./...`
in each snapshot directory. New measurements, generated summaries, profiles and
diagnostic logs belong in ignored `benchmark/runs/` or a temporary directory.
Only the report, environment metadata and raw samples are published here.
