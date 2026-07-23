package lc3513

func uniqueXorTriplets(nums []int) int {

	uniq := len(nums)

	if len(nums) < 3 {
		return uniq
	}

	pow2 := func(x int) int {
		return 1 << x
	}

	res := pow2(2)
	i := 2
	for res <= uniq {
		res = pow2(i)
		i++
	}

	return res
}
