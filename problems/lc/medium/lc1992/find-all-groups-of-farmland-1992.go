package lc1992

func findFarmland(land [][]int) [][]int {
	m := len(land)
	n := len(land[0])

	coord := []int{m, n, 0, 0}
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		land[row][col] = 2

		if row < coord[0] {
			coord[0] = row
		}
		if col < coord[1] {
			coord[1] = col
		}
		if row > coord[2] {
			coord[2] = row
		}
		if col > coord[3] {
			coord[3] = col
		}

		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && land[nr][nc] == 1 {
				dfs(nr, nc)
			}
		}
	}

	res := [][]int{}
	for i := range m {
		for j := range n {
			if land[i][j] == 1 {
				coord = []int{m, n, 0, 0}
				dfs(i, j)
				res = append(res, coord)
			}
		}
	}

	return res
}
