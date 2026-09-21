# Working on pretok

## Project boundary

Build `github.com/gosimd/pretok` as an HF-compatible pre-tokenization library
with a public compiled Pipeline API. Regex scanning is an internal component,
not a separate public interface. Follow the design and status in `README.md`.
The [supported-pattern registry](docs/SupportedPatterns.md) is the canonical
documentation list of implemented exact strings and tokenizer examples. Update it
when adding support, preserving exact decoded strings and pinned source links.
P01 Split uses gosimd/unicode v0.1.2 for bounded rune decoding and optional SIMD
UTF-8 processing; classification and boundary decisions remain scalar. Keep
compatibility claims scoped to registered patterns
and the pinned fixtures; do not claim unmeasured performance improvements.

Keep the library focused on pre-tokenization, including the implemented
Metaspace/ByteLevel transformations and Sequence composition. Do not add BPE,
vocabulary lookup, token IDs, normalization, special-token handling, Hub
downloads, or a full tokenizer pipeline without an explicit request.
Do not expand this into a general-purpose regex engine.

## Public API and ownership

- Keep a small root facade. The exact-string registry, plans, scratch, and CPU
  dispatch belong under `internal/`. Parser and AST/IR packages are deferred.
- Keep `CompileJSON`, `Pipeline`, `Piece`, `Span`, and their errors in the root
  facade. `Regexp`, `Compile`, `Cursor`, and isolated `Split` belong to
  `internal/regexp`. Preserve their internal contracts and existing coverage;
  expose regex functionality to consumers through Pipeline Split stages.
  Update documentation whenever the actual contract changes.
- Compile once into immutable plans. Concurrent scans must not share mutable
  execution state.
- Pipeline may transform text; `Piece.OriginalRange` maps back to source UTF-8
  byte offsets. Document split modes, gaps, empty input and UTF-8 validation.
  `Walk` pieces and `WalkTextUnsafe` strings are callback-scoped; `WalkText`
  strings may be retained. Preserve internal regex byte-span/append semantics.
- Do not retain input or destination buffers. Keep reusable scratch internal.
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

Keep public Pipeline tests, examples and benchmarks in the root package;
regex API tests and microbenchmarks belong in `internal/regexp`, regex report
infrastructure in `internal/regexp/bench`, and scanner tests in `internal/scan`.
All supported pre-tokenizer types are correctness and performance targets,
including built-ins, transformations and sequences, not just regex patterns.

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

Write release documentation for the target version without temporary publication
status labels such as "unpublished", "unreleased", "unreleased tag", or "next
release". The user creates the tag and publishes the changes after committing
the documentation. Do not add publication-status caveats merely because the tag
is not present when preparing README, changelog, or other release documentation.

Inspect the current checkout and preserve unrelated user changes. Once Go code
exists, format changed Go files and run relevant tests; use `go test ./...` for
the full suite and `go test -race ./...` when concurrency changes warrant it.
For P01 buffering changes, run both ordinary and `GOEXPERIMENT=simd` tests and race
checks, using the pinned published Unicode dependency with `GOWORK=off`.
Exercise scalar and available ISA paths using the controls actually implemented
in the repository. Cross-compilation checks buildability, not runtime parity.

For documentation-only changes, check the diff, links, examples, and consistency;
do not scaffold code or tests just to validate prose. Do not invent successful
checks for commands that cannot run in the current repository.

Report what changed, what was verified, and any remaining limits. Keep current
implementation, design proposals, externally verified compatibility, and measured
performance clearly distinguished.
