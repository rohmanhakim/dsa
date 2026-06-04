package easy

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func timeRequiredToBuy(tickets []int, k int) int {
	tickets[k] = tickets[k] * -1
	time := 0
	for len(tickets) > 0 {
		time++
		front := tickets[0]
		tickets = tickets[1:]
		if lib.Abs(front) > 0 {
			if front == -1 {
				break
			}
			val := lib.Abs(front) - 1
			if front < 0 {
				val = val * -1
			}

			if lib.Abs(val) > 0 {
				tickets = append(tickets, val)
			}
		}

	}

	return time
}

func TimeNeededToBuyTickets() {
	tickets := []int{2, 3, 2}
	k := 2
	// tickets := []int{5, 1, 1, 1}
	// k := 0
	out := timeRequiredToBuy(tickets, k)
	fmt.Println(out)
}
