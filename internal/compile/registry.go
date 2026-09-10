// Package compile selects scanner plans by exact, already JSON-decoded regex
// strings. It deliberately performs no parsing or normalization.
package compile

import "github.com/gosimd/pretok/internal/scan"

// QwenPattern is corpus P05. Do not register aliases without separate oracle
// validation and an explicit decision to expand the supported string set.
const QwenPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// QwenMarksPattern is the separately validated corpus P01 string.
const QwenMarksPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?[\p{L}\p{M}]+|\p{N}| ?[^\s\p{L}\p{M}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// QwenLiteralNewlinesPattern (P09) contains literal CR/LF after JSON decoding.
const QwenLiteralNewlinesPattern = "(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\\p{L}\\p{N}]?\\p{L}+|\\p{N}| ?[^\\s\\p{L}\\p{N}]+[\r\n]*|\\s*[\r\n]+|\\s+(?!\\S)|\\s+"

// QwenPossessivePattern (P13) shares P05 boundaries under the pinned profile.
const QwenPossessivePattern = `'(?i:[sdmt]|ll|ve|re)|[^\r\n\p{L}\p{N}]?+\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]++[\r\n]*|\s*[\r\n]|\s+(?!\S)|\s+`

// ASCIIDigitPattern is the exact corpus P16 string.
const ASCIIDigitPattern = `[0-9]`

// UnicodeTriplesPattern is the exact corpus P07 string.
const UnicodeTriplesPattern = `\p{N}{1,3}`

// QwenNumberTriplesPattern is the exact corpus P24 string.
const QwenNumberTriplesPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// QwenNumberRunsPattern is the exact corpus P20 string.
const QwenNumberRunsPattern = `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}+| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`

var qwen = scan.Qwen()
var qwenMarks = scan.QwenMarks()

// ASCIITriplesPattern is corpus P02; equivalent spellings are not aliases.
const ASCIITriplesPattern = `[0-9][0-9][0-9]`

var asciiTriples = scan.ASCIITriples()

// CaseWordsPattern is the separately validated corpus P04 string.
const CaseWordsPattern = `[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// CaseWordsSuffixPattern is the exact corpus P19 string.
const CaseWordsSuffixPattern = `[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+`

// CaseWordsSuffixTriplesPattern is the exact corpus P06 string.
const CaseWordsSuffixTriplesPattern = `[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+`

var caseWords = scan.CaseWords()
var caseWordsSuffix = scan.CaseWordsWithSuffix(1)
var caseWordsSuffixTriples = scan.CaseWordsWithSuffix(3)

// LineEndWhitespacePattern is corpus P03. Its anchor follows the pinned HF profile.
const LineEndWhitespacePattern = `\s+$`

var lineEndWhitespace = scan.LineEndWhitespace()

var asciiDigit = scan.ASCIINumbers(1, 1)
var unicodeTriples = scan.UnicodeNumbers(1, 3)
var qwenNumberTriples = scan.QwenNumbers(3)
var qwenNumberRuns = scan.QwenNumbers(-1)

func Lookup(pattern string) (*scan.Plan, bool) {
	if pattern == CaseWordsSuffixPattern {
		return &caseWordsSuffix, true
	}
	if pattern == CaseWordsSuffixTriplesPattern {
		return &caseWordsSuffixTriples, true
	}
	if pattern == ASCIIDigitPattern {
		return &asciiDigit, true
	}
	if pattern == UnicodeTriplesPattern {
		return &unicodeTriples, true
	}
	if pattern == QwenNumberTriplesPattern {
		return &qwenNumberTriples, true
	}
	if pattern == QwenNumberRunsPattern {
		return &qwenNumberRuns, true
	}

	if pattern == LineEndWhitespacePattern {
		return &lineEndWhitespace, true
	}
	if pattern == CaseWordsPattern {
		return &caseWords, true
	}
	if pattern == ASCIITriplesPattern {
		return &asciiTriples, true
	}
	if pattern == QwenPattern || pattern == QwenLiteralNewlinesPattern || pattern == QwenPossessivePattern {
		return &qwen, true
	}
	if pattern == QwenMarksPattern {
		return &qwenMarks, true
	}
	return nil, false
}
