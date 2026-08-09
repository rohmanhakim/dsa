package lc1140

import (
	"math"
)

func stoneGameII(piles []int) int {

	n := len(piles)

	if n == 1 {
		return piles[0]
	}
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	suffix := make([]int, n)
	suffix[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	for i := n - 1; i >= 0; i-- {
		for m := 1; m < n; m++ {
			if i+(2*m) >= n {
				dp[i][m] = suffix[i]
			} else {
				best := math.MinInt
				for x := 1; x <= (2 * m); x++ {
					take := suffix[i] - suffix[i+x] - dp[i+x][max(m, x)]
					if take > best {
						best = take
					}
				}
				dp[i][m] = best
			}
		}
	}

	return (suffix[0] + dp[0][1]) / 2
}
