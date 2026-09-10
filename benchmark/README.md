# Cross-engine benchmark

The offline suite compares the public scalar `pretok` API with native Oniguruma
and HF Tokenizers 0.22.2 for the exact P01, P02, P04, and P05 regexes. It contains five approximately
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
benchmark/bin/benchreport -pattern P02 -output benchmark/runs/p02-my-run
benchmark/bin/benchreport -pattern P04 -output benchmark/runs/p04-my-run
benchmark/bin/benchreport -pattern P05 -output benchmark/runs/p05-my-run
```

The default pattern is P05. Each invocation measures one pattern and records its
ID in raw samples and environment metadata.

Defaults are 10 samples, at least 500 ms each, all 21 corpus/size batches, eight
scan engine/mode rows (seven for P02) and three compilation rows. Expect roughly 15 minutes plus
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
limits measured corpora, but live HF references are generated and compared for
every corpus request before measurement. Results are in `REPORT.md`, `samples.jsonl`,
and `environment.json`; a failed run never sets `complete: true`.

The native runner uses a persistent JSON-line protocol. There is no per-match FFI,
Python binding, file I/O, compilation, or result serialization inside scan timing.
HF's native row includes construction of its input state and extraction of byte
offsets. Go allocation figures use separate 10-pass runs; native C/Rust allocation
figures are unavailable. UTF-8 validation is timed on both public input paths.
`onig-at` is an additional full-subject match-at-offset diagnostic with validation;
it has no dedicated internal-Go core comparison yet.

## Reference generation and correctness

Every normal run and `-verify-only` executes the same pipeline:

1. Load pinned input bytes and request ranges.
2. Check the native reference build against the existing regression fixtures.
3. Generate HF matches and isolated split spans in memory for each batch.
4. Compare every Oniguruma and Go boundary; fail before timing on any mismatch.
5. Discard reference arrays, retaining only counts and reusable-buffer capacity.
6. Run measurements (omitted with `-verify-only`).

No corpus reference arrays are stored, loaded, or decompressed. Native reference
construction, JSON transport, comparison, and reference disposal are outside the
scan timer. Both runners refuse to measure batches that have not passed their
reference gate. `matches` and `split` are separate oracle outputs and counters.
The match-at diagnostic remains limited to fully covering P01/P04/P05. P02 omits
`onig-at`: its gaps require a search operation. Iterator rows count matches;
Split rows count matches plus gaps, so P02 output counts can differ.

Small regression fixtures under `testdata/` remain independent of the mass-corpus
pipeline: they support offline Go tests and guard the pinned Unicode/reference
profile. Rust/HF dependencies are fixed by Cargo.lock; the actual binary hash is
recorded in each run. P02/P04 reuse the P05 fixtures only to guard the native build,
with their P05 expression. Their own HF references are generated live; no P02/P04
fixture archives are required.

## Rebuild corpus inputs

Only this explicit preparation step uses Python and the network. Use Python
3.9–3.12 with binary wheels. Dataset revisions and source shard paths are pinned
in the generator. HTTP ranges avoid downloading entire Wikipedia shards.

```sh
python3 -m venv /private/tmp/pretok-benchmark-python
/private/tmp/pretok-benchmark-python/bin/pip install -r benchmark/prepare/requirements.txt
/private/tmp/pretok-benchmark-python/bin/python benchmark/prepare/corpus.py
```

To rebuild request ranges and manifests from existing pools without network access:

```sh
/private/tmp/pretok-benchmark-python/bin/python benchmark/prepare/corpus.py --offline
```

Preparation records source, input, and generator hashes. It does not execute a
regex engine or generate expected output. The native build still has to pass the
native-build regression fixtures (26,820 for P01; the 7,738 P05 cases for P05/P02/P04), so a crate
version alone is insufficient to establish Unicode compatibility.

Source text uses deterministic gzip files; request ranges use plain JSON. Natural corpora are
approximately equal in UTF-8 bytes, not characters. Every 256 B, 4 KiB, 64 KiB, or
whole-pool request is a separate regex subject. Synthetic requests use predefined
run lengths around vector boundaries and are excluded from combined natural-text
ratios. All maintained inputs and tools live outside ignored `research/`.

Corpus preparation writes a single common manifest and input-range file per
corpus/size. The runner's `-pattern` selects the regex used to generate live
references; it does not select different input files. Add diagnostic scenarios to
this shared corpus so every pattern receives the same workloads. See the
[corpus format](corpus/README.md).
