package lc703

import (
	"container/heap"
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
	last := old[len(old)-1]
	*h = old[:len(old)-1]
	return last
}

type KthLargest struct {
	k int
	h minHeap
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k: k,
		h: []int{},
	}
	for _, num := range nums {
		heap.Push(&kthLargest.h, num)
		if kthLargest.h.Len() > k {
			heap.Pop(&kthLargest.h)
		}
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}

	return this.h[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
