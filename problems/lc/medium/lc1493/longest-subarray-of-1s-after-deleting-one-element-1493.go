package lc1493

import "fmt"

func longestSubarray(nums []int) int {
	left := 0
	right := 0
	curr := 0
	zero := 0
	max := 0
	for right < len(nums) {
		if nums[right] == 1 {
			curr++
		} else {
			if zero == left {
				zero = right
			} else {
				left = zero
				curr = right - zero - 1
				zero = right
			}
		}
		if curr > max {
			max = curr
		}
		right++
	}
	if len(nums) == max {
		max--
	}
	return max
}

func LongestSubarrayOf1sAfterDeletingOneElement() {
	// nums := []int{1, 1, 0, 1}
	// nums := []int{1, 1, 1}
	// nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	nums := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	out := longestSubarray(nums)
	fmt.Println(out)
}
