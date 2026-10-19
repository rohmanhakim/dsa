package lc994

func orangesRotting(grid [][]int) int {
	/*
		BFS Problem

		count fresh oranges upfront

		loop grid
			when encounter rotten orange put to queue
			when encounter fresh orange increase counter

		begin BFS
			dequeue front evaluate neighbors
			exhaust queue : loop to the number of queue length
		    level is minute elapsed
			fresh oranges on 4-directional cells are the neighbors
			if encounter fresh orange, decrease fresh orange counter
			at the end of loop advance level

		for possible case, fresh orange counter will be 0 in that case return level

		for impossible case:

			2 1 1
			0 1 1
			1 0 1

		fresh orange counter will not return 0 , then return -1
	*/

	freshOrange := 0

	m := len(grid)
	n := len(grid[0])

	queue := [][]int{}
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				freshOrange++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if len(queue) == 0 {
		if freshOrange > 0 {
			return -1
		} else {
			return 0
		}
	}

	level := -1

	dirs := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(queue) > 0 {

		for range len(queue) {
			front := queue[0]

			queue = queue[1:]

			for _, d := range dirs {
				nr := d[0] + front[0]
				nc := d[1] + front[1]

				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					freshOrange--
					grid[nr][nc] = 2
					queue = append(queue, []int{nr, nc})
				}
			}
		}

		level++
	}

	if freshOrange == 0 {
		return level
	}

	return -1
}
