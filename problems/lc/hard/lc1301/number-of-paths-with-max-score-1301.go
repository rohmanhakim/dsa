package lc1301

/*
You are given a square board of characters. You can move on the board starting at the bottom right square marked with the character 'S'.

You need to reach the top left square marked with the character 'E'. The rest of the squares are labeled either with a numeric character 1, 2, ..., 9 or with an obstacle 'X'. In one move you can go up, left or up-left (diagonally) only if there is no obstacle there.

Return a list of two integers: the first integer is the maximum sum of numeric characters you can collect, and the second is the number of such paths that you can take to get that maximum sum, taken modulo 10^9 + 7.

In case there is no path, return [0, 0].



Example 1:

Input: board = ["E23","2X2","12S"]
Output: [7,1]

Example 2:

Input: board = ["E12","1X1","21S"]
Output: [4,2]

Example 3:

Input: board = ["E11","XXX","11S"]
Output: [0,0]


Constraints:

2 <= board.length == board[i].length <= 100
*/

func pathsWithMaxScore(board []string) []int {
	const MOD = 10e8 + 7
	M := len(board)

	dirs := [][]int{{-1, 0}, {0, -1}, {-1, -1}}
	getNeighbors := func(row, col int) [][]int {
		neighbors := [][]int{}

		for _, d := range dirs {
			nr := row + d[0]
			nc := col + d[1]
			if nr >= 0 && nr < M && nc >= 0 && nc < M && board[nr][nc] != 'X' {
				neighbors = append(neighbors, []int{nr, nc})
			}
		}
		return neighbors
	}

	maxVal := 0
	pathCount := 0
	path := [][]int{}
	var dfs func(r, c, sum int)
	dfs = func(r, c, sum int) {
		if board[r][c] == 'E' {
			sum -= int('E' - '0')
			if sum > maxVal {
				if maxVal > 0 {
					pathCount = 0
				}
				maxVal = sum
			}
			if sum == maxVal {
				pathCount++
			}
			return
		}

		neighbors := getNeighbors(r, c)

		for _, n := range neighbors {
			val := int(board[n[0]][n[1]] - '0')
			sum += val % MOD
			path = append(path, []int{n[0], n[1]})
			dfs(n[0], n[1], sum)
			sum -= val % MOD
			path = path[:len(path)-1]
		}
	}

	dfs(M-1, M-1, 0)
	return []int{maxVal, pathCount}
}
