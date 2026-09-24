package lc3550

func smallestIndex(nums []int) int {

	for i := range nums {
		num := nums[i]

		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit
			num = num / 10
		}
		if sum == i {
			return i
		}
	}
	return -1
}
