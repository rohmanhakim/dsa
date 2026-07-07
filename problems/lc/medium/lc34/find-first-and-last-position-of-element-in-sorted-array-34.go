package lc34

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)

	res := []int{-1, -1}

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left >= len(nums) || nums[left] != target {
		return res
	}
	res[0] = left

	right = len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	res[1] = left - 1

	return res
}
