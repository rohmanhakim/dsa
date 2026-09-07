package lc940

func distinctSubseqII(s string) int {
	const MOD = 1000_000_007
	n := len(s)

	dp := 0

	last := make([]int, 26)
	for i := range last {
		last[i] = -1
	}

	for k := 0; k < n; k++ {
		prev := dp
		dp = ((2 * dp) - (last[s[k]-'a']) + MOD) % MOD
		last[s[k]-'a'] = prev
	}

	return dp
}
