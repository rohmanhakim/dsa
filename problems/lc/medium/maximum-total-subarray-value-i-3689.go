package medium

import (
	"fmt"
	"sort"
)

func maxTotalValue(nums []int, k int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return (int64(nums[len(nums)-1]) - int64(nums[0])) * int64(k)
}

func MaximumTotalSubarrayValueI() {
	nums := []int{1, 3, 2}
	k := 2
	out := maxTotalValue(nums, k)
	fmt.Println("")
	fmt.Println(out)
}
