package lc930

func numSubarraysWithSum(nums []int, goal int) int {
	n := len(nums)
	left := 0
	count := 0
	sum := 0
	prefixZero := 0
	for right := 0; right < n; right++ {
		curr := nums[right]
		sum += curr
		for (right > left && sum > goal) || (right > left && nums[left] == 0) {
			sum -= nums[left]
			if nums[left] == 0 {
				prefixZero++
			} else {
				prefixZero = 0
			}
			left++
		}
		if sum == goal {
			count += prefixZero + 1
		}
	}
	return count
}
