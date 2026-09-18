package lib

import "slices"

func Equal2DUnordered(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Make deep copies so we don't modify the inputs.
	ac := make([][]string, len(a))
	bc := make([][]string, len(b))

	for i := range a {
		ac[i] = slices.Clone(a[i])
		slices.Sort(ac[i])
	}

	for i := range b {
		bc[i] = slices.Clone(b[i])
		slices.Sort(bc[i])
	}

	// Sort the outer slice lexicographically.
	sort2D(ac)
	sort2D(bc)

	// Compare.
	for i := range ac {
		if !slices.Equal(ac[i], bc[i]) {
			return false
		}
	}

	return true
}

func sort2D(s [][]string) {
	slices.SortFunc(s, func(a, b []string) int {
		n := min(len(a), len(b))
		for i := 0; i < n; i++ {
			if a[i] < b[i] {
				return -1
			}
			if a[i] > b[i] {
				return 1
			}
		}

		switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
		default:
			return 0
		}
	})
}

func Equal1DUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// Track frequencies of elements in slice a
	counts := make(map[T]int, len(a))
	for _, val := range a {
		counts[val]++
	}

	// Subtract frequencies using slice b
	for _, val := range b {
		if counts[val] == 0 {
			return false // Element not found or occurs too many times
		}
		counts[val]--
	}

	return true
}
