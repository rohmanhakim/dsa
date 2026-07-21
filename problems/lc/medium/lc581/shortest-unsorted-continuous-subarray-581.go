package lc581

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	type item struct {
		val int
		idx int
	}
	dec := []item{}
	inc := []item{}

	left := math.MaxInt
	right := -1

	for i, num := range nums {
		for len(dec) > 0 && num < dec[len(dec)-1].val {
			left = min(left, dec[len(dec)-1].idx)
			dec = dec[:len(dec)-1]
		}
		dec = append(dec, item{
			val: num,
			idx: i,
		})
	}

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		for len(inc) > 0 && num > inc[len(inc)-1].val {
			right = max(right, inc[len(inc)-1].idx)
			inc = inc[:len(inc)-1]
		}
		inc = append(inc, item{
			val: num,
			idx: i,
		})
	}

	if left == math.MaxInt {
		return 0
	}
	return right - left + 1
}
