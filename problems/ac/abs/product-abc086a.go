package abs

import (
	"fmt"
)

func Product() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	even := (a*b)%2 == 0
	if even {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
