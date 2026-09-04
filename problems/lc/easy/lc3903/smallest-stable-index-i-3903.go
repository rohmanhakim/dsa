package lc3903

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	largest := nums[0]
	smallest := nums[n-1]

	prefixLargest := make([]int, n)
	prefixLargest[0] = largest

	suffixSmallest := make([]int, n)
	suffixSmallest[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if nums[i] > largest {
			prefixLargest[i] = nums[i]
			largest = nums[i]
		} else {
			prefixLargest[i] = largest
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] < smallest {
			suffixSmallest[i] = nums[i]
			smallest = nums[i]
		} else {
			suffixSmallest[i] = smallest
		}
	}

	for i := range nums {
		instability := prefixLargest[i] - suffixSmallest[i]
		if instability <= k {
			return i
		}
	}

	return -1
}
