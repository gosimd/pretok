# Reproducibility

All maintained inputs and generators needed for the release live in tracked
directories. Ignored `research/` is optional local exploration; it is not an
input to the library, tests, generators or benchmark runner.

## Ordinary tests

From a clean checkout, Go 1.27.0 or newer is sufficient:

```sh
GOWORK=off go test -mod=readonly ./...
GOWORK=off go vet -mod=readonly ./...
```

The tests read the existing P05 and P01 regression fixtures and pinned Unicode
profiles, and run inline regressions and independent differential checks for the
other patterns. They do not contact Hugging Face or invoke Python/Rust. The
[executable examples](../example_test.go) run as part of `go test`.

## Unicode tables and existing regression fixtures

Use a disposable clean checkout when regenerating: these commands write the
tables and manifests. They make no network requests. Installing their pinned
dependencies is a separate preparation step requiring network access or a local
wheel cache:

```sh
python3 -m venv /tmp/pretok-oracle
/tmp/pretok-oracle/bin/python -m pip install tokenizers==0.22.2
```

Go's `gofmt` must be on PATH. Run from the repository root, in this order:

```sh
/tmp/pretok-oracle/bin/python testdata/qwen/generate.py
/tmp/pretok-oracle/bin/python testdata/patterns/generate.py
/tmp/pretok-oracle/bin/python testdata/patterns/generate_case_tables.py
/tmp/pretok-oracle/bin/python testdata/patterns/generate_decimal_tables.py
/tmp/pretok-oracle/bin/python testdata/patterns/generate_punctuation_tables.py
git diff --exit-code -- internal/scan testdata
```

| Generator | Inputs and outputs |
| --- | --- |
| [P05](../testdata/qwen/generate.py) | Reads the tracked P05 manifest; regenerates L/N/whitespace ranges, contraction folds and the existing 7,738-case fixture |
| [P01](../testdata/patterns/generate.py) | Uses P05 fixture inputs; regenerates L\|M ranges and the existing 26,820-case fixture |
| [P04](../testdata/patterns/generate_case_tables.py) | Regenerates case-word classes and their profile |
| [P10](../testdata/patterns/generate_decimal_tables.py) | Regenerates the Nd profile shared by P10, P14 and P25 |
| [P22](../testdata/patterns/generate_punctuation_tables.py) | Regenerates the P\|S profile |

The generators enumerate all valid Unicode scalars using tokenizers 0.22.2's
native oracle. Go's `unicode.Version` and Python's Unicode categories do not
define scanner membership. Exact expressions, immutable model revisions, JSON
pointers, split settings and output hashes remain in the tracked manifests.
See [P05 provenance](../testdata/qwen/README.md) and the
[pattern profile guide](../testdata/patterns/README.md).

Byte-identical regeneration has been checked with Python 3.9.6, tokenizers
0.22.2's macOS arm64 native binary recorded in the manifests, and Go 1.27.0.
Version pinning alone does not pin a wheel across platforms: another native
binary changes `oracle_native_sha256`. Python/zlib versions can also change gzip
bytes. Inspect any differences rather than accepting updated hashes blindly;
compare decompressed fixtures and table membership as well. A changed shared
profile also requires reviewing dependent profile hashes, including P14/P25.
Reference upgrades are a compatibility decision under the
[release contract](Regexp.md#compatibility-policy-for-v01x).

## Shared benchmark input corpus

The [corpus](../benchmark/corpus/README.md) contains input text and request ranges,
not expected splits. To rebuild synthetic inputs, request ranges and the manifest
while reusing the checked-in natural text, install preparation dependencies and
run:

```sh
/tmp/pretok-oracle/bin/python -m pip install -r benchmark/prepare/requirements.txt
/tmp/pretok-oracle/bin/python benchmark/prepare/corpus.py --offline
git diff --exit-code -- benchmark/corpus
```

Offline regeneration validates the cached natural text against its source hashes;
it does not re-download or independently reconstruct the original dataset shards.
Omitting `--offline` explicitly downloads the pinned source shards and rebuilds
the natural pools. Source selection, revisions, article URLs and transformations
are documented in the manifests and [attribution notice](../benchmark/corpus/LICENSE.md).
Keep that notice and all `.source.json` files with redistributed corpus inputs.

## Live reference comparison

Build the native runner and Go harness using the
[benchmark instructions](../benchmark/README.md#build-and-run). Then run each
exact expression independently, choosing a new output directory for every run:

```sh
for id in P01 P02 P03 P04 P05 P06 P07 P08 P09 P10 P11 P12 P13 P14 P15 P16 P17 P18 P19 P20 P21 P22 P23 P24 P25; do
    benchmark/bin/benchreport -pattern "$id" -verify-only \
        -output "benchmark/runs/release-check-$id" || exit 1
done
```

HF generates matches and isolated pieces in memory for every shared request.
Oniguruma and Go must agree on every byte boundary. Existing P01/P05 fixtures
guard the native build separately from the selected pattern's live comparison.
No new expected-output archives are written. Ordinary Go tests are not a
substitute for this external compatibility check.

Timing is a separate operation: omit `-verify-only` to measure after the same
gate. Preserve only selected completed reports under the
[results policy](../benchmark/results/README.md); local runs stay ignored.
