package lc1926

import "fmt"

type mazeNode struct {
	row int
	col int
}

func nearestExit(maze [][]byte, entrance []int) int {

	queue := []mazeNode{{
		row: entrance[0],
		col: entrance[1],
	}}

	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", entrance[0], entrance[1])] = true

	isWall := func(row, col int) bool {
		if maze[row][col] == '+' {
			return true
		}
		return false
	}

	getNeighbors := func(row int, col int) []mazeNode {
		directions := [][]int{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		}
		neighbors := []mazeNode{}
		for _, d := range directions {
			nRow := row + d[0]
			nCol := col + d[1]
			inBound := nRow >= 0 && nRow < len(maze) && nCol >= 0 && nCol < len(maze[0])
			if !inBound {
				continue
			}
			notVisited := visited[fmt.Sprintf("%d,%d", nRow, nCol)] == false
			notWall := !isWall(nRow, nCol)
			if inBound && notVisited && notWall {
				neighbors = append(neighbors, mazeNode{
					row: nRow,
					col: nCol,
				})
			}
		}
		return neighbors
	}

	isGoal := func(row, col int) bool {
		inEdge := row == 0 || row == len(maze)-1 || col == 0 || col == len(maze[0])-1
		entrance := row == entrance[0] && col == entrance[1]

		if inEdge && !isWall(row, col) && !entrance {
			return true
		}

		return false
	}

	steps := 0
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			front := queue[0]
			queue = queue[1:]

			if isGoal(front.row, front.col) {
				return steps
			}

			neighbors := getNeighbors(front.row, front.col)
			for _, n := range neighbors {
				visited[fmt.Sprintf("%d,%d", n.row, n.col)] = true
				queue = append(queue, n)
			}
			fmt.Printf("%d,%d\n", front.row, front.col)
		}

		steps++
	}

	return -1
}

func nearestExitClaudeSonnet(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	type point struct{ r, c int }

	dirs := [4]point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Use 2D bool array instead of map — O(1) vs O(n) hash ops
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := []point{{entrance[0], entrance[1]}}
	visited[entrance[0]][entrance[1]] = true

	steps := 0
	for len(queue) > 0 {
		steps++
		for size := len(queue); size > 0; size-- {
			cur := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				r, c := cur.r+d.r, cur.c+d.c
				if r < 0 || r >= rows || c < 0 || c >= cols {
					continue
				}
				if maze[r][c] == '+' || visited[r][c] {
					continue
				}
				// Check goal at enqueue time — exit early
				isExit := r == 0 || r == rows-1 || c == 0 || c == cols-1
				if isExit {
					return steps
				}
				visited[r][c] = true
				queue = append(queue, point{r, c})
			}
		}
	}
	return -1
}

func NearestExitFromEntranceInMaze() {
	// maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	// entrance := []int{1, 2} //1

	maze := [][]byte{{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '.', '.', '+'}, {'+', '+', '+', '+', '.', '+', '.'}}
	entrance := []int{0, 1}
	for _, m := range maze {
		fmt.Println(m)
	}
	out := nearestExit(maze, entrance)
	fmt.Println(out)
}
