package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func pairSum(head *ListNode) int {
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

func pairSum(head *ListNode) int {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for head != nil {
		fmt.Printf("%d, ", head.Val)

		head = head.Next
	}

	return -1
}

func main() {
	head := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	// head := ListNode{
	// 	Val: 4,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 2,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 			},
	// 		},
	// 	},
	// }

	// head := ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 100000,
	// 	},
	// }

	out := pairSum(&head)

	fmt.Println(out)
}
