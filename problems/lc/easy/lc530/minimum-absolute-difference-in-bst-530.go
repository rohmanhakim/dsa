package lc530

import (
	"math"

	"github.com/rohmanhakim/leetcode/lib"
)

func getMinimumDifference(root *lib.TreeNode) int {

	/*
	                    236
	   		     104      701
	   			   227	    911
	*/

	/*
		var dfs func(curr *lib.TreeNode) int
		dfs = func(curr *lib.TreeNode) int {
			if curr.Left == nil && curr.Right == nil {
				fmt.Printf("%d %d\n", prev, curr.Val)
				diff := curr.Val - prev
				if diff < 0 {
					diff = diff * -1
				}
				return diff
			}

			prev = curr.Val

			left := math.MaxInt
			right := math.MaxInt

			if curr.Left != nil {
				left = dfs(curr.Left)
			}
			if curr.Right != nil {
				right = dfs(curr.Right)
			}
			fmt.Printf("%d min %d %d\n", curr.Val, left, right)
			return min(left, right)
		}

		res := dfs(root)
	*/

	prevVal := -1
	minDiff := math.MaxInt
	var traverse func(curr *lib.TreeNode)
	traverse = func(curr *lib.TreeNode) {
		if curr == nil {
			return
		}

		traverse(curr.Left)
		if prevVal > -1 {
			diff := curr.Val - prevVal
			if diff < minDiff {
				minDiff = diff
			}
		}
		prevVal = curr.Val

		traverse(curr.Right)
	}

	traverse(root)

	return minDiff
}
