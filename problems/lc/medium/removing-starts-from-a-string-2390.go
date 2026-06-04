package medium

import "fmt"

func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func RemovingStarsFromAString() {
	// s := "leet**cod*e" // lecoe
	s := "erase*****"
	out := removeStars(s)
	fmt.Println(out)
}
