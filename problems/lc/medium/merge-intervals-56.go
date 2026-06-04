package medium

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		newItem := intervals[i]
		if len(res) > 0 {
			if intervals[i][0] <= res[len(res)-1][1] {
				start := res[len(res)-1][0]
				end := intervals[i][1]
				if intervals[i][1] < res[len(res)-1][1] {
					end = res[len(res)-1][1]
				}
				newItem = []int{start, end}
				res[len(res)-1] = newItem
				continue
			}
		}
		res = append(res, newItem)

	}

	return res
}

func mergeClaudeSonnet(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		last := res[len(res)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			res = append(res, interval)
		}
	}

	return res
}

func MergeIntervals() {
	// intervals := [][]int{
	// 	{1, 3},
	// 	{2, 6},
	// 	{8, 10},
	// 	{15, 18},
	// }

	// intervals := [][]int{
	// 	{1, 4},
	// 	{4, 5},
	// }

	// intervals := [][]int{
	// 	{4, 7},
	// 	{1, 4},
	// }

	intervals := [][]int{
		{1, 4},
		{2, 3},
	}
	output := merge(intervals)
	fmt.Println(output)
}
