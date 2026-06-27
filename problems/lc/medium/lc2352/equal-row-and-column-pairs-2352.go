package lc2352

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	columns := make(map[int][]byte)
	row := make(map[string]int)

	for i := 0; i < len(grid); i++ {
		var sb strings.Builder
		for j := 0; j < len(grid); j++ {
			sb.WriteString(strconv.Itoa(grid[i][j]))
			sb.WriteByte(',')
			val := strconv.Itoa(grid[i][j])
			for _, c := range val {
				columns[j] = append(columns[j], byte(c))
			}
			columns[j] = append(columns[j], byte(','))
		}
		row[sb.String()]++
	}

	count := 0
	for _, v := range columns {
		if v, ok := row[string(v)]; ok {
			count += v
		}
	}

	return count
}

func equalPairsClaudeSonnet(grid [][]int) int {
	n := len(grid)
	rowCount := make(map[string]int)

	// Hash each row
	for _, r := range grid {
		rowCount[serialize(r)]++
	}

	// Check each column against the row map
	count := 0
	for j := 0; j < n; j++ {
		col := make([]int, n)
		for i := 0; i < n; i++ {
			col[i] = grid[i][j]
		}
		count += rowCount[serialize(col)]
	}

	return count
}

func serialize(arr []int) string {
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteByte(',')
	}
	return sb.String()
}

func EqualRowAndColumnPairs() {
	// grid := [][]int{
	// 	{3, 2, 1},
	// 	{1, 7, 6},
	// 	{2, 7, 7},
	// } // 1
	// grid := [][]int{
	// 	{3, 1, 2, 2},
	// 	{1, 4, 4, 5},
	// 	{2, 4, 2, 2},
	// 	{2, 4, 2, 2},
	// } // 3
	grid := [][]int{
		{11, 1},
		{1, 11},
	} // 2
	out := equalPairs(grid)
	fmt.Println(out)
}
