package lc209

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := math.MaxInt
	found := false

	left := 0
	right := 1

	sum := nums[0]
	for right <= n && left < n {
		sub := nums[left:right]

		if sum >= target {
			found = true
			if len(sub) < minLen {
				minLen = len(sub)
			}
			sum -= nums[left]
			left++
		} else {
			if right < n {
				sum += nums[right]
			}
			right++
		}
	}
	if found {
		return minLen
	}
	return 0
}
