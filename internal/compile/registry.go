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

// NewlinePattern is the exact decoded corpus P11 string.
const NewlinePattern = "[\r\n]"

var newline = scan.Newline()

// CJKPattern is the exact decoded corpus P21 string.
const CJKPattern = `[一-龥぀-ゟ゠-ヿ]+`

var cjk = scan.CJK()

// BroadRangesPattern is the exact decoded corpus P17 string.
const BroadRangesPattern = `[一-龥ࠀ-一가-퟿]+`

var broadRanges = scan.BroadRanges()

// PrefixedSymbolsPattern is the exact decoded corpus P18 string.
const PrefixedSymbolsPattern = `\s?[!-/:-~！-／：-～‘-‟　-。]+`

var prefixedSymbols = scan.PrefixedSymbols()

// PrefixedAlphabetPattern is the exact decoded corpus P08 string.
const PrefixedAlphabetPattern = `\s?[A-Za-zµÀ-ÖØ-öø-ƺƼ-ƿǄ-ʓʕ-ʯͰ-ͳͶͷͻ-ͽͿΆΈ-ΊΌΎ-ΡΣ-ϵϷ-ҁҊ-ԯԱ-ՖႠ-ჅᎠ-Ᏽᏸ-ᏽᲐ-ᲺᲽ-Ჿᴀ-ᴫᵫ-ᵷᵹ-ᶚḀ-ἕἘ-Ἕἠ-ὅὈ-Ὅὐ-ὗὙὛὝὟ-ώᾀ-ᾴᾶ-ᾼιῂ-ῄῆ-ῌῐ-ΐῖ-Ίῠ-Ῥῲ-ῴῶ-ῼℂℇℊ-ℓℕℙ-ℝℤΩℨK-ℭℯ-ℴℹℼ-ℿⅅ-ⅉⅎↃↄⰀ-ⱻⱾ-ⳤⳫ-ⳮⳲⳳꙀ-ꙭꚀ-ꚛꜢ-ꝯꝱ-ꞇꞋ-ꞎꭰ-ꮿﬀ-ﬆﬓ-ﬗＡ-Ｚａ-ｚ𐐀-𐑏𐒰-𐓓𐓘-𐓻𐲀-𐲲𐳀-𐳲𑢠-𑣟𞤀-𞥃]+`

var prefixedAlphabet = scan.PrefixedAlphabet()

// ComplementPattern is the exact decoded corpus P23 string.
const ComplementPattern = ` ?[^(\s|[.,!?…。，、।۔،])]+`

var complement = scan.Complement()

// ScriptGroupsPattern is the exact decoded corpus P12 string.
const ScriptGroupsPattern = `(?:[\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF\u3400-\u4DBF\uf900-\uFAFF\uFF65-\uFF9F\u2F00-\u2FDF]+|[\u0E00-\u0E7F]+|[\u0E80-\u0EFF]+|[\u1780-\u17FF]+|[\u1000-\u109F\uAA60-\uAA7F\uA9E0-\uA9FF]+|[\uAC00-\uD7AF\u1100-\u11FF]+)`

var scriptGroups = scan.ScriptGroups()

// HangulGroupsPattern is the exact decoded corpus P15 string.
const HangulGroupsPattern = `\p{Hangul}+|[！…“”‘’—：；，、-〿︰-﹏]+|[ㄅ-ㄯ]+|[一-龥぀-ゟ゠-ヿ]+`

var hangulGroups = scan.HangulGroups()

// PunctuationWordsPattern is exact corpus P22, including literal CR/LF.
const PunctuationWordsPattern = "[!\"#$%&'()*+,\\-./:;<=>?@\\[\\\\\\]^_`{|}~][A-Za-z]+|[^\r\n\\p{L}\\p{P}\\p{S}]?[\\p{L}\\p{M}]+| ?[\\p{P}\\p{S}]+[\r\n]*|\\s*[\r\n]+|\\s+(?!\\S)|\\s+"

var punctuationWords = scan.PunctuationWords()

// InitialDecimalRemainderPattern is the exact anchored corpus P10 string.
const InitialDecimalRemainderPattern = `\A\p{Nd}{1,2}(?=\p{Nd}{3}+\z)`

var initialDecimalRemainder = scan.InitialDecimalRemainder()

// ConsecutiveDecimalTriplesPattern is the exact anchored corpus P14 string.
const ConsecutiveDecimalTriplesPattern = `\A\p{Nd}{3}|\G\p{Nd}{3}`

var consecutiveDecimalTriples = scan.ConsecutiveDecimalTriples()

// RightAlignedDecimalsPattern is the exact corpus P25 string.
const RightAlignedDecimalsPattern = `\p{Nd}{1,510}(?=(?>\p{Nd}{510})*(?:\P{Nd}|$))|\G\p{Nd}{510}`

var rightAlignedDecimals = scan.RightAlignedDecimals()

func Lookup(pattern string) (*scan.Plan, bool) {
	if pattern == RightAlignedDecimalsPattern {
		return &rightAlignedDecimals, true
	}

	if pattern == ConsecutiveDecimalTriplesPattern {
		return &consecutiveDecimalTriples, true
	}

	if pattern == InitialDecimalRemainderPattern {
		return &initialDecimalRemainder, true
	}

	if pattern == PunctuationWordsPattern {
		return &punctuationWords, true
	}

	if pattern == HangulGroupsPattern {
		return &hangulGroups, true
	}

	if pattern == ScriptGroupsPattern {
		return &scriptGroups, true
	}

	if pattern == NewlinePattern {
		return &newline, true
	}
	if pattern == CJKPattern {
		return &cjk, true
	}
	if pattern == BroadRangesPattern {
		return &broadRanges, true
	}
	if pattern == PrefixedSymbolsPattern {
		return &prefixedSymbols, true
	}
	if pattern == PrefixedAlphabetPattern {
		return &prefixedAlphabet, true
	}
	if pattern == ComplementPattern {
		return &complement, true
	}

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
