/*
Given a string s consisting only of characters a, b and c.

Return the number of substrings containing at least one occurrence of all these characters a, b and c.



Example 1:

Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).

Example 2:

Input: s = "aaacb"
Output: 3
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb".

Example 3:

Input: s = "abc"
Output: 1


Constraints:

3 <= s.length <= 5 x 10^4
s only consists of a, b or c characters.
*/

package lc1358

import (
	"fmt"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	left := 0
	right := left

	count := 0

	abcCount := []int{0, 0, 0}
	for right < n {

		chRight := s[right]

		switch chRight {
		case 'a':
			abcCount[0]++
		case 'b':
			abcCount[1]++
		case 'c':
			abcCount[2]++
		}

		for abcCount[0] > 0 && abcCount[1] > 0 && abcCount[2] > 0 {
			count += n - right
			chLeft := s[left]

			switch chLeft {
			case 'a':
				abcCount[0]--
			case 'b':
				abcCount[1]--
			case 'c':
				abcCount[2]--
			}
			left++
		}

		right++

	}
	return count
}

func NumberofSubstringsContainingAllThreeCharacters() {
	s := "aaacb"
	out := numberOfSubstrings(s)
	fmt.Println(out)
}
