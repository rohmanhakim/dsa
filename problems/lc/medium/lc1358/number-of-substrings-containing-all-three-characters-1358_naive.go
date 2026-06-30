package lc1358

import (
	"strings"
)

func numberOfSubstringsNaive(s string) int {
	/*
		a
		ab
		abc
		abca
		abcab
		abcabc
		bcabc
		cabc
		abc
		bc
		c
		b
		bc
		bca
		bcab
		cab
		ab
		c
	*/
	n := len(s)
	left := 0
	right := 3
	half := n / 2
	rightBound := n

	count := 0
	for left < right-2 && right <= rightBound {
		sub := s[left:right]
		// fmt.Println(sub)
		if strings.Contains(sub, "a") && strings.Contains(sub, "b") && strings.Contains(sub, "c") {
			count++
		}
		if right == rightBound {
			if left == right-3 {
				half--
				rightBound--
				left = (n / 2) - half
				right = left + 3
			} else {
				left++
			}
		} else {
			right++
		}
	}
	return count
}
