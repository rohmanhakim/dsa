package lc835

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)

	translate := func(m [][]int, dir []int) [][]int {
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				resCol := dir[0] + i
				resRow := dir[1] + j
				if resCol >= 0 && resCol < n && resRow >= 0 && resRow < n {
					mat[resRow][resCol] = m[j][i]
				}
			}
		}

		return mat
	}

	overlay := func(a, b [][]int) int {
		count := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j]&b[i][j] == 1 {
					count++
				}
			}
		}
		return count
	}

	maxOverlap := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res := overlay(translate(img1, []int{i, j}), img2)
			if res > maxOverlap {
				maxOverlap = res
			}

			res = overlay(translate(img1, []int{-i, j}), img2)
			if res > maxOverlap {
				maxOverlap = res
			}

			res = overlay(translate(img1, []int{i, -j}), img2)
			if res > maxOverlap {
				maxOverlap = res
			}

			res = overlay(translate(img1, []int{-i, -j}), img2)
			if res > maxOverlap {
				maxOverlap = res
			}
		}
	}
	return maxOverlap
}
