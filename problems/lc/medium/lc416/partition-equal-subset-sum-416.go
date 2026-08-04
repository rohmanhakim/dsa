package lc416

func canPartition(nums []int) bool {

	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 == 1 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)

	dp[0] = true

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
