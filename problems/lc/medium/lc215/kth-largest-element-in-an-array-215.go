package lc215

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findKthLargest(nums []int, k int) int {

	h := minHeap([]int{})
	heap.Init(&h)

	for _, num := range nums {
		heap.Push(&h, num)

		if len(h) > k {
			heap.Pop(&h)
		}
	}

	return h[0]
}

func LC215() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	out := findKthLargest(nums, k)
	fmt.Println(out)
}
