package easy

import (
	"fmt"
)

func separateDigits(nums []int) []int {

	out := []int{}

	for _, num := range nums {
		res := []int{}
		x := num
		for x != 0 {
			digit := x % 10
			x = x / 10
			res = append(res, digit)
		}
		for len(res) != 0 {
			out = append(out, res[len(res)-1])
			res = res[:len(res)-1]
		}
	}

	return out
}

func separateDigitsClaudeSonnet(nums []int) []int {
	out := []int{}
	for _, num := range nums {
		start := len(out)
		for x := num; x != 0; x /= 10 {
			out = append(out, x%10) // appended in reverse
		}
		// reverse just the segment we added
		for i, j := start, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
	}
	return out
}

func SeparateTheDigitsInAnArray() {
	nums := []int{7, 1, 3, 9}

	out := separateDigits(nums)

	fmt.Println(out)
}
