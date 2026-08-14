package lc200

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	queue := [][]int{}

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	island := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				queue = append(queue, []int{i, j})
				grid[i][j] = '2'
				for len(queue) > 0 {
					front := queue[0]
					queue = queue[1:]

					for _, d := range dirs {
						nr := d[0] + front[0]
						nc := d[1] + front[1]

						if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == '1' {
							grid[nr][nc] = '2'
							queue = append(queue, []int{nr, nc})
						}
					}
				}
				island++
			}
		}
	}

	return island
}
