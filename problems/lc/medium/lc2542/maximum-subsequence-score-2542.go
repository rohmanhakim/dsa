package lc2542

import (
	"container/heap"
	"fmt"
	"sort"
)

type NumPair struct {
	n1 int
	n2 int
}

type SubseqMinHeap []NumPair

func (h SubseqMinHeap) Len() int           { return len(h) }
func (h SubseqMinHeap) Less(i, j int) bool { return h[i].n1 < h[j].n1 }
func (h SubseqMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *SubseqMinHeap) Push(x any) {
	*h = append(*h, x.(NumPair))
}
func (h *SubseqMinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	if k > len(nums1) {
		return -1
	}
	numPairs := []NumPair{}
	for i := range len(nums2) {
		numPairs = append(numPairs, NumPair{
			n1: nums1[i],
			n2: nums2[i],
		})
	}
	sort.Slice(numPairs, func(i, j int) bool {
		return numPairs[i].n2 > numPairs[j].n2
	})

	h := SubseqMinHeap{}
	heap.Init(&h)

	for i := 0; i < k-1; i++ {
		heap.Push(&h, numPairs[i])
	}

	currMax := int64(0)
	sum := 0
	temp := []NumPair{}
	for range h.Len() {
		val := heap.Pop(&h).(NumPair)
		sum += val.n1
		temp = append(temp, val)
	}
	for _, t := range temp {
		heap.Push(&h, t)
	}
	for i := k - 1; i < len(numPairs); i++ {

		heap.Push(&h, numPairs[i])
		sum += numPairs[i].n1

		prod := sum * numPairs[i].n2
		if currMax < int64(prod) {
			currMax = int64(prod)
		}
		worst := heap.Pop(&h).(NumPair)
		sum -= worst.n1

	}
	return currMax
}

func MaximumSubsequenceScore() {
	// nums1 := []int{1, 3, 3, 2}
	// nums2 := []int{2, 1, 3, 4}
	// k := 3 //12

	// nums1 := []int{4, 2, 3, 1, 1}
	// nums2 := []int{7, 5, 10, 9, 6}
	// k := 1 // 30

	// nums1 := []int{2, 1, 14, 12} // 2 0 1 3
	// nums2 := []int{11, 7, 13, 6}
	// k := 3 // 168

	nums1 := []int{22, 5, 25, 15, 28, 1}
	nums2 := []int{22, 30, 25, 25, 9, 10}
	k := 3 // 1364

	out := maxScore(nums1, nums2, k)
	fmt.Println(out)
}
