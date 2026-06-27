package lc232

import "fmt"

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	temp := []int{}
	for len(this.data) > 0 {
		temp = append(temp, this.data[len(this.data)-1])
		this.data = this.data[:len(this.data)-1]
	}
	out := temp[len(temp)-1]
	temp = temp[:len(temp)-1]
	for len(temp) > 0 {
		this.data = append(this.data, temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}
	return out
}

func (this *MyQueue) Peek() int {
	temp := []int{}
	for len(this.data) > 0 {
		temp = append(temp, this.data[len(this.data)-1])
		this.data = this.data[:len(this.data)-1]
	}
	out := temp[len(temp)-1]
	for len(temp) > 0 {
		this.data = append(this.data, temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}
	return out
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func ImplementQueueUsingStack() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
