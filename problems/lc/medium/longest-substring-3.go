package medium

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	charset := make(map[byte]int)

	max := 0
	num := 0
	first := 0
	last := 1
	for last < len(s) {
		charset[s[first]]++

		_, ok := charset[s[last]]

		if !ok {
			charset[s[last]]++
			num = last - first
			if num > max {
				max = num
			}
			last++
		} else {
			first++
			last = first + 1
			charset = make(map[byte]int)
		}
	}

	return max + 1
}

func lengthOfLongestSubstringClaudeSonnet(s string) int {
	charIndex := make(map[byte]int) // maps char -> its last seen index
	max := 0
	first := 0

	for last := 0; last < len(s); last++ {
		ch := s[last]
		// If ch was seen and its last occurrence is within the current window,
		// shrink the window from the left
		if idx, ok := charIndex[ch]; ok && idx >= first {
			first = idx + 1
		}
		charIndex[ch] = last
		if length := last - first + 1; length > max {
			max = length
		}
	}

	return max
}

func Longestsubstring() {
	s := "abba"

	fmt.Println(lengthOfLongestSubstring(s))
}
