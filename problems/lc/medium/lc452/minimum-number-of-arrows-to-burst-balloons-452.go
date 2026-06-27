package lc452

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	maxx := points[0][1]
	arrow := 1
	for i := 1; i < len(points); i++ {
		nextStart := points[i][0]
		nextEnd := points[i][1]
		if nextStart > maxx {
			arrow++
			maxx = nextEnd
			continue
		}
		if nextEnd < maxx {
			maxx = nextEnd
		}

	}
	return arrow
}

func findMinArrowShotsClaudeSonnet(points [][]int) int {
	/*
		Why sort by end instead of start?

		Your version sorts by start and tracks the shrinking overlap window (maxx).
		It works, but it's doing two things: detecting when to fire a new arrow, and eagerly
		tightening the window.

		The end-sorted version makes the invariant explicit: always shoot at the rightmost point
		that still covers the current group (i.e., the earliest end). When you sort by end,
		points[i][1] is already the optimal shot position — no need to track or update a window.

		The check points[i][0] > arrowPos is the single decision point.
		Both are O(n log n) / O(1). The end-sort version is the more "canonical" greedy framing
		for interval scheduling problems in general (it's the same invariant as the classic
		minimum interval cover and activity selection problems), so it's worth internalizing as
		the reusable pattern.
	*/

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1] // sort by END, not start
	})

	arrows := 1
	arrowPos := points[0][1] // shoot at the end of the first balloon

	for i := 1; i < len(points); i++ {
		if points[i][0] > arrowPos { // current balloon starts after last arrow
			arrows++
			arrowPos = points[i][1] // shoot at the end of this balloon
		}
	}
	return arrows
}

func MinimumNumberOfArrowsToBurstBalloons() {
	points :=
		[][]int{
			{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13},
		}

	// points := [][]int{
	// 	{5, 15},
	// 	{4, 14},
	// 	{3, 13},
	// 	{2, 12},
	// }

	// points := [][]int{
	// 	{2, 3},
	// 	{2, 3},
	// }

	// points := [][]int{
	// 	{1, 2},
	// 	{2, 3},
	// 	{3, 4},
	// 	{4, 5},
	// }
	// points := [][]int{
	// 	{1, 10},
	// 	{2, 11},
	// 	{3, 12},
	// 	{4, 13},
	// 	{5, 14},
	// }
	out := findMinArrowShots(points)
	fmt.Println(out)
}
