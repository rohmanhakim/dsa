package lc21

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func mergeTwoLists(list1 *lib.ListNode, list2 *lib.ListNode) *lib.ListNode {

	p1 := list1
	p2 := list2

	dummy := &lib.ListNode{}
	tail := dummy

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next = p1
			tail = tail.Next
			p1 = p1.Next
		} else {
			tail.Next = p2
			tail = tail.Next
			p2 = p2.Next
		}
	}

	if p1 == nil {
		tail.Next = p2
	} else if p2 == nil {
		tail.Next = p1
	}

	return dummy.Next
}
