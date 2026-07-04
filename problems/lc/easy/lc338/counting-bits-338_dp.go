package lc338

func countBitsDP(n int) []int {

	/*
		0
		10
		11
		100
		101

	*/

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = (dp[i/2]) + (i % 2)
	}

	return dp
}
