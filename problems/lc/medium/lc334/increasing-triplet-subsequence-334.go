package lc334

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	stack := []int{}
	i := 0
	minJ := math.MaxInt32

	for _, k := range nums {
		if i < minJ && minJ < k {
			return true
		}
		for len(stack) > 0 && k <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, k)

		if len(stack) > 1 {
			if k < minJ {
				minJ = k
			}
		}

		i = stack[0]

		if len(stack) == 3 {
			return true
		}
	}

	return false
}

func increasingTripletClaudeSonnet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32

	for _, n := range nums {
		if n <= first {
			first = n // new minimum
		} else if n <= second {
			second = n // new second minimum, greater than first
		} else {
			return true // n > second > first — triplet found
		}
	}

	return false
}

func IncreasingTripletSubsequence() {
	nums := []int{1, 2, 3, 4, 5} // true
	// nums := []int{5, 4, 3, 2, 1} // false
	// nums := []int{4, 3, 1, 7, 2, 1, 6, 4, 1, 8} // true
	// nums := []int{20, 100, 10, 12, 5, 13} // true
	out := increasingTriplet(nums)
	fmt.Println(out)
}
