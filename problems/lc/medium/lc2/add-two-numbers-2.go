package lc2

import (
	"github.com/rohmanhakim/leetcode/lib"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {

	curr1 := l1
	curr2 := l2

	res := &lib.ListNode{
		Val: 0,
	}

	head := res
	rem := 0
	for curr1 != nil || curr2 != nil {
		sum := rem
		if curr1 != nil {
			sum += curr1.Val
		}
		if curr2 != nil {
			sum += curr2.Val
		}
		if sum > 9 {
			sum = sum - 10
			rem = 1
		} else {
			rem = 0
		}
		res.Val = sum

		if curr1 != nil {
			curr1 = curr1.Next
		}
		if curr2 != nil {
			curr2 = curr2.Next
		}
		if curr1 != nil || curr2 != nil {
			res.Next = &lib.ListNode{}
			res = res.Next
		}
	}

	if rem > 0 {
		res.Next = &lib.ListNode{
			Val: rem,
		}
	}
	return head
}
