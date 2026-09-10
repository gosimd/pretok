// Package scan implements prewritten scalar scanner plans. Inputs are validated
// UTF-8 before entering this package; all positions are byte boundaries.
package scan

// Plan selects immutable matching code. It owns no per-input execution state.
// Match returns the endpoint of a match at pos, or pos if no branch matches.
type Plan struct {
	match func(src string, pos int) int
}

func (p *Plan) Match(src string, pos int) int { return p.match(src, pos) }

// Qwen builds the P05 plan. The template fixes alternative order; components
// implement only the constructions used by this exact expression.
func Qwen() Plan {
	return qwen(false)
}

// QwenMarks builds P01 with marks in words and excluded from symbol runs.
func QwenMarks() Plan { return qwen(true) }

func qwen(includeMarks bool) Plan {
	class := &letters
	if includeMarks {
		class = &lettersAndMarks
	}
	p := qwenPlan{words: prefixedLetters{class: class}, symbols: symbolRun{class: class}}
	return Plan{match: p.match}
}

type qwenPlan struct {
	contractions contractionSet
	words        prefixedLetters
	numbers      singleNumber
	symbols      symbolRun
	whitespace   whitespaceRun
}

func (p qwenPlan) match(src string, pos int) int {
	if end := p.contractions.match(src, pos); end > pos {
		return end
	}
	if end := p.words.match(src, pos); end > pos {
		return end
	}
	if end := p.numbers.match(src, pos); end > pos {
		return end
	}
	if end := p.symbols.match(src, pos); end > pos {
		return end
	}
	return p.whitespace.match(src, pos)
}
