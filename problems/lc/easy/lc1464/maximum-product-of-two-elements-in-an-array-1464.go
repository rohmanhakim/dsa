package lc1464

import (
	"sort"
)

func maxProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	return (nums[n-1] - 1) * (nums[n-2] - 1)
}
