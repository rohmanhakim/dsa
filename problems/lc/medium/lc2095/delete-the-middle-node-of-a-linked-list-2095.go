package lc2095

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *lib.ListNode) *lib.ListNode {
	if head.Next == nil {
		return nil
	}

	cur := head
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}

	mid := n / 2

	cur = head
	for range mid - 1 {
		cur = cur.Next
	}

	if cur.Next.Next != nil {
		cur.Next = cur.Next.Next
	} else {
		cur.Next = nil
	}

	return head
}

func deleteMiddleClaudeSonnet(head *lib.ListNode) *lib.ListNode {
	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next.Next // offset by 2 so slow stops at mid-1

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func DeleteTheMediumNodeOfALinkedList() {
	head := lib.ListNode{
		Val: 1,
		Next: &lib.ListNode{
			Val: 3,
			Next: &lib.ListNode{
				Val: 4,
				Next: &lib.ListNode{
					Val: 7,
					Next: &lib.ListNode{
						Val: 1,
						Next: &lib.ListNode{
							Val: 2,
							Next: &lib.ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}

	out := deleteMiddle(&head)
	for out != nil {
		fmt.Printf("%d, ", out.Val)
		out = out.Next
	}
}
