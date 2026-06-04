package medium

import (
	"fmt"
	"slices"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// close:
	// - same length of strings
	// - same distinct characters
	// - same number of unique occurences of characters
	freq1 := make(map[rune]int)
	freqRev1 := make(map[int]int)
	for _, c := range word1 {
		freq1[rune(c)]++
	}
	for _, v := range freq1 {
		freqRev1[v]++
	}

	freq2 := make(map[rune]int)
	freqRev2 := make(map[int]int)
	for _, c := range word2 {
		if _, ok := freq1[c]; !ok {
			return false
		}
		freq2[rune(c)]++
	}
	for _, v := range freq2 {
		freqRev2[v]++
	}
	for k, v := range freqRev2 {
		if val := freqRev1[k]; v != val {
			return false
		}
	}

	close := len(freq1) == len(freq2) && len(freqRev1) == len(freqRev2)

	return close
}

func closeStringsClaudeSonnet(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var freq1, freq2 [26]int
	for _, c := range word1 {
		freq1[c-'a']++
	}
	for _, c := range word2 {
		freq2[c-'a']++
	}

	for i := 0; i < 26; i++ {
		/*
			This checks that both frequencies are either **both zero or both non-zero** — in other words, that a character either exists in both strings or neither.

			Breaking it down:
			- `freq1[i] == 0` → character `i` is **absent** from word1
			- `freq2[i] == 0` → character `i` is **absent** from word2

			The `!=` between them means "these two absence-checks disagree", i.e. one is absent and the other isn't. If that's true for any character, the strings can't be close, so return false.

			A truth table makes it clearest:

			| freq1[i] | freq2[i] | (==0) != (==0) | meaning |
			|---|---|---|---|
			| 0 | 0 | false | both absent ✅ |
			| 3 | 5 | false | both present ✅ |
			| 0 | 4 | **true** | mismatch ❌ |
			| 2 | 0 | **true** | mismatch ❌ |

			It's a concise way of writing the equivalent:
			```go
			if (freq1[i] == 0 && freq2[i] != 0) || (freq1[i] != 0 && freq2[i] == 0) {
				return false
			}
			```
			Or even more readably, using XOR logic: one is zero but the other isn't.
		*/
		if (freq1[i] == 0) != (freq2[i] == 0) {
			return false
		}
	}

	f1 := freq1[:]
	f2 := freq2[:]
	sort.Ints(f1)
	sort.Ints(f2)

	return slices.Equal(f1, f2)
}

func DetermineIfTwoStringsAreClose() {
	// word1 := "abc"
	// word2 := "bca"
	// word1 := "a"
	// word2 := "aa"
	// word1 := "cabbba"
	// word2 := "abbccc"
	word1 := "uau"
	word2 := "ssx"
	// word1 := "aaabbbbccddeeeeefffff"
	// word2 := "aaaaabbcccdddeeeeffff"
	out := closeStrings(word1, word2)
	fmt.Println(out)
}
