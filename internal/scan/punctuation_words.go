package scan

// punctuationWordPlan implements P22's ordered ASCII-prefixed words, Unicode
// words, positive P|S runs, and the shared whitespace decision. Numbers have no
// branch of their own and can form gaps or one-scalar prefixes of words.
type punctuationWordPlan struct {
	words, symbols *exactSet
}

func PunctuationWords() Plan {
	p := punctuationWordPlan{
		words:   newExactSet(wordMarkRanges[:], false, false),
		symbols: newExactSet(punctuationSymbolRanges[:], false, false),
	}
	return Plan{match: p.match}
}

func asciiLetter(b byte) bool { return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' }

func (p punctuationWordPlan) match(src string, pos int) int {
	// Every ASCII punctuation/symbol is in P22's explicitly written prefix
	// class. This branch stops at the first byte outside [A-Za-z];
	// neither non-ASCII letters nor marks belong to its body.
	if src[pos] < 128 && p.symbols.ascii[src[pos]] && pos+1 < len(src) && asciiLetter(src[pos+1]) {
		end := pos + 2
		for end < len(src) && asciiLetter(src[end]) {
			end++
		}
		return end
	}

	r, size := runeAt(src, pos)
	if p.words.contains(r) {
		// Marks overlap the optional prefix and body; consuming a maximal
		// body gives the same endpoint, including a singleton mark.
		return p.words.runEnd(src, pos+size)
	}
	if r != '\r' && r != '\n' && !p.symbols.contains(r) && pos+size < len(src) {
		next, width := runeAt(src, pos+size)
		if p.words.contains(next) {
			return p.words.runEnd(src, pos+size+width)
		}
	}

	body := pos
	if src[body] == ' ' {
		body++
	}
	end := p.symbols.runEnd(src, body)
	if end > body {
		for end < len(src) && (src[end] == '\r' || src[end] == '\n') {
			end++
		}
		return end
	}
	return (whitespaceRun{}).match(src, pos)
}
