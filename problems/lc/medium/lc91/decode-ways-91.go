package lc91

import (
	"strconv"
)

/*
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n)

	if s[0] == '0' {
		return 0
	} else {
		dp[0] = 1
	}
	if n == 1 {
		return dp[0]
	}

	valid := func(s string) bool {
		if len(s) == 1 {
			return s != "0"
		}
		if s[0] == '0' {
			return false
		} else {
			num, _ := strconv.Atoi(s)
			if num > 26 {
				return false
			}
		}
		return true
	}
	if valid(s[:2]) && valid(s[1:2]) {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < n; i++ {
		ch := s[i]

		// fmt.Println(s[i-1:i+1], dp)
		if ch == '0' {

			if valid(s[i-1:i+1]) && i < n-1 {
				dp[i] = 1
			} else if valid(s[i-1:i+1]) && i == n-1 {
				dp[i] = dp[0]
			} else {
				return 0
			}
		} else {
			if valid(s[i-1 : i+1]) {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}

	fmt.Println(dp)

	return dp[n-1]
}
*/

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+2)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n+1; i++ {
		dp[i] = 0

		if s[i-2] != '0' {
			dp[i] += dp[i-1]
		}

		if i <= 2 {
			continue
		}

		sint, _ := strconv.Atoi(s[i-3 : i-1])

		if s[i-3] != '0' && 10 <= sint && sint <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n+1]
}
