package lc1979

import (
	"fmt"
	"slices"
)

func findGCD(nums []int) int {
	a := slices.Max(nums)
	b := slices.Min(nums)

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LC1979() {
	nums := []int{2, 5, 6, 9, 10}
	out := findGCD(nums)
	fmt.Println(out)
}
