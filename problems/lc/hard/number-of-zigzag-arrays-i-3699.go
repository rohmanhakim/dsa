package hard

import "fmt"

/*
You are given three integers n, l, and r.

A ZigZag array of length n is defined as follows:

Each element lies in the range [l, r].
No two adjacent elements are equal.
No three consecutive elements form a strictly increasing or strictly decreasing sequence.
Return the total number of valid ZigZag arrays.

Since the answer may be large, return it modulo 109 + 7.

A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

A sequence is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

Example 1:

Input: n = 3, l = 4, r = 5

Output: 2

Explanation:

There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:

[4, 5, 4]
[5, 4, 5]​​​​​​​

Example 2:

Input: n = 3, l = 1, r = 3

Output: 10

Explanation:

There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

[1, 2, 1], [1, 3, 1], [1, 3, 2]
[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
[3, 1, 2], [3, 1, 3], [3, 2, 3]
All arrays meet the ZigZag conditions.

Constraints:

3 <= n <= 2000
1 <= l < r <= 2000
*/

func zigZagArraysTopDown(n int, l int, r int) int {
	const MOD = 1000000007

	memoUp := make([][]int, n)
	memoDown := make([][]int, n)
	for i := range n {
		rangeLength := (r - l + 1)
		for j := 0; j <= rangeLength; j++ {
			memoUp[i] = append(memoUp[i], -1)
			memoDown[i] = append(memoDown[i], -1)
		}
	}

	var dp func(pos int, isUp bool, val int) int
	dp = func(pos int, isUp bool, val int) int {
		if pos == 0 {
			return 1
		}

		count := 0

		rangeStart, rangeEnd := l, val-1
		if !isUp {
			rangeStart, rangeEnd = val+1, r
		}
		for prevVal := rangeStart; prevVal <= rangeEnd; prevVal++ {
			if prevVal == val {
				continue
			}

			if pos-1 >= 1 {
				res := 0

				if isUp {
					if memoDown[pos-1][prevVal-l] != -1 {
						res = memoDown[pos-1][prevVal-l]
					} else {
						res = dp(pos-1, false, prevVal)
						memoDown[pos-1][prevVal-l] = res
					}
				} else {
					if memoUp[pos-1][prevVal-l] != -1 {
						res = memoUp[pos-1][prevVal-l]
					} else {
						res = dp(pos-1, true, prevVal)
						memoUp[pos-1][prevVal-l] = res
					}
				}
				count = (count + res) % MOD

			} else {

				count = (count + 1) % MOD
			}
		}
		return count
	}

	totalCount := 0
	for i := r; i >= l; i-- {
		totalCount = (totalCount + dp(n-1, true, i)) % MOD
		totalCount = (totalCount + dp(n-1, false, i)) % MOD
	}
	return totalCount
}

func zigZagArrays(n int, l int, r int) int {
	const MOD = 1000000007

	rangeLength := r - l + 1
	count := make([]int, rangeLength)
	for i := l; i <= r; i++ {
		count[i-l] = 1
	}

	var up, down []int

	for i := 1; i <= n-1; i++ {

		prefixDown := make([]int, (rangeLength + 1))
		prefixDown[0] = 0
		suffixUp := make([]int, (rangeLength + 1))
		suffixUp[rangeLength] = 0

		if i > 1 {
			for j := 1; j <= rangeLength; j++ {
				prefixDown[j] = (prefixDown[j-1] + down[j-1]) % MOD
			}

			for j := rangeLength - 1; j >= 0; j-- {
				suffixUp[j] = (suffixUp[j+1] + up[j]) % MOD
			}
		} else {
			for j := 1; j <= rangeLength; j++ {
				prefixDown[j] = prefixDown[j-1] + count[j-1]
			}

			for j := rangeLength - 1; j >= 0; j-- {
				suffixUp[j] = suffixUp[j+1] + count[j]
			}
		}

		newUp := make([]int, rangeLength)
		newDown := make([]int, rangeLength)
		for j := l; j <= r; j++ {
			newUp[j-l] = prefixDown[j-l]
			newDown[j-l] = suffixUp[j-l+1]
		}

		up, down = newUp, newDown
	}

	answer := 0
	for i := 0; i < rangeLength; i++ {
		answer = (answer + up[i] + down[i]) % MOD
	}
	return answer
}

func NumberofZigZagArraysI() {
	n := 3
	l := 1
	r := 3
	out := zigZagArrays(n, l, r)
	fmt.Println(out)
}
