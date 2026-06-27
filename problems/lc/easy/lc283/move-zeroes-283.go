package lc283

import "fmt"

func moveZeroes(nums []int) {
	b := 0
	for a := 0; a < len(nums); a++ {
		if nums[a] != 0 {
			nums[a], nums[b] = nums[b], nums[a]
			b++
		}
	}
}

func MoveZeroes() {
	// 0, 1, 0, 3, 12
	// ab
	// 1, 0, 0, 3, 12
	// b  a
	// 1, 0, 0, 3, 12
	//    b  a
	// 1, 0, 0, 3, 12
	//    b     a
	// 1, 3, 0, 0, 12
	//       b  a
	// 1, 3, 0, 0, 12
	//       b     a
	// 1, 3, 12, 0, 0

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
