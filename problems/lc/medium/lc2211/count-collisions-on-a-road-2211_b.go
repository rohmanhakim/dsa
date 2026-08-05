package lc2211

import "strings"

func countCollisionsAlternative(directions string) int {
	for len(directions) > 0 && directions[0] == 'L' {
		directions = strings.TrimPrefix(directions, "L")
	}

	for len(directions) > 0 && directions[len(directions)-1] == 'R' {
		directions = strings.TrimSuffix(directions, "R")
	}
	count := 0
	for _, ch := range directions {
		if ch != 'S' {
			count++
		}
	}
	return count
}
