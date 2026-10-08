package lc1021

func removeOuterParentheses(s string) string {

	res := []rune{}
	balance := 0
	for _, ch := range s {
		if ch == ')' {
			balance--
			if balance > 0 {
				res = append(res, ch)
			}
		} else {
			balance++
			if balance > 1 {
				res = append(res, ch)
			}
		}
	}

	return string(res)
}
