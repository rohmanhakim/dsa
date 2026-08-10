package lc1510

func winnerSquareGame(n int) bool {

	dp := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		for s := 1; s*s <= i; s++ {
			if !dp[i-s*s] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
