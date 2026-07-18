package lc2831

func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)

	freq := make(map[int]int)
	left := 0
	right := 0

	maxFreq := 1

	for left < n-1 && right < n {
		num := nums[right]

		freq[num]++

		maxFreq = max(maxFreq, freq[num])

		windowSize := right - left + 1

		for windowSize-maxFreq > k {
			freq[nums[left]]--
			left++
			windowSize--
		}
		right++
	}

	return maxFreq
}
