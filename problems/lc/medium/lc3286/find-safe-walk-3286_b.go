package lc3286

import (
	"container/heap"
	"math"
)

type Cell struct {
	row      int
	col      int
	distance int
	weight   int
}

type minHeap []Cell

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(Cell))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findSafeWalkDijkstra(grid [][]int, health int) bool {
	M := len(grid)
	N := len(grid[0])
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	cells := make([][]Cell, M)
	for i := range M {
		cells[i] = make([]Cell, N)
		for j := range N {
			cells[i][j].row = i
			cells[i][j].col = j
			cells[i][j].distance = math.MaxInt
			cells[i][j].weight = grid[i][j]
		}
	}

	getNeighbors := func(row, col int) [][]int {
		neighbors := [][]int{}
		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]

			if nr >= 0 && nr < M && nc >= 0 && nc < N {
				neighbors = append(neighbors, []int{nr, nc})
			}
		}
		return neighbors
	}

	h := &minHeap{}
	cells[0][0].distance = grid[0][0]
	heap.Push(h, cells[0][0])
	canReach := false
	for h.Len() > 0 {
		front := heap.Pop(h).(Cell)
		if front.distance > cells[front.row][front.col].distance {
			continue // stale entry, skip
		}

		currDistance := cells[front.row][front.col].distance

		if front.row == M-1 && front.col == N-1 {
			if currDistance < health {
				return true
			} else {
				return false
			}
		}

		neighbors := getNeighbors(front.row, front.col)

		for _, n := range neighbors {
			nVal := cells[n[0]][n[1]].weight

			if currDistance+nVal < cells[n[0]][n[1]].distance {
				cells[n[0]][n[1]].distance = currDistance + nVal

				heap.Push(h, cells[n[0]][n[1]])
			}
		}

	}

	return canReach
}
