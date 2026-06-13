package easy

import (
	"fmt"
	"strings"
)

func mapWordWeights(words []string, weights []int) string {
	var sb strings.Builder
	for _, w := range words {
		weight := 0
		for _, c := range w {
			weight += weights[c-'a']
		}
		res := 'z' - (weight % 26)
		sb.WriteRune(rune(res))
	}
	return sb.String()
}

func WeightedWordMapping() {
	words := []string{"abcd", "def", "xyz"}
	weights := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}
	out := mapWordWeights(words, weights)
	fmt.Println(out)
}
