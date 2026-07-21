package lc77

import (
	"slices"
)

func combine(n int, k int) [][]int {

	path := []int{}
	out := [][]int{}
	var dfs func(num int)

	dfs = func(num int) {
		if len(path) == k {
			out = append(out, slices.Clone(path))
			return
		}

		for i := num; i <= n; i++ {
			path = append(path, i)

			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return out
}
