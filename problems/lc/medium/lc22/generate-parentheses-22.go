package lc22

import "fmt"

func backtrack(n int, result *[]string, current string, openCount int, closeCount int) {
	if len(current) == n*2 {
		*result = append(*result, current)
		return
	}

	if openCount < n {
		backtrack(n, result, current+"(", openCount+1, closeCount)
	}

	if openCount > closeCount {
		backtrack(n, result, current+")", openCount, closeCount+1)
	}
}

func generateParenthesis(n int) []string {
	result := []string{}

	backtrack(n, &result, "", 0, 0)

	return result
}

func GenerateParentheses() {
	n := 3
	out := generateParenthesis(n)
	fmt.Println(out)
}
