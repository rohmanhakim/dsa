package medium

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	left := 0
	right := 0
	count := 0
	curr := chars[left]
	for right < len(chars) {
		if curr == chars[right] {
			count++
		}
		if curr != chars[right] {
			chars[left] = curr
			if count > 1 {
				cs := strconv.Itoa(count)
				for j := range len(cs) {
					chars[left+j+1] = cs[j]
				}
				left += len(cs)
			}
			curr = chars[right]
			left++
			count = 1
		}
		if right == len(chars)-1 {
			chars[left] = curr
			if count > 1 {
				cs := strconv.Itoa(count)
				for j := range len(cs) {
					chars[left+j+1] = cs[j]
				}
				left += len(cs)
			}
			curr = chars[right]
		}
		right++
	}
	return left + 1
}

func compressClaudeSonnet(chars []byte) int {
	left := 0 // write pointer
	i := 0    // read pointer

	for i < len(chars) {
		char := chars[i]
		count := 0

		// Count the full run of this character
		for i < len(chars) && chars[i] == char {
			i++
			count++
		}

		// Write character
		chars[left] = char
		left++

		// Write count if > 1
		if count > 1 {
			cs := strconv.Itoa(count)
			for _, c := range cs {
				chars[left] = byte(c)
				left++
			}
		}
	}

	return left
}

func StringCompression() {
	// chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	// chars := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}
	// chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	// chars := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	// chars := []byte{'a', 'b', 'c'}
	// chars := []byte{'a', 'a', 'a', 'a', 'a', 'b'}
	// chars := []byte{'a', 'a'}
	// chars := []byte{'a', 'a', 'a', 'a', 'b', 'a'}
	chars := []byte{'a'}
	out := compress(chars)
	fmt.Println(out)
}
