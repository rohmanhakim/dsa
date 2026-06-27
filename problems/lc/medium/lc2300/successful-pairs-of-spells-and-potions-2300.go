package lc2300

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	out := make([]int, len(spells))

	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})

	for i, s := range spells {
		idx := sort.Search(len(potions), func(i int) bool {
			return int64(s)*int64(potions[i]) >= success
		})

		out[i] = len(potions) - idx

		/* // ALternative manual binary search

		lo, hi := 0, len(potions)
		for lo < hi {
		    mid := (lo + hi) / 2
		    if int64(s)*int64(potions[mid]) >= success {
		        hi = mid
		    } else {
		        lo = mid + 1
		    }
		}
		out[i] = len(potions) - lo
		*/
	}
	return out
}

func SuccessfulPairsOfSpellsAndPotions() {
	// spells := []int{5, 1, 3}
	// potions := []int{1, 2, 3, 4, 5}
	// success := 7

	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	success := 16

	out := successfulPairs(spells, potions, int64(success))
	fmt.Println(out)
}
