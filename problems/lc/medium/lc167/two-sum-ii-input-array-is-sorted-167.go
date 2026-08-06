package lc167

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left := 0
	right := n - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			break
		}
	}
	return []int{left + 1, right + 1}
}
