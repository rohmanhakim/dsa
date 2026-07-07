package lc153

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// left sorted, minimum must be on the right
			left = mid + 1
		} else {
			// right sorted, minimum must be on the left (including mid)
			right = mid
		}
	}

	return nums[left]
}
