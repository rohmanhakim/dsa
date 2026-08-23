package lc1927

func sumGame(num string) bool {
	n := len(num)
	mid := n / 2

	q0 := 0
	q1 := 0
	n0 := 0
	n1 := 0
	for i, ch := range num {
		if i < mid {
			if ch == '?' {
				q0++
			} else {
				n0 += int(ch - '0')
			}
		} else {
			if ch == '?' {
				q1++
			} else {
				n1 += int(ch - '0')
			}
		}
	}

	if (q0+q1)%2 == 1 {
		return true
	} else {
		return 2*(n0-n1) != 9*(q1-q0)
	}
}
