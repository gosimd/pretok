# Regexp

`pretok.Regexp` represents one pre-tokenization regular expression compiled into
an immutable, specialized scanner. `pretok` is limited to regex-based boundary
detection; it does not implement a complete Hugging Face pre-tokenizer pipeline.

**Status:** the [Go API](../regexp.go) supports the 25 exact strings in the
[supported-pattern registry](SupportedPatterns.md), which lists their IDs,
expressions and tokenizer examples. Unregistered strings, including equivalent
spellings, return `ErrUnsupportedPattern`. P01 `Split` uses bounded rune decoding;
`GOEXPERIMENT=simd` enables supported SIMD UTF-8 kernels from gosimd/unicode.
Character classification and token-boundary selection remain scalar.

## v0.1.0 release contract

This document fixes the public contract for **v0.1.0**. The API, compilation,
iteration, splitting, and
ownership sections below are normative. Pattern-specific sections describe the
selected-match semantics. The release's [supported-pattern registry](SupportedPatterns.md)
and [pinned profiles](../testdata/patterns/README.md), read at the release revision,
define its exact 25 supported strings and reference behavior.

The public surface consists of `Regexp`, `Compile`, `Cursor`, `Cursor.Next`,
`Span`, `Regexp.Cursor`, `Regexp.Split`, and the three error sentinels
`ErrUnsupportedPattern`, `ErrUninitialized`, and `ErrInvalidUTF8`. `Cursor` names
both a type and a method of `Regexp`; there is no package-level `Cursor` or
`Split` function. Callers should use `errors.Is` to identify errors; their text
is diagnostic, not a compatibility key.

The observable guarantees are:

- Exact lookup of the caller's JSON-decoded pattern, with no normalization,
  syntax diagnostics, approximation, or fallback engine.
- Non-empty half-open UTF-8 byte spans into the supplied input, preserving
  ordered alternatives, the pinned character classes, and search assertions.
- Full UTF-8 validation before scanning, initialization errors taking priority,
  and an unchanged destination (including backing storage) on `Split` errors.
- Match-only iteration through `Cursor`; isolated matches and maximal unmatched
  gaps through `Split`. Empty input produces no matches and appends no spans.
- Append semantics and caller-owned destination reuse. Existing destination
  entries are preserved; appended offsets are relative to the current input.
- Immutable compiled expressions, independent cursor copies, and concurrent
  scans with separate cursors and destination buffers. A cursor borrows its
  input; `Regexp` and `Split` retain neither input nor destination buffers.

For `gosimd/bpe`, use `Split` when the next stage must receive all input pieces.
Using only `Cursor` would omit unmatched gaps. Regex support does not infer or
implement arbitrary HF `Split` behaviors, `invert` settings, or a Sequence.
ByteLevel mapping, special tokens, normalization, BPE and token IDs remain the
consumer's responsibility. The existing signatures are sufficient for this
boundary; no full pre-tokenizer pipeline is part of v0.1.0.

### Compatibility policy for v0.1.x

Although the module is pre-v1, the project adopts the following patch-release
policy:

- Preserve public signatures, sentinel error meanings, ownership, append/error
  guarantees, and the accepted exact-string set. New pattern registrations or
  intentional contract changes require a new minor version and migration notes
  where relevant.
- Preserve selected matches and split boundaries for supported strings.
  Refactoring or optimization must not silently change them. Do not update
  Unicode semantics implicitly with a Go toolchain or reference-engine upgrade.
- A correction of an implementation defect against the pinned contract may be
  released as a patch. If it changes output boundaries, release notes must name
  the affected patterns and show before/after input and spans, backed by a pinned
  reference regression test. An intentional new interpretation is not a patch
  correction.
- Internal scanner plans, cursor layout, and algorithms are implementation
  details. Benchmark timings and allocation counts are measured properties,
  not fixed numerical API guarantees; performance changes still require the
  project's validation and regression reporting.

## Scanner semantics

[P05 expression](SupportedPatterns.md#p05).

This exact string occurs in Qwen/Qwen2.5-0.5B, among other corpus models. See the
[fixture manifest](../testdata/qwen/manifest.json) for the model revision, original
file hash, JSON pointer, split settings, and oracle binary hash.

P09 and P13 are separately registered exact aliases of P05, with their own
[source profiles](../testdata/patterns/README.md) and live HF correctness gates.
P09 contains actual CR/LF characters in its decoded string where P05 uses regex
escapes. It must be passed unchanged after JSON decoding; no normalization occurs.
Its source uses Removed/invert=true, while the benchmark also verifies isolated
splits. Both outputs coincide for this fully covering expression.

[P13 expression](SupportedPatterns.md#p13).

For this exact shape the possessive prefix is disjoint from the word body, and
the possessive symbol body is disjoint from its CR/LF suffix. Giving characters
back cannot enable a match. The greedy `\s*[\r\n]` selects the same endpoint as
P05's `\s*[\r\n]+`: after the last CR/LF in the whitespace run. The factored
contractions retain the same pinned case folding and selected boundaries.
These facts justify sharing one immutable plan, without a general rule for
removing possessive quantifiers or registering other equivalent spellings.

[P01 expression](SupportedPatterns.md#p01).

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

The [scalar plan](../internal/scan/plan.go) composes the following private construction types:

| Component | Behavior |
| --- | --- |
| `contractionSet` | Standalone apostrophe contractions with pinned case folding |
| `prefixedLetters` | Optional single non-CR/LF/L/N prefix followed by letters (P05) or letters/marks (P01) |
| `singleNumber` / `numberRun` | One Unicode N scalar or a greedy run with scalar-count bounds |
| `symbolRun` | Optional ASCII space, symbols outside whitespace/L/N (also M for P01), CR/LF suffix |
| `whitespaceRun` | One whitespace scan and an ordered decision for the last three alternatives |

These types implement P01/P05 constructions; they are not a generic regex IR.
The whitespace scan records its end, the last rune's start, and the end of the
last CR/LF. It chooses the CR/LF cut first, otherwise EOF, otherwise leaves the last
whitespace scalar for the following match when possible. Higher-priority word
and symbol alternatives are always tried first. Cuts never subtract a byte from
a multibyte whitespace character.

P01/P05 cover all valid Unicode input, so their `Cursor` and `Split` outputs
coincide. [P02](SupportedPatterns.md#p02) exercises their distinct gap-handling contracts.

### P01 buffered splitting

On eligible SIMD builds, for P01 inputs of at least 256 bytes, `Cursor` creation validates the whole input
using `gosimd/unicode/utf8.ValidString`. `Split` then uses
[`QwenMarksCursor`](../internal/scan/rune_words.go) and `DecodeValidInto` to fill a
local 1,024-rune buffer without validating the same bytes again. The buffer stays
within the `Split` call; it is not stored in the compiled plan or public cursor.
No pool or caller-managed scratch is required. Destination allocation still
depends on caller capacity.

The buffer boundary is independent of token boundaries. Contraction lookahead
survives refills, and whitespace rollback reuses the current decoded window when
possible. Byte positions are reconstructed from rune widths. The pinned character
tables, contraction predicates and whitespace endpoint decision are shared with
the byte scanner. Full-input validation precedes every destination write, so a
malformed suffix leaves even unused destination backing storage unchanged.

Public `Cursor.Next` retains byte scanning and independent value-copy progress.
Short P01 inputs and other patterns retain their byte scanners. The Unicode
dependency selects SIMD kernels only in supported `GOEXPERIMENT=simd` builds;
other configurations retain byte scanning. The buffer is selected only when
Unicode has SIMD conversion available (ARM64 or AMD64 with AVX2). Direct internal
tests also exercise the buffered algorithm with portable conversion.

P02's `asciiTriplet` matches exactly three consecutive ASCII digits. It groups a
long digit run from the left, without consuming an incomplete group. Unicode
numbers outside ASCII do not match. The existing cursor searches past unmatched
scalars; the existing `Split` preserves each maximal gap. For `x123456y12`:

- `Cursor`: `[1,4)`, `[4,7)` — `123`, `456`.
- `Split`: `[0,1)`, `[1,4)`, `[4,7)`, `[7,10)` — `x`, `123`, `456`, `y12`.

`12a34` has no matches and produces one isolated piece. Empty input produces none.
The equivalent spelling `[0-9]{3}` is deliberately unregistered.
[P02 source metadata](../testdata/patterns/P02/manifest.json) pins its occurrence
in tiiuae/falcon-7b. Live HF references use Removed/invert=true for matches and
Isolated/invert=false for splits; no new output archive is stored.

Numeric variants:

| Pattern | Exact expression or change from P05 | Selected matches |
|---|---|---|
| [P16](../testdata/patterns/P16/manifest.json) | `[0-9]` | One ASCII digit |
| [P07](../testdata/patterns/P07/manifest.json) | `\p{N}{1,3}` | Greedy groups of 1–3 Unicode N scalars |
| [P24](../testdata/patterns/P24/manifest.json) | Replace only P05's `\p{N}` branch with `\p{N}{1,3}` | Qwen boundaries with bounded numeric groups |
| [P20](../testdata/patterns/P20/manifest.json) | Replace only P05's `\p{N}` branch with `\p{N}+` | Qwen boundaries with whole numeric runs |

P16/P07 use a numeric search plan: ASCII byte search is safe after UTF-8
validation, and Unicode search decodes each skipped scalar once. They leave unmatched gaps; `Cursor` skips them and `Split` preserves them.
P24/P20 cover all valid input. The internal `numberRun` stores a numeric class and
minimum/maximum scalar counts (`max == -1` means unbounded). ASCII `[0-9]`,
Unicode `Nd`, and Unicode `N` are not interchangeable: the latter includes `Nd`,
`Nl`, and `No`, so `１²Ⅳ١𐒠3` is two P07/P24 groups and one P20 group. P16
matches only its final ASCII `3`. Bounds count scalars, while spans remain byte
indices. Adjacent numbers from different scripts belong to the same run.

Qwen variants share every non-numeric construction and preserve alternative order.
Single-number plans retain their fast classifier; only longer numeric runs enter
the continuation loop. P02 retains its fixed three-ASCII-byte specialization.
Equivalent spellings such as `\d` or `[0-9]{1}` remain unsupported. Numeric
source manifests contain provenance/settings only; HF outputs are regenerated
in memory on the shared benchmark inputs before every full correctness gate.

P03:

```regex
\s+$
```

The pinned HF/Oniguruma profile treats `$` as a line-end assertion: it matches
before **any LF**, as well as at EOF. CR, NEL, U+2028 and U+2029 do not themselves
satisfy the anchor. Whitespace classification reuses the existing pinned tables.

`\s+` is greedy: a whitespace run reaching EOF is consumed whole, including LF.
Otherwise the match ends just before the last LF in that run, provided at least
one whitespace scalar precedes that LF. Examples:

| Input | Matches | Isolated pieces |
|---|---|---|
| `"a \nX"` | `" "` | `"a"`, `" "`, `"\nX"` |
| `"a\nX"` | none | `"a\nX"` |
| `"a \n\nX"` | `" \n"` | `"a"`, `" \n"`, `"\nX"` |
| `"a \n"` | `" \n"` | `"a"`, `" \n"` |
| `"a\r\nX"` | `"\r"` | `"a"`, `"\r"`, `"\nX"` |

The internal plan searches by whole whitespace runs. Shared run metadata records
the last LF position; failed runs are skipped entirely, keeping search linear.
The cursor delegates search to its immutable plan, retaining only its own position.
Each caller-supplied fragment defines a separate subject and its own EOF.

[P03 source metadata](../testdata/patterns/P03/manifest.json) pins its DeepSeek
occurrence and reference settings. `Cursor` skips gaps; `Split` preserves maximal
gaps. No matches are empty. `\s+\z` and `(?m:\s+$)` remain unregistered strings.

P04:

```regex
[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+
```

P04 orders two word alternatives: `prefix? A*B+`, then `prefix? A+B*`, where
`A = Lu|Lt|Lm|Lo|M`, `B = Ll|Lm|Lo|M`, and the optional prefix excludes CR/LF/L/N.
It has no contraction alternative or suffix; numbers are single Unicode N scalars.
Its symbol body excludes whitespace/L/N (marks remain symbols), and its greedy
suffix allows any mixture of CR, LF, and `/`.

The [word scanner](../internal/scan/case_words.go) records the end of the A run,
the last member also belonging to B, and the following B run. These observations
produce both candidate endpoints without rescanning the word for each alternative.
The first successful alternative wins, even if the second could consume more:
`AA中BB` becomes `AA中` + `BB`. The optional prefix must backtrack within the first
alternative before the second is tried: a leading combining mark followed by
`ABC` becomes the mark + `ABC`. This is not ordinary camelCase splitting.

The [P04 profile](../testdata/patterns/P04/manifest.json) pins its Nemotron source,
HF wheel and generated class-table hashes. Classes come from the pinned HF oracle
over every Unicode scalar; the runtime uses one class-mask lookup. Common
number, symbol, whitespace, cursor and gap handling are reused. P04 covers all
valid scalars, so matches and isolated splits coincide. References are generated
in memory on the common benchmark corpus, with no P04 output archive.

There is no normalization, ByteLevel mapping, or BPE processing here.

P19 and P06 extend the two ordered P04 word alternatives with the same optional
`(?i:'s|'t|'re|'ve|'m|'ll|'d)?` suffix on each. P19 retains single Unicode N
scalars; P06 uses N{1,3}. All other classes, alternative order, slash/CR/LF symbol
suffixes, and whitespace decisions remain those of P04. Exact strings and source
settings are pinned in the [P19](../testdata/patterns/P19/manifest.json) and
[P06](../testdata/patterns/P06/manifest.json) manifests.

Select the P04 word endpoint first, then try one existing `contractionSet` match
at that endpoint. The suffix is optional and its apostrophe cannot occur in the
word body, so it neither forces backtracking to an earlier body endpoint nor
justifies switching to the later word alternative. `we're` is one match;
`AA中BB's` is `AA中` plus `BB's`. A leading mark still participates in P04's
first-alternative prefix backtracking: `\u0301ABC's` becomes the mark plus `ABC's`.

Contractions attach only to words. `1's` remains `1` plus `'s`; `!'s` remains
`!'` plus `s`. Only one suffix attaches: `A'S'S` becomes `A'S` plus `'S`.
The existing scoped case-fold matcher handles `a'ſ` as one match. Both patterns
cover all valid UTF-8, so match-only and isolated spans coincide. This establishes
regex-stage compatibility with the pinned profile, not a complete HF pipeline.
No new Unicode tables or regex parsing are needed. P04 retains its existing
runtime entry point; the suffix plan reuses its word matcher and shared tail.

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
[corpus](../testdata/corpus_patterns.json) contains 25 unique strings. The
[supported-pattern registry](SupportedPatterns.md) records all 25 implemented
entries. Other exact strings remain unsupported.

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

## Exact literal sets: P08, P11, P17, P18, P21, P23

All six are partial matchers: `Cursor` skips gaps, while `Split` preserves maximal
unmatched gaps as separate spans. They use the same public contracts and UTF-8
validation. These literal sets do not track changes in Unicode letter categories.
Only `\s` membership uses the pinned reference whitespace table.

| Pattern | Body | Optional prefix |
| --- | --- | --- |
| P11 | One CR or LF; CRLF produces two matches | None |
| P21 | U+3040–30FF and U+4E00–9FA5, one maximal union run | None |
| P17 | U+0800–9FA5 and U+AC00–D7FF, including U+2003 | None |
| P18 | The exact punctuation/alphabet ranges in its manifest, including ASCII letters and U+3000 | One Unicode whitespace scalar |
| P08 | The exact extended alphabet ranges in its manifest, including supplementary-plane ranges | One Unicode whitespace scalar |
| P23 | Complement of whitespace, `(`, `)`, `|`, and `.,!?…。，、।۔،` | One ASCII space |

P23 follows the pinned Oniguruma nested-class interpretation. `[abc]` is one
match; `(abc)` matches only `abc`; hyphen, colon and semicolon are allowed.
P18's U+3000 belongs to both prefix and body. A singleton still matches because
its optional prefix can backtrack. The shared scanner consumes the maximal body
run and optionally includes the immediately preceding prefix; this preserves the
same endpoint even when the prefix belongs to the body. Two leading whitespace
scalars outside the body cannot both be attached.

The implementation stores sorted exact range unions and precomputes ASCII
membership once when constructing immutable plans. It searches gaps directly.
P11 can search bytes because UTF-8 continuation bytes cannot equal CR or LF.
There is no runtime parser, regex fallback, normalization, or new public API.
Each exact decoded expression and immutable source is in its
`testdata/patterns/PXX/manifest.json`; the common benchmark gate generates HF
matches and isolated splits live before timing.

## Ordered group runs: P12 and P15

[P12](../testdata/patterns/P12/manifest.json) contains six ordered alternatives:
CJK/kana ranges, Thai, Lao, Khmer, Myanmar ranges, and Hangul ranges. Each branch
has its own `+`. `中ก` therefore produces two adjacent matches, while `中あカ`
remains one. P12 uses its written ranges, not Unicode Script properties.

[P15](../testdata/patterns/P15/manifest.json) orders Hangul, a literal punctuation
set, Bopomofo, and P21's exact CJK/kana set. Its Hangul Script property is pinned
to HF Tokenizers 0.22.2; `benchmark/prepare/grouped.py` verifies its 14 ranges by
extracting every matching Unicode scalar from that reference. Go's current
Unicode tables do not determine this contract.

The first two P15 groups overlap at U+302E and U+302F. On the first scalar, the
scanner selects the first matching alternative, then keeps that exact set for
the whole run. Thus `〮、` gives `〮`, `、`, but `、〮` stays together;
`、〮가` gives `、〮`, `가`. A scalar does not have one permanent group ID.
Testing the selected set directly preserves overlaps without a membership mask.

The shared `groupedRun` uses existing `exactSet` membership and the same run
consumption loop as plain sets. Plans are immutable; group selection lives only
in the current search call. Both patterns can leave gaps: Cursor skips them and
Split preserves maximal gaps. UTF-8 validation, byte offsets, destination reuse,
and the public API are unchanged. Only the two exact decoded strings compile.

## P22: punctuation-prefixed words and positive symbol classes

The [P22 manifest](../testdata/patterns/P22/manifest.json) pins the exact decoded
string, including literal CR/LF, and immutable tokenizer source. Its order is:

1. One explicitly listed ASCII punctuation/symbol followed by ASCII letters.
2. An optional scalar outside CR/LF/L/P/S followed by one or more L|M scalars.
3. Optional ASCII space, a nonempty P|S run, then zero or more CR/LF bytes.
4. The existing ordered newline/trailing-whitespace/general-whitespace branches.

There is no contraction or numeric branch. `!abcé` matches `!abc`, then `é`;
`123abc` matches only `3abc` and leaves `12` as a gap. `!!abc` matches `!!`, then
`abc`: the first branch does not restart inside a symbol run. An initial mark can
belong to prefix and word body; a singleton mark still matches. Digits, format
characters and other controls can prefix a word, while P|S characters cannot.

`punctuationWordPlan` reuses the pinned L|M table, `exactSet` consumption and the
shared whitespace decision. P|S is a positive union, not the complement of L/N
or whitespace. Its 349 ranges are generated from HF Tokenizers 0.22.2 by
`testdata/patterns/generate_punctuation_tables.py`. Go 1.27's Unicode 17 P/S union
differs from this pinned oracle on 104 scalars, so current standard-library
categories cannot replace this table without changing the contract.

Cursor skips unmatched gaps and Split preserves maximal gaps. Public API,
UTF-8 validation, byte offsets and destination ownership are unchanged.

## P10: anchored decimal remainder

[P10](SupportedPatterns.md#p10) selects a leading group of one or two Nd scalars
only when the entire subject consists of that prefix and one or more triples.
`1234567` has one match, `1`; `123456`, `x1234` and `1234` followed by LF have none.
The lookahead leaves the rest as a single unmatched Split gap. See
[contextual decimal plans](ContextualNumbers.md) for search context, the pinned
meaning of `{3}+`, Unicode membership P14 continuation and P25 right-aligned groups.

## P14: consecutive decimal triples

[P14](SupportedPatterns.md#p14) consumes complete Nd triples starting at byte zero
and continuing only at the previous match's end. `123x456` yields only `123`;
Split retains `x456` as one gap. Failed partial triples never trigger a search
for later digits. The plan reuses P10's bounded decimal scanner, pinned class
and existing cursor exhaustion. See [continuation semantics](ContextualNumbers.md#p14-semantics).

## P25: right-aligned decimal groups

[P25](SupportedPatterns.md#p25) groups Nd runs from the right in blocks of 510
scalars: 512 digits become 2 + 510, and 1021 become 1 + 510 + 510. It searches
again after gaps. An internal, value-owned cursor state caches the byte end of
the current series, keeping the scan linear and cursor copies independent.
See [P25 semantics](ContextualNumbers.md#p25-semantics) for alternative priority,
state ownership and the shared pinned decimal class.
