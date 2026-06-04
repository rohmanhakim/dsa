package easy

import "fmt"

func shuffle(nums []int, n int) []int {
	res := []int{}
	x := 0
	y := len(nums) - n

	for i := range n {
		res = append(res, nums[x+i])
		res = append(res, nums[y+i])
	}

	return res
}

func ShuffleTheArray() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	out := shuffle(nums, n)
	fmt.Println(out)
}
