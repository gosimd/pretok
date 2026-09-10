# Prepared corpus

One shared input corpus for all supported patterns. Source attribution and licenses are in [LICENSE.md](LICENSE.md). No expected-output arrays are stored here.

| Pool | UTF-8 bytes | Scalars | ASCII bytes | Source |
|---|---:|---:|---:|---|
| en-wikitext2 | 1048576 | 1046437 | 99.67% | [records](en-wikitext2.source.json) |
| ru-wikipedia | 1048575 | 574857 | 9.77% | [records](ru-wikipedia.source.json) |
| ja-wikipedia | 1048575 | 379280 | 4.25% | [records](ja-wikipedia.source.json) |
| ar-wikipedia | 1048576 | 584575 | 11.51% | [records](ar-wikipedia.source.json) |
| zh-wikipedia | 1048574 | 390203 | 5.78% | [records](zh-wikipedia.source.json) |
| synthetic | 946670 | 548451 | 30.74% | Generated |

Each natural pool has 256-byte, 4096-byte, 65536-byte, and whole-pool batches. Boundaries are rounded down to UTF-8 scalar boundaries. Synthetic requests use explicit run lengths (including all short digit residues)
and exhaustive strings of length 0–5 over `0`, `9`, `a`, `１`, `²`, and `😀`. Case-word scenarios add exhaustive length 1–4 strings over `A`, `a`, `ǅ`, `中`,
`ʰ`, a combining mark, space, LF, `1`, and `/`; generated class-boundary neighbors;
2,000 seeded random inputs; and long runs. Every request is a separate regex subject.

[manifest.json](manifest.json), schema 3, records regex profiles (including the native-build guard expression and full-coverage flag), input hashes, source provenance, and request counts. `case_tables_sha256` identifies the table used to select
additional boundary inputs; the tables themselves contain classification data,
not expected split outputs. `.txt.gz` files hold only input text. Plain `*-SIZE.json` files hold only input ranges, for example:

```json
[{"start":0,"end":4}]
```

Before timing, the pinned native HF runner generates matches and isolated split spans for these exact inputs in memory. Go and Oniguruma compare every boundary. The generated arrays are discarded after checking; only output counts and buffer capacity remain for measurement. Nothing is saved as corpus reference output.

Run `python benchmark/prepare/corpus.py --offline` from the repository root to rebuild inputs and ranges using the pinned preparation dependencies. Add new scenarios to this common corpus; there are no per-pattern input or reference directories.
