package lc26

func removeDuplicates(nums []int) int {
	occ := make(map[int]bool)

	k := -1
	for _, num := range nums {
		if !occ[num] {
			k++
			nums[k] = num
			occ[num] = true
		}
	}
	return k + 1
}
