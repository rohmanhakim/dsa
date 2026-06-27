package lc645

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func findErrorNums(nums []int) []int {
	n := len(nums)

	set := make(map[int]bool)

	dup := 0
	sum := nums[0]
	set[nums[0]] = true

	sum2 := 0
	for i := 1; i < n; i++ {
		sum = sum + nums[i]
		sum2 = sum2 + i
		if _, ok := set[nums[i]]; ok {
			dup = nums[i]
		}
		set[nums[i]] = true

	}

	return []int{dup, (sum2 + n) - (sum - dup)}
}

func findErrorNumsClaudeSonnet(nums []int) []int {
	dup := 0
	missing := 0

	for _, num := range nums {
		idx := lib.Abs(num) - 1 // convert value to 0-based index
		if nums[idx] < 0 {
			dup = lib.Abs(num) // already negated → seen before
		} else {
			nums[idx] = -nums[idx] // mark as visited
		}
	}

	// the index still positive was never visited → missing number
	for i, num := range nums {
		if num > 0 {
			missing = i + 1
		}
	}

	return []int{dup, missing}
}

func SetMismatch() {
	// nums := []int{2, 2}//2,1
	nums := []int{1, 2, 2, 4} // [2,3]
	out := findErrorNums(nums)
	fmt.Println(out)
}
