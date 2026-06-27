package lc328

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func oddEvenList(head *lib.ListNode) *lib.ListNode {

	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

func OddEvenLinkedList() {
	// head := lib.ListNode{
	// 	Val: 1,
	// 	Next: &lib.ListNode{
	// 		Val: 2,
	// 		Next: &lib.ListNode{
	// 			Val: 3,
	// 			Next: &lib.ListNode{
	// 				Val: 4,
	// 				Next: &lib.ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	head := lib.ListNode{
		Val: 2,
		Next: &lib.ListNode{
			Val: 1,
			Next: &lib.ListNode{
				Val: 3,
				Next: &lib.ListNode{
					Val: 5,
					Next: &lib.ListNode{
						Val: 6,
						Next: &lib.ListNode{
							Val: 4,
							Next: &lib.ListNode{
								Val: 7,
								// Next: &lib.ListNode{
								// 	Val: 8,
								// },
							},
						},
					},
				},
			},
		},
	}
	// head := lib.ListNode{
	// 	Val: 1,
	// }
	out := oddEvenList(&head)

	for out != nil {
		fmt.Printf("%d, ", out.Val)
		out = out.Next
	}
}
