package lc154

import "fmt"

func findMinWithDup(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	min := 0

	for left < right {
		mid = left + (right-left)/2
		fmt.Printf("nums[%d]: %d, nums[%d]: %d, nums[%d]: %d\n", left, nums[left], mid, nums[mid], right, nums[right])
		if right-left == 1 {
			if nums[left] < nums[right] {
				min = left
				break
			} else {
				min = right
				break
			}
		}

		if nums[left] < nums[right] {
			right = mid
		} else if nums[left] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			if nums[left] < nums[mid] {
				left = mid
			} else if nums[left] > nums[mid] {
				right = mid
			} else {
				left++
			}
		} else if nums[mid] > nums[right] {
			if nums[left] < nums[mid] {
				left = mid
			} else if nums[left] > nums[mid] {
				right = mid
			} else {
				left++
			}
		} else {
			right--
		}
	}
	return nums[min]
}

func findMinWithDupClaudeSonnet(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// min is in right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// min is in left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right]: can't determine which side, shrink right
			right--
		}
	}

	return nums[left]
}

func FindMinimumInRotatedSortedArray2() {
	// nums := []int{4, 5, 6, 7, 0, 1, 4} // 0
	// nums := []int{1, 2, 3, 4, 5, 6, 7} // 1
	// nums := []int{10, 10, 1, 2, 10, 10, 10, 10, 10} // 1
	// nums := []int{10, 1, 10, 10, 10} // 1
	// nums := []int{3, 1, 2, 2, 2}
	// nums := []int{3, 4, 4, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, -10, -10, -10, -9, -8, -8, -8, -8, -8, -7, -7, -7, -7, -6, -6, -6, -6, -6, -6, -6, -5, -5, -5, -4, -4, -4, -4, -3, -3, -3, -3, -3, -3, -2, -2, -2, -2, -1, -1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3}
	// nums := []int{3, 4, 10, -10, -8, -4, -1, 0, 2} // -10
	nums := []int{13, 14, 20, 0, 2, 6, 9, 10, 12} // 0
	out := findMinWithDup(nums)
	fmt.Println(out)
}
