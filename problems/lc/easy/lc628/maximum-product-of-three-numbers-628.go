package lc628

import (
	"sort"
)

func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	return max(
		nums[n-1]*nums[0]*nums[1],
		nums[n-1]*nums[n-2]*nums[n-3],
	)
}
