package scan

// groupedRun selects the first matching alternative once, then keeps that set
// for the entire run. Overlapping classes must not switch the selected group.
type groupedRun struct{ groups []*exactSet }

func groupedPlan(ranges ...[]runeRange) Plan {
	p := groupedRun{groups: make([]*exactSet, len(ranges))}
	for i, rs := range ranges {
		p.groups[i] = newExactSet(rs, false, false)
	}
	return Plan{find: p.find}
}

func (p groupedRun) find(src string, pos int) (int, int) {
	for pos < len(src) {
		r, size := runeAt(src, pos)
		for _, group := range p.groups {
			if group.contains(r) {
				return pos, group.runEnd(src, pos+size)
			}
		}
		pos += size
	}
	return len(src), len(src)
}

// ScriptGroups builds P12's six ordered literal-set alternatives.
func ScriptGroups() Plan { return groupedPlan(scriptGroupRanges[:]...) }

// HangulGroups builds P15. The first two sets overlap at U+302E/U+302F;
// preserving the chosen set handles both initial priority and continuation.
func HangulGroups() Plan {
	return groupedPlan(hangulRanges[:], hangulPunctuationRanges[:],
		[]runeRange{{0x3105, 0x312F}}, cjkRanges[:])
}
