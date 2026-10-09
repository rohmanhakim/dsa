package lc463

func islandPerimeterAlternative(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	total := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				perimeter := 4

				dir := [][]int{
					{0, 1},
					{1, 0},
					{0, -1},
					{-1, 0},
				}

				for _, d := range dir {
					nr := i + d[0]
					nc := j + d[1]

					if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
						perimeter--
					}
				}

				total += perimeter
			}
		}
	}

	return total
}
