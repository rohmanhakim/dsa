package medium

import (
	"fmt"
)

/*
func search(nums []int, target int) int {
	loc := -1
	begin := 0
	end := len(nums) - 1

	for begin <= end {
		mid := begin + (end-begin)/2

		if target == nums[mid] {
			loc = mid
			break
		}
		if nums[begin] <= nums[mid] {
			if target >= nums[begin] && target <= nums[mid] {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				begin = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return loc
}
*/

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] { // right sorted
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // left sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func SearchInRotatedSortedArray() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// out := search(nums, 0)
	// nums := []int{8, 9, 3, 4, 5, 6, 7}
	// out := search(nums, 3)
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// out := search(nums, 4)
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// out := search(nums, 3)
	nums := []int{1}
	out := search(nums, 1)
	// nums := []int{1, 3}
	// out := search(nums, 3)
	fmt.Println(out)
}
