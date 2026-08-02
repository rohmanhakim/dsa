package lc877

import "fmt"

func stoneGame(piles []int) bool {

	n := len(piles)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for i := range n {
		dp[i][i] = piles[i]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[i]-dp[i][j-1])
		}
	}

	return dp[0][n-1] >= 0
}

func LC877() {
	piles := []int{5, 3, 4, 5}
	out := stoneGame(piles)
	fmt.Println(out)
}
