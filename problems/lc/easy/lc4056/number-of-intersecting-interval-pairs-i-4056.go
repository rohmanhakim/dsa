package lc4056

func countIntersectingIntervals(intervals [][]int) int {
	total := 0
	for i, in := range intervals {
		prevStart := in[0]
		prevEnd := in[1]
		for j := i + 1; j < len(intervals); j++ {
			start := intervals[j][0]
			end := intervals[j][1]
			if start <= prevEnd && end >= prevStart {
				total++
			}

		}
	}
	return total
}
