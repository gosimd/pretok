# Contextual decimal patterns

P10, P14 and P25 share the Unicode `Nd` class, but need different search rules.
All three exact strings are implemented. The public `Regexp`, `Cursor` and `Split` contracts stay unchanged.

## Context and ownership

The existing internal `Plan.Find(src, searchStart, state)` receives the entire supplied
subject and the start of the current search: zero initially, then the previous
match's end. A plan must preserve this value while trying candidate positions.
It may skip gaps internally or fail immediately. Failure exhausts the cursor;
the generic cursor must not advance and restart an anchored search.

`\A` refers to byte zero of `src`, `\z` to its exact end, and `\G` to the start
of the current search under the pinned reference. Do not implement these anchors
by slicing `src[searchStart:]`. Each call on a prior tokenizer stage's piece is
a new subject; retaining boundaries between Sequence stages belongs to the caller.

| Pattern | Search rule | Per-input data |
| --- | --- | --- |
| P10 | Only match at byte zero after checking the entire subject | Decimal run end and scalar count, local to the first search |
| P14 | Consume three Nd scalars at the current search start; stop on failure | Existing cursor position is sufficient |
| P25 | Find a decimal run and select groups aligned from its right end, preserving alternative priority | Cache the run end in cursor-owned internal state to avoid rescanning the suffix for every group |

The reusable scalar component is `scanDecimals(src, pos, limit)`, returning
`decimalRun{End, Count}`. `End` is a byte boundary; `Count` counts Unicode scalars.
A zero limit scans the whole run; a positive limit consumes at most that many
digits. It distinguishes `Nd` from both ASCII `[0-9]` and the broader `N` class.

P10 and P14 use local metadata. P25 additionally caches its current run end in
`scan.State`, an internal value owned by each cursor. Search plans receive state
by value and return its successor with the match boundaries. Stateless plans
keep their existing matching functions and leave the state unchanged.

Passing state by value avoids making a cursor escape through a function pointer
that accepts its address. It also preserves independent progress when a cursor
is copied. The shared compiled plan owns no mutable state. One integer is added
to each cursor; it contains no input or buffer references. Public signatures are
unchanged, and Cursor/reused Split must continue to allocate no memory.

## P10 semantics

The [exact P10 string](SupportedPatterns.md#p10) is interpreted using pinned HF
Tokenizers 0.22.2 and the native Oniguruma reference. In this syntax, `\p{Nd}{3}+`
repeats triples, equivalent here to `(?:\p{Nd}{3})+`; it does **not** mean a
possessive exact-three quantifier. The equivalent spelling is not registered.
This behavior was checked on 4, 5, 6, 7 and 8 digit subjects before implementation.

The plan scans the decimal run at zero. A match exists only when it reaches exact
EOF, contains at least four scalars, and its count modulo three is one or two.
The match consumes only that leading remainder. The lookahead consumes nothing,
so `Split` returns the prefix and one gap containing all remaining triples.
Subsequent cursor searches fail immediately because `\A` no longer holds.

| Input | Cursor matches | Isolated Split pieces |
| --- | --- | --- |
| `1234` | `1` | `1`, `234` |
| `12345` | `12` | `12`, `345` |
| `1234567` | `1` | `1`, `234567` |
| `123456` | none | `123456` |
| `x1234` or `1234` followed by LF | none | the whole input |
| `１٢𝟛4٥` | `１٢` (bytes 0–5) | `１٢`, `𝟛4٥` |

The scan is linear in the leading decimal run, with constant local storage.
Only the first one or two scalars are decoded again to locate the prefix endpoint.
Public entry points still validate all UTF-8 input before scanning, even on an
immediate anchor/class failure.

## P14 semantics

The [exact P14 string](SupportedPatterns.md#p14) matches three Nd scalars at the
current search start. On the first call, both alternatives select byte zero.
Later calls continue at the previous match's end through `\G`. If fewer than
three consecutive decimal digits remain at that position, iteration ends. A
later digit run after a gap must never restart this search.

`ConsecutiveDecimalTriples` calls the existing `scanDecimals(src, searchStart, 3)`.
A complete triple returns its byte boundaries; an incomplete triple returns the
existing search-failure sentinel. The cursor owns continuation and exhaustion,
while Split preserves the entire remaining suffix as one gap.

| Input | Cursor matches | Isolated Split pieces |
| --- | --- | --- |
| `1234567` | `123`, `456` | `123`, `456`, `7` |
| `123x456` | `123` | `123`, `x456` |
| `1234x567` | `123` | `123`, `4x567` |
| `x123456` | none | `x123456` |
| `１２٣٤𝟝67` | `１２٣` (bytes 0–8), `٤𝟝6` (8–15) | those two triples and `7` |

This differs from P02/P07, which search past unmatched gaps, and from P10, which
checks exact EOF and returns only a leading remainder. P14 does not require that
the whole subject consists of digits. No new Unicode tables, mutable plan state,
public methods or cursor fields are introduced.

[P14 provenance](../testdata/patterns/P14/manifest.json) pins its source and the
hash of the shared P10 Unicode profile. Offline tests include captured HF byte
boundaries, an independent RE2 match list restricted to its contiguous prefix,
all-scalar contexts, partial triples, EOF, cursor copies and failed continuation.
The live benchmark gate compares HF matches and isolated splits before timing.

## P25 semantics

The [exact P25 string](SupportedPatterns.md#p25) selects runs of Unicode Nd,
including after unmatched gaps, and groups them from the right in blocks of 510
scalars. A run of `n` digits starts with `(n-1)%510+1` digits and then zero or more
complete 510-digit groups. All positions remain UTF-8 byte offsets.

The first regex alternative always has a match at an Nd scalar: a prefix of
1–510 digits can leave a multiple of 510 digits before a non-Nd scalar or EOF.
Greediness chooses the longest such prefix. The `$` case adds no earlier endpoint
inside an Nd run: before LF, the non-Nd alternative already succeeds. The later
`\G` alternative therefore never overrides the selected match for this exact
pattern. This reasoning does not register other equivalent expressions.

| Run length | Selected group lengths |
| --- | --- |
| 509 | 509 |
| 510 | 510 |
| 511 | 1, 510 |
| 512 | 2, 510 |
| 1020 | 510, 510 |
| 1021 | 1, 510, 510 |

`RightAlignedDecimals` first finds a decimal run with the shared scanner, caches
its byte end and returns the initial group. Following calls consume only the next
510 scalars while inside the cached run. At its end the cache is invalidated and
search resumes after any gap. For example, `12²34` has two matches (`12`, `34`),
and Split preserves `²` as a separate piece. P14 would stop after its first failure.

Short runs of at most 510 digits return the already known endpoint directly.
Long runs are scanned at most twice overall: once to find their end, then once
across their groups. Runtime is linear in input size with constant state, instead
of repeatedly scanning the remaining suffix for each output group.

[P25 provenance](../testdata/patterns/P25/manifest.json) pins the exact source and
shared Nd profile. Tests cover right-to-left reference grouping, all scalars,
510/1020/1530 boundaries, mixed-width digits, gaps and exact EOF. Cursor copies
are tested before, within and after cached runs, alongside interleaved subjects,
concurrent scans, destination reuse and zero-allocation checks. The live native
gate independently verifies selected HF matches and isolated splits.

## Unicode and evidence

[P10 provenance](../testdata/patterns/P10/manifest.json) pins the tokenizer source,
oracle, generator and class-table hashes. The generator enumerates every valid
Unicode scalar against the pinned oracle, producing 71 Nd ranges (760 scalars).
Go 1.27's Unicode 17.0 table includes ten additional digits U+11DE0–U+11DE9 absent
from that profile. Using `unicode.Nd` unchanged would alter matches.

Offline tests include captured HF byte boundaries, an independent RE2 expression
with a captured prefix, all-scalar membership contexts, long runs, cursor copies,
destination reuse and malformed UTF-8. The common benchmark pipeline generates
HF matches and isolated splits in memory and compares every Go/Oniguruma boundary.
No new reference output arrays are stored.

The shared `decimal` diagnostic workload exercises successful P10/P14/P25 matches on
ASCII, Arabic-Indic, fullwidth, mathematical and mixed-width decimal subjects.
Report it separately from natural text, where P10 mostly rejects the subject and
throughput can predominantly measure UTF-8 validation. Neither is a full tokenizer
or BPE benchmark.
