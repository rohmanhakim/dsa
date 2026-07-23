package lc2487

import "github.com/rohmanhakim/leetcode/lib"

func removeNodes(head *lib.ListNode) *lib.ListNode {

	stack := []*lib.ListNode{}

	curr := head

	for curr != nil {
		for len(stack) > 0 && curr.Val > stack[len(stack)-1].Val {

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, curr)
		curr = curr.Next
	}

	curr = stack[0]

	for i := 1; i < len(stack); i++ {
		node := stack[i]

		curr.Next = node
		curr = curr.Next
	}
	return stack[0]
}
