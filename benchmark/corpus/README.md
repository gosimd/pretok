# Prepared corpus

One shared input corpus for all 25 supported exact patterns: **163,066 requests
across seven pools and 22 batches**. Source attribution and licenses are in
[LICENSE.md](LICENSE.md). No expected-output arrays are stored here.

| Pool | UTF-8 bytes | Scalars | ASCII bytes | Source |
| --- | ---: | ---: | ---: | --- |
| en-wikitext2 | 1048576 | 1046437 | 99.67% | [records](en-wikitext2.source.json) |
| ru-wikipedia | 1048575 | 574857 | 9.77% | [records](ru-wikipedia.source.json) |
| ja-wikipedia | 1048575 | 379280 | 4.25% | [records](ja-wikipedia.source.json) |
| ar-wikipedia | 1048576 | 584575 | 11.51% | [records](ar-wikipedia.source.json) |
| zh-wikipedia | 1048574 | 390203 | 5.78% | [records](zh-wikipedia.source.json) |
| synthetic | 5773708 | 2674572 | 17.37% | Generated |
| decimal | 166144 | 66465 | 10.01% | Generated |

Each natural pool has 256-byte, 4 KiB, 64 KiB and whole-pool batches, with
boundaries rounded down to UTF-8 scalar boundaries. Synthetic and decimal pools
use explicit request ranges. Every request is a separate regex subject.

Synthetic coverage includes exhaustive short strings, seeded mixtures, literal
and Unicode-class boundaries, overlapping script groups, whitespace/CR/LF,
contractions, word-case transitions, gaps and long runs. Decimal scenarios cover
ASCII versus Nd/N classes, bounded groups, anchors, continuation failures and
right-alignment around 510/1020/1530 boundaries.

The separate decimal pool contains 110 pure Nd subjects of 4–4,097 scalars in
ASCII, Arabic-Indic, fullwidth, mathematical and mixed-width layouts. Its selected
lengths have nonzero remainders modulo three, making successful P10 matches
measurable. Exact 510/1020 boundaries are covered by the synthetic pool.

[manifest.json](manifest.json), schema 3, records patterns, native-build guard
expressions, input hashes, sources and request counts. `.txt.gz` files contain
input bytes; `*-SIZE.json` files contain byte ranges, not expected matches:

```json
[{"start":0,"end":4}]
```

Before timing, the pinned HF runner generates matches and isolated pieces in
memory. Go and Oniguruma compare every boundary; the reference arrays are then
discarded. All patterns use the same inputs, with no per-pattern corpus directory.

From the repository root, rebuild synthetic inputs and all request ranges using
the checked-in natural text and pinned preparation dependencies:

```sh
python benchmark/prepare/corpus.py --offline
```

See [Reproducibility](../../docs/Reproducibility.md) for dependency setup,
regeneration checks and explicit online source reconstruction. Keep the corpus
license notice and source manifests with redistributed inputs. Historical reports
retain their own corpus fingerprints and request counts.
