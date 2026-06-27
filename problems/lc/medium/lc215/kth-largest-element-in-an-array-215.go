package lc215

import (
	"container/heap"
	"fmt"
)

type KMaxHeap []int

func (h KMaxHeap) Len() int {
	return len(h)
}

func (h KMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h KMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *KMaxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findKthLargest(nums []int, k int) int {

	h := &KMaxHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
	}

	res := 0
	for range k {
		res = heap.Pop(h).(int)
	}

	return res
}

func KthLargestElementInAnArray() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	out := findKthLargest(nums, k)
	fmt.Println(out)
}
