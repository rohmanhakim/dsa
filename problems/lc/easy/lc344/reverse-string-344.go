package lc344

import (
	"fmt"
)

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

func ReverseString() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Print("[")
	for i, c := range s {
		fmt.Printf("\"%c\"", c)
		if i < len(s)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]\n")
}
