package lc345

import "fmt"

func isVowel(r rune) bool {
	lower := r == 'a' || r == 'i' || r == 'u' || r == 'e' || r == 'o'
	upper := r == 'A' || r == 'I' || r == 'U' || r == 'E' || r == 'O'
	return lower || upper
}

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1

	sarr := []rune(s)

	for left < right {
		if isVowel(sarr[left]) {
			if isVowel(sarr[right]) {
				sarr[left], sarr[right] = sarr[right], sarr[left]
				right--
				left++
			} else {
				right--
			}
		} else {
			left++
		}
	}
	return string(sarr)
}

func ReverseVowelsOfAstring() {
	// s := "IceCreAm"
	s := "Live on evasions? No, I save no evil."
	out := reverseVowels(s)
	fmt.Println(out)
}
