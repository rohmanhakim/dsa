package lc139

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	n := len(s)

	dp := make([]bool, n+1)

	dp[0] = true
	for i := 1; i <= n; i++ {
		sub := s[0:i]

		val := false

		for _, w := range wordDict {
			idx := strings.LastIndex(sub, w)
			if idx > -1 && idx+len(w) == i {
				val = dp[i-len(w)]
			}
			if val {
				break
			}
		}

		dp[i] = val
	}
	return dp[n]
}
