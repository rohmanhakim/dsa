package lc1536

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	prefix := make([]int, n)
	prefix[0] = stoneValue[0]
	for i := 1; i < n; i++ {
		prefix[i] = stoneValue[i] + prefix[i-1]
	}

	suffix := make([]int, n)
	suffix[n-1] = stoneValue[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = stoneValue[i] + suffix[i+1]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(l, r int) int

	dp = func(l, r int) int {
		if r == l {
			return 0
		}
		if r-l == 1 {
			return min(stoneValue[l], stoneValue[r])
		}

		score := 0
		for i := l; i < r; i++ {

			leftSum := prefix[i]
			if l > 0 {
				leftSum -= prefix[l-1]
			}

			rightSum := suffix[i+1]
			if r < n-1 {
				rightSum -= suffix[r+1]
			}

			if leftSum < rightSum {
				if memo[l][i] != -1 {
					score = max(score, leftSum+memo[l][i])
				} else {
					memo[l][i] = dp(l, i)
					score = max(score, leftSum+memo[l][i])
				}
			} else if leftSum > rightSum {
				if memo[i+1][r] != -1 {
					score = max(score, rightSum+memo[i+1][r])
				} else {
					memo[i+1][r] = dp(i+1, r)
					score = max(score, rightSum+memo[i+1][r])
				}
			} else {
				fromLeft := 0
				if memo[l][i] != -1 {
					fromLeft = memo[l][i]
				} else {
					memo[l][i] = dp(l, i)
					fromLeft = memo[l][i]
				}

				fromRight := 0
				if memo[i+1][r] != -1 {
					fromRight = memo[i+1][r]
				} else {
					memo[i+1][r] = dp(i+1, r)
					fromRight = memo[i+1][r]
				}

				score = max(score, max(fromLeft, fromRight)+leftSum)
			}
		}
		return score
	}

	score := dp(0, n-1)

	return score
}
