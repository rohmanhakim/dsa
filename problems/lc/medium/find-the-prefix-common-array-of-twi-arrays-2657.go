package medium

import (
	"fmt"
)

func countCommonElements(a []int, b []int) int {
	occMap := make(map[int]int)

	p := 0
	for p < len(a) {
		occMap[a[p]]++
		occMap[b[p]]++
		p++
	}
	count := 0
	for _, v := range occMap {
		if v > 1 {
			count++
		}
	}
	return count
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = countCommonElements(A[:i+1], B[:i+1])
	}

	return res
}

func findThePrefixCommonArrayClaudeSonnet(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)
	freq := make(map[int]int)
	common := 0

	for i := 0; i < n; i++ {
		freq[A[i]]++
		if freq[A[i]] == 2 {
			common++
		}

		freq[B[i]]++
		if freq[B[i]] == 2 {
			common++
		}

		res[i] = common
	}

	return res
}

func FindThePrefixCommonArrayOfTwoArrays() {
	a := []int{1, 3, 2, 4}
	b := []int{3, 1, 2, 4} // 0,2,3,4

	// a := []int{2, 3, 1}
	// b := []int{3, 1, 2} // 0,1,3

	// a := []int{1, 2, 3}
	// b := []int{2, 3, 1} // 0,1,3

	out := findThePrefixCommonArray(a, b)

	fmt.Println(out)

}
