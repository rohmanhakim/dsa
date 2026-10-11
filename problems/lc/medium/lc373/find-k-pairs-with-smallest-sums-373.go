package lc373

import (
	"container/heap"
	"sort"
)

type item struct {
	a   int
	b   int
	row int
	col int
	sum int
}

type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	h := minHeap([]item{})

	m := len(nums1)
	n := len(nums2)

	res := []item{}

	for i := 0; i < m; i++ {
		heap.Push(&h, item{
			a:   nums1[i],
			b:   nums2[0],
			row: i,
			col: 0,
			sum: nums1[i] + nums2[0],
		})
	}

	for {
		smallest := heap.Pop(&h).(item)

		res = append(res, smallest)

		if len(res) == k {
			break
		}

		if smallest.col < n-1 {
			row := smallest.row
			col := smallest.col + 1

			heap.Push(&h, item{
				a:   nums1[row],
				b:   nums2[col],
				row: row,
				col: col,
				sum: nums1[row] + nums2[col],
			})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].sum != res[j].sum {
			return res[i].sum < res[j].sum
		}
		return res[i].a < res[j].a
	})

	ans := [][]int{}

	for _, r := range res {
		ans = append(ans, []int{r.a, r.b})
	}

	return ans
}
