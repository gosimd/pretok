# Regexp

`pretok.Regexp` represents one pre-tokenization regular expression compiled into
an immutable, specialized scanner. `pretok` is limited to regex-based boundary
detection; it does not implement a complete Hugging Face pre-tokenizer pipeline.

**Status:** the [Go API](../regexp.go) supports exact corpus strings P01 and P05
through a shared scalar template with immutable word/symbol parameters. The other
23 strings, including equivalent aliases, return `ErrUnsupportedPattern`.
SIMD is not implemented.

## Supported expressions

P05:

```regex
(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+
```

This exact string occurs in Qwen/Qwen2.5-0.5B, among other corpus models. See the
[fixture manifest](../testdata/qwen/manifest.json) for the model revision, original
file hash, JSON pointer, split settings, and oracle binary hash.

P01:

```regex
(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?[\p{L}\p{M}]+|\p{N}| ?[^\s\p{L}\p{M}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+
```

P01 includes Unicode marks in the word body and excludes them from symbol runs.
Its prefix still excludes only CR/LF/L/N, so marks overlap the prefix and body.
For example, `"a\u0301b"` stays one word, and a
standalone combining mark is also a word match. The shared scanner consumes a
leading mark as body: this selects the same endpoint as trying the greedy optional
prefix, then retrying without it if no body remains. Contraction priority is
unchanged: an initial `'s` still matches before a following combining mark.

See [P01 fixtures](../testdata/patterns/P01/manifest.json) for pinned source and
oracle details. Compatibility here is for matches and isolated splits; the source
model's complete pre-tokenizer pipeline is not implemented by this library.

The plan selects the L or L|M table at compilation time. The word and symbol
scanners share that class, with one Unicode lookup and a common ASCII fast path.

The [scalar plan](../internal/scan/plan.go) composes five private construction types:

| Component | Behavior |
| --- | --- |
| `contractionSet` | Standalone apostrophe contractions with pinned case folding |
| `prefixedLetters` | Optional single non-CR/LF/L/N prefix followed by letters (P05) or letters/marks (P01) |
| `singleNumber` | One Unicode N scalar, including Nd, Nl, and No |
| `symbolRun` | Optional ASCII space, symbols outside whitespace/L/N (also M for P01), CR/LF suffix |
| `whitespaceRun` | One whitespace scan and an ordered decision for the last three alternatives |

These types implement P01/P05 constructions; they are not a generic regex IR.
The whitespace scan records its end, the last rune's start, and the end of the
last CR/LF. It chooses the CR/LF cut first, otherwise EOF, otherwise leaves the last
whitespace scalar for the following match when possible. Higher-priority word
and symbol alternatives are always tried first. Cuts never subtract a byte from
a multibyte whitespace character.

Both regexes cover all valid Unicode input, so `Cursor` and `Split` currently emit
the same spans. Their distinct gap-handling contracts remain relevant for future
patterns. There is no normalization, ByteLevel mapping, or BPE processing here.

## API

```go
type Span struct {
    Start int
    End   int
}

func Compile(pattern string) (*Regexp, error)

func (r *Regexp) Cursor(src string) (Cursor, error)
func (c *Cursor) Next() (Span, bool)

func (r *Regexp) Split(src string, dst []Span) ([]Span, error)
```

The name `Regexp` identifies the compiled expression. Scanner algorithms are
implementation details. One `Regexp` represents one expression; composition of
multiple expressions belongs to the surrounding tokenizer.

## Compilation and errors

The first implementation of `Compile` uses an explicit registry mapping exact
regex strings to prewritten scanner plans and immutable parameters. The research
[corpus](../research/pretokenize_regex/all_patterns.json) currently contains 25
unique strings. P01 and P05 are implemented; the remaining strings are future
implementation targets.

**A syntactically valid general-purpose regex can still fail to compile.** Any
string absent from the supported registry returns an unsupported-pattern error.
This includes regexes outside the known pre-tokenization patterns, unregistered
equivalent spellings, and malformed expressions. The first implementation has no
parser and does not distinguish invalid syntax from unsupported semantics.

The lookup key is the exact regex value after the caller decodes its JSON source.
`Compile` performs no additional unescaping, whitespace trimming, AST parsing, or
normalization. Literal CR/LF and regex-escaped CR/LF spellings are different keys;
if both are supported, each must be registered explicitly.

Multiple keys may select one scanner plan when their selected matches have been
verified to be equivalent. Entries are enabled only after the corresponding
scalar behavior has been validated against the pinned reference. Presence in the
research corpus alone does not establish support. Patterns capable of producing
zero-length matches remain outside the initial supported subset.

There is no silent approximation or fallback to a general-purpose runtime regex
engine. Selected match boundaries must preserve alternative priority, greediness,
Unicode properties, and assertions. Shared-prefix decisions are implemented in
the scanner plans rather than discovered by a runtime compiler.

Future versions may introduce parsing and structural recognition without changing
`Compile(pattern) (*Regexp, error)`. Parsing and separate syntax diagnostics are
deferred. Errors are exposed as sentinels and can be checked with `errors.Is`:

- `ErrUnsupportedPattern`: `Compile` did not find the exact string; result is nil.
- `ErrUninitialized`: scanning was called on a nil or zero-value `Regexp`.
- `ErrInvalidUTF8`: scanning input is malformed UTF-8.

Scanning checks initialization before UTF-8. On either error, `Cursor` returns an
exhausted zero cursor and `Split` returns the original destination unchanged,
including its unused backing storage.

## Match iteration

`Cursor` validates the supplied UTF-8 string before iteration and returns an error
for malformed input. `Next` returns non-empty, non-overlapping matches in search
order. Unmatched gaps are skipped.

After exhaustion, `Next` returns `Span{}, false` on every call. The zero-value
cursor is exhausted. A cursor borrows its input until exhaustion or disposal and
holds mutable
iteration state; it must not be used concurrently. Copying a cursor copies its
progress; copies may be iterated independently.

## Splitting

`Split` appends each non-empty match and each maximal unmatched gap to `dst` as
separate pieces. This is isolated splitting with gaps retained. The appended spans
are ordered, non-overlapping, and together cover the exact input bytes.

- Pass `nil` to allocate output as needed, or `dst[:0]` to reuse existing capacity.
- Existing entries in `dst` remain a prefix of the result.
- Empty input adds no spans.
- Malformed UTF-8 returns an error and leaves `dst` unchanged.
- `Split` retains neither the input nor the destination storage.

Use `Cursor` when only matches are wanted. Other Hugging Face split behaviors are
not implied by this API or inferred from the expression.

## Offsets, ownership, and integration

`Span` is a half-open byte range `[Start, End)` into the exact string supplied to
`Cursor` or `Split`. Spans begin and end on UTF-8 character boundaries. They are
not rune indices; the corresponding text is `src[span.Start:span.End]`.

`Regexp` is immutable and safe for concurrent use with independent cursors and
destination buffers. Construct it with `Compile`; its zero value is not usable
for scanning. The pattern registry, execution plans, scratch, and CPU dispatch
stay internal.

When a surrounding tokenizer applies regex stages sequentially, it passes each
previous-stage piece separately to preserve existing boundaries. Offsets are then
relative to that piece; the caller adds its base offset when needed. Normalization,
special-token handling, ByteLevel mapping, BPE merges, and token IDs remain outside
`pretok`.
