package lc1520

import (
	"sort"
)

func maxNumOfSubstrings(s string) []string {
	indices := make([][]int, 26)

	for i := range indices {
		indices[i] = []int{}
	}

	for i, ch := range s {
		if len(indices[ch-'a']) == 0 {
			indices[ch-'a'] = []int{i, i}
		} else {
			indices[ch-'a'][1] = i
		}
	}

	for i := 0; i < 26; i++ {
		if len(indices[i]) == 0 {
			continue
		}

		left, right := indices[i][0], indices[i][1]

		changed := true

		for changed {
			changed = false
			for k := left; k <= right; k++ {
				ch := s[k] - 'a'
				if len(indices[ch]) > 0 {
					if indices[ch][0] < left {
						left = indices[ch][0]
						changed = true
					}
					if indices[ch][1] > right {
						right = indices[ch][1]
						changed = true
					}
				}
			}
		}
		indices[i][0], indices[i][1] = left, right
	}

	type interval struct {
		char  rune
		start int
		end   int
	}

	intervals := []interval{}

	for i, idx := range indices {
		if len(idx) > 0 {
			intervals = append(intervals, interval{
				char:  rune(i + 'a'),
				start: idx[0],
				end:   idx[1],
			})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	fintervals := []interval{}
	curr := interval{
		end: -1,
	}
	for i := 0; i < len(intervals); i++ {
		if intervals[i].start > curr.end {
			fintervals = append(fintervals, intervals[i])
			curr = intervals[i]
		}
	}

	res := []string{}

	for _, itv := range fintervals {
		res = append(res, s[itv.start:itv.end+1])
	}
	return res
}
