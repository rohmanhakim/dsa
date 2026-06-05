package d800

import (
	"fmt"
)

func Team() {
	var n, num, p, v, t int

	fmt.Scan(&n)
	for range n {

		fmt.Scan(&p, &v, &t)
		sure := p + v + t
		if sure >= 2 {
			num++
		}
	}
	fmt.Println(num)
}
