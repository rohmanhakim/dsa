package temp

import "fmt"

func backtrackSubset(nums []int, res *[][]int, path *[]int, start int) {
	tmp := append([]int(nil), (*path)...)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtrackSubset(nums, res, path, i+1)
		a := *path
		*path = a[:len(a)-1]
	}
}

func cartesian(pools [][]int, path *[]int) {
	if len(*path) == len(pools) {
		fmt.Println(*path)
		return
	}

	for _, item := range pools[len(*path)] {
		*path = append(*path, item)
		cartesian(pools, path)
		a := *path
		*path = a[:len(a)-1]
	}
}

func mainSubset() {
	// nums := []int{1, 2, 3}
	// res := [][]int{}
	// path := []int{}
	cpath := []int{}

	pools := [][]int{
		{1, 2},
		{4, 5},
	}

	// backtrackSubset(nums, &res, &path, 0)
	// fmt.Println(res)

	cartesian(pools, &cpath)
}
