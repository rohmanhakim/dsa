package lc2091

func minimumDeletions(nums []int) int {

	n := len(nums)

	minIdx, maxIdx := 0, 0

	for i := range n {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	if minIdx < maxIdx {
		res := min(maxIdx+1, n-minIdx)
		res = min(res, (minIdx+1)+(n-maxIdx))
		return res
	} else {
		res := min(minIdx+1, n-maxIdx)
		res = min(res, (maxIdx+1)+(n-minIdx))
		return res
	}
}
