package lc1631

import (
	"container/heap"
	"math"
)

type Item struct {
	col    int
	row    int
	effort int
}

type PQ []Item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].effort < h[j].effort }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *PQ) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func minimumEffortPath(heights [][]int) int {

	m := len(heights)
	n := len(heights[0])

	neighbors := func(row, col int) [][]int {
		dirs := [][]int{
			{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		}
		res := [][]int{}
		for _, d := range dirs {
			nr := d[0] + row
			nc := d[1] + col

			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				res = append(res, []int{nr, nc})
			}
		}
		return res
	}

	abs := func(num int) int {
		if num < 0 {
			return num * -1
		}
		return num
	}

	pq := PQ{{
		col:    0,
		row:    0,
		effort: 0,
	}}

	effort := make([][]int, m)

	for i := range effort {
		effort[i] = make([]int, n)

		for j := range effort[i] {
			effort[i][j] = math.MaxInt
		}
	}

	effort[0][0] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(Item)

		if curr.row == m-1 && curr.col == n-1 {
			return curr.effort
		}

		if curr.effort > effort[curr.row][curr.col] {
			continue
		}

		for _, neighbor := range neighbors(curr.row, curr.col) {
			nr := neighbor[0]
			nc := neighbor[1]
			ne := max(abs(heights[curr.row][curr.col]-heights[nr][nc]), curr.effort)

			if ne < effort[nr][nc] {
				effort[nr][nc] = ne

				heap.Push(&pq, Item{
					row:    nr,
					col:    nc,
					effort: ne,
				})
			}
		}
	}

	return -1
}
