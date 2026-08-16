package lc39

import (
	"slices"
)

func combinationSum(candidates []int, target int) [][]int {

	res := [][]int{}
	var dfs func(idx int, path []int, total int)
	dfs = func(idx int, path []int, total int) {
		if total == target {
			p := slices.Clone(path)
			res = append(res, p)
			return
		}

		if total < target {
			for i := idx; i < len(candidates); i++ {
				if total+candidates[i] <= target {
					path = append(path, candidates[i])
					total += candidates[i]
					dfs(i, path, total)
					path = path[:len(path)-1]
					total -= candidates[i]
				}
			}
		}
	}

	path := []int{}
	total := 0
	dfs(0, path, total)
	return res
}
