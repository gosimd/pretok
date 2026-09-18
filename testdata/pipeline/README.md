# HF pipeline oracle

`generate.py` explicitly uses **tokenizers==0.23.2**. Ordinary Go tests run
offline against `cases.json.gz` and verify its decompressed SHA-256 and case count
against `manifest.json`. The native oracle binary hash is also recorded.

The fixtures are synthetic and generated in this repository; they contain no
downloaded model or training text. Cases cover all supported built-ins, all 25
registered regexes, all five split behaviors and invert, literal multi-character
delimiters, prefix settings, nested transformations, Unicode numeric/word/space
classes, mixed scripts, combining marks, and empty strings. A fixed random seed
adds mixed short inputs and 1,000 composed pipelines. Long transformation cases
exercise buffer growth. They do not establish compatibility with entire models.

Expected pieces are obtained from the actual HF components. To verify
`OriginalRange`, the generator passes arbitrary non-empty byte ranges as token
offsets through HF `PreTokenizedString.tokenize` / `to_encoding`. Python's final
character offsets are independently converted to original UTF-8 byte offsets.
All byte subranges are covered on short pieces; longer pieces cover individual
bytes and the whole piece. This includes offsets inside UTF-8 encodings.

The built-in `word`, `numeric`, `whitespace`, and `punctuation` tables are obtained
by probing their actual HF pre-tokenizers on **every Unicode scalar**. This avoids
assuming Rust `regex`, Rust standard-library Unicode, `unicode_categories` and
Oniguruma share a Unicode version or membership definition.

Script tables are mechanically derived from the pinned HF 0.23.2
[`scripts.rs`](https://github.com/huggingface/tokenizers/blob/v0.23.2/tokenizers/src/pre_tokenizers/unicode_scripts/scripts.rs),
with the exact source SHA-256 checked against the manifest. The generator applies
HF's Han/Hiragana/Katakana/U+30FC and U+0020 adaptations. The upstream source is
Apache-2.0; see `LICENSE.tokenizers` and the repository `NOTICE`.

Reproduce in a separate Python environment, after explicitly installing the oracle
and obtaining that source file:

```sh
python -m pip install tokenizers==0.23.2
python testdata/pipeline/generate.py --scripts-source /path/to/scripts.rs
gofmt -w internal/pipeline/tables.go
GOWORK=off go test -mod=readonly -run TestPipelineHFOracle .
```

The generator itself performs no network access. `--fixtures-only` regenerates
the fixtures without rebuilding Unicode tables. Source or oracle changes require
an explicit profile update and review, not silently altered expected results.
