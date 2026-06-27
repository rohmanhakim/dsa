package lc373

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	u      int
	v      int
	column int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].u+h[i].v < h[j].u+h[j].v
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	h := &MinHeap{}

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{u: nums1[i], v: nums2[0], column: 0})
	}

	out := [][]int{}
	for len(out) < k {
		front := heap.Pop(h).(Pair)

		out = append(out, []int{front.u, front.v})

		if front.column+1 < len(nums2) {
			heap.Push(h, Pair{
				u:      front.u,
				v:      nums2[front.column+1],
				column: front.column + 1,
			})
		}
	}

	return out
}

func FindKPairsWithSmallestSums() {
	nums1 := []int{1, 7, 10, 11}
	nums2 := []int{2, 4, 6, 8}
	k := 6

	// nums1 := []int{1, 7, 11}
	// nums2 := []int{2, 4, 6}
	// k := 3

	// nums1 := []int{1, 2, 4, 5, 6}
	// nums2 := []int{3, 5, 7, 9}
	// k := 20
	out := kSmallestPairs(nums1, nums2, k)
	fmt.Println(out)

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2

	out = kSmallestPairs(nums1, nums2, k)
	fmt.Println(out)
}
