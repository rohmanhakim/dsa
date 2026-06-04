package easy

import (
	"fmt"
	"sort"
)

func isGood(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	max := nums[len(nums)-1]

	if len(nums) != max+1 {
		return false
	}
	nums = append(nums, max)
	for i := 0; i < max+1; i++ {
		if i > 0 && i < len(nums)-2 && nums[i]-nums[i-1] != 1 {
			return false
		}
	}
	return true
}

func CheckIfArrayIsGood() {
	nums := []int{1, 3, 3, 3}
	out := isGood(nums)

	fmt.Println(out)
}
