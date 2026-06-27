package lc238

import (
	"fmt"
	"slices"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	suffix := slices.Clone(prefix)
	suffix[n-1] = nums[n-1]
	product := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
		j := n - i - 1
		suffix[j] = suffix[j+1] * nums[j]
	}
	for i := 1; i < n-1; i++ {
		product[i] = suffix[i+1] * prefix[i-1]
	}
	product[0] = suffix[1]
	product[n-1] = prefix[n-2]
	return product
}

func ProductOfArrayExceptSelf() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{-1, 1, 0, -3, 3}
	nums := []int{4, 3, 2, 1, 2} // 12, 16, 24, 48, 24
	// 2 4 p[2] = 12
	// 1 12 p[1] = 24
	// 0 24 p[0] = 24
	out := productExceptSelf(nums)
	fmt.Println(out)
}
