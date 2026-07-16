package lc3867

import (
	"sort"
)

func gcdSum(nums []int) int64 {
	n := len(nums)
	mx := make([]int, n+1)
	prefix := make([]int, n)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	for i := 1; i <= n; i++ {
		mx[i] = max(nums[i-1], mx[i-1])
	}

	for i := 0; i < n; i++ {
		prefix[i] = gcd(nums[i], mx[i+1])
	}

	sort.Slice(prefix, func(i, j int) bool {
		return prefix[i] > prefix[j]
	})

	left := 0
	right := n - 1
	sum := int64(0)
	for left < right {
		gcdRes := gcd(prefix[left], prefix[right])
		sum += int64(gcdRes)
		left++
		right--
	}

	return sum
}
