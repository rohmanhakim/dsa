package lc151

import (
	"fmt"
	"strings"
	"unicode"
)

func isWhiteSpaceOnly(s string) bool {
	runes := []rune(s)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	left := 0
	right := len(words) - 1

	for left < right {
		for left < right && isWhiteSpaceOnly(words[left]) {
			left++
		}
		for left < right && isWhiteSpaceOnly(words[right]) {
			right--
		}
		if left < right && !isWhiteSpaceOnly(words[right]) && !isWhiteSpaceOnly(words[left]) {
			words[left], words[right] = words[right], words[left]
			right--
			left++
		}
	}

	var b strings.Builder
	for i, w := range words {
		if !isWhiteSpaceOnly(w) {
			b.WriteString(w)
			if i != len(words)-1 {
				b.WriteString(" ")
			}
		}

	}
	return strings.TrimSpace(b.String())
}

func reverseWordsClaudeSonnet(s string) string {
	words := strings.Fields(s) // splits on any whitespace, trims automatically
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}

func ReverseWordsInAString() {
	// s := "the sky is blue"
	// s := "  hello world  "
	// s := "a good   example"
	// s := "  Bob    Loves  Alice   "
	s := " 3c      2zPeO dpIMVv2SG    1AM       o       VnUhxK a5YKNyuG     x9    EQ  ruJO       0Dtb8qG91w 1rT3zH F0m n G wU"
	out := reverseWords(s)
	fmt.Println(out)
}
