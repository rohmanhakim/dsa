package lc452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	/*
		 1----6
		  2-----8
		       7----12
			      10----16
	*/
	xs := points[0][0]
	xe := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if points[i][1] <= xe {
			xe = points[i][1]
			if points[i][0] > xs {
				xs = points[i][0]
			}
		} else if points[i][0] > xe {
			count++
			xs = points[i][0]
			xe = points[i][1]
		}
	}
	return count
}
