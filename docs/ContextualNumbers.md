# Contextual decimal patterns

P10, P14 and P25 share the Unicode `Nd` class, but need different search rules.
P10 is implemented. P14 and P25 remain implementation targets, not registered
aliases. The public `Regexp`, `Cursor` and `Split` contracts stay unchanged.

## Context and ownership

The existing internal `Plan.Find(src, searchStart)` receives the entire supplied
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
| P14, planned | Consume three Nd scalars at the current search start; stop on failure | Existing cursor position is sufficient |
| P25, planned | Find a decimal run and select groups aligned from its right end, preserving alternative priority | Cache the run end in cursor-owned internal state to avoid rescanning the suffix for every group |

The reusable scalar component is `scanDecimals(src, pos, limit)`, returning
`decimalRun{End, Count}`. `End` is a byte boundary; `Count` counts Unicode scalars.
A zero limit scans the whole run; a positive limit consumes at most that many
digits. It distinguishes `Nd` from both ASCII `[0-9]` and the broader `N` class.

P10 uses local metadata, so it needs no additional cursor fields or stateful plan
hook. When implementing P25, place its cached run boundary in an internal value
owned by each cursor, never in the shared immutable plan. Cursor copies must
have independent progress and caches. Add that execution state only when P25
uses it; the public API and the input/search-start distinction above remain fixed.

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

The shared `decimal` diagnostic workload exercises successful P10 matches on
ASCII, Arabic-Indic, fullwidth, mathematical and mixed-width decimal subjects.
Report it separately from natural text, where P10 mostly rejects the subject and
throughput can predominantly measure UTF-8 validation. Neither is a full tokenizer
or BPE benchmark.
