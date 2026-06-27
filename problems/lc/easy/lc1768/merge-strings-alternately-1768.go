package lc1768

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {

	var sb strings.Builder
	sb.Grow(len(word1) + len(word2)) // pre-allocate exact capacity

	for p := 0; p < len(word1) || p < len(word2); p++ {
		if p < len(word1) {
			sb.WriteByte(word1[p])
		}
		if p < len(word2) {
			sb.WriteByte(word2[p])
		}
	}

	return sb.String()
}

func mergeAlternatelyTwoPointers(word1 string, word2 string) string {
	var sb strings.Builder
	p1, p2 := 0, 0 // ← two explicit pointers

	for p1 < len(word1) && p2 < len(word2) {
		sb.WriteByte(word1[p1])
		sb.WriteByte(word2[p2])
		p1++
		p2++
	}

	// append remainders
	// for p1 < len(word1) {
	// 	sb.WriteByte(word1[p1])
	// 	p1++
	// }
	// for p2 < len(word2) {
	// 	sb.WriteByte(word2[p2])
	// 	p2++
	// }

	return sb.String()
}

func MergeStringsAlternately() {
	// word := "abc"
	// word2 := "pqr"
	word := "abcd"
	word2 := "pq"
	out := mergeAlternatelyTwoPointers(word, word2)
	fmt.Println(out)
}
