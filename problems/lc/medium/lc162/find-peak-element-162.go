package lc162

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums)
	mid := -1
	for left < right {
		mid = left + (right-left)/2

		if mid > 0 && mid < len(nums)-1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if mid == 0 {
			return 0
		} else if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return mid
}
