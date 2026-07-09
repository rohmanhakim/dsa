package lc56

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= prev[1] {
			if intervals[i][1] > prev[1] {
				prev[1] = intervals[i][1]
			}
		} else {
			res = append(res, prev)
			prev = intervals[i]
		}
	}
	res = append(res, prev)
	return res
}
