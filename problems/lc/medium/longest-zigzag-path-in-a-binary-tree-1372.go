package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func checkZigzag(root *lib.TreeNode, dir int, count *int) {
	if root == nil {
		return
	}

	// fmt.Printf("-> v: %d, dir: %d, pdir: %d\n", root.Val, dir, pdir)

	*count++

	if dir == -1 {
		checkZigzag(root.Right, 1, count)
	} else {
		checkZigzag(root.Left, -1, count)
	}
}

func zigzagDFS(root *lib.TreeNode, maxZigzag *int) {

	if root == nil {
		return
	}

	count := 0
	checkZigzag(root, -1, &count)
	*maxZigzag = max(count, *maxZigzag)
	count = 0
	checkZigzag(root, 1, &count)
	*maxZigzag = max(count, *maxZigzag)

	// fmt.Printf("root: %d, zcount: %d\n", root.Val, count)

	zigzagDFS(root.Left, maxZigzag)
	zigzagDFS(root.Right, maxZigzag)
}

func dfsZigzag(root *lib.TreeNode, dir int, count int, maxCount *int) {
	if root == nil {
		return
	}

	if count > *maxCount {
		*maxCount = count
	}

	if dir == -1 {
		dfsZigzag(root.Left, 1, count+1, maxCount)
		dfsZigzag(root.Right, -1, 1, maxCount)
	} else {
		dfsZigzag(root.Right, -1, count+1, maxCount)
		dfsZigzag(root.Left, 1, 1, maxCount)
	}
}

func longestZigZag(root *lib.TreeNode) int {
	maxCount := 0
	dfsZigzag(root, -1, 0, &maxCount)
	dfsZigzag(root, 1, 0, &maxCount)
	return maxCount
}

func LongestZigZagPathInABinaryTree() {
	// root := lib.SliceToBinaryTreeNodes([]*int{
	// 	lib.Ptr(1),
	// 	nil,
	// 	lib.Ptr(2),
	// 	lib.Ptr(3),
	// 	lib.Ptr(4),
	// 	nil,
	// 	nil,
	// 	lib.Ptr(5),
	// 	lib.Ptr(6),
	// 	nil,
	// 	lib.Ptr(7),
	// 	nil,
	// 	nil,
	// 	nil,
	// 	lib.Ptr(8),
	// }) //3

	// root := lib.SliceToBinaryTreeNodes([]*int{
	// 	lib.Ptr(1),
	// 	lib.Ptr(2),
	// 	lib.Ptr(3),
	// 	nil,
	// 	lib.Ptr(4),
	// 	nil,
	// 	nil,
	// 	lib.Ptr(5),
	// 	lib.Ptr(6),
	// 	nil,
	// 	lib.Ptr(7),
	// }) // 4

	// root := lib.SliceToBinaryTreeNodes([]*int{
	// 	lib.Ptr(1),
	// 	lib.Ptr(2),
	// 	lib.Ptr(3),
	// 	lib.Ptr(4),
	// 	nil,
	// 	nil,
	// 	lib.Ptr(5),
	// }) // 1

	// root := lib.SliceToBinaryTreeNodes([]*int{
	// 	lib.Ptr(1),
	// 	lib.Ptr(2),
	// 	nil,
	// 	lib.Ptr(3),
	// 	lib.Ptr(4),
	// 	nil,
	// 	nil,
	// 	lib.Ptr(5),
	// 	lib.Ptr(6),
	// 	lib.Ptr(7),
	// }) // 3

	root := lib.SliceToBinaryTreeNodes([]*int{
		lib.Ptr(1),
		nil,
		lib.Ptr(2),
		lib.Ptr(3),
		lib.Ptr(4),
		nil,
		nil,
		lib.Ptr(5),
		lib.Ptr(6),
		nil,
		lib.Ptr(7),
		nil,
		nil,
		nil,
		lib.Ptr(8),
	}) // 4

	out := longestZigZag(root)
	fmt.Println(out)
	lib.PrintTreeASCII(root)
}
