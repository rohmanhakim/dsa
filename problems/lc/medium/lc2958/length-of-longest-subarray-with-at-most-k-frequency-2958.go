package lc2958

func maxSubarrayLength(nums []int, k int) int {

	n := len(nums)

	left := 0

	freq := make(map[int]int)

	longest := 0
	for right := 0; right < n; right++ {
		curr := nums[right]

		freq[curr]++

		for freq[curr] > k {
			freq[nums[left]]--
			left++
		}

		wLen := right - left + 1
		if wLen > longest {
			longest = wLen
		}
	}
	return longest
}
