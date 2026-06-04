package medium

import (
	"fmt"
)

func expand(s string, left int, right int, start *int, max *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		length := right - left + 1

		if length > *max {
			*max = length
			*start = left
		}
		left--
		right++
	}
}

func longestPalindrome(s string) string {
	start := 0
	max := 0

	for i := range s {
		expand(s, i, i, &start, &max)
		expand(s, i, i+1, &start, &max)
	}

	return s[start : start+max]
}

func LongestPalindromicSubstring() {
	s := "cbbd"

	out := longestPalindrome((s))

	fmt.Printf("%s\n", out)
}
