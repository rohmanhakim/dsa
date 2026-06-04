package easy

import (
	"fmt"
	"sort"
)

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	p := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[p] {
			p++
		}
		if p == len(s) {
			return true
		}
	}
	return false
}

// Preprocesses t into a map of char -> sorted list of indices
func buildIndex(t string) map[byte][]int {
	index := make(map[byte][]int)
	for i := 0; i < len(t); i++ {
		index[t[i]] = append(index[t[i]], i)
	}
	return index
}

// Checks if s is a subsequence of t using the prebuilt index.
// For each char in s, binary search for the smallest index in t
// that is strictly greater than our current position.
func isSubsequenceWithIndex(s string, index map[byte][]int) bool {
	pos := -1 // current position in t; next char must appear after this
	for i := 0; i < len(s); i++ {
		positions, ok := index[s[i]]
		if !ok {
			return false // char doesn't exist in t at all
		}

		// Find the first index in positions that is > pos
		j := sort.SearchInts(positions, pos+1)
		if j == len(positions) {
			return false // no valid position found
		}
		pos = positions[j] // advance our pointer in t
	}
	return true
}

func IsSubsequence() {
	// s := "axc"
	// t := "ahbgdc"
	s := "b"
	t := "b"
	out := isSubsequence(s, t)
	fmt.Println(out)

	t = "ahbgdc"
	index := buildIndex(t) // preprocess once

	queries := []string{"abc", "axc", "a", "z", ""}
	for _, s := range queries {
		fmt.Printf("isSubsequence(%q, %q) = %v\n", s, t, isSubsequenceWithIndex(s, index))
	}
}
