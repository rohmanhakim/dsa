/*
This problem gives us a general **undirected tree** (any node can have any number of children)

For general trees (and graphs) given as edge pairs, the standard approach is to convert `edges [][]int` into an **adjacency list** — a map from node → list of neighbors:

```
edges = [[1,2], [1,3], [3,4], [3,5]]

adjacency list:
1 → [2, 3]
2 → [1]
3 → [1, 4, 5]
4 → [3]
5 → [3]
```

This is the "lightweight" representation. we don't allocate tree nodes with pointers — we just use a `map[int][]int` (or `[][]int` if we know n). This is the standard way to represent trees/graphs in competitive programming.

## Why adjacency list is sufficient

To find max depth we need a traversal (BFS or DFS). But both BFS and DFS work directly on the adjacency list — we just need to **track visited nodes** (to avoid going back up the undirected edges) and **track the parent of each node** (to reconstruct the path if needed).

## Now, the second layer of the problem

Before we code, ask ourselves: what do We actually *need* from the traversal?

1. The **depth** of the deepest node (distance from node 1)
2. The **path** from node 1 to that deepest node — or at least its **length** (number of edges)

And then ask: once We have a path of length `d` with edges each weighted 1 or 2, how many assignments make the sum odd?

This is a combinatorics sub-problem: we have `d` positions, each gets value 1 or 2. we need an odd number of 1s. The count of ways to choose an odd number of positions out of `d` is `2^(d-1)`. (Half of all 2^d assignments have odd sum, half have even.)

## So the chain of thought is:

1. Convert edges → adjacency list
2. BFS/DFS from node 1 → find max depth `d`
3. Return `2^(d-1) mod 10^9+7`

No tree struct needed. The adjacency list + BFS/DFS is the general-purpose pattern for any edge-list tree problem. Does this mental model click? Toggle to Act mode when we're ready to implement.
*/

/*
The key insight behind the combinatorics formula, built up from first principles:

## The setup

we have `d` edges, each assigned weight 1 or 2. Total possible assignments = `2^d`.

## The question: how many of those have an odd sum?

## The insight: exactly half

Consider any single assignment, say `(1, 2, 1, 2, ...)` for d edges. Now **flip just the first edge**: if it was 1, make it 2; if it was 2, make it 1. This toggles the parity of the total sum (odd ↔ even).

This gives we a **one-to-one pairing**: for every assignment with an odd sum, there's exactly one counterpart with an even sum (and vice versa) — they differ only in the first edge.

```
(1, 2, 1) → sum = 4 (even)
(2, 2, 1) → sum = 5 (odd)   ← flipped first edge

(1, 1, 1) → sum = 3 (odd)
(2, 1, 1) → sum = 4 (even)  ← flipped first edge
```

Since every odd-sum assignment pairs with exactly one even-sum assignment, and together they exhaust all `2^d` possibilities:

```
odd_count = even_count = 2^d / 2 = 2^(d-1)
```

## That's the whole insight

No need to enumerate cases or sum binomial coefficients. The **pairing argument** (also called an **involution proof**) instantly tells we that half of all assignments have odd sum.

This pattern comes up a lot: whenever we have a binary choice (like 1 vs 2, or 0 vs 1) and we care about parity of the sum, the answer is usually `2^(n-1)`.
*/

package medium

import "fmt"

const MOD_10_9_7 = 1000000007

func getMaxDepth(tree map[int][]int, node, parent int) int {
	max := 0
	for _, neighbor := range tree[node] {
		if neighbor == parent {
			continue
		}
		depth := getMaxDepth(tree, neighbor, node)
		if depth > max {
			max = depth
		}
	}
	return max + 1
}

func assignEdgeWeights(edges [][]int) int {
	tree := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		tree[edges[i][0]] = append(tree[edges[i][0]], edges[i][1])
		tree[edges[i][1]] = append(tree[edges[i][1]], edges[i][0])
	}

	length := getMaxDepth(tree, 1, -1) - 1

	fmt.Println(length)
	return modpow(2, length-1, MOD_10_9_7)
}

func modpow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

func NumberOfWaysToAssignEdgeWeightsI() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}}
	out := assignEdgeWeights(edges)
	fmt.Println(out)
}
