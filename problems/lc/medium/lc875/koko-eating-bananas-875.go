package lc875

import (
	"slices"
)

func minEatingSpeed(piles []int, h int) int {

	/*
		3 6 7 11

		sum = 3 + 6 + 7 + 11 = 27

		1 -> 1 * sum() = 1 * 27
		11 -> len(piles) = 4
		12 -> len(piles) = 4
		13 -> len(piles) = 4
		...

		k = between 1 to 11

		wants: slowest possible to get <= 8

		brute force :

		````
		minK = math.MaxInt

		for k = 1 to max(piles) {
			res = hoursToEatAllPiles(k)
			if res <= h  && k < minK {
				minK = k
			}
		}

		return minK
		```

		optimization: binary search from i to max(piles)
		keep evaluate hoursToEatAllPiles(mid) until left > right
		return the last valid mid
		```
		left := 1
		right := slices.Max(piles)
		minK = right
		for left <= right {
			mid := left + (((right - left)/2))

			res := hoursToEatAllPiles(mid)

			if res <= h {
				minK = mid
				right = mid-1
			} else {
				left = mid+1
			}
		}

		return minK
		```
	*/

	intCeilDiv := func(a, b int) int {
		return (a + b - 1) / b
	}

	hoursToEatAllPiles := func(k int) int {
		total := 0
		for _, p := range piles {
			/*
				divide p with k to get rate then round up
			*/

			rate := intCeilDiv(p, k)
			total += rate
		}
		return total
	}

	// -- optimized --

	left := 1
	right := slices.Max(piles)
	minK := right
	for left <= right {
		mid := left + ((right - left) / 2)

		res := hoursToEatAllPiles(mid)
		if res <= h {
			minK = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return minK

	/// --- brute force ---
	// minK := math.MaxInt

	// for k := 1; k <= slices.Max(piles); k++ {
	// 	res := hoursToEatAllPiles(k)

	// 	if res <= h && k < minK {
	// 		minK = k
	// 	}
	// }
	// return minK
}
