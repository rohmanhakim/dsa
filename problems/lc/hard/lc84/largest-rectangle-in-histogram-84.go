package lc84

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := []int{}

	maxArea := 0
	left := -1
	for i := 0; i < len(heights); i++ {
		if heights[i] > 0 {
			for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
			for j := len(stack) - 1; j > 0; j-- {
				w := (stack[j] - stack[j-1] + stack[len(stack)-1] - stack[j]) //(len(stack) - j - 1))
				area := w * heights[stack[j]]
				if area > maxArea {
					maxArea = area
				}
			}
			area := heights[stack[0]] * (i - left)

			if area > maxArea {
				maxArea = area
			}
		} else {
			left = i
			stack = []int{}
		}
	}
	return maxArea
}

func largestRectangleAreaClaudeSonnet(heights []int) int {
	// Append a sentinel 0 to flush remaining stack at the end
	heights = append(heights, 0)
	stack := []int{} // stores indices, monotonically increasing by height
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		// Pop whenever current bar is shorter than stack top
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// Width: from the new stack top (exclusive) to i (exclusive)
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			area := 0
			if area = h * w; area > maxArea {
				maxArea = area
			}

			if area == 10 {
				fmt.Printf("===> i: %d, w: %d\n", i, w)
			}
		}
		stack = append(stack, i)
		fmt.Println(stack)
	}

	return maxArea
}

func LargestRectangeInHistogram() {
	heights := []int{2, 1, 5, 6, 2, 3}
	// heights := []int{3, 2, 6, 5, 1, 2}
	// heights := []int{2, 4}
	// heights := []int{1, 2, 3, 4, 5, 6, 7}
	// heights := []int{7, 6, 5, 4, 3, 2, 1}
	// heights := []int{1, 1, 1, 1, 1, 1}
	// heights := []int{2, 3}
	// heights := []int{4, 2, 0, 3, 2, 4, 3, 4}
	// 0,2,3,4
	// heights := []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}
	out := largestRectangleAreaClaudeSonnet(heights)
	fmt.Println(out)
}
