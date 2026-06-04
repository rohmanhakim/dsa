package medium

import "fmt"

func buildArray(target []int, n int) []string {
	sPush := "Push"
	sPop := "Pop"
	stack := []int{}
	res := []string{}
	targetPointer := 0
	for i := range n {
		currTarget := target[targetPointer]
		stack = append(stack, i+1)
		res = append(res, sPush)
		top := stack[len(stack)-1]
		if top == currTarget {
			targetPointer++
			if targetPointer == len(target) {
				break
			}
		} else {
			stack = stack[:len(stack)-1]
			res = append(res, sPop)
		}
	}

	return res
}

func buildArrayClaudeSonnet(target []int, n int) []string {
	res := []string{}
	ti := 0 // pointer into target

	for i := 1; i <= n && ti < len(target); i++ {
		res = append(res, "Push")
		if i != target[ti] {
			res = append(res, "Pop")
		} else {
			ti++
		}
	}

	return res
}

func BuildAnArrayWithCtackOperations() {
	target := []int{1, 2}
	n := 4
	out := buildArrayClaudeSonnet(target, n)
	fmt.Println(out)
}
