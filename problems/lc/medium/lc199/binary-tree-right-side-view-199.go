package lc199

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func rightSideView(root *lib.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	curr := root

	queue := []*lib.TreeNode{curr}

	res := []int{curr.Val}
	for len(queue) > 0 {
		level := len(queue)

		for range level {
			front := queue[0]

			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}

		}
		if len(queue) > 0 {
			res = append(res, queue[len(queue)-1].Val)
		}
	}
	return res
}
