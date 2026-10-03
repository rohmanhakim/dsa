package lc32

import (
	"slices"
)

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n+1)

	dp[0] = 0

	stack := []int{}
	for i := 1; i <= n; i++ {
		ch := s[i-1]

		if ch == '(' {
			stack = append(stack, i-1)
		} else {
			if len(stack) > 0 {
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				pLen := i - j

				dp[i] = pLen
				dp[i] += dp[j]
			}
		}
	}

	return slices.Max(dp)
}
