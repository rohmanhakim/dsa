package lc57

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := [][]int{}

	nin := newInterval

	prev := []int{}
	merged := false
	for i := 0; i < len(intervals); {
		in := intervals[i]

		if len(prev) > 0 && nin[0] > prev[1] {
			if nin[1] < in[0] {
				res = append(res, nin)
				merged = true
			}
		}

		if nin[1] < in[0] {
			res = append(res, in)
			i++
		} else if nin[0] > in[1] {
			res = append(res, in)
			i++
		} else if in[1] >= nin[0] {
			left := min(in[0], nin[0])
			right := max(in[1], nin[1])

			for i < len(intervals) && intervals[i][0] <= nin[1] {
				if intervals[i][1] > right {
					right = intervals[i][1]
				}
				i++
			}
			res = append(res, []int{left, right})
			merged = true
		}
		prev = in
	}

	if !merged {
		if nin[1] < res[0][0] {
			res = append([][]int{nin}, res...)
		} else {
			res = append(res, nin)
		}
	}

	return res
}
