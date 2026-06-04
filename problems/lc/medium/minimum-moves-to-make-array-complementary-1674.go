package medium

import (
	"fmt"
)

// func  minMoves(nums []int, limit int) int {
// 	minMove := 0
// 	left := 0
// 	right := len(nums) - 1
// 	sums := make(map[int]int)
// 	mostFrequentSum := 0
// 	highestSum := 0
// 	for left <= len(nums)-1 {
// 		sum := nums[left] + nums[right]
// 		if left < right {
// 			sums[sum]++
// 			if sums[sum] > mostFrequentSum {
// 				mostFrequentSum = sum
// 			}
// 			if sum > highestSum {
// 				highestSum = sum
// 			}
// 		}
// 		if left > right {
// 			fmt.Printf("most freq sum: %d\n", mostFrequentSum)
// 			fmt.Printf("left: %d, right: %d, nums[left]: %d, nums[right]: %d\n", left, right, nums[left], nums[right])
// 			diff := 0
// 			if sum > mostFrequentSum {
// 				diff = sum - mostFrequentSum
// 				for diff > 0 {
// 					minMove++
// 					fmt.Printf("--> diff %d\n", diff)
// 					if nums[left]-diff > 0 {
// 						break
// 					} else if nums[right]-diff > 0 {
// 						break
// 					}
// 					diff--
// 				}
// 			} else if sum < mostFrequentSum {
// 				diff = mostFrequentSum - sum
// 				for diff > 0 {
// 					minMove++
// 					if nums[left]+diff <= limit {
// 						break
// 					} else if nums[right]+diff <= limit {
// 						break
// 					}
// 					diff--
// 				}
// 			}
// 		}

// 		left++
// 		right--
// 	}
// 	return minMove
// }

// 1, 2, 4, 3

// 2, 3, 4, 5, 6, 7, 8
// 2, 1, 2, 2, 2, 2, 2
// 2, 4

func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, 2*limit+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		lo, hi := min(a, b), max(a, b)
		diff[2] += 2
		diff[2*limit+1] -= 2

		diff[lo+1] -= 1

		diff[a+b] -= 1

		diff[a+b+1] += 1
		diff[limit+hi+1] += 1

	}

	minM, total := n, 0
	for t := 2; t <= 2*limit; t++ {
		total += diff[t]
		if total < minM {
			minM = total
		}
	}

	return minM
}

func MinimumMovesToMakeArrayComplementary() {
	// nums := []int{50, 1, 2, 1, 2, 3, 4, 2, 3, 52}
	// limit := 100

	// nums := []int{1, 2, 2, 1}
	// limit := 2

	// nums := []int{1, 2, 4, 3}
	// limit := 4

	nums := []int{28, 50, 76, 80, 64, 30, 32, 84, 53, 8}
	limit := 84

	out := minMoves(nums, limit)

	// correct answer is 4 moves
	fmt.Printf("%d\n", out)
}
