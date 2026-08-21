package lc3116

import (
	"math/bits"
	"slices"
)

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcd(a, b)) * b
}

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)

	count := func(x int64) int64 {
		ans := int64(0)
		for mask := 1; mask < (1 << n); mask++ {
			l := int64(1)
			valid := true

			for i := 0; i < n; i++ {
				if mask&(1<<i) == 0 {
					continue
				}

				l = lcm(l, int64(coins[i]))

				if l > x {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			cnt := x / l

			if bits.OnesCount(uint(mask))%2 == 1 {
				ans += cnt
			} else {
				ans -= cnt
			}
		}
		return ans
	}

	left := int64(slices.Min(coins))
	right := int64(k) * left

	for left < right {
		mid := left + ((right - left) / 2)

		if count(mid) < int64(k) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(left)
}
