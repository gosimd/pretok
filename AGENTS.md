# Working on pretok

## Project boundary

Build `github.com/gosimd/pretok` as a compiler for common pre-tokenization regex
patterns into specialized scanners. Follow the design and status in `README.md`.
The root API supports P01/P05 through a shared scalar template, P02 through
an ASCII-triplet plan, P03 through a line-end whitespace search, and P04 through
an ordered case-word plan. Plans share cursor/gap handling and common
constructions. P09/P13 are independently validated exact aliases of the P05 plan.
The other 18 strings and SIMD remain unsupported. Keep compatibility claims scoped to these patterns
and the pinned fixtures; do not claim unmeasured performance improvements.

Keep the library focused on pre-token boundaries. Do not add BPE, vocabulary
lookup, token IDs, normalization, special-token handling, ByteLevel mapping, Hub
downloads, or a configurable full tokenizer pipeline without an explicit request.
Do not expand this into a general-purpose regex engine.

## Public API and ownership

- Keep a small root facade. The exact-string registry, plans, scratch, and CPU
  dispatch belong under `internal/`. Parser and AST/IR packages are deferred.
- Keep the public names `Regexp`, `Compile`, `Cursor`, and `Split` consistent with
  `regexp.go`. Preserve the implemented contracts and their tests.
  Update documentation whenever the actual contract changes. Do not introduce a
  full `PreTokenizer` or `CompileSequence` pipeline without an explicit request.
- Compile once into immutable plans. Concurrent scans must not share mutable
  execution state.
- Return byte spans into the supplied input. Preserve its bytes and document the
  split mode, unmatched-gap handling, empty-input behavior, and UTF-8 policy.
- Support caller-owned destination reuse with explicit append semantics. Do not
  retain input or destination buffers. Keep any reusable scratch internal.
- In the first implementation, return an unsupported-pattern error for any string
  absent from the supported registry, including malformed syntax. Do not add a
  parser merely for syntax diagnostics. Never approximate or use a regex fallback.

## Compiler correctness

Use an explicit registry of exact regex strings for the first implementation.
The current corpus contains 25 unique strings; it is an implementation target,
not a support claim. Keys are regex values after JSON decoding by the caller.
Do not trim, unescape again, parse, or normalize the supplied expression.

Several entries may share a prewritten scanner plan after their selected matches
and split boundaries have been validated against the pinned reference. Register
only implemented and validated entries. Keep unregistered equivalent spellings
unsupported rather than attempting structural recognition.

AST parsing and broader recognition are future work. If introduced later, define
the accepted dialect first and prove each normalization preserves selected matches
and split boundaries, not just the language recognized by the expression.

Preserve alternative order, quantifier bounds and greediness, flag scope,
character-class meaning, Unicode properties, lookaround, and EOF behavior.
Do not reorder alternatives, discard whitespace branches as redundant, equate
ASCII digits with Unicode numbers, or broaden literal spaces to whitespace
without proof under the chosen semantics. Pin Unicode/reference versions where
they affect results.

Keep compilation outside the hot path. Runtime should execute dedicated scanner
logic, not walk the regex AST or run a general regex virtual machine. Shared
prefixes and character classification can be factored only when branch priority
and consumption remain identical.

Initially reject patterns that can produce zero-length matches. Runtime entry
points must reject malformed UTF-8 before scanning; `Split` must leave `dst`
unchanged on that error. Do not silently normalize or replace input bytes.

## Corpus and compatibility evidence

Grow support from real tokenizer configurations:

- Record model ID, immutable revision, file path, exact regex, relevant surrounding
  pre-tokenizer settings, and reference implementation version for each fixture.
- Preserve exact patterns and record explicit aliases that share a plan. Do not
  infer equivalence from similar spelling or a family from a model name alone.
- Record how the corpus was selected. Claims about popularity, family counts, or
  coverage require actual collected evidence.
- Keep fixtures small, reproducible, and suitable for redistribution. Ordinary
  tests must run offline; corpus collection and oracle generation are separate
  explicit steps.
- Limit compatibility claims to verified patterns and split configurations.
  Regex support alone is not full Hugging Face tokenizer compatibility.

## Implementation sequence

For a new pattern family, first capture reference pieces and byte offsets using a
pinned upstream implementation with the intended split settings. Implement a
clear scalar scanner and validate it against that external oracle. The scalar
scanner then becomes the executable baseline for optimized backends; agreement
between two local implementations alone does not establish compatibility.

Profile before adding SIMD. Keep traversal and boundary decisions shared where
practical, isolate ISA-specific classification/matching, and preserve a portable
scalar path. CPU dispatch must never execute instructions unsupported by the host.
Avoid public backend knobs unless a demonstrated contract requires them.

Use the layout in the README as a starting point, creating packages only when
there is code to justify them. Prefer readable implementations and small changes
over speculative abstractions or broad refactoring.

## Tests

For compiler and scanner changes, add focused regression tests and appropriate
differential or fuzz coverage:

- Exact reference-versus-scalar pieces and byte offsets.
- Scalar-versus-SIMD parity for each available backend.
- Registered exact spellings and aliases that should compile, plus unregistered
  equivalent spellings, malformed expressions, and near misses that must fail.
- Ordered alternatives, scoped case folding, bounded number runs, contractions,
  Unicode letters/marks/numbers, punctuation, and mixed scripts.
- Leading, internal, and trailing whitespace; tabs, CR, LF, CRLF, Unicode
  whitespace, lookahead contexts, and EOF.
- Empty and short inputs, long runs, multibyte characters around vector/tail
  boundaries, destination append/reuse, and concurrent scanner reuse.

Check that spans remain in bounds, ordered, non-overlapping, and reconstruct the
input under the documented split mode. Test malformed UTF-8 once its policy is
defined. Never weaken reference expectations merely to make an optimization pass.

## Performance evidence

Measure compile time separately from scanning. For scanning, benchmark both fresh
destinations and reused capacity, short prompts and long texts, ASCII and
multilingual inputs, and representative whitespace/punctuation distributions.
Report `ns/op`, throughput where meaningful, `B/op`, and `allocs/op`.

Compare scalar and optimized paths on the same machine, toolchain, corpus, and
benchmark settings. Repeat noisy measurements, record CPU/OS/Go/backend details,
and report real-workload results as well as core microbenchmarks. Inspect generated
assembly for SIMD changes to confirm intended instructions and identify unwanted
calls or spills. Keep an optimization only when its measured tradeoff is justified;
include regressions rather than combining unrelated best-case measurements.

## Verification and reporting

Inspect the current checkout and preserve unrelated user changes. Once Go code
exists, format changed Go files and run relevant tests; use `go test ./...` for
the full suite and `go test -race ./...` when concurrency changes warrant it.
Exercise scalar and available ISA paths using the controls actually implemented
in the repository. Cross-compilation checks buildability, not runtime parity.

For documentation-only changes, check the diff, links, examples, and consistency;
do not scaffold code or tests just to validate prose. Do not invent successful
checks for commands that cannot run in the current repository.

Report what changed, what was verified, and any remaining limits. Keep current
implementation, design proposals, externally verified compatibility, and measured
performance clearly distinguished.
