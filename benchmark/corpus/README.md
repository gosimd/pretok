# Prepared corpus

One shared input corpus for all supported patterns. Source attribution and licenses are in [LICENSE.md](LICENSE.md). No expected-output arrays are stored here.

| Pool | UTF-8 bytes | Scalars | ASCII bytes | Source |
|---|---:|---:|---:|---|
| en-wikitext2 | 1048576 | 1046437 | 99.67% | [records](en-wikitext2.source.json) |
| ru-wikipedia | 1048575 | 574857 | 9.77% | [records](ru-wikipedia.source.json) |
| ja-wikipedia | 1048575 | 379280 | 4.25% | [records](ja-wikipedia.source.json) |
| ar-wikipedia | 1048576 | 584575 | 11.51% | [records](ar-wikipedia.source.json) |
| zh-wikipedia | 1048574 | 390203 | 5.78% | [records](zh-wikipedia.source.json) |
| synthetic | 1050469 | 617864 | 32.68% | Generated |

Each natural pool has 256-byte, 4096-byte, 65536-byte, and whole-pool batches. Boundaries are rounded down to UTF-8 scalar boundaries. Synthetic requests use explicit run lengths (including all short digit residues)
and exhaustive strings of length 0–5 over `0`, `9`, `a`, `１`, `²`, and `😀`. Case-word scenarios add exhaustive length 1–4 strings over `A`, `a`, `ǅ`, `中`,
`ʰ`, a combining mark, space, LF, `1`, and `/`; generated class-boundary neighbors;
2,000 seeded random inputs; and long runs. Line-anchor inputs add short whitespace/CR/LF/Unicode mixtures and runs up to
129 scalars in several end contexts (in addition to existing shared 4096-scalar
runs). Every request is a separate regex subject.

[manifest.json](manifest.json), schema 3, records regex profiles (including the native-build guard expression and full-coverage flag), input hashes, source provenance, and request counts. `case_tables_sha256` identifies the table used to select
additional boundary inputs; the tables themselves contain classification data,
not expected split outputs. `.txt.gz` files hold only input text. Plain `*-SIZE.json` files hold only input ranges, for example:

```json
[{"start":0,"end":4}]
```

Before timing, the pinned native HF runner generates matches and isolated split spans for these exact inputs in memory. Go and Oniguruma compare every boundary. The generated arrays are discarded after checking; only output counts and buffer capacity remain for measurement. Nothing is saved as corpus reference output.

Run `python benchmark/prepare/corpus.py --offline` from the repository root to rebuild inputs and ranges using the pinned preparation dependencies. Add new scenarios to this common corpus; there are no per-pattern input or reference directories.

Numeric diagnostics also cover 1/2/3/4/5/6/7-length boundaries, long runs,
mixed Nd/Nl/No scalars, astral numbers, and neighbors of every pinned N range.
These inputs are shared by all patterns. P16/P07 may leave unmatched gaps;
P24/P20 cover the entire subject. No per-pattern corpus or expected-output
archive is stored.

Case-word suffix diagnostics cover both ordered word branches, leading marks,
scoped case folds, partial/repeated contractions, non-word suffix contexts, and
long words near common boundary sizes. These inputs also belong to the shared
synthetic pool and are verified for every selected pattern.

The shared synthetic pool also covers P08/P11/P17/P18/P21/P23 literal-range endpoints and neighbors, supplementary-plane characters, U+3000 prefix/body overlap, CRLF, and P23 nested-class delimiters. These are input scenarios for every pattern; references are generated live.

P12/P15 share this corpus as well: ordered script transitions, exact range boundaries, long group runs, and both directions through U+302E/U+302F overlaps are included. Reference spans are generated live for each selected pattern.

The grouped-pattern expansion contains 118,090 requests across the same 21 batches. Offline regeneration is byte-identical; live correctness comparison passes for all 21 supported patterns.

P22 extends the shared synthetic inputs with P|S boundaries and neighbors, every ASCII prefix, Unicode marks, controls, number gaps and ordered word/symbol transitions. All patterns use these same inputs; reference outputs are generated live.

The P22 expansion contains 141,415 shared requests across 21 batches. Table and corpus regeneration reproduce the same bytes.

P10 adds Nd boundary contexts, exhaustive short inputs and exact-EOF failures to
`synthetic`. A second shared diagnostic pool, `decimal`, contains 110 pure decimal
subjects of 4–4,097 scalars, all with a nonzero remainder modulo three. It covers
ASCII, Arabic-Indic, fullwidth, mathematical and mixed-width digits. This makes
successful P10 matches measurable separately from natural-text rejection. The
corpus now has seven pools and 22 batches, all available to every pattern.

P14 adds 32 continuation-failure cases to the common synthetic pool: complete or
partial triples followed by an ASCII letter, LF, NUL or a non-decimal number,
then another digit run. It reuses the existing decimal pool and all 22 batches.
References are generated live for every selected pattern.

P25 adds 240 shared synthetic subjects around 510/1020/1530 group boundaries,
with ASCII, Arabic-Indic, fullwidth and mathematical digits, mixed-width suffixes
and intervening gaps. It reuses the existing decimal pool and all 22 batches.
All inputs remain shared, and expected output spans are generated live.
