# Prepared corpus

Pinned text pools and expected P05 byte spans. Source text attribution and licenses are in [LICENSE.md](LICENSE.md); exact source records and preparation details are in the linked manifests. These counts describe the checked-in snapshot.

| Pool | UTF-8 bytes | Scalars | ASCII bytes | Whole-pool spans | Mean span bytes | Max span bytes | Source |
|---|---:|---:|---:|---:|---:|---:|---|
| en-wikitext2 | 1048576 | 1046437 | 99.67% | 214795 | 4.88 | 31 | [records](en-wikitext2.source.json) |
| ru-wikipedia | 1048575 | 574857 | 9.77% | 102447 | 10.24 | 45 | [records](ru-wikipedia.source.json) |
| ja-wikipedia | 1048575 | 379280 | 4.25% | 56514 | 18.55 | 228 | [records](ja-wikipedia.source.json) |
| ar-wikipedia | 1048576 | 584575 | 11.51% | 123493 | 8.49 | 35 | [records](ar-wikipedia.source.json) |
| zh-wikipedia | 1048574 | 390203 | 5.78% | 62146 | 16.87 | 156 | [records](zh-wikipedia.source.json) |
| synthetic | 270224 | 173784 | 42.90% | 38801 | 6.96 | 57912 | Generated |

Each natural pool has 256-byte, 4096-byte, 65536-byte, and whole-pool request batches. Boundaries are rounded down to UTF-8 scalar boundaries. Each request has its own reference spans, since an artificial EOF may change the result. Synthetic requests are specified runs rather than uniform-size chunks; their whole-pool span count above is descriptive only and is not the benchmark request total.

[manifest.json](manifest.json) records file hashes, per-batch request/span counts, oracle identity, and character-class frequencies. `.txt.gz` files contain input bytes; `*-SIZE.json.gz` files contain request ranges and expected spans. Both runners verify them before timing. Ordinary execution does not contact Hugging Face.
