# P05 provenance and regeneration

[manifest.json](manifest.json) pins the exact P05 regex, Qwen model revision,
tokenizer file hash, JSON pointer, split settings, native oracle binary hash and
generated artifact hashes. The generator reads this tracked manifest directly;
it does not download a tokenizer or depend on local research files.

With tokenizers==0.22.2 and `gofmt` available, run from the repository root:

```sh
python testdata/qwen/generate.py
```

[generate.py](generate.py) enumerates every valid Unicode scalar to extract
L, N and whitespace membership and contraction-letter case folds. It regenerates
[unicode_tables.go](../../internal/scan/unicode_tables.go) and the existing
`cases.jsonl.gz` regression fixture, then records their hashes in the manifest.
The generated inputs are synthetic, not copied tokenizer training text.

The 7,738 cases cover exhaustive short strings, class boundaries and neighbors,
seeded random mixtures, contractions, combining marks, long runs and whitespace.
Matches use `Removed, invert=true`; isolated pieces use `Isolated, invert=false`.
Python character offsets are converted to UTF-8 byte offsets. Gzip uses a fixed
timestamp. Full tokenizer behavior, ByteLevel and BPE are outside these fixtures.

For the environment, generation order, byte reproducibility limits and the
separate live-reference gate, see [Reproducibility](../../docs/Reproducibility.md).
The [registry](../../docs/SupportedPatterns.md) is the canonical list of supported
patterns; the name of this fixture directory does not imply full Qwen support.
