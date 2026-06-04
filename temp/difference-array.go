package temp

import "fmt"

func mainDifferenceArray() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	diff := make([]int, 17)

	diff[0] += 2
	diff[2] -= 2

	diff[8] += 7
	diff[12] -= 7

	diff[13] += 3
	diff[15] -= 3

	running := 0
	for i := 0; i < len(nums); i++ {
		running += diff[i]
		nums[i] += running
	}

	fmt.Println(nums)
}
