package lc3523

func maximumPossibleSize(nums []int) int {
	stack := []int{}

	length := 0
	for _, num := range nums {
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				length++
			}
		}
		stack = append(stack, num)
	}

	return length + 1
}
