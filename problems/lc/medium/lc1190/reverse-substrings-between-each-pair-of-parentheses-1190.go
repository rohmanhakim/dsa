package lc1190

func reverseParentheses(s string) string {

	stack := []byte{}

	for _, ch := range s {
		if ch == ')' {
			reversed := []byte{}

			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if top != '(' {
					reversed = append(reversed, byte(top))
				} else {
					break
				}
			}

			for _, chr := range reversed {
				stack = append(stack, chr)
			}
		} else {
			stack = append(stack, byte(ch))
		}
	}

	return string(stack)
}
