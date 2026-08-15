package lc3702

func longestSubsequence(nums []int) int {

	n := len(nums)

	totalXor := nums[0]
	anyNonZero := false
	if nums[0] != 0 {
		anyNonZero = true
	}

	for i := 1; i < n; i++ {
		totalXor ^= nums[i]
		if nums[i] != 0 {
			anyNonZero = true
		}
	}

	if totalXor != 0 {
		return n
	} else if anyNonZero {
		return n - 1
	} else {
		return 0
	}
}
