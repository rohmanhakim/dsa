package medium

import (
	"fmt"
)

func isZeroArray(nums []int, queries [][]int) bool {

	diff := make([]int, len(nums)+1)

	for _, q := range queries {
		diff[q[0]] -= 1
		diff[q[1]+1] += 1
	}

	running := 0

	sum := 0
	for i := 0; i < len(nums); i++ {
		running += diff[i]
		res := nums[i] + running
		if res <= 0 {
			nums[i] = 0
		} else {
			nums[i] += running
		}
		sum += nums[i]
	}

	return sum == 0
}

func ZeroArrayTransformationI() {
	nums := []int{1, 0, 1}
	queries := [][]int{
		{0, 2},
	}

	// nums := []int{4, 3, 2, 1}
	// queries := [][]int{
	// 	{1, 3},
	// 	{0, 2},
	// }

	out := isZeroArray(nums, queries)

	fmt.Println(out)
}
