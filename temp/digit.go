package temp

import (
	"fmt"
)

func mainDigit() {
	x := 5678
	for x != 0 {
		digit := x % 10

		x /= 10

		fmt.Printf("%d %d\n", x, digit)
	}
}
