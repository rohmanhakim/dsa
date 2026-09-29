package lc1614

func maxDepth(s string) int {
	stack := []rune{}

	depth := 0
	maxDepth := 0
	for _, ch := range s {
		if ch == ')' {
			depth--
			for len(stack) > 0 {
				if stack[len(stack)-1] == '(' {
					break
				}
				stack = stack[:len(stack)-1]
			}
		} else {
			if ch == '(' {
				depth++
				if depth > maxDepth {
					maxDepth = depth
				}
			}
			stack = append(stack, ch)
		}
	}
	return maxDepth
}
