package lc70

import "fmt"

func climbStairsTopDown(n int) int {

	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		} else {
			if memo[n] != -1 {
				return memo[n]
			}
			res := dp(n-2) + dp(n-1)
			memo[n] = res
			return res
		}
	}

	return dp(n)
}

func climbStairsBottomUp(n int) int {
	// start from initial state, iterate to the final state

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	one, two := 1, 2
	climb := 0
	for i := 2; i < n; i++ {
		climb = one + two
		one, two = two, climb
	}

	return climb
}

func climbStairs(n int) int {
	out := climbStairsBottomUp(n)
	return out
}

func ClimbingStairs() {
	n := 44
	out := climbStairs(n)
	fmt.Println(out)
}
