package lc669

import (
	"github.com/rohmanhakim/leetcode/lib"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *lib.TreeNode, low int, high int) *lib.TreeNode {

	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	left := trimBST(root.Left, low, high)
	right := trimBST(root.Right, low, high)

	root.Left = left
	root.Right = right

	return root
}
