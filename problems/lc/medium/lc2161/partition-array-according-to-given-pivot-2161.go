package lc2161

import "fmt"

func pivotArray(nums []int, pivot int) []int {

	nums = append([]int{0}, nums...)

	read := 1
	write := 0
	n := len(nums)
	pivotCount := 0
	right := []int{}
	for read < n {
		if nums[read] > pivot {
			right = append(right, nums[read])
		} else if nums[read] < pivot {
			nums[write] = nums[read]
			write++
		} else {
			pivotCount++
		}
		read++
	}
	nums = nums[:write]
	for range pivotCount {
		nums = append(nums, pivot)
	}
	for _, n := range right {
		nums = append(nums, n)
	}

	return nums
}

func pivotArrayClaudeSonnet(nums []int, pivot int) []int {
	less, equal, greater := []int{}, []int{}, []int{}

	for _, n := range nums {
		switch {
		case n < pivot:
			less = append(less, n)
		case n == pivot:
			equal = append(equal, n)
		default:
			greater = append(greater, n)
		}
	}

	return append(append(less, equal...), greater...)
}

func PartitionArrayAccordingToGivenPivot() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	out := pivotArray(nums, pivot)
	fmt.Println(out)
}
