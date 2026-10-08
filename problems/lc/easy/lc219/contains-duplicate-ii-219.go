package lc219

func containsNearbyDuplicate(nums []int, k int) bool {

	freq := make(map[int]int)
	n := len(nums)
	left := 0

	for right := 0; right < n; right++ {
		freq[nums[right]]++

		for right-left > k {
			freq[nums[left]]--
			left++
		}

		if freq[nums[right]] > 1 {
			return true
		}
	}
	return false
}
