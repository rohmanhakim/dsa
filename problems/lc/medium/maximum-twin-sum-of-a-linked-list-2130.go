package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

/*
func pairSum(head *lib.ListNode) int {
	num := []int{}

	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}

	left := 0
	right := len(num) - 1
	max := 0
	for left < right {
		sum := num[left] + num[right]
		if sum > max {
			max = sum
		}
		left++
		right--
	}

	return max
}
*/

func pairSum(head *lib.ListNode) int {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *lib.ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	left := head
	right := prev
	max := 0
	for left != nil && right != nil {
		sum := left.Val + right.Val
		if sum > max {
			max = sum
		}
		left = left.Next
		right = right.Next
	}

	for prev != nil {
		fmt.Printf("%d, ", prev.Val)
		prev = prev.Next
	}

	return max
}

func MaximumTwinSumOfALinkedList() {
	// head := lib.ListNode{
	// 	Val: 5,
	// 	Next: &lib.ListNode{
	// 		Val: 4,
	// 		Next: &lib.ListNode{
	// 			Val: 2,
	// 			Next: &lib.ListNode{
	// 				Val: 1,
	// 			},
	// 		},
	// 	},
	// }

	head := lib.ListNode{
		Val: 4,
		Next: &lib.ListNode{
			Val: 2,
			Next: &lib.ListNode{
				Val: 2,
				Next: &lib.ListNode{
					Val: 3,
				},
			},
		},
	}

	// head := lib.ListNode{
	// 	Val: 1,
	// 	Next: &lib.ListNode{
	// 		Val: 100000,
	// 	},
	// }

	out := pairSum(&head)

	fmt.Println(out)
}
