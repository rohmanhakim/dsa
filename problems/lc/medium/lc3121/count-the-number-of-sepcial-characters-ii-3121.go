package lc3121

import (
	"fmt"
)

func numberOfSpecialChars(word string) int {

	lower := make([]int, 26)
	upper := make([]int, 26)

	for i, c := range word {
		if c >= 'a' && c <= 'z' {
			if upper[c-'a'] == 0 {
				lower[c-'a'] = (i + 1)
			} else {
				lower[c-'a'] = -1
			}
		} else {
			upper[c-'A'] = (i + 1)
		}
	}

	count := 0

	for i, idx := range lower {
		if idx <= 0 || upper[i] <= 0 {
			continue
		}
		if idx < upper[i] {
			count++
		}
	}

	return count
}

func numberOfSpecialCharsClaudeSonnet(word string) int {
	lastLower := make([]int, 26)
	firstUpper := make([]int, 26)

	// Initialize to -1 (unseen)
	for i := range lastLower {
		lastLower[i] = -1
		firstUpper[i] = -1
	}

	for i, c := range word {
		if c >= 'a' && c <= 'z' {
			lastLower[c-'a'] = i // always update — we want the last one
		} else {
			if firstUpper[c-'A'] == -1 {
				firstUpper[c-'A'] = i // only record the first occurrence
			}
		}
	}

	count := 0
	for i := 0; i < 26; i++ {
		if lastLower[i] != -1 && firstUpper[i] != -1 && lastLower[i] < firstUpper[i] {
			count++
		}
	}

	return count
}

func CountTheNumberOfSpecialCharactersII() {
	// word := "aaAbcBC"
	// word := "abc"
	// word := "AbBCab"
	// word := "cCceDC"
	word := "abDBAbb"
	// word := "aAbBaAbB"
	// word := "aAbBaAbB"
	out := numberOfSpecialCharsClaudeSonnet(word)
	fmt.Println(out)
}
