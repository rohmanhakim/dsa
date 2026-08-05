package lc2211

func countCollisions(directions string) int {

	stack := []rune{}

	count := 0
	for _, ch := range directions {

		if ch == 'R' {
			stack = append(stack, ch)
		}
		if ch == 'L' {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch top {
				case 'R':
					count += 2
				case 'S':
					count += 1
				}
				stack = append(stack, 'S')
			}
		}
		if ch == 'S' {
			for len(stack) > 0 && stack[len(stack)-1] == 'R' {
				count++
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 'S')
		}
	}

	var top rune
	for len(stack) > 0 && stack[len(stack)-1] == 'R' {
		stack = stack[:len(stack)-1]
	}
	for len(stack) > 0 && stack[len(stack)-1] == 'S' {
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	if top == 'S' && len(stack) > 0 {
		for _, ch := range stack {
			if ch == 'R' {
				count++
			}
		}
	}
	return count
}
