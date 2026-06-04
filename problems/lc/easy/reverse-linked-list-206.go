package easy

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func reverseList(head *lib.ListNode) *lib.ListNode {

	p := head

	var prev *lib.ListNode
	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}

	return prev
}

func ReverseLinkedList() {
	head := lib.ListNode{
		Val: 1,
		Next: &lib.ListNode{
			Val: 2,
			Next: &lib.ListNode{
				Val: 3,
				Next: &lib.ListNode{
					Val: 4,
					Next: &lib.ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	out := reverseList(&head)

	for out != nil {
		fmt.Printf("%d, ", out.Val)
		out = out.Next
	}
}
