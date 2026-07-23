package lc213

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	computeDP := func(subNums []int) int {
		n := len(subNums)

		dp := make([]int, n)

		dp[0] = subNums[0]
		if len(subNums) == 1 {
			return subNums[0]
		}
		dp[1] = subNums[1]

		if subNums[0] > subNums[1] {
			dp[1] = subNums[0]
		}

		for i := 2; i < n; i++ {
			toRob := subNums[i] + dp[i-2]
			toSkip := dp[i-1]

			dp[i] = max(toRob, toSkip)
		}

		return dp[n-1]
	}

	withFirstHouse := computeDP(nums[:len(nums)-1])
	withoutFirstHouse := computeDP(nums[1:])

	return max(withFirstHouse, withoutFirstHouse)
}
