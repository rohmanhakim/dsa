package lc3

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	left := 0
	right := 1

	sub := ""
	for right <= len(s) {

		curr := s[right-1]

		if strings.Contains(sub, string(curr)) {
			left = left + strings.Index(sub, string(curr)) + 1
		}

		sub = s[left:right]

		length := len(sub)
		if length > maxLength {
			maxLength = length
		}

		right++
	}
	return maxLength
}
