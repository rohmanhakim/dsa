package lc20

func isValid(s string) bool {
	stack := []rune{}

	popAndCheckTop := func(topShould rune) bool {
		if len(stack) > 0 && stack[len(stack)-1] == topShould {
			stack = stack[:len(stack)-1]
			return true
		} else {
			return false
		}
	}

	for _, ch := range s {
		switch ch {
		case ')':
			if !popAndCheckTop('(') {
				return false
			}
		case ']':
			if !popAndCheckTop('[') {
				return false
			}
		case '}':
			if !popAndCheckTop('{') {
				return false
			}
		default:
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
