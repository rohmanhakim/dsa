package medium

import (
	"fmt"
)

/*

875. Koko Eating Bananas

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.



Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4

Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30

Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23



Constraints:

    1 <= piles.length <= 10^4
    piles.length <= h <= 10^9
    1 <= piles[i] <= 10^9

*/

// RECOGNITION FRAMEWORK - Binary Search on Answer
//
// STEP 1: First 30 seconds - Check these signals:
// □ "Find minimum speed/rate/capacity/time" OR "Find maximum minimum"
// □ Input size is large (10^4 to 10^9) - suggests O(n log n) or O(log n) needed
// □ Brute force would be O(n * range) which is too slow
// □ There's a constraint like "must finish within X hours/days"
//
// STEP 2: The Monotonicity Test (2 minutes)
// Ask: "If I try a candidate answer k, and it works, would a larger k also work?"
//   - If YES → Binary search on answer applies
//   - Ask: "If k works, does k+1 work?" (For minimization problems)
//   - Ask: "If k fails, does k-1 fail?" 
//
// Example monotonicity check for this problem:
//   "If eating speed 10 works (finishes ≤ h hours), does speed 11 work?"
//   Yes! Faster speed only helps. Monotonicity confirmed.
//
// STEP 3: Define the Feasibility Function (3 minutes)
// Write a helper: canFinish(candidate) -> bool
// - Must be O(n) or better
// - Calculate total time/resources needed at this candidate rate
// - Return true if within constraints, false otherwise
//
// STEP 4: Set Search Boundaries (1 minute)
// - Low bound: Usually 1, or min possible value (e.g., max element for "at least one per day")
// - High bound: Usually max(input), or sum(input), or some physical constraint
// - Ask: "What's the worst case I need to handle?"
//
// STEP 5: Binary Search Template (5 minutes)
// left, right := low, high
// result := high  // Start with worst case
// for left <= right {
//     mid := left + (right-left)/2
//     if feasible(mid) {
//         result = mid        // Found valid, try better (smaller for min, larger for max)
//         right = mid - 1     // For minimization: go left
//     } else {
//         left = mid + 1      // Not valid, need more resources/speed
//     }
// }
// return result
//
// COMMON MISTAKES TO AVOID:
// 1. Using sort.BinarySearch on unsorted conceptual space (use manual binary search)
// 2. Integer overflow in mid calculation: use left + (right-left)/2, not (left+right)/2
// 3. Off-by-one in boundary conditions - always test with 1-element arrays
// 4. Forgetting that the feasibility check might need ceiling division
// 5. Setting right = max(piles) when you actually need sum(piles) (for capacity problems)

// ceilDiv calculates ceiling division using integer arithmetic.
// Formula: (work + rate - 1) / rate
// 
// Why this works:
// - We want ceil(work / rate)
// - In integer arithmetic, (a + b - 1) / b is equivalent to ceiling(a / b)
// - Avoids floating point operations (faster, no precision issues with large integers up to 10^9)
//
// Example: ceil(7 / 3) = ceil(2.33) = 3
//          (7 + 3 - 1) / 3 = 9 / 3 = 3 ✓
func ceilDiv(work int, rate int) int {
	return (work + rate - 1) / rate
}

// hoursNeededToEat calculates hours needed to finish one pile at speed k.
// If pile has 7 bananas and k = 3, she needs 3 hours (3, 3, 1).
func hoursNeededToEat(pile int, k int) int {
	return ceilDiv(pile, k)
}

// totalHoursToEatAll sums up hours needed for all piles at speed k.
// Time Complexity: O(n) where n = len(piles)
func totalHoursToEatAll(piles []int, k int) int {
	total := 0
	for _, p := range piles {
		hour := hoursNeededToEat(p, k)
		total += hour
	}
	return total
}

// minEatingSpeed uses BINARY SEARCH ON ANSWER pattern.
//
// KEY CONCEPT 1: Binary Search on Answer
// - We're not searching the input array (piles isn't sorted)
// - We're searching the "answer space": possible values of k (1 to max(piles))
// - Recognition signal: "Find minimum maximum" or "Find maximum minimum" problems
//
// KEY CONCEPT 2: Monotonicity Property (Feasibility Function)
// - If speed k works (finishes ≤ h hours), then ANY speed > k also works
// - If speed k fails (takes > h hours), then ANY speed < k also fails
// - This creates a "boundary" between false and true that binary search can find
// - This property is what makes binary search on answer valid
//
// KEY CONCEPT 3: Decision vs Optimization
// - Original problem: Optimization (find minimum k)
// - Converted to: Decision (can we finish in h hours with speed k?)
// - If we can check feasibility in O(n) and search space is large (10^9),
//   binary search gives us O(n log(max)) instead of O(n * max)
//
// Algorithm:
// - left = 1 (slowest possible, eating 1 banana/hour)
// - right = max(piles) (fastest needed, one pile per hour)
// - Binary search for the smallest k where totalHoursToEatAll(piles, k) <= h
//
// Time Complexity: O(n log(max(piles)))
// - log(max(piles)) iterations of binary search
// - Each iteration does O(n) work to calculate total hours
func minEatingSpeed(piles []int, h int) int {
	// Find the maximum pile to set the upper bound
	// Eating at speed = max(pile) guarantees we finish each pile in 1 hour
	// Total time = len(piles) hours, which is always <= h (given constraint: len(piles) <= h)
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	left := 1
	right := maxPile
	result := maxPile

	for left <= right {
		mid := left + (right-left)/2
		hours := totalHoursToEatAll(piles, mid)

		if hours <= h {
			// This speed works, but we want the MINIMUM valid speed
			// Since we want minimum, try slower speeds (lower k) on the left side
			result = mid
			right = mid - 1
		} else {
			// Too slow (takes too many hours), need faster speed
			// Search right side for larger k values
			left = mid + 1
		}
	}

	return result
}

// Similar problems to practice this pattern:
// - 1011. Capacity To Ship Packages Within D Days (exact same pattern)
// - 2187. Minimum Time to Complete Trips (binary search on time)
// - 410. Split Array Largest Sum (binary search on maximum subarray sum)
// - Aggressive Cows (binary search on minimum distance)
func KokoEatingBananas() {
	// piles := []int{3, 6, 7, 11}
	// h := 8 // 4
	piles := []int{30, 11, 23, 4, 20}
	h := 5 // 30
	// piles := []int{30,11,23,4,20}
	// h := 6 // 23
	out := minEatingSpeed(piles, h)
	// fmt.Println(totalHoursToEatAll(piles, 4))
	fmt.Println(out)
}
