package lc1004

import "fmt"

func longestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	max := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		length := right - left + 1

		if length > max {
			max = length
		}
	}
	return max
}

func longestOnesClaudeSonnet(nums []int, k int) int {
	/*
		you use an inner while loop to shrink the window until it's valid again.
		This is fine, but you can simplify it to a single if — instead of
		restoring validity, just slide the window forward by 1, keeping the
		window size the same or growing it.

	    The key insight: you never need a window smaller than your current max,
		so you can just shift the left pointer forward by 1 unconditionally
		whenever zeros > k.


	*/
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}
		// Instead of shrinking to valid, just slide forward by 1
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		// No need to track max — window only ever grows or shifts
	}
	return len(nums) - left

	/*
		Why this works:
		- The window [left, right] only grows when a valid extension is found,
		and slides (same size) otherwise — it never shrinks.
		- Since you only care about the maximum window ever seen, and the
		- window never shrinks below that max, len(nums) - left at the end
		equals that max length.
		- k is repurposed to track remaining flips (goes negative when over
		budget), eliminating the zeros variable.
	*/
}

func MaxConsecutiveOnesIII() {
	// nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	// k := 3
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	out := longestOnes(nums, k)
	fmt.Println(out)
}
