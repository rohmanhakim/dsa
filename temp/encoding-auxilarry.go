package temp

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func encodingAuxilarry(nums []int) []int {
	for _, n := range nums {
		idx := lib.Abs(n) - 1
		fmt.Printf("n: %d, idx: %d, nums[idx]: %d\n", n, idx, nums[idx])

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	return nums
}

func mainEncodingAuxilarry() {
	nums := []int{5, 1, 4, 2, 1}

	for _, n := range nums {
		idx := lib.Abs(n) - 1
		fmt.Printf("n: %d, idx: %d, nums[idx]: %d\n", n, idx, nums[idx])

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	out := encodingAuxilarry(nums)
	// -4 -1, -3, 0, 0
	fmt.Println(out)
}
