package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	remainder := 0
	res := &lib.ListNode{
		Val: 0,
	}

	n1 := l1
	n2 := l2
	nr := res

	for nr != nil {

		n1Val := 0
		n2Val := 0

		if n1 != nil {
			n1Val = n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			n2Val = n2.Val
			n2 = n2.Next
		}

		sum := n1Val + n2Val + remainder

		if sum >= 10 {
			remainder = sum / 10
			sum = sum - 10
		} else {
			remainder = 0
		}

		nr.Val = sum

		if n1 != nil || n2 != nil || remainder > 0 {
			nr.Next = &lib.ListNode{
				Val: 0,
			}
		}
		nr = nr.Next
	}

	return res
}

func AddTwoNumbers() {
	l1 := &lib.ListNode{
		Val: 2,
		Next: &lib.ListNode{
			Val: 4,
			Next: &lib.ListNode{
				Val: 3,
			},
		},
	}

	l2 := &lib.ListNode{
		Val: 5,
		Next: &lib.ListNode{
			Val: 6,
			Next: &lib.ListNode{
				Val: 4,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	fmt.Print("[")
	curr := result
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Print(",")
		}
		curr = curr.Next
	}
	fmt.Print("]")

}
