package lc856

func scoreOfParentheses(s string) int {

	type item struct {
		ch    rune
		value int
	}

	stack := []item{}

	for _, ch := range s {
		if ch == ')' {
			score := 0
			empty := true
			for len(stack) > 0 {
				top := stack[len(stack)-1]

				if top.ch == '(' {
					stack = stack[:len(stack)-1]
					value := 2 * score
					if empty {
						value = 1
					}
					stack = append(stack, item{
						ch:    '-',
						value: value,
					})
					break
				} else {
					empty = false
					stack = stack[:len(stack)-1]
					score += top.value
				}
			}
		} else {
			stack = append(stack, item{
				ch:    ch,
				value: 0,
			})
		}
	}

	res := 0
	for _, it := range stack {
		res += it.value
	}
	return res
}
