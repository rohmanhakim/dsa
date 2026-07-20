package lc1260

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	for range k {
		temp := grid[m-1][n-1]
		for i := range m {
			last := grid[i][n-1]
			copy(grid[i][1:], grid[i])

			grid[i][0], temp = temp, last
		}
	}

	return grid
}
