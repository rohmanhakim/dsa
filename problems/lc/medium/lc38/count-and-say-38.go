package lc38

import "fmt"

func rle(s string) string {

	freq := make(map[rune]int)
	res := ""
	lastChar := rune(s[0])
	for i, char := range s {
		freq[char]++
		if lastChar != char {
			res += fmt.Sprintf("%d%c", freq[lastChar], lastChar)
			freq[lastChar] = 0
		}

		if i == len(s)-1 {
			res += fmt.Sprintf("%d%c", freq[char], char)
		}
		lastChar = char
	}

	return res
}

func countAndSay(n int) string {
	res := "1"
	for range n - 1 {
		res = rle(res)
	}

	return res
}

func CountAndSay() {
	n := 4
	out := countAndSay(n)

	fmt.Println(out)
}
