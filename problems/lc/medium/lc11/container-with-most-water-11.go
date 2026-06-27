package lc11

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxVal := 0
	for left < right {
		lh := height[left]
		rh := height[right]
		w := right - left

		area := w * min(lh, rh)
		if area > maxVal {
			maxVal = area
		}

		if lh > rh {
			right--
		} else {
			left++
		}
	}
	return maxVal
}

func ContainerWithMostWater() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{8, 7, 2, 1}
	out := maxArea(height)
	fmt.Println(out)
}
