package easy

import (
	"fmt"
	"math"
)

func minElement(nums []int) int {
	min := math.MaxInt32

	for _, n := range nums {
		// digits := []int{}
		sum := 0
		for n > 0 {
			digit := n % 10
			sum = sum + digit
			// digits = append(digits, digit)
			n = n / 10
		}

		if sum < min {
			min = sum
		}

		// fmt.Println(sum)

	}
	return min
}

func MinimumElementAfterReplacementWithdigitSum() {
	// nums := []int{10, 12, 13, 14} // 1
	// nums := []int{1, 2, 3, 4} // 1
	nums := []int{999, 19, 199} // 10
	out := minElement(nums)
	fmt.Println(out)
}
