package lc3870

import "fmt"

func countCommas(n int) int {
	if n < 1000 {
		return 0
	}

	return n - 1000 + 1
}

func LC3870() {
	n := 1002
	out := countCommas(n)
	fmt.Println(out)
}
