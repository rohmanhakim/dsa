package lc53

func maxSubArray(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	answer := nums[0]
	prev := nums[0]
	for i := 1; i < n; i++ {
		extend := prev + nums[i]
		fresh := nums[i]
		prev = max(extend, fresh)
		answer = max(answer, prev)
	}

	return answer
}
