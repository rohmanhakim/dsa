package lc350

import (
	"slices"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	res := []int{}
	sorted := slices.Clone(nums2)
	sort.Ints(sorted)
	for _, a := range nums1 {
		left := 0
		right := len(sorted) - 1

		for left <= right {
			mid := left + ((right - left) / 2)

			if sorted[mid] == a {
				res = append(res, a)
				sorted = slices.Delete(sorted, mid, mid+1)
				break
			} else if sorted[mid] < a {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return res
}
