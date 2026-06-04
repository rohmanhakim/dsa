package medium

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	q := 0
	negativeDivisor := false
	negativeDividend := false
	negativeResult := false
	if divisor < 0 {
		negativeDivisor = true
	}
	if dividend < 0 {
		negativeDividend = true
	}

	negativeResult = !(negativeDividend == negativeDivisor)

	if !negativeResult && dividend-divisor == 0 || negativeResult && dividend+divisor == 0 {
		if negativeResult {
			return -1
		}

		return 1
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}

		return -dividend
	}

	if dividend == 0 {
		return 0
	}

	if dividend > divisor {
		if negativeDivisor {
			divisor = -divisor
		}
		for dividend >= divisor {
			dividend -= divisor
			q++
		}
	} else {
		if negativeDivisor {
			divisor = -divisor
		}
		for dividend < 0 {
			dividend += divisor
			if dividend > 0 {
				break
			}
			q++
		}
	}

	if negativeResult {
		return 0 - q
	}
	return q
}

func divideClaudeSonnet(dividend int, divisor int) int {
	// Handle the one overflow case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Determine sign, then work in negatives (avoids overflow)
	negative := (dividend < 0) != (divisor < 0)
	a, b := toNeg(dividend), toNeg(divisor)

	result := 0
	// a and b are both negative; we want a <= b (closer to 0 means smaller magnitude)
	for a <= b {
		temp, mult := b, -1
		// Double temp until it would overshoot a
		for temp >= math.MinInt32/2 && a <= temp+temp {
			temp += temp // temp doubles (stays negative)
			mult += mult // tracks how many divisors we've consumed
		}
		a -= temp      // subtract the largest valid chunk
		result -= mult // accumulate (mult is negative, so this adds)
	}

	if negative {
		return -result
	}
	return result
}

func toNeg(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func DivideTwoIntegers() {
	dividend := math.MaxInt32
	divisor := -1

	out := divide(dividend, divisor)

	fmt.Println(out)
}
