package lc1406

import (
	"slices"
)

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	dp := make([]int, n+3)

	for i := n - 1; i >= 0; i-- {
		takes := []int{}
		takes = append(takes, stoneValue[i]-dp[i+1])
		if n-i > 1 {
			takes = append(takes, stoneValue[i]+stoneValue[i+1]-dp[i+2])
		}
		if n-i > 2 {
			takes = append(takes, stoneValue[i]+stoneValue[i+1]+stoneValue[i+2]-dp[i+3])
		}

		dp[i] = slices.Max(takes)
	}

	switch {
	case dp[0] < 0:
		return "Bob"
	case dp[0] == 0:
		return "Tie"
	default:
		return "Alice"
	}

}
