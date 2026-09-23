package lc1658

func minOperations(nums []int, x int) int {

	n := len(nums)

	total := 0

	for _, num := range nums {
		total += num
	}

	target := total - x

	if target < 0 {
		return -1
	}
	if target == 0 {
		return n
	}

	left := 0
	sum := 0
	longest := 0
	for right := 0; right < n; right++ {
		curr := nums[right]
		sum += curr

		for sum > target {
			sum -= nums[left]
			left++
		}

		wlen := right - left + 1
		if sum == target {
			if wlen > longest {
				longest = wlen
			}
		}
	}

	if longest == 0 {
		return -1
	}
	return n - longest
}
