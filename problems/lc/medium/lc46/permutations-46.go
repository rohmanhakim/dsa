package lc46

import (
	"fmt"
	"slices"
)

func bt(nums []int, res *[][]int, path *[]int, visited []bool) {

	if len(*path) == len(nums) {
		*res = append(*res, slices.Clone(*path))
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		*path = append(*path, nums[i])
		bt(nums, res, path, visited)
		visited[i] = false
		a := *path
		*path = a[:len(*path)-1]
	}
}

func permute(nums []int) [][]int {

	res := [][]int{}
	path := []int{}
	visited := make([]bool, len(nums))

	bt(nums, &res, &path, visited)

	return res
}

func Permutations() {
	// nums := []int{1, 2, 3}
	// nums := []int{0, 1}
	nums := []int{1}
	out := permute(nums)
	fmt.Println(out)
}
