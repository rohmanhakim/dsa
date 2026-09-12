package lc3414

import (
	"slices"
	"sort"
)

func maximumWeight(intervals [][]int) []int {

	type interval struct {
		i int
		l int
		r int
		w int
	}

	sorted := []interval{}

	for i, itv := range intervals {
		sorted = append(sorted, interval{
			i: i,
			l: itv[0],
			r: itv[1],
			w: itv[2],
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].r < sorted[j].r
	})

	dp := make([][]int, len(sorted)+1)
	for i := range dp {
		dp[i] = make([]int, 5)
	}

	indices := make([][][]int, len(sorted)+1)
	for i := range indices {
		indices[i] = make([][]int, 5)
		for j := range indices[i] {
			indices[i][j] = []int{}
		}
	}

	for i := 1; i <= len(sorted); i++ {
		curr := sorted[i-1]
		idx := sort.Search(len(sorted), func(j int) bool {
			return sorted[j].r >= curr.l
		})

		for j := 1; j <= 4; j++ {
			dp[i][j] = dp[i-1][j]
			indices[i][j] = indices[i-1][j]
			if idx > 0 {
				p := idx
				if dp[p][j-1]+curr.w > dp[i][j] {
					dp[i][j] = dp[p][j-1] + curr.w
					indices[i][j] = append(indices[p][j-1], curr.i)
				} else if dp[p][j-1]+curr.w == dp[i][j] {
					var candidate []int
					if j == 1 {
						candidate = []int{curr.i}
					} else {
						candidate = slices.Clone(indices[p][j-1])
						candidate = append(candidate, curr.i)
					}
					sort.Ints(candidate)
					current := slices.Clone(indices[i][j])
					sort.Ints(current)
					if slices.Compare(candidate, current) < 0 {
						indices[i][j] = candidate
					}
				}

			}
		}
		if idx == 0 {
			for j := 1; j <= 4; j++ {
				if curr.w > dp[i][j] {
					dp[i][j] = curr.w
					indices[i][j] = []int{curr.i}
				} else if curr.w == dp[i][j] {
					candidate := []int{curr.i}
					current := slices.Clone(indices[i][j])
					sort.Ints(current)
					if slices.Compare(candidate, current) < 0 {
						indices[i][j] = candidate
					}
				}
			}

		}
	}

	bestScore := dp[len(sorted)][1]
	bestIndices := slices.Clone(indices[len(sorted)][1])

	for j := 2; j <= 4; j++ {
		if dp[len(sorted)][j] > bestScore {
			bestScore = dp[len(sorted)][j]
			bestIndices = slices.Clone(indices[len(sorted)][j])
		} else if dp[len(sorted)][j] == bestScore {
			candidate := slices.Clone(indices[len(sorted)][j])
			sort.Ints(candidate)
			current := slices.Clone(bestIndices)
			sort.Ints(current)
			if slices.Compare(candidate, current) < 0 {
				bestIndices = candidate
			}
		}
	}
	sort.Ints(bestIndices)
	return bestIndices
}
