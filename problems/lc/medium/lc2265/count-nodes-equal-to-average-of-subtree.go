package lc2265

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func averageOfSubtree(root *lib.TreeNode) int {

	var dfs func(curr *lib.TreeNode) (int, int, int)
	dfs = func(curr *lib.TreeNode) (int, int, int) {
		if curr.Left == nil && curr.Right == nil {
			return curr.Val, 1, 1
		}

		sumLeft := 0
		sumRight := 0
		countLeft := 0
		countRight := 0
		validLeft := 0
		validRight := 0

		if curr.Left != nil {
			sumLeft, countLeft, validLeft = dfs(curr.Left)
		}
		if curr.Right != nil {
			sumRight, countRight, validRight = dfs(curr.Right)
		}

		sum := sumLeft + sumRight + curr.Val
		count := countLeft + countRight + 1
		avg := sum / count
		valid := validLeft + validRight
		if curr.Val == avg {
			valid++
		}

		return sum, count, valid
	}

	_, _, result := dfs(root)

	return result
}
