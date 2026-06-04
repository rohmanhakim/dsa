package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func rightSideView(root *lib.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*lib.TreeNode{root}
	rightSide := []int{}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			// fmt.Printf("%d ", queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i == n-1 {
				rightSide = append(rightSide, queue[i].Val)
			}
		}

		queue = queue[n:]

		// fmt.Println("")
	}

	return rightSide
}

func BinaryTreeRightSideView() {
	root := lib.SliceToBinaryTreeNodes([]*int{
		lib.Ptr(1),
		lib.Ptr(2),
		lib.Ptr(3),
		nil,
		lib.Ptr(5),
		nil,
		lib.Ptr(4),
	})

	out := rightSideView(root)
	fmt.Println(out)
}
