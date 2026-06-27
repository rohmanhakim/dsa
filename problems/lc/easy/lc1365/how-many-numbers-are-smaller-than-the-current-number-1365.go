package lc1365

import (
	"fmt"
	"slices"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	sorted := slices.Clone(nums)
	mins := make(map[int]int)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	last := -1
	lastMin := 0
	for i, n := range sorted {
		if n != last {
			mins[n] = i
			lastMin = i
		} else {
			mins[n] = lastMin
		}
		last = sorted[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = mins[nums[i]]
	}

	return nums
}

func HowManyNumbersAreSmallerThanTheCurrentNumber() {
	nums := []int{8, 1, 2, 2, 3}
	out := smallerNumbersThanCurrent(nums)

	fmt.Println(out)
}
