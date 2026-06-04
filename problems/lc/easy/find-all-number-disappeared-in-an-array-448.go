package easy

import (
	"fmt"
	"sort"

	"github.com/rohmanhakim/leetcode/lib"
)

func findDisappearedNumbers(nums []int) []int {
	appears := make([]int, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for _, n := range nums {
		appears[n-1] = 1
	}
	res := []int{}
	for i, n := range appears {
		if n == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func findDisappearedNumbersClaudeSonnet(nums []int) []int {
	// Pass 1: mark visited indices by negating the value there
	for _, n := range nums {
		idx := lib.Abs(n) - 1 // map value → index (handle already-negated values)
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// Pass 2: indices still positive were never visited → missing numbers
	res := []int{}
	for i, n := range nums {
		if n > 0 {
			res = append(res, i+1)
		}
	}
	return res

}

func FindAllNumbersDisappearedInAnArray() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	out := findDisappearedNumbersClaudeSonnet(nums)
	fmt.Println(out)
}
