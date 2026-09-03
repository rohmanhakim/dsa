package lc3876

import (
	"math"
)

func uniformArray(nums1 []int) bool {

	n := len(nums1)

	countOdd, countEven := 0, 0

	smallest := math.MaxInt
	for _, num := range nums1 {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
		if num < smallest {
			smallest = num
		}
	}

	if countEven == n || countOdd == n {
		return true
	}

	if smallest%2 == 0 && countOdd > 0 {
		return false
	}

	return true
}
