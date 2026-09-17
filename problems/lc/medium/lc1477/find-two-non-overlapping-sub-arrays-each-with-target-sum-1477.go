package lc1477

import (
	"math"
)

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)

	dp := make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt
	}

	left := 0
	sum := 0

	best := math.MaxInt

	for right := 0; right < n; right++ {
		sum += arr[right]

		for sum > target {
			sum -= arr[left]
			left++
		}

		if sum == target {
			wlen := right - left + 1
			candidate := wlen
			if left > 0 && dp[left-1] < math.MaxInt {
				candidate = dp[left-1] + wlen
				if candidate < best {
					best = candidate
				}
			}

			if right > 0 {
				dp[right] = min(dp[right-1], wlen)
			} else {
				dp[right] = wlen
			}
		} else {
			if right > 0 {
				dp[right] = dp[right-1]
			}
		}
	}

	if best > n {
		return -1
	}
	return best
}
