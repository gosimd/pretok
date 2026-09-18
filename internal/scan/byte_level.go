package scan

// ByteLevel selects HF's built-in GPT-2 expression. It is deliberately not an
// additional alias in the public exact-string regex registry.
func ByteLevel() Plan { return Plan{match: byteLevelMatch} }

func byteLevelMatch(src string, pos int) int {
	if src[pos] == '\'' {
		for _, suffix := range [...]string{"s", "t", "re", "ve", "m", "ll", "d"} {
			end := pos + 1 + len(suffix)
			if end <= len(src) && src[pos+1:end] == suffix {
				return end
			}
		}
	}
	start := pos
	if src[start] == ' ' && start+1 < len(src) {
		start++
	}
	r, _ := runeAt(src, start)
	class := 0
	if letters.contains(r) {
		class = 1
	} else if number(r) {
		class = 2
	} else if !space(r) {
		class = 3
	}
	if class != 0 {
		end := start
		for end < len(src) {
			c, w := runeAt(src, end)
			match := class == 1 && letters.contains(c) || class == 2 && number(c) || class == 3 && !space(c) && !letters.contains(c) && !number(c)
			if !match {
				break
			}
			end += w
		}
		return end
	}
	// GPT-2 has only the two whitespace alternatives, without Qwen's CR/LF
	// branch. Before a following non-space, the lookahead leaves one scalar.
	run := (whitespaceRun{}).observe(src, pos)
	if run.end == len(src) {
		return run.end
	}
	if run.lastRuneStart > pos {
		return run.lastRuneStart
	}
	return run.end
}
