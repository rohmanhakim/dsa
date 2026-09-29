package lc2267

func hasValidPath(grid [][]byte) bool {

	m := len(grid)
	n := len(grid[0])
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	dp := make([][]map[int]bool, m)

	for i := range dp {
		dp[i] = make([]map[int]bool, n)
	}

	dp[0][0] = make(map[int]bool)
	dp[0][0][1] = true

	for j := 1; j < n; j++ {
		dp[0][j] = make(map[int]bool)

		for k := range dp[0][j-1] {
			if grid[0][j] == '(' {
				dp[0][j][k+1] = true
			} else if k-1 >= 0 {
				dp[0][j][k-1] = true
			}
		}
	}

	for i := 1; i < m; i++ {
		dp[i][0] = make(map[int]bool)

		for k := range dp[i-1][0] {
			if grid[i][0] == '(' {
				dp[i][0][k+1] = true
			} else if k-1 >= 0 {
				dp[i][0][k-1] = true
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = make(map[int]bool)

			top := dp[i-1][j]
			left := dp[i][j-1]

			for k := range top {
				if grid[i][j] == '(' {
					dp[i][j][k+1] = true
				} else if k-1 >= 0 {
					dp[i][j][k-1] = true
				}
			}
			for k := range left {
				if grid[i][j] == '(' {
					dp[i][j][k+1] = true
				} else if k-1 >= 0 {
					dp[i][j][k-1] = true
				}
			}
		}
	}

	if ok := dp[m-1][n-1][0]; ok {
		return true
	}
	return false
}
