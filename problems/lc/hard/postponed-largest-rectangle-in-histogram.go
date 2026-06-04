package hard

/*
func largestRectangleArea(heights []int) int {
	incStack := []int{}
	decStack := []int{}

	// 2      2          2
	// 1      2 1        1 -> 2
	// 5      5 -> 2, 1  1, 5            (5 * 1), (3 * 1)
	// 6      6 -> 5     1, 5, 6         (6 * 1), (4 * 1), (5 * 2)
	// 2      6, 2       2 -> 1, 5, 6    (2 * 1), (5 * 1), (2 * 3)
	// 3      3 -> 6,2   2, 3            (3 * 1), (6 * 1), (2 * 4)

	// 3      3                3
	// 2      3, 2             2 -> 3
	// 6      6 -> 3, 2        2, 6
	// 5      6, 5             5 -> 2, 6
	// 1      6, 5, 1          1 -> 5
	// 2      2 -> 6, 5, 1     1, 2

	// 1     1           1
	// 2     2 -> 1      1, 2
	// 3     3 -> 2      1, 2, 3
	// 4     4 -> 3      1, 2, 3, 4,
	// 5     5 -> 4      1, 2, 3, 4, 5   (5 * 1), (5 * 1), ()

	max := 0
	currMinIndex := 0
	currMin := heights[currMinIndex]
	for i, height := range heights {
		if height <= currMin {
			currMin = height
			currMinIndex = i
		}
		a := height * 1
		b := (i + 1) * currMin
		c := 0

		// c := ((i + 1) - currMinIndex) * height

		for len(decStack) > 0 && heights[decStack[len(decStack)-1]] > height {
			topIndex := decStack[len(decStack)-1]
			decStack = decStack[:len(decStack)-1]
			c = (i - topIndex + 1) * height
		}
		// fmt.Println(incStack)
		for len(incStack) > 0 && heights[incStack[len(incStack)-1]] <= height {
			topIndex := incStack[len(incStack)-1]
			incStack = incStack[:len(incStack)-1]
			temp := (i - topIndex + 1) * heights[topIndex]

			if i == 4 {
				fmt.Printf("i: %d, height: %d, topIndex: %d, heights[top]: %d, minIndex: %d, min: %d, max: %d, temp: %d\n", i, height, topIndex, heights[topIndex], currMinIndex, currMin, max, temp)
			}
			if temp > c {
				c = temp
			}
		}

		candidates := []int{a, b, c}
		fmt.Println(candidates)
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i] > candidates[j]
		})
		// fmt.Println(candidates)

		if candidates[0] > max {
			max = candidates[0]
		}

		incStack = append(incStack, i)
	}
	return max
}
*/

// 2	2			2
// 1	2, 1		1 -> 2
// 5	5 -> 2,1	1, 5
// 6    6 ->5		1, 6 -> 5
// 2    6, 2		1, 2 -> 6
// 3    6, 3 -> 2   1, 3 -> 2

// 1	1			1
// 2	2 -> 1		1, 2
// 3	3 -> 2		1, 2, 3
// 4 	4 -> 3		1, 2, 3, 4
// 5	5 -> 4		1, 2, 3, 4, 5

func LargestRectangleInHistogram() {
	// heights := []int{2, 1, 5} // 5s
	// heights := []int{2, 1, 5, 6, 2, 3} // 10
	// heights := []int{2, 1, 2, 6, 2, 3} // 8
	// heights := []int{2, 1, 5, 5, 5, 5, 6, 2, 3} // 25
	// heights := []int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 6, 2, 3} // 17
	// heights := []int{1, 2, 2} // 4
	// heights := []int{1, 2, 3, 4, 5}
	// heights := []int{6, 5, 4, 3, 2}

	// out := largestRectangleArea(heights)
	// fmt.Println(out)
}
