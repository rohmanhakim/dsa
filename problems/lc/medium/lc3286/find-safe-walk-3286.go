package lc3286

/*
---
title: "Find a Safe Walk Through a Grid"
source: "https://leetcode.com/problems/find-a-safe-walk-through-a-grid/description/?envType=daily-question&envId=2026-07-02"
author:
published:
created: 2026-07-02
description: "Can you solve this real interview question? Find a Safe Walk Through a Grid - You are given an m x n binary matrix grid and an integer health.You start on the upper-left corner (0, 0) and would like to get to the lower-right corner (m - 1, n - 1).You can move up, down, left, or right from one cell to another adjacent cell as long as your health remains positive.Cells (i, j) with grid[i][j] = 1 are considered unsafe and reduce your health by 1.Return true if you can reach the final cell with a health value of 1 or more, and false otherwise. Example 1:Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1Output: trueExplanation:The final cell can be reached safely by walking along the gray cells below.[https://assets.leetcode.com/uploads/2024/08/04/3868_examples_1drawio.png]Example 2:Input: grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3Output: falseExplanation:A minimum of 4 health points is needed to reach the final cell safely.[https://assets.leetcode.com/uploads/2024/08/04/3868_examples_2drawio.png]Example 3:Input: grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5Output: trueExplanation:The final cell can be reached safely by walking along the gray cells below.[https://assets.leetcode.com/uploads/2024/08/04/3868_examples_3drawio.png]Any path that does not go through the cell (1, 1) is unsafe since your health will drop to 0 when reaching the final cell. Constraints: * m == grid.length * n == grid[i].length * 1 <= m, n <= 50 * 2 <= m * n * 1 <= health <= m + n * grid[i][j] is either 0 or 1."
tags:
  - "clippings"
---
[3286\. Find a Safe Walk Through a Grid](/problems/find-a-safe-walk-through-a-grid/)

Attempted

Medium

Topics

![premium lock icon](/_next/static/images/lock-a6627e2c7fa0ce8bc117c109fb4e567d.svg)Companies

Hint

You are given an `m x n` binary matrix `grid` and an integer `health`.

You start on the upper-left corner `(0, 0)` and would like to get to the lower-right corner `(m - 1, n - 1)`.

You can move up, down, left, or right from one cell to another adjacent cell as long as your health *remains* **positive**.

Cells `(i, j)` with `grid[i][j] = 1` are considered **unsafe** and reduce your health by 1.

Return `true` if you can reach the final cell with a health value of 1 or more, and `false` otherwise.

**Example 1:**

**Input:** grid = \[\[0,1,0,0,0\],\[0,1,0,1,0\],\[0,0,0,1,0\]\], health = 1

**Output:** true

**Explanation:**

The final cell can be reached safely by walking along the paths below.

```
0,0
1,0
2,0
2,1
2,2
1,2
0,2
0,3
0,4
1,4
2,4
```

**Example 2:**

**Input:** grid = \[\[0,1,1,0,0,0\],\[1,0,1,0,0,0\],\[0,1,1,1,0,1\],\[0,0,1,0,1,0\]\], health = 3

**Output:** false

**Explanation:**

A minimum of 4 health points is needed to reach the final cell safely


**Example 3:**

**Input:** grid = \[\[1,1,1\],\[1,0,1\],\[1,1,1\]\], health = 5

**Output:** true

**Explanation:**

The final cell can be reached safely by walking along the paths below.

```
0,0
0,1
1,1
2,1
2,2
```

Any path that does not go through the cell `(1, 1)` is unsafe since your health will drop to 0 when reaching the final cell.

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 50`
- `2 <= m * n`
- `1 <= health <= m + n`
- `grid[i][j]` is either 0 or 1.
*/

import (
	"fmt"
	"math"
)

func findSafeWalk(grid [][]int, health int) bool {
	M := len(grid)
	N := len(grid[0])
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	type state struct {
		visited  bool
		distance int
	}

	stateGrid := make([][]state, M)
	for i := range M {
		stateGrid[i] = make([]state, N)
		for j := range N {
			stateGrid[i][j].distance = math.MaxInt
		}
	}

	getNeighbors := func(row, col int) [][]int {
		neighbors := [][]int{}
		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]

			if nr >= 0 && nr < M && nc >= 0 && nc < N && !stateGrid[nr][nc].visited {
				neighbors = append(neighbors, []int{nr, nc})
			}
		}
		return neighbors
	}

	deque := [][]int{{0, 0}}
	stateGrid[0][0].visited = true
	stateGrid[0][0].distance = grid[0][0]
	canReach := false
	for len(deque) > 0 {
		front := deque[0]
		deque = deque[1:]
		stateGrid[front[0]][front[1]].visited = true
		currDistance := stateGrid[front[0]][front[1]].distance

		if front[0] == M-1 && front[1] == N-1 {
			if currDistance < health {
				return true
			} else {
				return false
			}
		}

		neighbors := getNeighbors(front[0], front[1])

		for _, n := range neighbors {
			nVal := grid[n[0]][n[1]]

			if currDistance+nVal < stateGrid[n[0]][n[1]].distance {
				stateGrid[n[0]][n[1]].distance = currDistance + nVal

				if nVal == 0 {
					temp := deque
					deque = append([][]int{{n[0], n[1]}}, temp...)
				} else {
					deque = append(deque, n)
				}
			}
		}

	}

	for i := range M {
		for j := range N {
			fmt.Printf("%d ", stateGrid[i][j].distance)
		}
		fmt.Println("")
	}
	return canReach
}

func FindaSafeWalkThroughaGrid() {
	grid := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	health := 5
	out := findSafeWalkDijkstra(grid, health)
	fmt.Println(out)
}
