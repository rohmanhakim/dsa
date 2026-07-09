package lc198

func rob(nums []int) int {
	maxCash := 0

	oddCash := nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	evenCash := max(oddCash, nums[1])
	maxCash = max(evenCash, oddCash)
	for i := 2; i < len(nums); i++ {
		even := (i+1)%2 == 0
		if even {
			stratE := evenCash + nums[i]
			stratO := oddCash
			maxCash = max(stratO, stratE)
			evenCash = maxCash
		} else {
			stratO := oddCash + nums[i]
			stratE := evenCash
			maxCash = max(stratO, stratE)
			oddCash = maxCash
		}
	}

	return maxCash
}
