package lc1929

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}

func ConcatenatinOfArray() {
	nums := []int{1, 2, 1}
	out := getConcatenation(nums)

	fmt.Println(out)
}
