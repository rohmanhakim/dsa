package hard

import (
	"fmt"
)

/*
You are given two integers num1 and num2 representing an inclusive range [num1, num2].

The waviness of a number is defined as the total count of its peaks and valleys:

    A digit is a peak if it is strictly greater than both of its immediate neighbors.
    A digit is a valley if it is strictly less than both of its immediate neighbors.
    The first and last digits of a number cannot be peaks or valleys.
    Any number with fewer than 3 digits has a waviness of 0.

Return the total sum of waviness for all numbers in the range [num1, num2].


Example 1:
Input: num1 = 120, num2 = 130
Output: 3
Explanation:
- In the range [120, 130]:
    - 120: middle digit 2 is a peak, waviness = 1.
    - 121: middle digit 2 is a peak, waviness = 1.
    - 130: middle digit 3 is a peak, waviness = 1.
    - All other numbers in the range have a waviness of 0.
- Thus, total waviness is 1 + 1 + 1 = 3.

Example 2:
Input: num1 = 198, num2 = 202
Output: 3
Explanation:
- In the range [198, 202]:
    - 198: middle digit 9 is a peak, waviness = 1.
    - 201: middle digit 0 is a valley, waviness = 1.
    - 202: middle digit 0 is a valley, waviness = 1.
    - All other numbers in the range have a waviness of 0.
- Thus, total waviness is 1 + 1 + 1 = 3.

Example 3:
Input: num1 = 4848, num2 = 4848
Output: 2
Explanation:
- Number 4848: the second digit 8 is a peak, and the third digit 4 is a valley, giving a waviness of 2.

Constraints:
    1 <= num1 <= num2 <= 10^15​​​​​​​
*/

type state struct {
	pos          int
	last         int
	secondToLast int
	tight        int
	started      int
}

// We now store both the total waviness and the count of valid branches
type dpResult struct {
	waveSum int64
	count   int64
}

func extractDigits(num int64) []int {
	// build the LSF order
	digits := []int{}
	for num > 0 {
		digit := num % 10
		digits = append(digits, int(digit))
		num = num / 10
	}

	// reverse to MSF order
	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}

	return digits
}

func solve(num int64) int64 {
	if num <= 0 {
		return 0
	}

	digits := extractDigits(num)
	memo := make(map[state]dpResult)

	res, _ := dp(state{
		pos:          0,
		last:         -1,
		secondToLast: -1,
		tight:        1,
		started:      0,
	}, digits, memo)

	return res
}

func dp(s state, digits []int, memo map[state]dpResult) (int64, int64) {
	// Base case: 0 waviness accumulated here, 1 valid suffix formed
	if s.pos == len(digits) {
		return 0, 1
	}

	if v, ok := memo[s]; ok {
		return v.waveSum, v.count
	}

	limit := 9
	if s.tight == 1 {
		limit = digits[s.pos]
	}

	var totalWave int64 = 0
	var totalCount int64 = 0

	for d := 0; d <= limit; d++ {
		newTight := 0
		if s.tight == 1 && d == limit {
			newTight = 1
		}

		if s.started == 0 && d == 0 {
			// Still in leading zero territory
			childWave, childCount := dp(state{
				pos:          s.pos + 1,
				last:         -1,
				secondToLast: -1,
				tight:        newTight,
				started:      0,
			}, digits, memo)

			totalWave += childWave
			totalCount += childCount
			continue
		}

		var wave int64 = 0
		if s.started == 1 && s.secondToLast != -1 {
			// BUG 1 FIXED: The middle digit is `s.last`
			if s.last > s.secondToLast && s.last > d {
				wave = 1
			} else if s.last < s.secondToLast && s.last < d {
				wave = 1
			}
		}

		childWave, childCount := dp(state{
			pos:          s.pos + 1,
			last:         d,
			secondToLast: s.last,
			tight:        newTight,
			started:      1,
		}, digits, memo)

		// BUG 2 FIXED: Multiply the current wave by the childCount
		totalWave += childWave + (wave * childCount)
		totalCount += childCount
	}

	memo[s] = dpResult{waveSum: totalWave, count: totalCount}
	return totalWave, totalCount
}

func totalWaviness(num1 int64, num2 int64) int64 {
	return solve(num2) - solve(num1-1)
}

func TotalWavinessOfNumbersInRangeII() {
	num1 := 120
	num2 := 130
	out := totalWaviness(int64(num1), int64(num2))

	fmt.Println(out)

	// fmt.Println(120 % 100)
}
