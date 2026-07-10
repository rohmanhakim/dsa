package lc141

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func hasCycle(head *lib.ListNode) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
