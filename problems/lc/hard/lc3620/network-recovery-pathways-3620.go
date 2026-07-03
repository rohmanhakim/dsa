package lc3620

/*

You are given a directed acyclic graph of `n` nodes numbered from 0 to `n − 1`. This is represented by a 2D array `edges` of length `m`, where `edges[i] = [u<sub>i</sub>, v<sub>i</sub>, cost<sub>i</sub>]` indicates a one‑way communication from node `u<sub>i</sub>` to node `v<sub>i</sub>` with a recovery cost of `cost<sub>i</sub>`.

Some nodes may be offline. You are given a boolean array `online` where `online[i] = true` means node `i` is online. Nodes 0 and `n − 1` are always online.

A path from 0 to `n − 1` is **valid** if:

- All intermediate nodes on the path are online.
- The total recovery cost of all edges on the path does not exceed `k`.

For each valid path, define its **score** as the minimum edge‑cost along that path.

Return the **maximum** path score (i.e., the largest **minimum**\-edge cost) among all valid paths. If no valid path exists, return -1.

**Example 1:**

**Input:** edges = \[\[0,1,5\],\[1,3,10\],\[0,2,3\],\[2,3,4\]\], online = \[true,true,true,true\], k = 10

**Output:** 3

**Explanation:**

![](https://assets.leetcode.com/uploads/2025/06/06/graph-10.png)

- The graph has two possible routes from node 0 to node 3:
	1. Path `0 → 1 → 3`
		- Total cost = `5 + 10 = 15`, which exceeds k (`15 > 10`), so this path is invalid.
		2. Path `0 → 2 → 3`
		- Total cost = `3 + 4 = 7 <= k`, so this path is valid.
				- The minimum edge‐cost along this path is `min(3, 4) = 3`.
- There are no other valid paths. Hence, the maximum among all valid path‐scores is 3.

**Example 2:**

**Input:** edges = \[\[0,1,7\],\[1,4,5\],\[0,2,6\],\[2,3,6\],\[3,4,2\],\[2,4,6\]\], online = \[true,true,true,false,true\], k = 12

**Output:** 6

**Explanation:**

![](https://assets.leetcode.com/uploads/2025/06/06/graph-11.png)

- Node 3 is offline, so any path passing through 3 is invalid.
- Consider the remaining routes from 0 to 4:
	1. Path `0 → 1 → 4`
		- Total cost = `7 + 5 = 12 <= k`, so this path is valid.
				- The minimum edge‐cost along this path is `min(7, 5) = 5`.
		2. Path `0 → 2 → 3 → 4`
		- Node 3 is offline, so this path is invalid regardless of cost.
		3. Path `0 → 2 → 4`
		- Total cost = `6 + 6 = 12 <= k`, so this path is valid.
				- The minimum edge‐cost along this path is `min(6, 6) = 6`.
- Among the two valid paths, their scores are 5 and 6. Therefore, the answer is 6.

**Constraints:**

- `n == online.length`
- `2 <= n <= 5 * 10<sup>4</sup>`
- `0 <= m == edges.length <= ``min(10<sup>5</sup>, n * (n - 1) / 2)`
- `edges[i] = [u<sub>i</sub>, v<sub>i</sub>, cost<sub>i</sub>]`
- `0 <= u<sub>i</sub>, v<sub>i</sub> < n`
- `u<sub>i</sub> != v<sub>i</sub>`
- `0 <= cost<sub>i</sub> <= 10<sup>9</sup>`
- `0 <= k <= 5 * 10<sup>13</sup>`
- `online[i]` is either `true` or `false`, and both `online[0]` and `online[n − 1]` are `true`.
- The given graph is a directed acyclic graph.
*/

import (
	"fmt"
	"math"
)

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	type edge struct {
		v int
		w int
	}

	maxEdgeWeight := 0
	adj := make([][]edge, len(online))
	for _, e := range edges {
		if !online[e[0]] || !online[e[1]] {
			continue
		}
		if adj[e[0]] == nil {
			adj[e[0]] = []edge{
				{
					v: e[1],
					w: e[2],
				},
			}
		} else {
			adj[e[0]] = append(adj[e[0]], edge{
				v: e[1],
				w: e[2],
			})
		}
		if e[2] > maxEdgeWeight {
			maxEdgeWeight = e[2]
		}
	}

	visited := make([]bool, len(online))
	stack := []int{}

	var dfs func(u, t int)
	dfs = func(u, t int) {
		visited[u] = true

		for _, e := range adj[u] {
			if !visited[e.v] && online[e.v] && e.w >= t {
				dfs(e.v, t)
			}
		}

		stack = append(stack, u)
	}

	feasible := func(t int) bool {
		visited = make([]bool, len(online))
		stack = []int{}

		// topo sort
		vertexCount := len(online)

		for i := 0; i < vertexCount; i++ {
			if !visited[i] {
				dfs(i, t)
			}
		}

		order := []int{}

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			order = append(order, top)
			stack = stack[:len(stack)-1]
		}

		// DP
		dp := make([]int64, vertexCount)
		for i := range dp {
			dp[i] = math.MaxInt
		}
		dp[0] = 0
		for _, u := range order {
			for _, e := range adj[u] {
				if dp[u] != math.MaxInt && e.w >= t {
					dp[e.v] = min(dp[e.v], dp[u]+int64(e.w))
				}
			}
		}

		return dp[vertexCount-1] <= k
	}

	ans := -1
	lo, hi := 0, maxEdgeWeight
	for lo <= hi {
		mid := (lo + hi) / 2
		if feasible(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func NetworkRecoveryPathways() {
	/*
		0 -> 1	5
		1 -> 3	10
		0 -> 2	3
		2 -> 3	4

		3, 1, 2, 0

		3	0
		1	10
		2	4
		0	min(3+4, 5+10) = 7
	*/
	// edges := [][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}}
	// k := 10
	// online := []bool{true, true, true, true}

	edges := [][]int{{0, 1, 7}, {1, 4, 5}, {0, 2, 6}, {2, 3, 6}, {3, 4, 2}, {2, 4, 6}}
	k := 12
	online := []bool{true, true, true, false, true}

	out := findMaxPathScore(edges, online, int64(k))
	fmt.Println(out)
}
