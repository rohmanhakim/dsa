package lc856

func scoreOfParenthesesAlternative(s string) int {
	depth := 0
	score := 0
	for i, ch := range s {
		if ch == '(' {
			depth++
		} else {
			if i > 0 && s[i-1] == '(' {
				score += 1 << (depth - 1)
			}
			depth--
		}
	}

	return score
}
