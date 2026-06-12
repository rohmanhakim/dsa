package temp

import "fmt"

func UnionFind() {

	// graph := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	// n := len(graph)
	n := 6

	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i // each node is its own root initially
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // path compression
		}
		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px == py {
			return
		}
		// union by rank
		if rank[px] < rank[py] {
			parent[px] = py
		} else if rank[px] > rank[py] {
			parent[py] = px
		} else {
			parent[py] = px
			rank[px]++
		}
	}

	union(0, 1)
	union(1, 2)
	fmt.Println(parent)
	fmt.Println(rank)
}
