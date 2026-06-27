package lc700

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func bstDFS(root *lib.TreeNode, val int) *lib.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	left := bstDFS(root.Left, val)
	right := bstDFS(root.Right, val)

	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	return nil
}

func searchBST(root *lib.TreeNode, val int) *lib.TreeNode {
	res := bstDFS(root, val)
	return res
}

func SearchInBinarySearchTree() {
	root := lib.SliceToBinaryTreeNodes([]*int{
		lib.Ptr(4),
		lib.Ptr(2),
		lib.Ptr(7),
		lib.Ptr(1),
		lib.Ptr(3),
	})
	val := 2 // [2,1,3]

	out := searchBST(root, val)

	fmt.Println(out.Val)

}
