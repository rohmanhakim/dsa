package main

import (
	"fmt"
)

func Team() {
	var n, num, p, v, t int

	fmt.Scanf("%d", &n)
	for range n {

		fmt.Scanf("%d %d %d", &p, &v, &t)
		sure := p + v + t
		if sure >= 2 {
			num++
		}
	}
	fmt.Println(num)
}

func main() {
	Team()
}
