package lc3731

import (
	"slices"
)

func findMissingElements(nums []int) []int {
	freq := make(map[int]bool)

	for i := range nums {
		freq[nums[i]] = true
	}
	smallest := slices.Min(nums)
	largest := slices.Max(nums)

	res := []int{}
	for i := smallest; i <= largest; i++ {
		if !freq[i] {
			res = append(res, i)
		}
	}

	return res
}
