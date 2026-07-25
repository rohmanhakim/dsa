package lc2467

import "math"

// Clean version: Two DFS passes
// Phase 1: DFS from Bob to find his path to node 0 (parent-tracking, no visited map)
// Phase 2: DFS from root exploring all leaf paths (parent-tracking, track max profit)
//
// Key simplifications vs original:
//   - Standard 3-line adjacency list (append always works on nil slices)
//   - Parent-tracking (node, parent) eliminates visited maps entirely for trees
//   - Bob's DFS returns path directly (no closure/global state)
//   - Alice's DFS explores all leaf paths in one pass

func mostProfitablePathClean(edges [][]int, bob int, amount []int) int {
	n := len(amount)

	// Build adjacency list — 3 lines, not 14
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// Phase 1: Find Bob's path to node 0 using DFS with parent-tracking
	bobPath := findBobPath(bob, -1, adj)

	// Build time array: bobTime[node] = time step when Bob reaches node
	// Initialize to n (greater than any possible time) = "Bob never reaches this node"
	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = n
	}
	for t, node := range bobPath {
		bobTime[node] = t
	}

	// Phase 2: DFS from root, explore all leaf paths, return max profit
	maxProfit := math.MinInt
	var dfs func(node, parent, time, profit int)
	dfs = func(node, parent, time, profit int) {
		// Compute income for this node based on Bob's arrival time
		if bobTime[node] > time {
			profit += amount[node] // Alice arrives first: full amount
		} else if bobTime[node] == time {
			profit += amount[node] / 2 // Simultaneous: half
		}
		// else Bob arrived first: nothing (gain = 0)

		// Check if leaf (only neighbor is parent, or root with single child)
		isLeaf := true
		for _, nb := range adj[node] {
			if nb != parent {
				isLeaf = false
				dfs(nb, node, time+1, profit)
			}
		}

		if isLeaf && profit > maxProfit {
			maxProfit = profit
		}
	}

	dfs(0, -1, 0, 0)
	return maxProfit
}

// findBobPath: DFS from start to target, returns the unique path.
// Uses parent-tracking instead of visited map.
func findBobPath(node, parent int, adj [][]int) []int {
	if node == 0 {
		return []int{0}
	}
	for _, nb := range adj[node] {
		if nb != parent {
			if path := findBobPath(nb, node, adj); path != nil {
				return append([]int{node}, path...)
			}
		}
	}
	return nil // no path found (shouldn't happen in a valid tree)
}
