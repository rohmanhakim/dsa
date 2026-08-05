package lc2904

import (
	"math"
)

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	left := 0
	count := 0
	shortestLen := math.MaxInt
	best := ""
	for right := range n {
		count += int(s[right] - '0')

		for (s[left] == '1' && count > k) || (s[left] == '0' && count == k) {
			count -= int(s[left] - '0')
			left++
		}
		if count == k {
			sub := s[left : right+1]
			if len(sub) < shortestLen || (len(sub) == shortestLen && sub < best) {
				shortestLen = len(sub)
				best = sub
			}
		}
	}
	return best
}
