package lc33

import (
	"sort"
)

func search(nums []int, target int) int {

	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] < nums[0]
	})

	bsearch := func(left, right int) int {
		for left <= right {
			mid := left + ((right - left) / 2)

			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				return mid
			}
		}
		return -1
	}

	if idx < len(nums) {
		if target >= nums[idx] && target <= nums[len(nums)-1] {
			return bsearch(idx, len(nums)-1)
		} else {
			return bsearch(0, idx-1)
		}
	} else {
		return bsearch(0, len(nums)-1)
	}
}
