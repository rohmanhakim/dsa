package lc3312

import (
	"slices"
	"sort"
)

func gcdValues(nums []int, queries []int64) []int {

	comb2 := func(n int) int {
		return n * (n - 1) / 2
	}

	maxG := slices.Max(nums)

	divCount := make([]int, maxG+1)
	cnt := make([]int, maxG+1)
	for _, num := range nums {
		cnt[num]++
	}
	for i := 1; i <= maxG; i++ {
		for j := i; j <= maxG; j += i {
			divCount[i] += cnt[j]
		}
	}

	pairsWithGcdAtLeastG := make([]int, maxG+1)

	for i := 1; i <= maxG; i++ {
		pairsWithGcdAtLeastG[i] = comb2(divCount[i])
	}

	exact := make([]int, maxG+1)

	for i := maxG; i > 0; i-- {
		val := pairsWithGcdAtLeastG[i]
		for j := i; j <= maxG && val > 0; j += i {
			val -= exact[j]
		}
		exact[i] = val
	}

	prefix := make([]int64, maxG+1)

	for i := 1; i < len(exact); i++ {
		prefix[i] = int64(exact[i]) + prefix[i-1]
	}

	ans := make([]int, len(queries))

	for i := range ans {
		idx := sort.Search(len(prefix), func(j int) bool {
			return prefix[j] > queries[i]
		})
		ans[i] = idx
	}
	return ans
}
