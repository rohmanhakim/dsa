package medium

import (
	"fmt"
)

func countMajoritySubarrays(nums []int, target int) int {
	pref := make([]int, len(nums)+1)
	if nums[0] == target {
		pref[0] = 1
	}
	for i := 1; i < len(nums); i++ {
		count := 0
		if nums[i] == target {
			count = 1
		}
		pref[i] = pref[i-1] + count
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sublen := j - i + 1
			numBefore := 0
			if i > 0 {
				numBefore = pref[i-1]
			}
			numOfTarget := pref[j] - numBefore
			if numOfTarget > sublen/2 {
				count++
			}
		}
	}

	return count
}

func CountSubarraysWithMajorityElementI() {
	nums := []int{1, 2, 2, 3}
	target := 2
	out := countMajoritySubarrays(nums, target)
	fmt.Println(out)
}
