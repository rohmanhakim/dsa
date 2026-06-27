package lc739

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	out := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[top] = i - top
		}
		stack = append(stack, i)
	}

	return out
}

func DailyTemperatures() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73} // 1,1,4,2,1,1,0,0
	// temperatures := []int{30, 40, 50, 60} // 1,1,1,0
	// temperatures := []int{30, 60, 90} // 1,1,0
	out := dailyTemperatures(temperatures)
	fmt.Println(out)
}
