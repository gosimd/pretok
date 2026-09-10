# Prepared corpus

One shared corpus for P01 and P05: identical input text, request boundaries, and scenarios. Source attribution and licenses are in [LICENSE.md](LICENSE.md). These counts describe the checked-in snapshot.

| Pool | UTF-8 bytes | Scalars | ASCII bytes | P01 whole-pool spans | P05 whole-pool spans | Source |
|---|---:|---:|---:|---:|---:|---|
| en-wikitext2 | 1048576 | 1046437 | 99.67% | 214777 | 214795 | [records](en-wikitext2.source.json) |
| ru-wikipedia | 1048575 | 574857 | 9.77% | 102344 | 102447 | [records](ru-wikipedia.source.json) |
| ja-wikipedia | 1048575 | 379280 | 4.25% | 56514 | 56514 | [records](ja-wikipedia.source.json) |
| ar-wikipedia | 1048576 | 584575 | 11.51% | 118537 | 123493 | [records](ar-wikipedia.source.json) |
| zh-wikipedia | 1048574 | 390203 | 5.78% | 62146 | 62146 | [records](zh-wikipedia.source.json) |
| synthetic | 270224 | 173784 | 42.90% | 38769 | 38801 | Generated |

Each natural pool has 256-byte, 4096-byte, 65536-byte, and whole-pool request batches. Boundaries are rounded down to UTF-8 scalar boundaries. Each request is a separate regex subject; its artificial EOF may change selected matches. Synthetic requests are specified runs rather than uniform-size chunks; whole-pool span counts above are descriptive only.

[manifest.json](manifest.json), schema 2, contains all pattern profiles, file hashes, oracle identity, shared request counts, and per-pattern span statistics. Text lives in `.txt.gz` files. Each `*-SIZE.json.gz` stores request ranges once, with expected spans keyed by pattern ID. For example, a request containing `"a\u0301b"` has:

```json
{"start":0,"end":4,"spans":{"P01":[[0,4]],"P05":[[0,1],[1,4]]}}
```

Both runners select the appropriate references and verify them before timing. Missing or incorrect references are errors, never a fallback to another pattern. Ordinary execution does not contact Hugging Face.

Regenerate all pattern references together with `python benchmark/prepare/corpus.py --offline` from the repository root, using the pinned Python dependencies. New scenarios belong in the common generator and corpus; do not create per-pattern corpus directories.
