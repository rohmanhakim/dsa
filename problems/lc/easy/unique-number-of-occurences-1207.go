package easy

import "fmt"

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)

	for _, n := range arr {
		freq[n]++
	}

	freqRev := make(map[int]int)
	for k, v := range freq {
		freqRev[v] = k
	}

	return len(freqRev) == len(freq)
}

func UniqueNumberOfOccurences() {
	// arr := []int{1, 2, 2, 1, 1, 3}
	// arr := []int{1, 2}
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	out := uniqueOccurrences(arr)
	fmt.Println(out)
}
