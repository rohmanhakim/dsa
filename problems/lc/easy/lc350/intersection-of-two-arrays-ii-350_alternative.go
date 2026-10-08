package lc350

import (
	"slices"
	"sort"
)

func intersectAlternative(nums1 []int, nums2 []int) []int {

	/* assume both inputs sorted */
	sorted1 := slices.Clone(nums1)
	sort.Ints(sorted1)
	sorted2 := slices.Clone(nums2)
	sort.Ints(sorted2)

	i := 0
	j := 0

	res := []int{}
	for i < len(sorted1) && j < len(sorted2) {
		if sorted1[i] < sorted2[j] {
			i++
		} else if sorted1[i] > sorted2[j] {
			j++
		} else if sorted1[i] == sorted2[j] {
			res = append(res, sorted1[i])
			i++
			j++
		}
	}

	return res
}
