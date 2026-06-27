package lc1046

import (
	"container/heap"
	"fmt"
)

/*
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})

	for len(stones) > 1 {

		y := stones[0]
		x := stones[1]
		if x > y {
			x, y = y, x
		}
		res := y - x
		stones = stones[2:]
		if res > 0 {
			stones = append(stones, res)
			sort.Slice(stones, func(i, j int) bool {
				return stones[i] > stones[j]
			})
		}
	}
	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}
*/

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		res := y - x
		fmt.Printf("y: %d, x: %d, res: %d\n", y, x, res)
		if res > 0 {
			heap.Push(&h, res)
		}
	}

	if h.Len() == 1 {
		return h[0]
	}

	// for h.Len() > 0 {
	// 	fmt.Printf("%d, ", h.Pop())
	// }

	// fmt.Println("")

	return 0
}

func LastStoneWeight() {
	stones := []int{2, 7, 4, 1, 8, 1}
	// stones := []int{2, 2}
	// stones := []int{9, 10, 1, 7, 3}
	out := lastStoneWeight(stones)

	fmt.Println(out)
}
