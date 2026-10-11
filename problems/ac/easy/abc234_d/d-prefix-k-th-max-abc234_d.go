package abc234d

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

func dPrefixKthMax(n int, k int, perm []int) {
	h := minHeap([]int{})
	heap.Init(&h)

	for i := 0; i < k-1; i++ {
		heap.Push(&h, perm[i])
	}

	for i := k; i <= n; i++ {
		heap.Push(&h, perm[i-1])
		if len(h) > k {
			heap.Pop(&h)
		}

		fmt.Println(h[0])
	}
}

func dPrefixKthMaxAlternative(n int, k int, perm []int) {
	h := minHeap([]int{})
	heap.Init(&h)

	for i := 0; i < k; i++ {
		heap.Push(&h, perm[i])
	}

	fmt.Println(h[0])

	for i := k; i < n; i++ {
		if perm[i] > h[0] {
			h[0] = perm[i]
			heap.Fix(&h, 0)
		}

		fmt.Println(h[0])
	}
}

func DPrefixKthMax() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	perm := []int{}
	for i := 1; i <= n; i++ {
		var s int
		fmt.Scanf("%d", &s)
		perm = append(perm, s)
	}

	dPrefixKthMaxAlternative(n, k, perm)
}
