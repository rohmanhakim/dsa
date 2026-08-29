package lc2948

import (
	"slices"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {

	type number struct {
		val int
		idx int
	}
	sorted := []number{}
	for i, num := range nums {
		sorted = append(sorted, number{
			val: num,
			idx: i,
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].val < sorted[j].val
	})

	segments := [][]number{}
	left := 0
	right := 1
	for right < len(sorted) {
		if sorted[right].val-sorted[right-1].val > limit {
			segments = append(segments, sorted[left:right])
			left = right

		}
		right++
		if right == len(sorted) {
			segments = append(segments, sorted[left:right])
		}
	}

	res := slices.Clone(nums)
	for _, segment := range segments {
		sortedIdx := make([]int, len(segment))
		for i, se := range segment {
			sortedIdx[i] = se.idx
		}
		sort.Ints(sortedIdx)
		for i, idx := range sortedIdx {
			res[idx] = segment[i].val
		}
	}

	return res
}
