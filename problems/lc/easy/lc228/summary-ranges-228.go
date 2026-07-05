package lc228

/*
You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"


Constraints:

0 <= nums.length <= 20
-2^31 <= nums[i] <= 2^31 - 1
All the values of nums are unique.
nums is sorted in ascending order.
*/

import (
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	rangeStart := 0

	res := []string{}
	for i := 1; i < len(nums); i++ {
		var sb strings.Builder

		if nums[i]-nums[i-1] > 1 {
			sb.WriteString(strconv.Itoa(nums[rangeStart]))
			if nums[i-1] != nums[rangeStart] {
				sb.WriteString("->")
				sb.WriteString(strconv.Itoa(nums[i-1]))
			}
			res = append(res, sb.String())
			rangeStart = i
		}
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(nums[rangeStart]))
	if rangeStart < len(nums)-1 {
		sb.WriteString("->")
		sb.WriteString(strconv.Itoa(nums[len(nums)-1]))
	}
	res = append(res, sb.String())
	return res
}
