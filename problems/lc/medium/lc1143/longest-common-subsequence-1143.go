package lc1143

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.



Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.


Constraints:

1 <= text1.length, text2.length <= 1000
text1 and text2 consist of only lowercase English characters.
*/

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, n+1)
	}

	/*
		     a  b  c  b  a
		  0  0  0  0  0  0
		a 0  1  1  1  1  1
		b 0  1  2  2  2  2
		c 0  1  2  3  3  3
		b 0  1  2  3  4  4
		c 0  1  2  3  4  4
		b 0  1  2  3  4  4
		a 0  1  2  3  4  5
	*/

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func LongestCommonSubsequence() {
	// text1 := "abcde"
	// text2 := "ace" // 3
	// text1 := "bl"
	// text2 := "yby" // 1
	// text1 := "bsbininm"
	// text2 := "jmjkbkjkv" //1
	text1 := "abcba"
	text2 := "abcbcba"

	out := longestCommonSubsequence(text1, text2)
	fmt.Println(out)
}
