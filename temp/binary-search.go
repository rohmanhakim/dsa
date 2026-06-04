package temp

import "fmt"

func mainBinarySearch() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 8

	left := 0
	right := len(nums) - 1
	mid := 0
	index := -1
	for left < right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			index = mid
			break
		}

		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Println(index)
}

func mainBinarySearchRotated() {

	nums := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	target := 6

	left := 0
	right := len(nums) - 1
	mid := 0
	index := -1
	for left < right {
		mid = left + (right-left)/2
		fmt.Printf("l: %d, m: %d, r: %d\n", left, mid, right)
		if nums[mid] == target {
			index = mid
			break
		}

		if nums[mid] < nums[right] {
			// take right
			if target > nums[mid] && target < nums[right] {
				left = mid
			} else {
				//take left
				right = mid
			}
		} else {
			// left must be ordered
			// take left
			if target < nums[mid] && target > nums[left] {
				right = mid
			} else {
				// take right
				left = mid
			}
		}
	}

	fmt.Println(index)
}

func minInRotatedSorted() {
	nums := []int{2, 2, 2, 4, 0, 1}
	min := -1

	left := 0
	right := len(nums) - 1
	mid := 0
	for left < right {
		mid = left + (right-left)/2

		if nums[mid] < nums[right] {
			if right-left == 1 {
				min = left
				break
			}
			right = mid
		} else {
			if right-left == 1 {
				min = right
				break
			}
			left = mid
		}
	}

	fmt.Println(nums[min])
}
