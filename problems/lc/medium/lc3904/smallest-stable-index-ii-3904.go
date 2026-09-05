package lc3904

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	smallestSuffix := make([]int, n)

	smallestSuffix[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		smallestSuffix[i] = min(smallestSuffix[i+1], nums[i])
	}

	largest := nums[0]
	for i := range n {
		if nums[i] > largest {
			largest = nums[i]
		}

		instability := largest - smallestSuffix[i]

		if instability <= k {
			return i
		}
	}

	return -1
}
