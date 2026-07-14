package lc3336

func subsequencePairCount(nums []int) int {
	const MOD = 1_000_000_007

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	M := 0

	for _, num := range nums {
		if num > M {
			M = num
		}
	}

	prev := make([][]int, M+1)
	for i := range M + 1 {
		prev[i] = make([]int, M+1)
	}

	prev[0][0] = 1

	for _, num := range nums {
		curr := make([][]int, M+1)
		for i := range M + 1 {
			curr[i] = make([]int, M+1)
		}
		for j := 0; j <= M; j++ {
			gcdj := gcd(j, num)
			for k := 0; k <= M; k++ {
				gcdk := gcd(k, num)

				curr[gcdj][k] = (curr[gcdj][k] + prev[j][k]) % MOD
				curr[j][gcdk] = (curr[j][gcdk] + prev[j][k]) % MOD
				curr[j][k] = (curr[j][k] + prev[j][k]) % MOD
			}
		}
		prev = curr
	}

	sum := 0
	for g := 1; g <= M; g++ {
		sum = (sum + prev[g][g]) % MOD
	}

	return sum
}
