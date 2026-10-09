package lc463

func islandPerimeter(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	bfs := func(row, col int) int {

		queue := [][]int{{row, col}}

		total := 0

		dir := [][]int{
			{0, 1},
			{1, 0},
			{-1, 0},
			{0, -1},
		}

		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]

			fr := front[0]
			fc := front[1]

			grid[fr][fc] = 2
			neighbors := [][]int{}

			perimeter := 4
			for _, d := range dir {
				nr := fr + d[0]
				nc := fc + d[1]

				if nr >= 0 && nr < m && nc >= 0 && nc < n {
					if grid[nr][nc] == 1 || grid[nr][nc] == 2 {
						perimeter--
					}
					if grid[nr][nc] == 1 {
						grid[nr][nc] = 2
						neighbors = append(neighbors, []int{nr, nc})
					}
				}
			}

			total += perimeter
			for _, ne := range neighbors {
				queue = append(queue, ne)
			}
		}

		return total
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res := bfs(i, j)
				return res
			}
		}
	}
	return 0
}
