package lc2144

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	minCost := 0
	a := 0
	b := 0
	for len(cost) > 0 {
		if a == 0 {
			a = cost[0]
			cost = cost[1:]
		} else if b == 0 {
			b = cost[0]
			cost = cost[1:]
		}
		if a > 0 && b > 0 {
			minCost += a + b
			a = 0
			b = 0
			if len(cost) > 0 {
				cost = cost[1:]
			}
		}
		if len(cost) == 1 {
			minCost += cost[0]
		}
	}

	return minCost
}

func minimumCostClaudeSonnet(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	total := 0
	for i, c := range cost {
		if (i+1)%3 != 0 {
			total += c
		}
	}
	return total
}

func MinimumOfBuyingCandiesWithDiscount() {
	// cost := []int{1, 2, 3}
	// cost := []int{5, 5}
	cost := []int{3, 3, 3, 1}
	// cost := []int{1}

	// cost := []int{6, 5, 7, 9, 2, 2}
	// 100, 99, 1

	out := minimumCost(cost)
	fmt.Println(out)
}
