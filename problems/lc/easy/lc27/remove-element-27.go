package lc27

func removeElement(nums []int, val int) int {

	/*
		3 2 2 3 -> 2 2 _ _

		loop through nums
		maintain pointer to read the current nums[i]
		maintain pointer to write started at index 0
		constantly read nums at read pointer
		if nums[read] == val : ignore
		if nums[read] != val :
			advance counter for k
			write nums[read] to nums[write]
			advance write pointer by 1



		read := 0
		write := 0
		k: 0
		for read < len(nums) {
			curr := nums[read]

			if curr != val {
		    	k++
				nums[write] = curr
				write++
			}
			read++
		}
	*/

	read, write := 0, 0

	k := 0
	for read < len(nums) {
		curr := nums[read]

		if curr != val {
			k++
			nums[write] = curr
			write++
		}
		read++
	}

	return k
}
