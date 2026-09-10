package scan

// P04 orders prefix? A*B+ before prefix? A+B*, where
// A=Lu|Lt|Lm|Lo|M and B=Ll|Lm|Lo|M. Greedy A may need to give
// characters back to B. These are selected endpoints, not longest-word choices.
type prefixedCaseWords struct{}

type caseWordEnds struct{ first, second int }

func observeCaseWord(src string, pos int) caseWordEnds {
	end, lastLowerEnd := pos, pos
	for end < len(src) {
		r, size := runeAt(src, end)
		kind := caseKind(r)
		if kind&caseUpper == 0 {
			break
		}
		end += size
		if kind&caseLower != 0 {
			lastLowerEnd = end
		}
	}
	upperEnd := end
	for end < len(src) {
		r, size := runeAt(src, end)
		if caseKind(r)&caseLower == 0 {
			break
		}
		end += size
	}
	first := lastLowerEnd
	if end > upperEnd {
		first = end
	}
	second := pos
	if upperEnd > pos {
		second = end
	}
	return caseWordEnds{first: first, second: second}
}

func (prefixedCaseWords) match(src string, pos int) int {
	r, size := runeAt(src, pos)
	body := pos
	if r != '\r' && r != '\n' && !letters.contains(r) && !number(r) {
		body += size
	}
	ends := observeCaseWord(src, body)
	if ends.first > body {
		return ends.first
	}
	// The optional prefix must backtrack within the FIRST alternative before
	// trying the second. A leading M belongs to both the prefix and A/B.
	// If the prefixed A*B+ failed, the remaining A run has no B member;
	// retrying without the prefix therefore ends at that leading mark.
	if body > pos && caseKind(r)&caseLower != 0 {
		return body
	}
	if ends.second > body {
		return ends.second
	}
	return pos
}
