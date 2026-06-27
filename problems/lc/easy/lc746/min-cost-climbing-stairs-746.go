package lc746

import (
	"fmt"
	"math"
)

func getBestCostTopDown(cost []int, step int) int {

	// fmt.Printf("step[%d] ", step)

	if step < 2 {
		fmt.Println("")
		return 0
	}

	stepCostSingleClimb := cost[step-1]
	stepCostDoubleClimb := cost[step-2]

	// fmt.Printf("sc1: %d, sc2: %d\n", stepCostSingleClimb, stepCostDoubleClimb)

	stratSingleClimb := getBestCostTopDown(cost, step-1) + stepCostSingleClimb
	stratDoubleClimb := getBestCostTopDown(cost, step-2) + stepCostDoubleClimb

	// fmt.Printf("step[%d] sc1: %d, sc2: %d, single: %d, double %d\n", step, stepCostSingleClimb, stepCostDoubleClimb, stratSingleClimb, stratDoubleClimb)

	return min(stratSingleClimb, stratDoubleClimb)
}

func getBestCostTopDownWithMemo(cost []int) int {

	memo := make([]int, len(cost)+1)
	for i := range memo {
		memo[i] = math.MaxInt32
	}

	var dp func(step int) int
	dp = func(step int) int {
		if step < 2 {
			return 0
		}

		if memo[step] != math.MaxInt32 {
			return memo[step]
		}

		stepCostSingleClimb := cost[step-1]
		stepCostDoubleClimb := cost[step-2]

		stratSingleClimb := dp(step-1) + stepCostSingleClimb
		stratDoubleClimb := dp(step-2) + stepCostDoubleClimb

		res := min(stratSingleClimb, stratDoubleClimb)
		memo[step] = res
		return res
	}
	return dp(len(cost))
}

func minCostClimbingStairs(cost []int) int {
	// out := getBestCostTopDown(cost, len(cost))
	out := getBestCostTopDownWithMemo(cost)
	return out
}

func MinCostClimbingStairs() {
	// cost := []int{10, 15, 20} // 15
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1} // 6
	out := minCostClimbingStairs(cost)
	fmt.Println(out)
}
