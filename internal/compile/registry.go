// Package compile selects scanner plans by exact, already JSON-decoded regex
// strings. It deliberately performs no parsing or normalization.
package compile

import "github.com/gosimd/pretok/internal/scan"

// QwenPattern is corpus P05. Do not register aliases without separate oracle
// validation and an explicit decision to expand the supported string set.
const QwenPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// QwenMarksPattern is the separately validated corpus P01 string.
const QwenMarksPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?[\p{L}\p{M}]+|\p{N}| ?[^\s\p{L}\p{M}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

var qwen = scan.Qwen()
var qwenMarks = scan.QwenMarks()

// ASCIITriplesPattern is corpus P02; equivalent spellings are not aliases.
const ASCIITriplesPattern = `[0-9][0-9][0-9]`

var asciiTriples = scan.ASCIITriples()

// CaseWordsPattern is the separately validated corpus P04 string.
const CaseWordsPattern = `[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+`

var caseWords = scan.CaseWords()

func Lookup(pattern string) (*scan.Plan, bool) {
	if pattern == CaseWordsPattern {
		return &caseWords, true
	}
	if pattern == ASCIITriplesPattern {
		return &asciiTriples, true
	}
	if pattern == QwenPattern {
		return &qwen, true
	}
	if pattern == QwenMarksPattern {
		return &qwenMarks, true
	}
	return nil, false
}
