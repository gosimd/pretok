# P01 reference fixtures

`P01/manifest.json` pins the exact decoded regex, immutable source revision,
source settings, tokenizers 0.22.2 wheel hash, generated Unicode table hash, and
fixture hash. Ordinary Go tests are offline and verify these hashes.

Regenerate with the Python environment from `benchmark/prepare/requirements.txt`:

```sh
python generate.py
```

Run this command from this directory. The script performs no network access. It
reads the P01 manifest for provenance and the original P05 cases as shared input
seeds. It enumerates the pinned oracle's L|M class over every Unicode scalar,
then generates independent matches and isolated split spans for 26,820 cases:
word-class boundaries, exhaustive short strings, mixed scripts and punctuation,
contractions, whitespace, long runs, and the existing randomized Qwen inputs.
Python scalar offsets are converted to UTF-8 byte offsets before writing fixtures.

`internal/scan/word_tables.go` contains the precomputed L|M union needed by P01. N,
whitespace and contraction-fold tables are shared with P05. The template selects
the P05 L table or P01 L|M table once; both use the same ASCII path and a single
Unicode table lookup. P02–P04 remain unsupported and have no fixtures in this directory.
