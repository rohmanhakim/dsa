package acabc081a

import (
	"fmt"
)

func PlacingMarbles() {
	var inp string
	fmt.Scanf("%s", &inp)

	count := 0
	left := 0
	right := len(inp) - 1
	for left < right {
		if inp[left] == '1' {
			count++
		}
		if inp[right] == '1' {
			count++
		}
		left++
		right--
	}
	fmt.Println(count)
}
