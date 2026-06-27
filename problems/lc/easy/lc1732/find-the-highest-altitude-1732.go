package lc1732

import "fmt"

func largestAltitude(gain []int) int {

	alt := 0
	maxAlt := 0
	for i := 0; i < len(gain); i++ {
		alt = alt + gain[i]
		if alt > maxAlt {
			maxAlt = alt
		}
	}
	return maxAlt
}

func FindTheHighestAltitude() {
	gain := []int{-5, 1, 5, 0, -7}
	out := largestAltitude(gain)
	fmt.Println(out)
}
