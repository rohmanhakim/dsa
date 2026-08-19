package lc1386

import (
	"sort"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	sort.Slice(reservedSeats, func(i, j int) bool {
		return reservedSeats[i][0] < reservedSeats[j][0]
	})

	count := 0
	i := 1
	row := reservedSeats[0][0]
	curr := 0
	for i <= n && curr < len(reservedSeats) {
		if i != row {
			i++
			count += 2
			continue
		}

		seats := make([]int, 10)

		for j := curr; j < len(reservedSeats); j++ {
			if reservedSeats[j][0] == row {
				seats[reservedSeats[j][1]-1] = 1
			} else {
				row = reservedSeats[j][0]
				break
			}
			curr++
		}

		leftBlock := 0
		rightBlock := 0
		midBlock := 0

		for j := 1; j <= 4; j++ {
			if seats[j] == 0 {
				leftBlock++
			} else {
				leftBlock = 0
				break
			}
		}
		for j := 3; j <= 6; j++ {
			if seats[j] == 0 {
				midBlock++
			} else {
				midBlock = 0
				break
			}
		}

		for j := 5; j <= 8; j++ {
			if seats[j] == 0 {
				rightBlock++
			} else {
				rightBlock = 0
				break
			}
		}

		available := 0
		if leftBlock > 0 || rightBlock > 0 {
			if leftBlock > 0 {
				available++
			}
			if rightBlock > 0 {
				available++
			}
		} else if midBlock > 0 {
			available++
		}

		count += available
		i++
	}

	if i > row && i <= n {
		count += (n - row) * 2
	}

	return count
}
