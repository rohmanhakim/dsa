package lc3524

import (
	"slices"
)

func resultArray(nums []int, k int) []int64 {
	dp := make([]int64, k)

	res := make([]int64, k)
	for _, num := range nums {
		prev := slices.Clone(dp)
		dp = make([]int64, k)

		for r := range prev {
			if prev[r] > 0 {
				dp[(r*num)%k] += prev[r]
			}
		}
		dp[num%k] += 1
		for i := range k {
			res[i] += dp[i]
		}
	}

	return res
}
