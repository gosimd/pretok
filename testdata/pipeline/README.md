# HF pipeline oracle

`generate.py` explicitly uses **tokenizers==0.23.2**. Ordinary Go tests run
offline against `cases.json.gz` and verify its decompressed SHA-256 and case count
against `manifest.json`. The native oracle binary hash is also recorded.

Each case has a `name` field with three slash-separated components:
`<component>/<settings>/<input-case>`. The Go test uses it directly, for example:

```text
TestPipelineHFOracle/Digits/individual=true/mixed_number_categories
TestPipelineHFOracle/Metaspace/prepend=first,split=true,replacement=U+2581/leading_repeated_trailing_spaces
TestPipelineHFOracle/Sequence/Split_ByteLevel-9c89680975f8/letters_digits_accent
```

Names use ASCII without whitespace or extra slashes; delimiters and replacements
use Unicode code points. Split regexes use the stable P01–P25 registry IDs.
Non-empty sequences list stage types followed by the first 12 hex digits of
SHA-256 over their configuration encoded as UTF-8 JSON with sorted keys, literal
Unicode and compact separators. This distinguishes stage settings without putting
the full JSON or regex into the test name; the exact configuration is still stored
in `config`. Empty sequences use `Sequence/empty`.

Fixed inputs have descriptive names; seeded inputs use `random_0001`–`random_0012`,
and random compositions use `random_composition_0001`–`random_composition_1000`.
Registry-sweep cases use the `registry_` input prefix to distinguish P02/P16 from
their separate behavior-matrix cases. Long inputs use `long_transform_buffer_growth`.
The generator and Go test reject duplicate names rather than relying on Go's
automatic `#NN` suffixes. Failure output includes the configuration and input.

The [scenario coverage matrix](../../docs/Pipeline.md#scenario-coverage) records
the existing tests, configuration variants, benchmarks and gaps. Naming metadata
does not add cases: all 4,475 inputs, configurations and expected pieces/ranges
were preserved when names were introduced. The fixture hash includes the names.

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
