package lc1872

func stoneGameVIII(stones []int) int {
	n := len(stones)

	prefix := make([]int, n)
	prefix[0] = stones[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	dp := make([]int, n)

	dp[n-1] = prefix[n-1]
	for i := n - 2; i >= 0; i-- {
		best := max(dp[i+1], prefix[i]-dp[i+1])
		dp[i] = best
	}

	return dp[1]
}
