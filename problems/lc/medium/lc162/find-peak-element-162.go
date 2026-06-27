package lc162

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + ((hi - lo) / 2)

		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func FindPeakElement() {
	// nums := []int{1, 2, 3, 1} // 2
	nums := []int{1, 2, 1, 3, 5, 6, 4} // 1 or 5
	// nums := []int{1, 2, 3} // 2
	// nums := []int{1, 2, 1, 3, 2}
	out := findPeakElement(nums)
	fmt.Println(out)
}

/// 1,2,3,4,5,6
