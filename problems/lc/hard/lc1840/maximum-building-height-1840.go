package lc1840

import (
	"fmt"
	"sort"
)

/*
You want to build n new buildings in a city. The new buildings will be built in a line and are labeled from 1 to n.

However, there are city restrictions on the heights of the new buildings:

The height of each building must be a non-negative integer.
The height of the first building must be 0.
The height difference between any two adjacent buildings cannot exceed 1.
Additionally, there are city restrictions on the maximum height of specific buildings. These restrictions are given as a 2D integer array restrictions where restrictions[i] = [id(i), maxHeight(i)] indicates that building id(i) must have a height less than or equal to maxHeight(i).

It is guaranteed that each building will appear at most once in restrictions, and building 1 will not be in restrictions.

Return the maximum possible height of the tallest building.

Example 1:

Input: n = 5, restrictions = [[2,1],[4,1]]
Output: 2
Explanation: We can build the buildings with heights [0,1,2,1,2], and the tallest building has a height of 2.

Example 2:

Input: n = 6, restrictions = []
Output: 5
Explanation: We can build the buildings with heights [0,1,2,3,4,5], and the tallest building has a height of 5.

Example 3:

Input: n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
Output: 5
Explanation: We can build the buildings with heights [0,1,2,3,3,4,4,5,4,3], and the tallest building has a height of 5.

Constraints:

2 <= n <= 10^9
0 <= restrictions.length <= min(n - 1, 10^5)
2 <= id(i) <= n
id(i) is unique.
0 <= maxHeighti <= 10^9
*/

func maxBuilding(n int, restrictions [][]int) int {

	restrictions = append(restrictions, []int{1, 0})
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	// Left-to-right pass: propagate constraints from left
	// Each building's effective height is limited by the previous one + distance
	for i := 1; i < len(restrictions); i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]
		restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1]+dist)
	}

	// Right-to-left pass: propagate constraints from right
	for i := len(restrictions) - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]
		restrictions[i][1] = min(restrictions[i][1], restrictions[i+1][1]+dist)
	}

	// Compute max height in gaps between consecutive restrictions
	maxHeight := 0
	for i := 1; i < len(restrictions); i++ {
		hL := restrictions[i-1][1]
		hR := restrictions[i][1]
		d := restrictions[i][0] - restrictions[i-1][0]
		maxInGap := (hL + hR + d) / 2
		if maxInGap > maxHeight {
			maxHeight = maxInGap
		}
	}

	// Check max height achievable beyond the last restriction
	lastRestriction := restrictions[len(restrictions)-1]
	maxBeyond := lastRestriction[1] + (n - lastRestriction[0])

	return max(maxHeight, maxBeyond)
}

func MaximumBuildingHeight() {
	n := 5
	restrictions := [][]int{{2, 1}, {4, 1}} // 2

	// n := 10
	// restrictions := [][]int{{2, 5}, {5, 3}, {7, 4}, {10, 3}} // 5

	out := maxBuilding(n, restrictions)
	fmt.Println(out)
}
