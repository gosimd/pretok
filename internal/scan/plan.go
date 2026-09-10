// Package scan implements prewritten scalar scanner plans. Inputs are validated
// UTF-8 before entering this package; all positions are byte boundaries.
package scan

// Plan selects immutable matching code. It owns no per-input execution state.
// Most plans match at a position; a search plan can skip failed runs wholesale.
type Plan struct {
	match func(src string, pos int) int
	find  func(src string, pos int) (int, int)
}

func (p *Plan) Match(src string, pos int) int { return p.match(src, pos) }
func (p *Plan) HasSearch() bool               { return p.find != nil }

// Find returns the next match or len(src), len(src). Requires HasSearch().
func (p *Plan) Find(src string, pos int) (int, int) { return p.find(src, pos) }

// ASCIITriples builds P02 using the cursor's shared search past unmatched scalars;
// Split owns preserving maximal gaps, including short digit-run residues.
func ASCIITriples() Plan { return Plan{match: asciiTriplet{}.match} }

// LineEndWhitespace builds P03 with the pinned reference's line-end anchor.
func LineEndWhitespace() Plan { return Plan{find: lineEndWhitespace{}.find} }

// Qwen builds the P05 plan. The template fixes alternative order; components
// implement only the constructions used by this exact expression.
func Qwen() Plan {
	return qwen(false, 1)
}

// QwenMarks builds P01 with marks in words and excluded from symbol runs.
func QwenMarks() Plan { return qwen(true, 1) }

// QwenNumbers varies only the Unicode-number bound: 3 for P24, -1 for P20.
func QwenNumbers(max int) Plan { return qwen(false, max) }

func ASCIINumbers(min, max int) Plan {
	return Plan{find: numberRun{ascii: true, min: min, max: max}.find}
}
func UnicodeNumbers(min, max int) Plan { return Plan{find: numberRun{min: min, max: max}.find} }

func qwen(includeMarks bool, max int) Plan {
	class := &letters
	if includeMarks {
		class = &lettersAndMarks
	}
	p := qwenPlan{words: prefixedLetters{class: class}, tail: wordTail{numberMax: max, symbols: symbolRun{class: class}}}
	return Plan{match: p.match}
}

type qwenPlan struct {
	contractions contractionSet
	words        prefixedLetters
	tail         wordTail
}

// wordTail preserves the common ordered number, symbol and whitespace branches.
type wordTail struct {
	numberMax  int
	symbols    symbolRun
	whitespace whitespaceRun
}

func (p qwenPlan) match(src string, pos int) int {
	if end := p.contractions.match(src, pos); end > pos {
		return end
	}
	if end := p.words.match(src, pos); end > pos {
		return end
	}
	return p.tail.match(src, pos)
}

func (p wordTail) match(src string, pos int) int {
	if end := (singleNumber{}).match(src, pos); end > pos {
		if p.numberMax != 1 {
			return (numberRun{min: 1, max: p.numberMax}).rest(src, pos, end)
		}
		return end
	}
	if end := p.symbols.match(src, pos); end > pos {
		return end
	}
	return p.whitespace.match(src, pos)
}

// CaseWords builds P04, with no contraction branch or attached suffix.
func CaseWords() Plan {
	p := caseWordPlan{tail: wordTail{numberMax: 1, symbols: symbolRun{class: &letters, slashSuffix: true}}}
	return Plan{match: p.match}
}

type caseWordPlan struct {
	words prefixedCaseWords
	tail  wordTail
}

func (p caseWordPlan) match(src string, pos int) int {
	if end := p.words.match(src, pos); end > pos {
		return end
	}
	return p.tail.match(src, pos)
}
