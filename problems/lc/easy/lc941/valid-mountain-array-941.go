package lc941

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

```
moutnain array:
0 2 3 4 5 2 1 0
strictly increasing up to arr[4] ->
strictly decreasing from arr[4] ->

not mountain array:
0 2 3 3 5 2 1 0
increasing but not strictly from arr[0] to arr[4]->
strictly decreasing from arr[4]
```

Example 1:

Input: arr = [2,1]
Output: false
Example 2:

Input: arr = [3,5,5]
Output: false
Example 3:

Input: arr = [0,3,2,1]
Output: true


Constraints:

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^4
*/

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}

	prev := arr[0]
	idx := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < prev {
			if i < 2 {
				return false
			}
			idx = i
			break
		}
		if arr[i] == prev {
			return false
		}
		prev = arr[i]
	}

	if idx == -1 {
		return false
	}
	prev = arr[idx-1]
	for i := idx; i < len(arr); i++ {
		if arr[i] >= prev {
			return false
		}
		prev = arr[i]
	}
	return true
}

func ValidMountainArray() {

	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	out := validMountainArray(arr)
	fmt.Println(out)
}
