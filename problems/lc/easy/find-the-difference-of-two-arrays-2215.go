package easy

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	appear1 := make(map[int]int)
	appear2 := make(map[int]int)
	for _, n := range nums1 {
		appear1[n] = 1
	}
	for _, n := range nums2 {
		appear2[n] = 2
	}
	only1 := []int{}
	only2 := []int{}
	for k := range appear1 {
		if _, ok := appear2[k]; !ok {
			only1 = append(only1, k)
		}
	}
	for k := range appear2 {
		if _, ok := appear1[k]; !ok {
			only2 = append(only2, k)
		}
	}
	return [][]int{only1, only2}
}

func FindTheDifferenceOfTwoArrays() {
	// nums1 := []int{1, 2, 3}
	// nums2 := []int{2, 4, 6}
	// nums1 := []int{1, 2, 3, 3}
	// nums2 := []int{1, 1, 2, 2}
	nums1 := []int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}
	nums2 := []int{-1, -40, -44, 41, 10, -43, 69, 10, 2}

	out := findDifference(nums1, nums2)
	fmt.Println(out)
}
