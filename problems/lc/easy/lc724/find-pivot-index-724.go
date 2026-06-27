package lc724

import "fmt"

func pivotIndex(nums []int) int {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		left = left + nums[i]
		if sum == left {
			return i
		}
		sum = sum - nums[i]
	}
	return -1
}

func FindPivotIndex() {
	// nums := []int{1, 7, 3, 6, 5, 6}
	// nums := []int{2, 1, -1}
	nums := []int{1, 2, 3}
	out := pivotIndex(nums)
	fmt.Println(out)
}
