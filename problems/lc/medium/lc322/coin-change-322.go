package lc322

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		val := math.MaxInt
		for _, c := range coins {
			if i-c < 0 || dp[i-c] == math.MaxInt {
				continue
			}
			if dp[i-c]+1 < val {
				val = dp[i-c] + 1
			}
		}
		dp[i] = val
	}
	if dp[amount] < math.MaxInt {
		return dp[amount]
	}
	return -1
}
