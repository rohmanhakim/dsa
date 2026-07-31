package lc3016

import (
	"sort"
)

func minimumPushes(word string) int {
	type char struct {
		val  int
		freq int
	}
	freqs := make(map[int]int)
	chars := []char{}
	for _, ch := range word {
		freqs[int(ch-'a')]++
	}

	for k, v := range freqs {
		chars = append(chars, char{
			val:  k,
			freq: v,
		})
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i].freq > chars[j].freq
	})

	mult := 1
	batch := 8
	count := 0
	for i := 1; i <= len(chars); i++ {
		if i > batch {
			mult++
			batch = mult * 8
		}
		ch := chars[i-1]
		count += ch.freq * mult
	}

	return count
}
