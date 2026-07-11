package lc104

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func maxDepth(root *lib.TreeNode) int {
	// curr := root

	// if root == nil {
	// 	return 0
	// }
	// depth := 0
	// for curr != nil {
	// 	fmt.Println(curr.Val)
	// 	// if curr.Val == 3 {
	// 	// 	fmt.Printf("---> ")
	// 	// 	fmt.Println(curr)
	// 	// }
	// 	if curr.Left != nil {
	// 		curr = curr.Left
	// 		continue
	// 	}
	// 	if curr.Right != nil {
	// 		curr = curr.Right
	// 	}
	// 	depth++
	// }
	// return depth

	var traverse func(curr *lib.TreeNode, depth int) int
	traverse = func(curr *lib.TreeNode, depth int) int {
		if curr == nil {
			return depth
		}

		left, right := 0, 0
		left = traverse(curr.Left, depth+1)
		right = traverse(curr.Right, depth+1)
		return max(left, right)
	}
	return traverse(root, 0)
}
