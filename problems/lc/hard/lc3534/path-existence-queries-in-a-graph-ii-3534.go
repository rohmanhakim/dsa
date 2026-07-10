package lc3534

import (
	"sort"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	LOG := 0
	for (1 << LOG) <= n {
		LOG++
	}

	sortedIdx := make([]int, n)
	for i := range n {
		sortedIdx[i] = i
	}

	sort.Slice(sortedIdx, func(i, j int) bool {
		return nums[sortedIdx[i]] < nums[sortedIdx[j]]
	})

	reach := make([]int, n)
	for i := 0; i < n; i++ {
		idx := sort.Search(n, func(j int) bool {
			return nums[sortedIdx[j]]-nums[sortedIdx[i]] > maxDiff
		})
		reach[i] = idx - 1
	}

	jump := make([][]int, LOG)
	jump[0] = reach

	for k := 1; k < LOG; k++ {
		jump[k] = make([]int, n)
		for i := 0; i < n; i++ {
			jump[k][i] = jump[k-1][jump[k-1][i]]
		}
	}

	breaks := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[sortedIdx[i]]-nums[sortedIdx[i-1]] > maxDiff {
			breaks[i-1] = 1
		}
	}

	prefixBreaks := make([]int, len(nums))
	for i := 1; i < len(breaks); i++ {
		prefixBreaks[i] = breaks[i-1] + prefixBreaks[i-1]
	}

	pos := make([]int, n)
	for sortedPos, origIdx := range sortedIdx {
		pos[origIdx] = sortedPos
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		u := q[0]
		v := q[1]
		if u == v {
			res[i] = 0
		} else if prefixBreaks[pos[u]] != prefixBreaks[pos[v]] {
			res[i] = -1
		} else {
			a := min(pos[u], pos[v])
			b := max(pos[u], pos[v])
			steps := 0
			for k := LOG - 1; k >= 0; k-- {
				if jump[k][a] < b {
					a = jump[k][a]
					steps += 1 << k
				}
			}
			res[i] = steps + 1
		}

	}

	return res
}
