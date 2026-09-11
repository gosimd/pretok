# Benchmark design

**Status:** implemented for scalar P01–P13 and P15–P24, with five pinned natural-language pools,
two shared synthetic pools, Go measurement/orchestration, native Oniguruma/HF runners,
exact-span correctness gates, and Markdown/JSON reports. Build and run instructions
are in [benchmark/README.md](../benchmark/README.md). SIMD and an internal Go
match-at core comparison remain future work.

The benchmark answers two separate questions:

1. How much does a specialized scanner improve boundary discovery over a compatible
   general regex engine?
2. How much does the public regex-splitting API cost compared with HF Tokenizers'
   implementation of the same splitting stage?

Select P01–P13 or P15–P24 with `-pattern`. Each run uses one exact supported string and its
own [P01](../testdata/patterns/P01/manifest.json),
[P02](../testdata/patterns/P02/manifest.json),
[P03](../testdata/patterns/P03/manifest.json),
[P04](../testdata/patterns/P04/manifest.json),
[P05](../testdata/qwen/manifest.json),
[P09](../testdata/patterns/P09/manifest.json), or
[P13](../testdata/patterns/P13/manifest.json) provenance manifest. The numeric profiles
[P16](../testdata/patterns/P16/manifest.json), [P07](../testdata/patterns/P07/manifest.json),
[P24](../testdata/patterns/P24/manifest.json), and [P20](../testdata/patterns/P20/manifest.json)
use the same reference-generation pipeline. The case-word suffix profiles
[P19](../testdata/patterns/P19/manifest.json) and [P06](../testdata/patterns/P06/manifest.json)
also generate their references live. P02–P04/P06–P09/P11–P13/P15–P24 store no reference
outputs; their native-build guard uses the existing P05 fixtures with the P05 regex. Compile once before timing scans.
Use the exact same decoded regex, input bytes, subject boundaries, and output
semantics in every comparison. Broader tokenizer pipelines are a later integration
benchmark in `gosimd/bpe`.

## Baselines

| Level | pretok | Reference | Meaning |
| --- | --- | --- | --- |
| Compilation | `Compile(pattern)` | Native Oniguruma compilation; HF Split construction separately | One-time setup cost |
| Match iteration | `Cursor` and repeated `Next` | Native Oniguruma `find_iter` | Complete ordered sequence of match spans |
| Position matching, diagnostic | Internal Go counterpart deferred | Native Oniguruma match-at-offset loop, including input validation | Additional P01/P04/P05/P06/P09/P13/P19/P20/P24 measurement; no core speedup ratio yet |
| Public splitting | `Split(src, dst[:0])` and `Split(src, nil)` separately | Oniguruma iteration plus isolated-gap adapter and span output | Same output, with reused and fresh destination policies |
| HF splitting stage | Public `Split` | Native HF `Split(Regex(pattern), Isolated, invert=false)` plus byte-offset extraction | Complete boundary-stage API cost, including HF representation overhead |
| Future optimized backend | Public API using SIMD | The same API using the scalar plan | Backend-only improvement on the same semantics and allocation policy |

The primary regex engine is **Oniguruma**, through the same resolved `onig` /
`onig_sys` dependency versions and build features as the native HF reference.
HF Tokenizers v0.22.2's Oniguruma wrapper uses `find_iter` to collect matches and
gaps. Its Cargo manifest enables the `onig` backend by default, but the harness
must pin the actual dependency resolution with a lockfile and report it.
[HF wrapper](https://github.com/huggingface/tokenizers/blob/v0.22.2/tokenizers/src/utils/onig.rs),
[HF dependency configuration](https://github.com/huggingface/tokenizers/blob/v0.22.2/tokenizers/Cargo.toml).

Go's standard `regexp` package uses RE2 syntax, which cannot directly execute P01/P05's
negative lookahead. Removing or rewriting that branch would change the workload.
It is therefore not a primary baseline for this expression. Other engines may be
added as explicitly versioned, correctness-checked competitors; a .NET/PCRE-style
engine is not automatically an HF semantic reference.
[Go regexp](https://pkg.go.dev/regexp),
[RE2 syntax](https://github.com/google/re2/wiki/Syntax).

### The MatchPrefix proposal

Treat `MatchPrefix` as the operation "match this complete regex at this byte
position", not a loop over separately compiled alternatives. P01/P04/P05/P06/P09/P13/P19/P20/P24 cover every
valid scalar, so the following diagnostic is valid for these patterns:

```text
pos = 0
while pos < len(subject):
    length = match_at_offset(compiled_regex, subject, pos)
    require length > 0
    consume(pos, pos + length)
    pos += length
```

Oniguruma exposes this operation through `match_with_options`; its result is a
byte length, not an absolute endpoint. Pass the full original subject and an
explicit offset. Reuse match-region storage when the selected API requires it.
[Oniguruma Rust API](https://docs.rs/onig/6.5.3/onig/struct.Regex.html#method.match_with_options).

Do not repeatedly pass `subject[pos:]`: that changes the subject boundary and can
change anchor/lookbehind semantics when more patterns are supported. For patterns
with gaps, an anchored-prefix loop is not equivalent to search iteration. The
primary baseline remains `find_iter`, whose search behavior also follows HF's
implementation. Any engine error or unexpected zero-length match invalidates the
sample; it must not be counted as a fast result.

### The HF proposal

The principal HF row runs only the configured **regex Split stage**. For each
input, the timed native Rust operation includes:

1. Construction of a fresh `PreTokenizedString` from the source text.
2. Execution of the already constructed `Split` pre-tokenizer.
3. Extraction of spans with `OffsetReferential::Original` and `OffsetType::Byte`.
4. Consumption of the result and release of per-call temporary state.

Input-container construction is part of this public-stage cost. Repeatedly
splitting an already split object is not the intended workload. HF's source
exposes explicit byte-offset extraction; use it rather than measuring a Python
character-to-byte conversion.
[HF Split](https://github.com/huggingface/tokenizers/blob/v0.22.2/tokenizers/src/pre_tokenizers/split.rs),
[HF PreTokenizedString](https://github.com/huggingface/tokenizers/blob/v0.22.2/tokenizers/src/tokenizer/pre_tokenizer.rs).

HF retains richer per-piece state than `[]Span`. That cost stays in the HF row and
is identified as a public-stage comparison, not attributed entirely to the regex
engine. Its existing allocation policy is reported as `native`; do not pretend it
has the same destination-reuse API. Keep the direct Oniguruma rows to separate
engine work from HF representation costs.

Python `pre_tokenize_str` may be an additional application-level row. It includes
binding calls and Python result objects, so it must not be the only HF baseline.
Full `Tokenizer.encode`, ByteLevel mapping, BPE merges, special tokens, and
post-processing do not belong in a scanner speedup ratio. Once integrated in BPE,
compare two otherwise identical full tokenizers with only regex splitting changed.

## Input corpus

The initial natural-language suite contains five text pools:

| ID | Language | Proposed source |
| --- | --- | --- |
| `en-wikitext2` | English | `Salesforce/wikitext`, `wikitext-2-raw-v1`, train |
| `ru-wikipedia` | Russian / Cyrillic | `wikimedia/wikipedia`, `20231101.ru`, train |
| `ja-wikipedia` | Japanese | `wikimedia/wikipedia`, `20231101.ja`, train |
| `ar-wikipedia` | Arabic | `wikimedia/wikipedia`, `20231101.ar`, train |
| `zh-wikipedia` | Chinese | `wikimedia/wikipedia`, `20231101.zh`, train |

Use WikiText's **raw** variant to avoid vocabulary-based unknown-token replacement.
Raw WikiText is still a prepared dataset; do not describe it as untouched Wikipedia
prose. Wikimedia's dataset provides cleaned article text by language and dump date.
Use supplied text as-is, preserving punctuation, whitespace, marks, and case.
[WikiText dataset card](https://huggingface.co/datasets/Salesforce/wikitext),
[Wikimedia dataset card](https://huggingface.co/datasets/wikimedia/wikipedia).

Select deterministic source records, pin their dataset commit and record IDs, and
store the exact resulting bytes. The implemented preparation reads the pinned first shard in stored row order,
keeps non-empty text,
and take at most 32 KiB per record at a UTF-8 boundary. Join records with exactly
`\n\n` and keep a manifest of source ranges and inserted separators. This bounds
one article's contribution without rewriting its content. Do not loop a short
excerpt to fill the corpus.

Target **1 MiB of UTF-8 per language**, taking the longest valid UTF-8 prefix no
larger than that budget. Actual sizes may differ by at most three bytes. Do not
pad with spaces to force equal size. Report actual processed bytes. A later larger
corpus is a new dataset revision, not an unrecorded change to the same benchmark.

English WikiText and the other four pools have different preparation histories.
They measure workload-dependent performance; language alone is not a controlled
variable. Add `en-wikipedia` from the same dump as an optional sixth control before
attributing differences specifically to script or language. Russian represents
one Cyrillic-language workload, not all Cyrillic text.

Record per pool:

- Source repository, immutable revision, subset, split, shard paths, row/article
  IDs, original article URLs where available, and extraction-script revision.
- Exact joining/truncation rules, file hash, byte length, scalar count, and pattern-specific span
  count. The source article attributions and license notice accompany the samples;
  dataset text must not silently inherit the library's MIT license.
- ASCII-byte fraction, Unicode-whitespace and number frequencies, average span
  byte length, and maximum span byte length.

The dataset cards describe their respective attribution/share-alike licensing;
retain their notices and provenance with distributed excerpts.
[WikiText license information](https://huggingface.co/datasets/Salesforce/wikitext#licensing-information),
[Wikimedia license information](https://huggingface.co/datasets/wikimedia/wikipedia#licensing-information).

## Request sizes and synthetic cases

A corpus pool is not automatically one application request. Derive deterministic
request batches at approximately **256 B, 4 KiB, and 64 KiB**, plus the complete
approximately **1 MiB** pool as one throughput subject. Cut only at UTF-8 boundaries,
record the resulting ranges, and treat each request as a separate subject in all
engines. A chunk's EOF may change regex decisions; generate/check expected spans
for each actual request instead of slicing results for the whole pool.

One timed operation processes the complete ordered batch for the selected size.
All engines receive the identical batch. Report both time per batch and time per
request; one corpus pass processes the same actual byte total at each size.
These are deterministic text excerpts, not claimed to be a sampled distribution
of production prompts. Rotate across the batch rather than repeatedly timing only
its first short string. Distinguish steady-state cache reuse from any future
explicit cold-cache benchmark.

Natural text will not adequately exercise every construction. Keep a small
separate diagnostic suite: ASCII and Unicode letter runs, marks, numbers, symbols,
contractions, whitespace before a word, whitespace at EOF, CR/LF suffixes, and runs
of 1/2/3/4-byte scalars. Include sizes around prospective vector boundaries and
long-run scaling cases. Do not combine these synthetic results with the natural
language headline average.

## Correctness gate

Before timing each engine on a batch:

- Validate input UTF-8, pattern identity, oracle profile, and manifest hashes.
- Generate HF matches and isolated splits in memory for each batch, outside timing.
- Compare every ordered `(Start, End)` byte span against that live oracle for
  that exact request. Counts and checksums alone do not prove equality.
- Verify non-empty spans, valid scalar boundaries, gap policy, and reconstruction
  for isolated splitting. P01/P04/P05/P06/P09/P13/P19/P20/P24's match-only and isolated outputs coincide because
  it covers the input; retain the policy labels for future patterns.
- Confirm dependency-resolved native Oniguruma behavior against existing fixtures;
  a crate version alone does not establish parity with the recorded Python wheel.

On mismatch, print the engine, request ID, first differing span, and nearby escaped
text, then stop. Do not publish a speedup for an incorrect result or modify the
pattern to make the comparison compile. Correctness, hashes, and corpus statistics
are computed outside the measured interval. Corpus reference arrays are never
persisted. Discard them after checking; retain only counts and buffer capacities.

For P02/P03/P07/P08/P11/P12/P15/P16/P17/P18/P21/P22/P23, omit the match-at diagnostic: unmatched gaps are expected. Match-only
rows count matches, while isolated-split rows include maximal gaps. A subject
with no match still measures a complete scan; its match throughput is zero.

## Timing boundaries and runners

Use a Go benchmark executable for pretok and a native Rust executable for the
Oniguruma/HF references, sharing corpus and result manifests. Time loops inside
each process. Process startup, loading, decompression, compilation, fixture
verification, JSON serialization, and subprocess communication are outside scan
timing. Do not call a foreign function or Python once per match; if a later FFI
adapter is measured, label that integration cost separately.

Public pretok scanning includes its existing UTF-8 validation. For matching work
with the same byte-input contract, the Rust wrapper validates each request with
`from_utf8` inside the timed public operation before scanning. Label that wrapper
explicitly. A separate core diagnostic may assume prevalidated input on both sides
and call the internal Go plan; it must not introduce a public skip-validation API.
Preserve the regex engine's normal internal checks and report build options.

For both direct regex and pretok span collection, measure `reused` and `fresh`
output separately. In `reused`, reserve sufficient capacity during warm-up and
reset length each request; in `fresh`, allocate a new destination each request.
Include normal destruction/GC effects in steady-state measurements; do not disable
GC for one competitor or hide output cleanup outside every timed operation.

Iterators consume all boundaries using the same small rolling accumulator so the
compiler cannot remove work. Span-output benchmarks make the returned buffer
observable through appropriate sinks/black-box operations. Full equality checks
and cryptographic hashes remain outside the timer. Inspect benchmark codegen if
implausibly low setup times suggest constant folding or dead-code elimination;
load the regex string from the manifest and retain compiled results when measuring
compilation.

Start with one worker, Go `GOMAXPROCS=1` / `-cpu=1`, Rust release builds, no competing
benchmark processes, and identical CPU/power conditions. Warm up each binary,
calibrate batches to at least 500 ms per sample, then collect at least 10 samples.
Alternate competitor order between rounds, including AB/BA for two-way comparisons.
Run each engine separately rather than concurrently. Parallel scaling is a later,
separately labelled suite.

Record Go and Rust versions, package commits, dirty diff hash, Cargo.lock hash,
resolved native versions, compiler flags, CPU, OS, architecture, thread count,
allocator, enabled ISA/backend, and the manifest hash. Keep all raw samples.

## Metrics and report

Each row identifies pattern, corpus, request size, engine, level, validation policy,
output policy, and backend. Report:

- Median ns/batch and derived ns/request, with sample spread (IQR) and sample count.
- Throughput `actual input bytes / seconds / 1e6` in MB/s (decimal units).
- Unicode scalars/s and pre-token spans/s; use "spans", not BPE "tokens".
- Go B/op and allocs/op with their denominator. Rust allocation measurements, if
  instrumented, are separate runs and include native C allocations; otherwise show
  unavailable rather than zero. RSS is not an allocation-count substitute.
- Speedup `reference time / pretok time` only for identical workload and declared
  comparison level. Values below 1 are regressions and stay in the report.

Keep separate tables for regex-engine and HF-stage ratios. Show all five languages
and every request size before any aggregate. For an equal-byte combined workload,
compute `sum(reference times) / sum(pretok times)` over corresponding batches;
do not average MB/s or individual speedup ratios arithmetically. A geometric mean,
if shown as a secondary summary, must be labelled as such. Report compiler setup
separately; do not amortize it over an arbitrary document count in scanner results.

## Repository layout and first implementation

The maintained benchmark files are:

```text
benchmark/
    corpus/              prepared text pools, request ranges, attribution, manifest
    prepare/             corpus.py, pinned preparation dependencies
    native/              Rust Oniguruma + HF runner, Cargo.toml and Cargo.lock
    runs/                ignored local runs, including smoke and verification
    results/<pattern>/   selected completed runs by date, machine, and backend
cmd/benchreport/          sequential runner orchestration and report generation
```

Keep maintained benchmark sources and required manifests outside ignored
`research/`. Keep runtime Go dependencies free of native/Python benchmark tooling.
Existing unit/oracle tests remain small and offline. Corpus preparation is an
explicit network-enabled step; benchmark execution itself is offline. Missing
corpus data must produce an actionable error in a requested benchmark run, not a
silently empty report.

The implementation includes the five pools, synthetic cases, correctness gate,
public pretok cursor/split rows, direct native Oniguruma iterator/collection and
match-at-offset rows, native HF Split, and compilation rows. SIMD comparisons and
an internal Go match-at core row remain deferred. Existing small Go benchmarks
remain smoke checks and do not replace this cross-engine report.


Follow the [publication policy](../benchmark/results/README.md): retain only
selected completed results in Git, keyed by exact regex identity rather than
model name. All patterns share text pools, request files, and one corpus manifest.
Those files store only inputs. Each run generates reference spans for the chosen
pattern in memory before checking all results and starting timers.
Add any new scenarios to the common corpus. The runner selects P01–P13 or P15–P24 per
invocation. Broader support requires
implementing and validating each additional pattern.

P08/P11/P17/P18/P21/P23 use the same 21 batches. They are partial patterns, so match-at-offset is omitted. Each default complete run contains 1,500 samples (seven scan engines over 21 batches plus three compile engines, ten repetitions each, at least 500 ms per sample).

The P12/P15 profiles are partial grouped alternatives. They share the existing corpus and omit match-at-offset. Their source manifests pin exact decoded strings and HF settings; `-verify-only` generates live reference matches and splits before comparing every request.

P22 uses the shared 21 batches and is partial: its unmatched numeric/control gaps require distinct match-only and isolated Split outputs. The match-at diagnostic is omitted. A default complete run contains 1,500 samples including separate compilation timing.

P10 adds a shared `decimal` workload for successful anchored decimal remainders,
bringing the common corpus to 22 batches. The mixed `synthetic` workload also
includes Nd boundary neighbors, short exhaustive decimal contexts and anchored
failures. These inputs are available to every pattern. P10 is partial, so the
match-at diagnostic is omitted; a full run has 1,570 samples (seven scan engines
over 22 batches plus three compile engines, ten repetitions each). Natural text
mostly exercises rejection and UTF-8 validation; report successful decimal
subjects separately. Previous completed reports retain their original 21 batches.
