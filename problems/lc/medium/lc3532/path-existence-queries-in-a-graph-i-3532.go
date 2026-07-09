package lc3532

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {

	res := make([]bool, len(queries))

	breaks := make([]int, n)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > maxDiff {
			breaks[i-1] = 1
		}
	}

	prefixBreaks := make([]int, n)

	for i := 1; i < n; i++ {
		prefixBreaks[i] = breaks[i-1] + prefixBreaks[i-1]
	}

	for i, q := range queries {
		if prefixBreaks[q[1]]-prefixBreaks[q[0]] == 0 {
			res[i] = true
		}
	}

	return res
}
