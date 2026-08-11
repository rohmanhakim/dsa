package lc2996

func missingInteger(nums []int) int {
	n := len(nums)

	sum := nums[0]
	set := make(map[int]bool)
	stop := false
	set[nums[0]] = true
	for i := 1; i < n; i++ {
		set[nums[i]] = true
		if nums[i]-nums[i-1] != 1 {
			stop = true
		}
		if !stop {
			sum += nums[i]
		}
	}

	for {
		if set[sum] == false {
			return sum
		}
		sum++
	}
}
