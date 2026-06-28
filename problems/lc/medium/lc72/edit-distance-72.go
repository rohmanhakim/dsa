package lc72

/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*/

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, n+1)
	}

	/*
		          r  o  s
		       0  1  2  3
			h  1  1  2  3
			o  2  2  1  2
			r  3  2  2  2
			s  4  3  3  2
			e  5  4  4  3


	*/

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				opReplace := dp[i-1][j-1] + 1
				opDelete := dp[i-1][j] + 1
				opInsert := dp[i][j-1] + 1
				best := min(opInsert, opDelete)

				if opReplace < best {
					best = opReplace
				}
				dp[i][j] = best
			}
		}
	}

	return dp[m][n]
}

func EditDistance() {
	word1 := "horse"
	word2 := "ros"

	out := minDistance(word1, word2)
	fmt.Println(out)
}
