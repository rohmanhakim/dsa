package lc1846

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	arr[0] = 1

	next := 2
	for i := 1; i < n; i++ {
		if abs(arr[i]-arr[i-1]) > 1 {
			arr[i] = next
		} else {
			next = arr[i]
		}
		next++
	}
	return arr[n-1]
}

func maximumElementAfterDecrementingAndRearrangingNoSort(arr []int) int {
	n := len(arr)

	count := make([]int, n+1)

	for i := 1; i <= n; i++ {
		count[min(n, arr[i-1])]++
	}

	ans := 1

	for i := 2; i <= n; i++ {
		ans = min(ans+count[i], i)
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func MaximumElementAfterDecreasingandRearranging() {
	arr := []int{1, 2, 1000, 4, 100}
	out := maximumElementAfterDecrementingAndRearrangingNoSort(arr)
	fmt.Println(out)
}
