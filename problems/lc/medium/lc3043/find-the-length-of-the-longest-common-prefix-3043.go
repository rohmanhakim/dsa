package lc3043

import (
	"fmt"
	"strconv"
)

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := make(map[string]int)

	for i := 0; i < len(arr1); i++ {
		num := strconv.Itoa(arr1[i])
		key := ""
		for _, d := range num {
			key += string(d)
			prefixes[key]++
		}
	}

	max := 0
	for i := 0; i < len(arr2); i++ {
		num := strconv.Itoa(arr2[i])
		key := ""
		for _, d := range num {
			key += string(d)
			if _, ok := prefixes[key]; ok {
				if len(key) > max {
					max = len(key)
				}
			}
		}
	}

	return max
}

func FindTheLengthOfTheLongestCommonPrefix() {
	arr1 := []int{1, 10, 100}
	// arr2 := []int{1000, 10000}
	// arr1 := []int{1, 2, 3}
	arr2 := []int{4, 4, 4}
	// arr1 := []int{6, 32}
	// arr2 := []int{8, 36}
	out := longestCommonPrefix(arr1, arr2)

	fmt.Println(out)
	// fmt.Println(commonPrefix(32, 36))
}
