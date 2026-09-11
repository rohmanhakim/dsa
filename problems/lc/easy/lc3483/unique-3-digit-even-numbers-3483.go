package lc3483

import (
	"fmt"
	"strconv"
	"strings"
)

/*
func totalNumbers(digits []int) int {

	countEvent := 0
	countZero := 0

	seen := make([]bool, 10)

	uniquePerm := func(n, r, cap int) int {
		if r > n {
			return 0
		}

		res := 1
		for i := 0; i < r; i++ {
			res *= n - i
			if res >= cap {
				return cap
			}
		}
		return res
	}

	for _, d := range digits {
		if d == 0 {
			countZero++
		} else if d%2 == 0 {
			countEvent++
		}
		seen[d] = true
	}

	factN := uniquePerm(len(digits)-1-countZero, 3-1-countZero, math.MaxInt)
	return factN * countEvent
}
*/

func totalNumbers(digits []int) int {

	res := make(map[string]int)
	path := []int{}
	visited := make(map[int]bool)
	var dfs func(curr int)

	dfs = func(curr int) {
		visited[curr] = true

		if len(path) == 3 {
			var sb strings.Builder
			for i := range 3 {
				sb.WriteString(strconv.Itoa(digits[path[i]]))
			}
			res[sb.String()]++
			return
		}

		for i, d := range digits {
			if visited[i] != true {
				if len(path) < 2 || (len(path) == 2 && d%2 == 0) {
					path = append(path, i)
					dfs(i)
					path = path[:len(path)-1]
					visited[i] = false

				}
			}
		}
	}

	for i, d := range digits {
		if d != 0 {
			clear(visited)
			path = []int{i}
			dfs(i)
		}
	}

	return len(res)
}

func LC3483() {
	/*
		314
		214
		324
		124
		134
		234


		132
		142
		312
		342
		412
		432
	*/
	digits := []int{1, 2, 2, 2}
	// digits := []int{6, 6, 6}
	out := totalNumbers(digits)
	fmt.Println(out)
}
