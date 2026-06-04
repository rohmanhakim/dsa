package medium

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var digits []int
	sum := 0
	negative := false
	if x < 0 {
		x = x * -1
		negative = true
	}

	if x == 0 {
		digits = []int{0}
	}

	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	for i, val := range digits {
		n := 1
		for range len(digits) - i - 1 {
			n = n * 10
		}
		sum = sum + val*n
		if sum > math.MaxInt32 {
			return 0
		}
	}

	fmt.Print("\n")

	if negative {
		sum = sum * -1
	}
	return sum
}

func reverseClaudeSonnet(x int) int {
	result := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	for x != 0 {
		digit := x % 10
		x /= 10

		// Check overflow BEFORE multiplying/adding
		// equivalent to: result*10 + digit > MaxInt32
		if result > (math.MaxInt32-digit)/10 {
			return 0
		}

		result = result*10 + digit
	}

	return result * sign
}

func reverse2(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}
	remaining := x * sign
	sum := 0
	for remaining != 0 {
		currentDigit := remaining % 10

		remaining = remaining / 10

		if sum > math.MaxInt32/10-currentDigit {
			return 0
		}

		sum = sum*10 + currentDigit
	}
	return sum * sign
}

func ReverseInteger() {
	// 2147483641
	x := 1463847412
	fmt.Printf("%d\n", reverse2(x))
}
