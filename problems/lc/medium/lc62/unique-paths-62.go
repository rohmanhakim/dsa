package lc62

import "fmt"

func uniquePathsTopDown(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(y, x int) int

	dp = func(y, x int) int {
		if y == 0 && x == 0 {
			return 1
		}
		if memo[y][x] != -1 {
			return memo[y][x]
		}

		ways := 0
		if x > 0 {
			ways += dp(y, x-1)
		}
		if y > 0 {
			ways += dp(y-1, x)
		}

		memo[y][x] = ways
		return ways
	}

	count := dp(m-1, n-1)
	return count
}

func uniquePathsBottomUp(m int, n int) int {
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range n {
		dp[0][i] = 1
	}

	for i := range m {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func UniquePaths() {
	m := 3
	n := 7
	out := uniquePathsBottomUp(m, n)
	fmt.Println(out)
}
