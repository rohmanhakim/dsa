package lc3718

import (
	"sort"
)

func missingMultiple(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	target := k
	for {

		left := 0
		right := n - 1
		found := false
		for left <= right {
			mid := left + ((right - left) / 2)

			if nums[mid] > target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				found = true
				break
			}
		}
		if !found {
			return target
		}

		target = target + k
	}
}
