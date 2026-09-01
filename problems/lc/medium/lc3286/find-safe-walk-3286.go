package lc3286

import (
	"fmt"
	"math"
)

func findSafeWalk(grid [][]int, health int) bool {
	M := len(grid)
	N := len(grid[0])
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	type state struct {
		visited  bool
		distance int
	}

	stateGrid := make([][]state, M)
	for i := range M {
		stateGrid[i] = make([]state, N)
		for j := range N {
			stateGrid[i][j].distance = math.MaxInt
		}
	}

	getNeighbors := func(row, col int) [][]int {
		neighbors := [][]int{}
		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]

			if nr >= 0 && nr < M && nc >= 0 && nc < N && !stateGrid[nr][nc].visited {
				neighbors = append(neighbors, []int{nr, nc})
			}
		}
		return neighbors
	}

	deque := [][]int{{0, 0}}
	stateGrid[0][0].visited = true
	stateGrid[0][0].distance = grid[0][0]
	canReach := false
	for len(deque) > 0 {
		front := deque[0]
		deque = deque[1:]
		stateGrid[front[0]][front[1]].visited = true
		currDistance := stateGrid[front[0]][front[1]].distance

		if front[0] == M-1 && front[1] == N-1 {
			if currDistance < health {
				return true
			} else {
				return false
			}
		}

		neighbors := getNeighbors(front[0], front[1])

		for _, n := range neighbors {
			nVal := grid[n[0]][n[1]]

			if currDistance+nVal < stateGrid[n[0]][n[1]].distance {
				stateGrid[n[0]][n[1]].distance = currDistance + nVal

				if nVal == 0 {
					temp := deque
					deque = append([][]int{{n[0], n[1]}}, temp...)
				} else {
					deque = append(deque, n)
				}
			}
		}

	}

	for i := range M {
		for j := range N {
			fmt.Printf("%d ", stateGrid[i][j].distance)
		}
		fmt.Println("")
	}
	return canReach
}

func FindaSafeWalkThroughaGrid() {
	grid := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	health := 5
	out := findSafeWalkDijkstra(grid, health)
	fmt.Println(out)
}
