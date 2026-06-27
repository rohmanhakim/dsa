package lc1189

import (
	"fmt"
	"math"
)

func maxNumberOfBalloons(text string) int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	query := "balloon"
	i := 0
	nextChar := rune(query[i])
	count := 0
	for {
		if val, ok := freq[nextChar]; !ok || val == 0 {
			return count
		} else {
			freq[nextChar]--
			if i >= len(query)-1 {
				count++
				i = 0
			} else {
				i++
			}
			nextChar = rune(query[i])
		}
	}
}

func maxNumberOfBalloonsClaudeSonnet(text string) int {
	freq := make(map[rune]int)
	for _, ch := range text {
		freq[ch]++
	}

	// balloon = b:1, a:1, l:2, o:2, n:1
	needed := map[rune]int{
		'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1,
	}

	result := math.MaxInt32
	for ch, count := range needed {
		result = min(result, freq[ch]/count)
	}
	return result
}

func MaximumNumberOfBaloons() {

	text := "nlaebolko"
	out := maxNumberOfBalloons(text)
	fmt.Println(out)
}
