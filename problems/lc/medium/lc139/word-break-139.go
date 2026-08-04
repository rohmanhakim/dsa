package lc139

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		sub := s[:i]
		for _, w := range wordDict {
			if len(sub) >= len(w) && sub[len(sub)-len(w):] == w {
				if dp[i-len(w)] {
					dp[i] = true
					break
				}
			}
		}

	}

	return dp[len(s)]
}
