package lc1331

import (
	"slices"
)

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	sorted := slices.Clone(arr)
	slices.Sort(sorted)
	ranks := make(map[int]int)
	rank := 0
	for _, a := range sorted {
		if _, ok := ranks[a]; !ok {
			ranks[a] = rank + 1
			rank++
		}
	}
	res := make([]int, n)
	for i, a := range arr {
		res[i] = ranks[a]
	}
	return res
}
