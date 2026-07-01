/*
You are given a **0-indexed** 2D matrix `grid` of size `n x n`, where `(r, c)` represents:

- A cell containing a thief if `grid[r][c] = 1`
- An empty cell if `grid[r][c] = 0`

You are initially positioned at cell `(0, 0)`. In one move, you can move to any adjacent cell in the grid, including cells containing thieves.

The **safeness factor** of a path on the grid is defined as the **minimum** manhattan distance from any cell in the path to any thief in the grid.

Return *the **maximum safeness factor** of all paths leading to cell* `(n - 1, n - 1)`*.*

An **adjacent** cell of cell `(r, c)`, is one of the cells `(r, c + 1)`, `(r, c - 1)`, `(r + 1, c)` and `(r - 1, c)` if it exists.

The **Manhattan distance** between two cells `(a, b)` and `(x, y)` is equal to `|a - x| + |b - y|`, where `|val|` denotes the absolute value of val.

**Example 1:**

```
0 0 1
0 0 0
0 0 0

path: any
```


```
Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
Output: 0
Explanation: All paths from (0, 0) to (n - 1, n - 1) go through the thieves in cells (0, 0) and (n - 1, n - 1).
```

**Example 2:**

```
0 0 1
0 0 0
0 0 0

path: [0,0,1], [0,0,1], [0,0,0]
```

```
Input: grid = [[0,0,1],[0,0,0],[0,0,0]]
Output: 2
Explanation: The path depicted in the picture above has a safeness factor of 2 since:
- The closest cell of the path to the thief at cell (0, 2) is cell (0, 0). The distance between them is | 0 - 0 | + | 0 - 2 | = 2.
It can be shown that there are no other paths with a higher safeness factor.
```

**Example 3:**

```
0 0 0 1
0 0 0 0
0 0 0 0
1 0 0 0

path: [0,0], [0,1], [1,1], [1,2], [2,2], [3,2], [3,3]
```

```
Input: grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
Output: 2
Explanation: The path depicted in the picture above has a safeness factor of 2 since:
- The closest cell of the path to the thief at cell (0, 3) is cell (1, 2). The distance between them is | 0 - 1 | + | 3 - 2 | = 2.
- The closest cell of the path to the thief at cell (3, 0) is cell (3, 2). The distance between them is | 3 - 3 | + | 0 - 2 | = 2.
It can be shown that there are no other paths with a higher safeness factor.
```

**Constraints:**

- `1 <= grid.length == n <= 400`
- `grid[i].length == n`
- `grid[i][j]` is either `0` or `1`.
- There is at least one thief in the `grid`.
*/

package lc2812

import (
	"fmt"
	"sort"
)

func maximumSafenessFactor(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	thieves := [][]int{}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 1 {
				if (row == 0 && col == 0) || (row == m-1 && col == n-1) {
					return 0
				}
				thieves = append(thieves, []int{row, col})
				visited[row][col] = 1
			}
		}
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	safetyGrid := make([][]int, m)
	for i := range m {
		safetyGrid[i] = make([]int, n)
	}

	getNeighbors := func(row, col int) [][]int {
		neighbors := [][]int{}
		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && visited[nr][nc] == 0 {
				neighbors = append(neighbors, []int{nr, nc})
			}
		}
		return neighbors
	}

	level := 0
	for len(thieves) > 0 {
		q := len(thieves)

		for i := 0; i < q; i++ {
			front := thieves[0]
			thieves = thieves[1:]

			neighbors := getNeighbors(front[0], front[1])

			for _, n := range neighbors {
				visited[n[0]][n[1]] = 1
				safetyGrid[n[0]][n[1]] = level + 1
				thieves = append(thieves, n)
			}
		}
		level++
	}

	var canReach func(row, col, treshold int) bool
	canReach = func(row, col, treshold int) bool {
		visited[row][col] = 1
		if safetyGrid[row][col] < treshold {
			return false
		}
		if row == m-1 && col == n-1 {
			return true
		}

		neighbors := getNeighbors(row, col)
		res := false
		for _, n := range neighbors {
			if safetyGrid[n[0]][n[1]] < treshold {
				continue
			}
			visited[n[0]][n[1]] = 1
			res = canReach(n[0], n[1], treshold)
			if res == true {
				break
			}
		}
		return res
	}

	level--
	tresholds := []int{}
	for i := range level {
		tresholds = append(tresholds, level-i)
	}
	idx := sort.Search(level, func(i int) bool {
		visited = make([][]int, m)
		for i := range visited {
			visited[i] = make([]int, n)
		}
		return canReach(0, 0, tresholds[i])
	})

	/*
		3 2 1 0
		2 3 2 1
		1 2 3 2
		0 1 2 3

		3 2 1 0
		2 1 0 1
		1 0 1 2
		0 1 2 3

		0 0 0
		0 0 0
		1 1 0


		1 0 0
		2 1 1
		3 2 2
	*/

	if idx >= len(tresholds) {
		return 0
	}
	return tresholds[idx]
}

func FindtheSafestPathinaGrid() {
	// grid := [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}
	// grid := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	grid := [][]int{{0, 1, 1}, {0, 0, 0}, {0, 0, 0}}
	// grid := [][]int{{0, 0, 1}, {0, 1, 1}, {0, 0, 0}}

	// grid := [][]int{
	// 	{1, 1, 1, 1, 1, 1, 1, 1},
	// 	{1, 0, 1, 0, 0, 0, 0, 1},
	// 	{1, 0, 1, 0, 1, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 0, 0, 1},
	// 	{1, 0, 1, 0, 1, 0, 1, 1},
	// 	{1, 0, 0, 0, 1, 0, 0, 1},
	// 	{1, 1, 1, 1, 1, 1, 1, 1},
	// }
	out := maximumSafenessFactor(grid)
	fmt.Println(out)
}
