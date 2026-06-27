package lc1665

import (
	"fmt"
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		a1, b1 := tasks[i][0], tasks[i][1]
		a2, b2 := tasks[j][0], tasks[j][1]

		left := max(b1, a1+b2)
		right := max(b2, a2+b1)

		if left != right {
			return left < right
		}

		return b1 < b2
	})

	sum := 0
	energyRequired := 0
	largestMinimum := 0
	for i := 0; i < len(tasks); i++ {
		sum = sum + tasks[i][0]
		sumTask := sum
		if i < len(tasks)-1 {
			sumTask = sum + tasks[i+1][1]
		}
		if largestMinimum < tasks[i][1] {
			largestMinimum = tasks[i][1]
		} else if largestMinimum == tasks[i][1] {
			largestMinimum++
		}
		if energyRequired < sumTask {
			energyRequired = sumTask
		}
	}

	if energyRequired < largestMinimum {
		return largestMinimum
	}
	return energyRequired
}

func MinimumInitialEnergyToFinishTasks() {
	// tasks := [][]int{
	// 	{1, 2}, // 6
	// 	{2, 4}, // 4
	// 	{4, 8}, //
	// } // 8

	// tasks := [][]int{
	// 	{1, 3},
	// 	{2, 4},
	// 	{10, 11},
	// 	{10, 12},
	// 	{8, 9},
	// } // 32

	// tasks := [][]int{
	// 	{1, 7},
	// 	{2, 8},
	// 	{3, 9},
	// 	{4, 10},
	// 	{5, 11},
	// 	{6, 12},
	// } // 27

	// tasks := [][]int{
	// 	{3, 3},
	// 	{1, 6},
	// 	{1, 10},
	// 	{1, 2},
	// 	{4, 7},
	// 	{2, 4},
	// } //12

	// tasks := [][]int{
	// 	{1, 1},
	// 	{1, 3},
	// } // 3

	// tasks := [][]int{
	// 	{1, 1},
	// 	{1, 4},
	// 	{1, 4},
	// } // 5

	tasks := [][]int{
		{1, 2},
		{2, 6},
		{4, 8},
	} // 10

	// tasks := [][]int{
	// 	{1, 1},
	// 	{1, 4},
	// 	{1, 4},
	// } // 5

	output := minimumEffort(tasks)

	fmt.Printf("%d\n", output)
}
