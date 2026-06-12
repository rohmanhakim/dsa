package medium

import "fmt"

func countProvinces(cities [][]int, city int, visited []bool) {
	visited[city] = true
	stack := []int{city}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[current] = true

		for neighbor := range cities[current] {
			if cities[current][neighbor] == 1 && !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))

	count := 0
	for city := range isConnected {
		if !visited[city] {
			countProvinces(isConnected, city, visited)
			count++
		}
	}

	return count
}

func findCircleNumUnionFind(isConnected [][]int) int {
	n := len(isConnected)
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

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // upper triangle only, matrix is symmetric
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	// count distinct roots
	count := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			count++
		}
	}
	return count
}

func NumberOfProvinces() {
	// isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	isConnected := [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}

	out := findCircleNum(isConnected)
	fmt.Println(out)
}
