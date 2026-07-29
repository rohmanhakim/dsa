package lib

// nCr
func Comb(n, r, cap int) int {
	if r > n-r {
		r = n - r
	}
	res := 1
	for i := 1; i <= r; i++ {
		res = res * (n - r + i) / i // exact integer division
		if res >= cap {
			return cap
		}
	}
	return res
}
