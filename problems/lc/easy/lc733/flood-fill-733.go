package lc733

import (
	"fmt"
	"slices"
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	img := slices.Clone(image)
	m := len(img)
	n := len(img[0])

	getNeighbors := func(row, col, color int) [][]int {
		dirs := [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		res := [][]int{}
		for _, d := range dirs {
			nrow := row + d[0]
			ncol := col + d[1]

			if nrow >= 0 && nrow < m && ncol >= 0 && ncol < n && img[nrow][ncol] == color {
				res = append(res, []int{nrow, ncol})
			}
		}
		return res
	}

	var dfs func(row, col, oldColor int)

	dfs = func(row, col, oldColor int) {

		img[row][col] = color

		neighbors := getNeighbors(row, col, oldColor)

		for _, n := range neighbors {
			img[n[0]][n[1]] = color
			dfs(n[0], n[1], oldColor)
		}
	}

	oldColor := img[sr][sc]
	if oldColor != color {
		dfs(sr, sc, img[sr][sc])
	}

	return img
}

func LC733() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	color := 2
	out := floodFill(image, sr, sc, color)
	fmt.Println(out)
}
