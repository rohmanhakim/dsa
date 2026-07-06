package lc1288

/*
Given an array intervals where intervals[i] = [lᵢ, rᵢ] represent the interval [lᵢ, rᵢ), remove all intervals that are covered by another interval in the list.

The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.

Return the number of remaining intervals.



Example 1:

Input: intervals = [[1,4],[3,6],[2,8]]
Output: 2
Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
Example 2:

Input: intervals = [[1,4],[2,3]]
Output: 1


Constraints:

1 <= intervals.length <= 1000
intervals[i].length == 2
0 <= lᵢ < rᵢ <= 10^5
All the given intervals are unique.
*/

import (
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	/*
		1--4
		 2-----8
		  3--6
	*/

	count := 0
	best := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > best[1] {
			if intervals[i][0] == best[0] {
				count++
			}
			best = intervals[i]
		} else if intervals[i][1] <= best[1] {
			count++
		}
	}

	return len(intervals) - count
}
