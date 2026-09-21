# Compiled pre-tokenizer pipelines

This is the **unreleased public API**, defined in [pretok.go](../pretok.go),
on the `pretok-api` branch. It implements
pre-tokenization, not a complete tokenizer. The former root `Regexp`, `Compile`
and `Cursor` API has moved to [internal/regexp](../internal/regexp); regex functionality is exposed
through Pipeline's Split stage. The 25-string regex registry is unchanged.

```go
func CompileJSON(config []byte) (*Pipeline, error)
func (p *Pipeline) Walk(text string, emit func(Piece) error) error
func (p *Pipeline) WalkText(text string, emit func(string) error) error
func (p *Pipeline) WalkTextUnsafe(text string, emit func(string) error) error
func (p Piece) Text() string
func (p Piece) OriginalRange(start, end int) (Span, error)
```

## Compilation and supported configuration

Pass the decoded JSON **value of `pre_tokenizer`**, not an entire `tokenizer.json`
document. Compile once and reuse the immutable result. JSON `null` and an empty
Sequence mean identity: emit the non-empty input once. Empty input always emits
nothing. A nil or zero Pipeline is uninitialized.

| Type | Supported settings and defaults |
| --- | --- |
| `BertPreTokenizer` | Remove whitespace, isolate each ASCII/Unicode punctuation character |
| `CharDelimiterSplit` | Required `delimiter`: one Unicode scalar, removed from output |
| `Digits` | `individual_digits: false`; HF numeric class `Nd/Nl/No` |
| `Metaspace` | `replacement: "▁"`, `prepend_scheme: "always"`, `split: true` |
| `Punctuation` | `behavior: "Isolated"`; ASCII punctuation plus HF Unicode punctuation |
| `Sequence` | Required `pretokenizers` array; nested sequences are supported |
| `UnicodeScripts` | HF's pinned script profile and Any/Han adaptations |
| `Whitespace` | Rust-regex word/non-word runs, whitespace discarded |
| `WhitespaceSplit` | Split on HF whitespace, discard whitespace |
| `Split` | Required `pattern`; `behavior: "Isolated"`, `invert: false` |
| `ByteLevel` | `add_prefix_space: true`, `trim_offsets: true`, `use_regex: true` |

Split patterns must be exactly one of `{"Regex":"..."}` and
`{"String":"..."}`. Regex accepts only the [25 registered strings](SupportedPatterns.md).
String is a literal, with no regex interpretation, and must be non-empty.
All five behaviors are supported: `Removed`, `Isolated`, `Contiguous`,
`MergedWithPrevious`, `MergedWithNext`. Names are case-sensitive. `invert` flips
match/gap roles before applying the behavior.

Metaspace replaces U+0020 only. It supports `always`, `first`, and `never` prefix
schemes and `split: false`. `first` depends on the original source position, even
after earlier transformations and with either text-only method. Legacy `add_prefix_space`
and redundant `str_rep` Metaspace fields are not accepted; use the current schema.

ByteLevel's built-in GPT-2 scanner is internal and does not register another
accepted spelling for regex Split stages. `trim_offsets` is accepted and checked as a
boolean but **does not run Encoding postprocessing**. OriginalRange returns
pre-tokenization alignment without trimming original whitespace.

Compilation rejects malformed JSON, invalid UTF-8, unpaired surrogate escapes,
duplicate keys, unknown fields, wrong field types, and unsupported settings.
Unspecified documented options use the defaults above; this deliberately accepts
minimal objects in addition to HF's serialized objects. Nesting is limited to
32 Sequence levels and execution to 256 flattened stages. Null is only accepted
as the whole configuration, not as a Sequence member. These are explicit library
limits, not claims that every JSON accepted by HF is accepted here.

`errors.Is` identifies `ErrInvalidConfig`, `ErrUnsupportedPreTokenizer`, or the
existing `ErrUnsupportedPattern`. Errors include the configuration path where
possible. Regex spelling is never trimmed, normalized, or unescaped a second time.

## Execution and ownership

Walk invokes the callback once per **final** piece, in order, after all stages.
Stages act independently within previous-stage boundaries. They never join across
these boundaries. Original source ranges can have gaps (removed separators) and
can overlap (multiple output pieces derived from one input character).

Pipeline supports concurrent calls with separate execution state, including
reentrant calls made from a callback. The complete input is validated before the
first callback. Error precedence is `ErrUninitialized`, `ErrNilCallback`, then
`ErrInvalidUTF8`. Callback errors are returned unchanged, execution stops
immediately, and no later callbacks run. Prior consumer side effects cannot be
rolled back. Panics propagate; private workspace cleanup still executes.

**Piece, its Text string, and its alignment are valid only during that callback.**
Keeping a copy of Piece does not extend its lifetime. To retain text, call
`strings.Clone`; resolve source offsets before returning. Callbacks must not access
these views asynchronously after return. The library does not store source strings
or callbacks after completion.

**WalkText strings may be retained and used after the callback returns**, including
in asynchronous processing. This method defensively copies each emitted string.

**WalkTextUnsafe avoids defensive copying, but its strings are valid only until
the callback returns.** Subsequent execution may overwrite their backing memory.
Do not retain them or pass them to asynchronous processing without first copying
them with `strings.Clone` inside the callback.

Both text-only methods perform the same transformations and emit identical texts,
without building the alignment map for the consumer. They retain only the stage
context needed for decisions such as Metaspace `first`. A BPE consumer can use
WalkTextUnsafe to process each string synchronously inside its callback. No BPE
dependency or integration is added by this API.

## Offset mapping

All positions are UTF-8 **byte offsets**. OriginalRange maps a non-empty range in
Piece.Text to the smallest covering source range in the original Walk input.
Endpoints within a multibyte encoding are allowed: the result covers complete
original scalars, matching HF's per-byte character alignment. Empty, reversed,
negative, and out-of-bounds ranges return `ErrInvalidRange`; the zero Piece has
no valid ranges. Empty-range behavior is intentionally not HF's low-level
NormalizedString empty-range convention.

For `ab12é`, Split `[0-9]` followed by ByteLevel without prefix or regex emits:

| Text | Original byte range |
| --- | --- |
| `ab` | `[0,2)` |
| `1` | `[2,3)` |
| `2` | `[3,4)` |
| `Ã©` | `[4,6)` |

Both `[0,2)` and `[2,4)` inside `Ã©` map to `[4,6)` in the original string.
Added prefix characters map to the first original character of their input piece,
as in HF. Transformations compose this mapping through every subsequent stage.

The source referential is exactly the string supplied to Walk. If the caller
normalizes text first, offsets refer to that prepared text. This library neither
normalizes text nor restores positions before an external normalization step.

## Memory and performance

Execution streams final pieces through callbacks and does not build a complete
array of output or intermediate pieces. Each active transformation stage may
buffer its entire current input piece, so this is not constant-memory streaming
over chunks of a source document. UTF-8 validation also examines the full input.

Identity and boundary-only stages use source substrings with arithmetic offsets.
Transforming stages reuse private byte buffers; Walk currently stores a source
range for each transformed UTF-8 byte. Neither text-only method constructs those
maps. WalkText copies final text even when it comes from a source substring.
Buffers are exclusive to a call, including nested calls. A private sync.Pool
reuses workspaces; workspaces exceeding 1 MiB of buffer capacity are discarded
after the call. Pools may be cleared by GC and do not bound total process memory.
No public Scratch is exposed. Walk and WalkTextUnsafe expose zero-copy string
views under the callback lifetime contract above.

No unconditional zero-allocation guarantee is made. Buffer growth, cold calls,
pool eviction and large transformations can allocate. WalkText additionally
allocates a defensive copy for each emitted piece. Run the included compile and
execution benchmarks separately: Walk maps offsets, WalkText consumes owned text,
and WalkTextUnsafe consumes borrowed text. Their ownership and work differ; the
text-only comparison measures the cost of defensive copying. There are no new
SIMD performance claims.

```sh
GOWORK=off go test -mod=readonly ./...
GOWORK=off go test -mod=readonly -race ./...
GOWORK=off go test -mod=readonly -run '^$' -bench 'Benchmark(Pipeline|CompileJSON)' -benchmem -count=5 .
```

## Compatibility evidence and exclusions

The [offline oracle fixtures and generator](../testdata/pipeline/README.md) pin
HF Tokenizers 0.23.2. They compare all three walking methods, whole-piece ranges and
internal ranges through HF's token-offset conversion. Built-in Unicode classes
come from those exact components; they are separate from the existing HF 0.22.2
regex profile. Compatibility is scoped to the documented settings and evidence.

### Scenario coverage

The existing [oracle test](../pipeline_test.go) runs all **4,475 cases** through
Walk, WalkText and WalkTextUnsafe. It compares transformed text, source ranges,
internal byte-range mappings and retained WalkText strings against the recorded
HF 0.23.2 expectations. Names follow
`TestPipelineHFOracle/<component>/<settings>/<case>`.

The table counts top-level fixture configurations; component occurrences inside
Sequence are additional coverage, not included in the individual component rows.
Counts describe existing evidence, not exhaustive coverage of every input.
Test groups below are relative to `TestPipelineHFOracle/`; benchmark groups are
relative to [`BenchmarkPipeline/`](../pipeline_bench_test.go).

| Scenario / test group | HF cases / distinct configurations | Settings covered by oracle fixtures | Existing execution benchmarks |
| --- | --- | --- | --- |
| `BertPreTokenizer/` | 31 / 1 | Default behavior | None |
| `CharDelimiterSplit/` | 93 / 3 | Delimiters `x`, `é`, U+0020 | None |
| `Digits/` | 62 / 2 | `individual=false` and `true` | None |
| `Metaspace/` | 559 / 18 | Three prepend schemes × two split modes × replacements `▁`, `_`, `é` | `Metaspace/prepend=always,split=true,replacement=U+2581/` only |
| `Punctuation/` | 155 / 5 | All five split behaviors | None |
| `Sequence/` | 1,310 / 843 | Empty sequence, nine fixed compositions, 1,000 random compositions; no nested Sequence oracle cases | `Sequence/Split_ByteLevel-9c89680975f8/` only |
| `UnicodeScripts/` | 31 / 1 | Default behavior | None |
| `Whitespace/` | 31 / 1 | Default behavior | `Whitespace/default/` |
| `WhitespaceSplit/` | 31 / 1 | Default behavior | None |
| `Split/` | 2,015 / 63 | All 25 registered patterns in isolated mode; five behaviors × invert for two regexes and two literals | Only inside the sequence above |
| `ByteLevel/` | 126 / 6 | Four prefix/regex combinations with trim=true; two long cases with trim=false | `ByteLevel/prefix=true,trim=true,regex=true/` only |
| `Identity/` | 31 / 1 | JSON null | None |

Every listed execution benchmark has the same four existing inputs
(`short_ascii`, `short_mixed`, `long_ascii`, `long_mixed`) and three methods
(`Walk`, `WalkText`, `WalkTextUnsafe`), giving 48 measured combinations.
Compilation has one existing benchmark:
`BenchmarkCompileJSON/Sequence/Split_ByteLevel-9c89680975f8`.
Renaming these cases does not add workloads or change their timed operations.

Examples of individual oracle cases:

```text
TestPipelineHFOracle/BertPreTokenizer/default/punctuation_symbols_emoji
TestPipelineHFOracle/CharDelimiterSplit/delimiter=U+00E9/repeated_accent
TestPipelineHFOracle/Digits/individual=true/mixed_number_categories
TestPipelineHFOracle/Sequence/Split_ByteLevel-9c89680975f8/letters_digits_accent
```

Select scenarios from the repository root:

```sh
GOWORK=off go test -mod=readonly -run '^TestPipelineHFOracle/Digits/' -v .
GOWORK=off go test -mod=readonly -run '^TestPipelineHFOracle/Metaspace/' .
GOWORK=off go test -mod=readonly -run '^$' -bench '^BenchmarkPipeline/Whitespace/' -benchmem .
```

The [fixture guide](../testdata/pipeline/README.md) defines the naming scheme and
regeneration process. Full configuration and input remain available in failure
output, including for sequences identified by a short settings hash.

Existing gaps remain visible: most standalone component inputs are at most 28
UTF-8 bytes, with long fixtures only for Metaspace and ByteLevel. Execution
fuzzing covers three configurations; concurrency, cancellation and text ownership
checks are concentrated on the Split/ByteLevel example. Six of the nine core
built-ins have no dedicated execution benchmark. The broader
[regex benchmark suite](../benchmark/README.md#regex-cross-engine-suite) does not
substitute for those Pipeline measurements.

### Integration limits

No normalization, AddedToken matching, token IDs, BPE, decoder, postprocessor,
Hub access, or full tokenizer.json loading is provided. In particular, calling
Walk separately on segments around AddedTokens resets the source context. A
future integration needs a separate segment-context contract for original positions
and prefix behavior; this API does not claim that integration.
