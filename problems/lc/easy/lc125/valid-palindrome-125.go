package lc125

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 10^5
s consists only of printable ASCII characters.
*/

import (
	"strings"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	s = strings.ToLower(s)

	for left < right {
		a := s[left]
		b := s[right]

		if (a < 'a' || a > 'z') && (a < '0' || a > '9') {
			left++
			continue
		}
		if (b < 'a' || b > 'z') && (b < '0' || b > '9') {
			right--
			continue
		}

		if a != b {
			return false
		}
		left++
		right--
	}
	return true
}
