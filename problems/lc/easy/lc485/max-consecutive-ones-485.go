package lc485

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	new := 0
	for i, n := range nums {
		if n == 1 {
			new++
		}
		if n == 0 || i == len(nums)-1 {
			fmt.Printf("--> %d,  %d\n", i, new)
			if new >= max {
				max = new
			}
			new = 0
		}
	}
	return max
}

func MaxConsecutiveOnes() {
	// nums := []int{1, 1, 0, 1, 1, 1} // 3
	nums := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1} // 2
	out := findMaxConsecutiveOnes(nums)
	fmt.Println(out)
}
