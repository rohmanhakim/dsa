package lc27

func removeElement(nums []int, val int) int {

	n := len(nums)
	read := 0
	write := 0
	total := 0
	for read < n {
		if nums[read] == val {
			total++
		} else {
			nums[write] = nums[read]
			write++
		}
		read++
	}

	return n - total
}
