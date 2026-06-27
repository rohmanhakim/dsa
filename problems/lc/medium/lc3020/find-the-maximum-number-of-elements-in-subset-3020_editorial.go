package lc3020

func maximumLengthEditorial(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	best := 1

	// Special case: x = 1 (squaring keeps it at 1 forever)
	// The whole chain is [1,1,...,1], length must be odd.
	if f := freq[1]; f >= 1 {
		length := f
		if f%2 == 0 {
			length = f - 1
		}
		best = max(best, length)
	}

	for x := range freq {
		if x == 1 || freq[x] < 2 {
			continue
		}

		pairs := 0
		cur := x
		for {
			sq := cur * cur
			if freq[sq] >= 2 {
				pairs++
				cur = sq
			} else {
				// cur is highest level with freq >= 2
				if freq[sq] == 1 {
					// sq can serve as peak
					best = max(best, 2*(pairs+1)+1)
				} else {
					// cur itself is the peak (use 1 of its copies)
					best = max(best, 2*pairs+1)
				}
				break
			}
		}
	}

	return best
}
