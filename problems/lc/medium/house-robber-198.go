package medium

import (
	"fmt"
	"math"
)

/*
The key lesson

after conversation with an LLM, my key lesson is "Model the constraint into the transition, not the state."

My progression:

 • Initial mistake: Carried full history (robbedHouses map) to validate adjacency, creating
   exponential state space
 • Correct model: State = house index only; adjacency enforced by transitions (i-2 vs i-1)

The pattern for future DP problems:

 1 Identify the decision: At each step, what binary (or n-ary) choice do i make? (Rob vs Skip)
 2 Determine minimal state: Ask "What information from the past actually affects future
   decisions?"
    • Not: "Which houses did I rob?" (too much)
    • But: "Where am I?" (index) — the constraint is handled by where i can go next
 3 Bake constraints into transitions: Instead of checking "Is this move legal?" carry the
   consequence forward:
    • Rob → next state is i-2 (forced skip)
    • Skip → next state is i-1 (free choice)
 4 Trust the recurrence: The overlapping subproblems will naturally emerge once the transition
   graph is correct; then add memoization

Red flags to watch for:

 • If my state includes data structures tracking history (maps, sets, lists), i'm probably
   over-specifying
 • If i'm validating moves by querying history, the constraint belongs in the state transition
   instead
 • If the recurrence doesn't have a clear "take vs not-take" (or similar) branching, i haven't
   identified the decision point

This applies to House Robber, Climbing Stairs, Longest Increasing Subsequence, and most linear
sequence DP problems.
*/

func robTopDown(nums []int) int {

	var dp func(house int) int

	memo := make([]int, len(nums))

	for i := range memo {
		memo[i] = math.MinInt32
	}

	dp = func(house int) int {
		if house < 0 {
			return 0
		}

		memoizedVal := memo[house]
		if memoizedVal != math.MinInt32 {
			return memoizedVal
		}

		value := nums[house]

		stratRob := dp(house-2) + value

		stratSkip := dp(house - 1)

		res := max(stratRob, stratSkip)
		memo[house] = res
		return res
	}

	return dp(len(nums) - 1)
}

func rob(nums []int) int {
	res := robTopDown(nums)
	return res
}

func HouseRobber() {
	// nums := []int{1, 2, 3, 1}
	nums := []int{2, 7, 9, 3, 1}
	out := rob(nums)
	fmt.Println(out)

}
