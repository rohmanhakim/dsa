package temp

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head  *Node
	count int
}

func (ll *LinkedList) append(data int) {
	newNode := &Node{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
	ll.count++
}

func (ll *LinkedList) popFront() *Node {
	head := ll.head
	ll.head = ll.head.next
	head.next = nil
	ll.count--
	return head
}

func (ll *LinkedList) print() {
	curr := ll.head
	size := 0
	for range ll.count {
		curr = curr.next
		fmt.Println(size)
		size++
	}
}

func (ll *LinkedList) shift(num int) {
	for range num {
		head := ll.popFront()
		ll.append(head.data)
	}
}
