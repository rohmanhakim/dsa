package lib

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConstructListNode(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nodes[0],
	}
	curr := head
	for i := 1; i < len(nodes); i++ {
		node := &ListNode{
			Val: nodes[i],
		}
		curr.Next = node
		curr = curr.Next
	}
	return head
}

func LinkTail(head *ListNode, pos int) {
	if pos >= 0 {
		curr := head
		var to *ListNode
		for range pos {
			curr = curr.Next
		}
		to = curr
		curr = head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = to
	}
}

func Link(a *ListNode, b *ListNode) {
	a.Next = b
}

func PrintList(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("")
}
