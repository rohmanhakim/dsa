package lc20

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []rune{}

	pair := func(a, b rune) bool {
		br1 := a == '(' && b == ')'
		br2 := a == '[' && b == ']'
		br3 := a == '{' && b == '}'
		return br1 || br2 || br3
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if !pair(top, ch) {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
