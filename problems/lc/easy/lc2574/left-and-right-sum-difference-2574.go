package lc2574

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func leftRightDifference(nums []int) []int {

	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	answer := make([]int, len(nums))
	leftSum[0] = 0
	rightSum[len(nums)-1] = 0
	for i, j := 1, len(nums)-2; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[j] = rightSum[j+1] + nums[j+1]
	}
	for i, n := range leftSum {
		answer[i] = abs(n - rightSum[i])
	}
	return answer
}

func leftRightDifferenceClaudeSonnet(nums []int) []int {
	total := 0
	for _, n := range nums {
		total += n
	}

	answer := make([]int, len(nums))
	left := 0
	for i, n := range nums {
		right := total - left - n
		answer[i] = abs(left - right)
		left += n
	}
	return answer
}

func LeftAndRightSumDifferences() {
	nums := []int{10, 4, 8, 3}
	out := leftRightDifference(nums)

	fmt.Println(out)
}
