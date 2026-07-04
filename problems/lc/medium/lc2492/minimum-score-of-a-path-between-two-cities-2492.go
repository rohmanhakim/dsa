package lc2492

/*
You are given a positive integer `n` representing `n` cities numbered from `1` to `n`. You are also given a **2D** array `roads` where `roads[i] = [a<sub>i</sub>, b<sub>i</sub>, distance<sub>i</sub>]` indicates that there is a **bidirectional** road between cities `a<sub>i</sub>` and `b<sub>i</sub>` with a distance equal to `distance<sub>i</sub>`. The cities graph is not necessarily connected.

The **score** of a path between two cities is defined as the **minimum** distance of a road in this path.

Return *the **minimum** possible score of a path between cities* `1` *and* `n`.

**Note**:

- A path is a sequence of roads between two cities.
- It is allowed for a path to contain the same road **multiple** times, and you can visit cities `1` and `n` multiple times along the path.
- The test cases are generated such that there is **at least** one path between `1` and `n`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2022/10/12/graph11.png)

**Input:** n = 4, roads = \[\[1,2,9\],\[2,3,6\],\[2,4,5\],\[1,4,7\]\]
**Output:** 5
**Explanation:** The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5.
It can be shown that no other path has less score.

**Example 2:**

![](https://assets.leetcode.com/uploads/2022/10/12/graph22.png)

**Input:** n = 4, roads = \[\[1,2,2\],\[1,3,4\],\[3,4,7\]\]
**Output:** 2
**Explanation:** The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.

**Constraints:**

- `2 <= n <= 10<sup>5</sup>`
- `1 <= roads.length <= 10<sup>5</sup>`
- `roads[i].length == 3`
- `1 <= a<sub>i</sub>, b<sub>i</sub> <= n`
- `a<sub>i</sub> != b<sub>i</sub>`
- `1 <= distance<sub>i</sub> <= 10<sup>4</sup>`
- There are no repeated edges.
- There is at least one path between `1` and `n`.
*/

import (
	"fmt"
	"math"
)

func minScore(n int, roads [][]int) int {

	type road struct {
		to   int
		dist int
	}
	adj := make([][]road, n)
	for i := range len(roads) {
		r := roads[i]
		if adj[r[0]-1] == nil {
			adj[r[0]-1] = []road{{
				to:   r[1] - 1,
				dist: r[2],
			}}
		} else {
			adj[r[0]-1] = append(adj[r[0]-1], road{
				to:   r[1] - 1,
				dist: r[2],
			})
		}

		if adj[r[1]-1] == nil {
			adj[r[1]-1] = []road{{
				to:   r[0] - 1,
				dist: r[2],
			}}
		} else {
			adj[r[1]-1] = append(adj[r[1]-1], road{
				to:   r[0] - 1,
				dist: r[2],
			})
		}
	}

	visited := make([]bool, n)

	queue := []int{0}
	visited[0] = true
	minScore := math.MaxInt
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		neighbors := adj[front]
		for _, v := range neighbors {
			if v.dist < minScore {
				minScore = v.dist
			}

			if !visited[v.to] {
				visited[v.to] = true
				queue = append(queue, v.to)
			}
		}
	}

	return minScore
}

func MinimumScoreofaPathBetweenTwoCities() {
	// n := 4
	// roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}

	n := 4
	roads := [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}
	out := minScore(n, roads)
	fmt.Println(out)
}
