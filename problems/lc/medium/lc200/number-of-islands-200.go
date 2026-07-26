package lc200

import (
	"fmt"
	"strconv"
	"strings"
)

func numIslands(grid [][]byte) int {

	lands := make(map[string]bool)

	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				lands[fmt.Sprintf("%d,%d", i, j)] = true
			}
		}
	}

	getNeighbors := func(row, col int, visited [][]bool) [][]int {
		dirs := [][]int{
			{0, 1},
			{0, -1},
			{-1, 0},
			{1, 0},
		}

		res := [][]int{}
		for _, d := range dirs {
			nrow := row + d[0]
			ncol := col + d[1]

			if nrow >= 0 && nrow < m && ncol >= 0 && ncol < n && !visited[nrow][ncol] && grid[nrow][ncol] == '1' {
				res = append(res, []int{nrow, ncol})
			}
		}
		return res
	}

	var dfs func(row, col int)

	dfs = func(row, col int) {

		visited[row][col] = true
		delete(lands, fmt.Sprintf("%d,%d", row, col))
		neighbors := getNeighbors(row, col, visited)

		for _, n := range neighbors {
			visited[n[0]][n[1]] = true
			delete(lands, fmt.Sprintf("%d,%d", n[0], n[1]))
			dfs(n[0], n[1])
		}
	}

	count := 0
	for k := range lands {
		coordStr := strings.Split(k, ",")
		row, _ := strconv.Atoi(coordStr[0])
		col, _ := strconv.Atoi(coordStr[1])
		dfs(row, col)
		count++
	}

	return count
}
