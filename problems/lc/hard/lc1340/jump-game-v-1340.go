package lc1340

import (
	"fmt"
	"sort"
)

func maxJumps(arr []int, d int) int {
	n := len(arr)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return arr[indexes[i]] < arr[indexes[j]]
	})

	maxJump := make([]int, n)
	for _, i := range indexes {
		maxJump[i] = 1
		for x := 1; x <= d && i+x < n; x++ {
			if arr[i+x] >= arr[i] {
				break
			}

			if maxJump[i+x]+1 > maxJump[i] {
				maxJump[i] = maxJump[i+x] + 1
			}
		}
		for x := 1; x <= d && i-x >= 0; x++ {
			if arr[i-x] >= arr[i] {
				break
			}

			if maxJump[i-x]+1 > maxJump[i] {
				maxJump[i] = maxJump[i-x] + 1
			}
		}

	}

	max := 0
	for i := 0; i < len(arr); i++ {
		if maxJump[i] > max {
			max = maxJump[i]
		}
	}

	return max
}

func JumpGameV() {
	arr := []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}
	d := 2
	out := maxJumps(arr, d)
	fmt.Println(out)
}
