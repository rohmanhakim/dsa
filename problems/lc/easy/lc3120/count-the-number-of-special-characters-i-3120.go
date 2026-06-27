package lc3120

import (
	"fmt"
	"strings"
)

func numberOfSpecialChars(word string) int {
	appears := make(map[string]int)
	count := 0
	for _, c := range word {
		cs := string(c)
		lower := strings.ToLower(cs)
		isLowerCase := lower == cs
		if val, ok := appears[lower]; ok {
			if val == 0 { // lowercase
				if !isLowerCase {
					count++
					appears[lower] = 2 // mark as visited and counted
				}
			} else if val == 1 { // uppercase
				if isLowerCase {
					count++
					appears[lower] = 2
				}
			}
		} else {
			if isLowerCase {
				appears[lower] = 0
			} else {
				appears[lower] = 1
			}
		}
	}
	return count
}

func CountTheNumberOfSpecialCharactersI() {
	// word := "aaAbcBC"// 3
	// word := "abc" // 0
	// word := "abBCab" // 1
	word := "aAA" // 1
	// word := "Cc"
	out := numberOfSpecialChars(word)
	fmt.Println(out)
}
