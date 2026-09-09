package lc3871

func countCommas(n int64) int64 {

	if n < 1000 {
		return 0
	}

	pow := func(x, y int64) int64 {
		res := int64(1)
		for range y {
			res *= x
		}
		return res
	}

	mult := 1
	limit := n
	res := int64(0)
	for ; mult < 15; mult++ {
		limit = pow(1000, int64(mult+1))
		if n < int64(limit) {
			break
		}
		res += n - limit
	}

	return n - 1000 + int64(mult) + int64(res)
}
