package easy

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

func RemoveDuplicatesFromSortedArray() {
	nums := []int{1, 1, 2}

	k := removeDuplicates(nums)

	fmt.Printf("%d, [", k)
	for _, num := range nums[:k] {
		fmt.Printf("%d,", num)
	}
	fmt.Print("]\n")
}
