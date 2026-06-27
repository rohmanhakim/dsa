package lc994

import "fmt"

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	type orange struct {
		row int
		col int
	}
	queue := []orange{}
	fresh := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch grid[i][j] {
			case 2:
				queue = append(queue, orange{row: i, col: j})
			case 1:
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	getNeighbors := func(row, col int) []orange {
		neighbors := []orange{}
		for _, d := range dirs {
			nRow := row + d[0]
			nCol := col + d[1]

			if nRow >= 0 && nRow < rows && nCol >= 0 && nCol < cols && grid[nRow][nCol] == 1 { //&& !rotten[nRow][nCol] {
				neighbors = append(neighbors, orange{row: nRow, col: nCol})
			}
		}
		return neighbors
	}

	elapsed := -1 // trick start at -1
	for len(queue) > 0 {
		elapsed++ // increment elapsed here, when the level actually enqueues something
		level := len(queue)

		for i := 0; i < level; i++ {
			front := queue[0]
			queue = queue[1:]

			neighbors := getNeighbors(front.row, front.col)

			for _, n := range neighbors {
				grid[n.row][n.col] = 2
				fresh--
				queue = append(queue, n)
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return elapsed
}

func RottingOranges() {
	grid := [][]int{{0}}
	for _, g := range grid {
		fmt.Println(g)
	}
	fmt.Println("")
	out := orangesRotting(grid)
	fmt.Println(out)
}
