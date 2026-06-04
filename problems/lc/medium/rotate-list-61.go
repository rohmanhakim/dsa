package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func rotateRight(head *lib.ListNode, k int) *lib.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	size := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		size++
	}

	cycle := k % size
	if cycle == 0 {
		return head
	}

	newTail := head
	for i := 0; i < size-cycle-1; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	tail.Next = head

	return newHead
}

func RotateList() {
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

	head = rotateRight(head, 2)

	curr := head

	for curr != nil {
		fmt.Printf("%d, ", curr.Val)
		curr = curr.Next
	}
}
