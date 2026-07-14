package temp

import "fmt"

func Gcd(a, b int) int {
	for b != 0 {
		fmt.Printf("%d / %d = %d, remaining: %d\n", a, b, a/b, a%b)
		a, b = b, a%b
	}
	return a
}
