package easy

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	remaining := n
	for i := 0; i < len(flowerbed); i++ {

		leftFree := true
		righFree := true
		if flowerbed[i] == 0 {
			if i > 0 {
				if flowerbed[i-1] == 1 {
					leftFree = false
				}
			}
			if i < len(flowerbed)-1 {
				if flowerbed[i+1] == 1 {
					righFree = false
				}
			}
			if leftFree && righFree {
				flowerbed[i] = 1
				remaining--
			}
		}
		if remaining == 0 {
			return true
		}
	}

	return false
}

func canPlaceFlowersClaudeSonnet(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		leftFree := i == 0 || flowerbed[i-1] == 0
		rightFree := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if flowerbed[i] == 0 && leftFree && rightFree {
			flowerbed[i] = 1
			n--
		}
	}
	return n == 0
}

func CanPlaceFlowers() {
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 1 // true

	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 2 // false

	// flowerbed := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	// n := 2 // true

	// flowerbed := []int{0, 0, 1, 0, 1}
	// n := 1 // true

	// flowerbed := []int{1, 0, 0, 0, 1, 0, 0}
	// n := 2 // true

	flowerbed := []int{0, 0, 0, 0, 0, 1, 0, 0}
	n := 0 // true

	out := canPlaceFlowers(flowerbed, n)
	fmt.Println(out)
}
