package lc1111

func maxDepthAfterSplit(seq string) []int {
	depth := 0
	layers := make([]int, len(seq))

	for i, ch := range seq {
		if ch == '(' {
			depth++
			layers[i] = depth
		} else {
			layers[i] = depth
			depth--
		}
	}

	res := make([]int, len(seq))

	for i, l := range layers {
		if l%2 == 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}
