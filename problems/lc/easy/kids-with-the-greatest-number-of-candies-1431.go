package easy

import (
	"fmt"
	"slices"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))

	candiesCopy := slices.Clone(candies)

	sort.Slice(candiesCopy, func(i, j int) bool {
		return candiesCopy[i] > candiesCopy[j]
	})
	greatest := candiesCopy[0]

	for i, c := range candies {
		if c+extraCandies >= greatest {
			res[i] = true
		}
	}

	return res
}

func kidsWithCandiesClaudeSonnet(candies []int, extraCandies int) []bool {
	greatest := slices.Max(candies) // single O(n) pass, no copy needed

	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= greatest
	}
	return res
}

func KidsWithTheGreatesNumberOfCandies() {
	// candies := []int{2, 3, 5, 1, 3}
	// extraCandies := 3 // t,t,t,f,t
	// candies := []int{4, 2, 1, 1, 1}
	// extraCandies := 1 // t,f,f,f,f
	candies := []int{12, 1, 12}
	extraCandies := 10 // t,f,t
	out := kidsWithCandiesClaudeSonnet(candies, extraCandies)

	fmt.Println(out)
}
