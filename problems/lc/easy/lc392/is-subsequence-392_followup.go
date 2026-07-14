package lc392

import (
	"sort"
)

func isSubsequenceFollowUp(stream []string, t string) []bool {
	nt := len(t)

	idx := make(map[rune][]int)

	for i := 0; i < nt; i++ {
		ch := rune(t[i])
		if _, ok := idx[ch]; !ok {
			idx[ch] = []int{i}
		} else {
			idx[ch] = append(idx[ch], i)
		}
	}

	res := make([]bool, len(stream))
	for i, s := range stream {

		prev := -1
		point := 0
		for _, ch := range s {
			if val, ok := idx[ch]; ok {
				idx := sort.Search(len(val), func(j int) bool {
					return val[j] > prev
				})
				if idx < len(val) {
					prev = val[idx]
					point++
				}
			}
		}
		res[i] = point == len(s)
	}

	return res
}
