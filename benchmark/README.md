# Cross-engine benchmark

The offline suite compares the public scalar `pretok` API with native Oniguruma
and HF Tokenizers 0.22.2 for the exact P01 and P05 regexes. It contains five approximately
1 MiB natural text pools and a separate synthetic diagnostic pool. See
[the methodology](../docs/Benchmark.md) and [corpus licenses](corpus/LICENSE.md).

The [P01 Apple M5 report](results/P01/2026-09-10-apple-m5-scalar/REPORT.md)
contains 1,710 samples with raw measurements and environment fingerprints. For
4 KiB requests across the five natural corpora, summed median batch times give
3.79× for Oniguruma span collection versus reused `pretok.Split`, and 6.00× for
native HF Split versus reused `pretok.Split`. The latter includes HF stage
bookkeeping; neither comparison measures BPE or a full tokenizer. These results
apply to scalar P01, Go 1.27.0, and this pinned corpus on Apple M5.

The [original P05 report](results/P05/2026-09-09-apple-m5-scalar/REPORT.md) is retained
with its original source fingerprints. It predates the shared-class refactor.

## Build and run

Requirements: Go 1.27, a Rust toolchain compatible with the checked-in Cargo.lock
(tested with Rust 1.90.0), and a C compiler for Oniguruma. Python is not needed for
ordinary offline benchmark execution. The first Cargo build downloads dependencies;
subsequent builds can use `--offline`.

From the repository root:

```sh
cargo build --release --locked --manifest-path benchmark/native/Cargo.toml
mkdir -p benchmark/bin
go build -o benchmark/bin/benchreport ./cmd/benchreport
benchmark/bin/benchreport -pattern P01 -output benchmark/runs/p01-my-run
benchmark/bin/benchreport -pattern P05 -output benchmark/runs/p05-my-run
```

The default pattern is P05. Each invocation measures one pattern and records its
ID in raw samples and environment metadata.

Defaults are 10 samples, at least 500 ms each, all 21 corpus/size batches, eight
scan engine/mode rows and three compilation rows. Expect roughly 15 minutes plus
correctness gates and warm-up. Engines execute sequentially, with reversed order
on alternate rounds. All raw samples and environment hashes are retained. Existing
output directories are never overwritten; choose a new `-output` for each run.
The default output is ignored `benchmark/runs/local`. Keep all local experiments
under `benchmark/runs/`; only selected complete runs are published under
`benchmark/results/<pattern-id>/<date>-<machine>-<backend>/`, following the
[results policy](results/README.md).

Useful shorter commands:

```sh
benchmark/bin/benchreport -verify-only -output benchmark/runs/verification
benchmark/bin/benchreport -samples 1 -min-time 10ms -output benchmark/runs/smoke
benchmark/bin/benchreport -filter ru-wikipedia -output benchmark/runs/russian
```

Both runners store each corpus contiguously and pass request views to scanners.
Use `-cpu-label` with a verified host description if sandbox restrictions block
automatic CPU discovery.

A smoke run is a functionality check, not stable performance evidence. `-filter`
limits measured corpora, but all original fixtures and corpus request spans are
still checked by the native reference. Results are in `REPORT.md`, `samples.jsonl`,
and `environment.json`; a failed run never sets `complete: true`.

The native runner uses a persistent JSON-line protocol. There is no per-match FFI,
Python binding, file I/O, compilation, or result serialization inside scan timing.
HF's native row includes construction of its input state and extraction of byte
offsets. Go allocation figures use separate 10-pass runs; native C/Rust allocation
figures are unavailable. UTF-8 validation is timed on both public input paths.
`onig-at` is an additional full-subject match-at-offset diagnostic with validation;
it has no dedicated internal-Go core comparison yet.

## Rebuild corpus and oracle data

Only this explicit preparation step uses Python and the network. Use Python
3.9–3.12 with binary wheels. Dataset revisions and source shard paths are pinned
in the generator. HTTP ranges avoid downloading entire Wikipedia shards.

```sh
python3 -m venv /private/tmp/pretok-benchmark-python
/private/tmp/pretok-benchmark-python/bin/pip install -r benchmark/prepare/requirements.txt
/private/tmp/pretok-benchmark-python/bin/python benchmark/prepare/corpus.py
```

To regenerate expected spans from already prepared pools without network access:

```sh
/private/tmp/pretok-benchmark-python/bin/python benchmark/prepare/corpus.py --offline
```

Preparation records exact Python-wheel binary and output hashes. The native build
then has to pass the selected pattern's wheel-oracle fixtures (26,820 for P01; 7,738 for P05) and every generated
request before timing; matching a package version alone is insufficient. Unicode
range changes require deliberate reference review, not automatic fixture updates.

Source text and reference spans are deterministic gzip files. Natural corpora are
approximately equal in UTF-8 bytes, not characters. Every 256 B, 4 KiB, 64 KiB, or
whole-pool request is a separate regex subject. Synthetic requests use predefined
run lengths around vector boundaries and are excluded from combined natural-text
ratios. All maintained inputs and tools live outside ignored `research/`.

Corpus preparation regenerates all supported patterns together in `corpus/`.
There is one manifest and one request file per corpus/size. Each request stores its
input range once and reference spans keyed by pattern ID. The runner's `-pattern`
selects the expected spans and regex; it does not select different inputs or files.
Add any new diagnostic scenarios to this shared corpus so all patterns receive
the same workloads. See the [corpus format](corpus/README.md).
