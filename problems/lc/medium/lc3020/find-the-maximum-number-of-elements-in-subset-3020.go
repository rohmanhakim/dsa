package lc3020

import (
	"fmt"
)

/*
You are given an array of positive integers nums.

You need to select a subset of nums which satisfies the following condition:

You can place the selected elements in a 0-indexed array such that it follows the pattern: [x, x^2, x^4, ..., x^(k/2), x^k, x^(k/2), ..., x^4, x^2, x] (Note that k can be be any non-negative power of 2). For example, [2, 4, 16, 4, 2] and [3, 9, 3] follow the pattern while [2, 4, 8, 4, 2] does not.
Return the maximum number of elements in a subset that satisfies these conditions.



Example 1:

Input: nums = [5,4,1,2,2]
Output: 3
Explanation: We can select the subset {4,2,2}, which can be placed in the array as [2,4,2] which follows the pattern and 2^2 == 4. Hence the answer is 3.
Example 2:

Input: nums = [1,3,2,4]
Output: 1
Explanation: We can select the subset {1}, which can be placed in the array as [1] which follows the pattern. Hence the answer is 1. Note that we could have also selected the subsets {2}, {3}, or {4}, there may be multiple subsets which provide the same answer.


Constraints:

2 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
*/

func maximumLength(nums []int) int {
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	// fmt.Println(nums)
	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
	}

	max := 1
	for x := range freq {
		m := 1
		length := 1
		if freq[x] < 2 {
			continue
		}
		if x == 1 {
			if freq[x]%2 == 0 {
				length = freq[x] - 1
			} else {
				length = freq[x]
			}
		} else {
			for {
				k := powInt(2, m)
				next := powInt(x, k)

				if freq[next] == 0 {
					// fmt.Printf("x: %d, k: %d, m: %d, next: %d ", x, k, m, next)
					break
				}
				// if freq[next] >= 2 {
				// fmt.Printf("x: %d, k: %d, m: %d, next: %d ", x, k, m, next)
				length = 2*(m) + 1
				if freq[next] < 2 {
					break
				}

				fmt.Println(length)
				// break
				// }

				// length = length + 2
				m++
			}
		}

		fmt.Printf("x: %d, length: %d\n", x, length)
		if length > max {
			max = length
		}
	}
	return max
}

func powInt(x, e int) int {
	res := x
	for range e - 1 {
		res = res * x
	}
	return res
}

func FindTheMaximumNumberOfElementsInSubset() {
	// nums := []int{5, 4, 1, 2, 2} // 3
	nums := []int{1, 3, 2, 4} //1
	// nums := []int{1, 1} //1
	// nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024} // 9
	// nums := []int{1, 16, 49, 16, 121}//1
	// nums := []int{14, 14, 196, 196, 38416, 38416} // 5
	// nums := []int{65025, 312481, 107584, 148996, 322624, 194481, 570025, 15376, 123904, 848241, 88804, 47961, 117649, 66564, 295936, 271441, 16900, 474721, 27556, 285156, 11236, 175561, 917764, 968256, 16, 38025, 312481, 426409, 354025, 8464, 522729, 60516, 210681, 378225, 638401, 101124, 697225, 427716, 262144, 940900, 988036, 324900, 151321, 309136, 178929, 168921, 189225, 4, 301401, 659344, 786769, 964324, 15625, 302500, 56644, 61504, 31684, 369664, 345744, 19321, 59049, 5041, 40000, 147456, 372100, 708964, 171396, 214369, 707281, 484, 49729, 82944, 100489, 103684, 58564, 208849, 946729, 84100, 4, 600625, 334084, 683929, 9604, 245025, 97969, 147456, 160801, 434281, 223729, 294849, 166464, 432964, 518400, 376996, 17424, 315844, 256, 737881, 10000, 632025} // 3

	out := maximumLength(nums)
	fmt.Println(out)
}
