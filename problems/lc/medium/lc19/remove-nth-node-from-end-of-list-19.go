package lc19

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	dummy := &lib.ListNode{
		Next: head,
	}

	slow := dummy
	fast := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func RemoveNthNodeFromEndOfList() {

	head := &lib.ListNode{
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
	n := 2

	output := removeNthFromEnd(head, n)

	for output != nil {
		fmt.Printf("%d,", output.Val)
		output = output.Next
	}
}
