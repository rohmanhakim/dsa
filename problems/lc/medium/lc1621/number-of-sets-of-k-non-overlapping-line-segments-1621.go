package lc1621

func numberOfSets(n int, k int) int {
	const MOD = 1000_000_007

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= k; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= k; i++ {
		runningSum := 0
		for j := 1; j < n; j++ {
			runningSum = (runningSum + dp[i-1][j-1]) % MOD

			dp[i][j] = (dp[i][j-1] + runningSum) % MOD
		}
	}

	return dp[k][n-1]
}

func numberOfSetsONSpace(n int, k int) int {
	const MOD = 1000_000_007

	dp := make([]int, n)
	prefixSum := make([]int, n+1)

	for i := range dp {
		dp[i] = 1
		prefixSum[i+1] = (prefixSum[i] + dp[i]) % MOD
	}

	for i := 1; i <= k; i++ {
		dp[0] = 0
		for j := 1; j < n; j++ {
			dp[j] = (dp[j-1] + prefixSum[j]) % MOD
		}

		for j := 0; j < n; j++ {
			prefixSum[j+1] = (prefixSum[j] + dp[j]) % MOD
		}
	}

	return dp[n-1]
}
