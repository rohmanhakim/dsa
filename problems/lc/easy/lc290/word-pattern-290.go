package lc290

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	char2word := make(map[rune]string)
	word2char := make(map[string]rune)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, ch := range pattern {
		if val, charExists := char2word[ch]; !charExists {
			if _, wordExists := word2char[words[i]]; !wordExists {
				char2word[ch] = words[i]
				word2char[words[i]] = ch
			} else {
				return false
			}
		} else if val != words[i] {
			return false
		}
	}

	return true
}
