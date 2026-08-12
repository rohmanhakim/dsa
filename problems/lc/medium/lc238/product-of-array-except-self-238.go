package lc238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		pref := 1
		if i > 0 {
			pref = prefix[i-1]
		}
		suf := 1
		if i < n-1 {
			suf = suffix[i+1]
		}

		res[i] = pref * suf
	}
	return res
}
