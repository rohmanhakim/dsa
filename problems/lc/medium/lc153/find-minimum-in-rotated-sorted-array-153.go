package lc153

import "fmt"

func binarySearchFindMinRotatedArray(nums []int, begin int, end int, min *int) {
	diff := end - begin
	mid := begin + (end-begin)/2

	if nums[begin] < nums[end] {
		*min = nums[begin]
	} else if nums[begin] > nums[mid] { // rotated, go left
		if diff == 1 {
			*min = nums[begin]
		} else {
			binarySearchFindMinRotatedArray(nums, begin, mid, min)
		}
	} else {
		if diff == 1 {
			*min = nums[end]
		} else {
			binarySearchFindMinRotatedArray(nums, mid, end, min)
		}
	}
}

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	min := 0
	binarySearchFindMinRotatedArray(nums, 0, len(nums)-1, &min)

	return min
}

func findMinClaudeSonnet(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// min is in the right half
			left = mid + 1
		} else {
			// min is in the left half (including mid)
			right = mid
		}
	}

	return nums[left]
}

func FindMinimumInRotatedSortedArray() {

	nums := []int{11, 13, 15, 17} // 11
	// nums := []int{4, 5, 6, 7, 0, 1, 2} // 0
	// nums := []int{5, 1, 2, 3, 4} // 1
	out := findMin(nums)

	fmt.Println(out)
}
