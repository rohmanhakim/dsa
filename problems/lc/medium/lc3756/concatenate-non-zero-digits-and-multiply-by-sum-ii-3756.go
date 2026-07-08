package lc3756

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007
	res := make([]int, len(queries))

	xPref := make([]int, len(s)+1)
	nzPref := make([]int, len(s)+1)
	sumPref := make([]int, len(s)+1)

	xPref[0] = 0
	nzPref[0] = 0
	sumPref[0] = 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '0' {
			xPref[i] = xPref[i-1]
			nzPref[i] = nzPref[i-1]
			sumPref[i] = sumPref[i-1]
		} else {
			xPref[i] = ((10 * xPref[i-1]) + int(s[i-1]-'0')) % MOD
			nzPref[i] = nzPref[i-1] + 1
			sumPref[i] = sumPref[i-1] + int(s[i-1]-'0')
		}

	}

	mods := []int{1}

	for i := 1; i <= nzPref[len(nzPref)-1]; i++ {
		mods = append(mods, (mods[i-1]*10)%MOD)
	}

	for i, q := range queries {
		l := q[0]
		r := q[1]

		nz := nzPref[r+1] - nzPref[l]
		x := (xPref[r+1] - xPref[l]*mods[nz]%MOD + MOD) % MOD
		sum := sumPref[r+1] - sumPref[l]
		res[i] = (x * sum) % MOD
	}

	return res
}
