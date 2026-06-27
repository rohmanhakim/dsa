package lc3700

import "fmt"

/*
	You are given three integers n, l, and r.

	A ZigZag array of length n is defined as follows:

	Each element lies in the range [l, r].
	No two adjacent elements are equal.
	No three consecutive elements form a strictly increasing or strictly decreasing sequence.
	Return the total number of valid ZigZag arrays.

	Since the answer may be large, return it modulo 10^9 + 7.

	A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

	A sequence is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

	Example 1:

	Input: n = 3, l = 4, r = 5

	Output: 2

	Explanation:

	There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:

	[4, 5, 4]
	[5, 4, 5]

	Example 2:

	Input: n = 3, l = 1, r = 3

	Output: 10

	Explanation:

	​​​​​​​There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

	[1, 2, 1], [1, 3, 1], [1, 3, 2]
	[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
	[3, 1, 2], [3, 1, 3], [3, 2, 3]
	All arrays meet the ZigZag conditions.


	Constraints:

	3 <= n <= 10^9
	1 <= l < r <= 75​​​​​​​
*/

func zigZagArraysII(n int, l int, r int) int {
	const MOD = 1e9 + 7

	rangeLength := r - l + 1
	m := make([][]int, 2*rangeLength)
	for i := range len(m) {
		m[i] = make([]int, 2*rangeLength)
	}

	for i := 0; i < rangeLength; i++ {
		for j := 0; j < i; j++ {
			m[i][rangeLength+j] = 1
		}
		for j := i + 1; j < rangeLength; j++ {
			m[rangeLength+i][j] = 1
		}
	}

	mult := func(a, b [][]int) [][]int {
		size := len(a)
		result := make([][]int, size)
		for i := range result {
			result[i] = make([]int, size)
		}

		for i := 0; i < size; i++ {
			for k := 0; k < size; k++ {
				if a[i][k] == 0 { // skip zero rows — optimization
					continue
				}
				for j := 0; j < size; j++ {
					result[i][j] = (result[i][j] + a[i][k]*b[k][j]) % MOD
				}
			}
		}
		return result
	}

	result := make([][]int, 2*rangeLength)
	for i := range len(result) {
		result[i] = make([]int, 2*rangeLength)
		result[i][i] = 1
	}

	exp := n - 2

	for exp > 0 {
		if exp%2 == 1 {
			result = mult(result, m)
		}
		m = mult(m, m)
		exp = exp / 2
	}

	s := make([]int, 2*rangeLength)
	for i := 0; i < rangeLength; i++ {
		s[i] = i
		s[rangeLength+i] = rangeLength - 1 - i
	}

	finalMatrix := make([]int, 2*rangeLength)
	for i := 0; i < 2*rangeLength; i++ {
		for j := 0; j < 2*rangeLength; j++ {
			finalMatrix[i] = (finalMatrix[i] + result[i][j]*s[j]) % MOD
		}
	}

	answer := 0
	for i := 0; i < 2*rangeLength; i++ {
		answer = (answer + finalMatrix[i]) % MOD
	}

	return answer
}

func NumberOfZigZagArraysII() {
	n := 3
	l := 4
	r := 5
	out := zigZagArraysII(n, l, r)
	fmt.Println(out)
}
