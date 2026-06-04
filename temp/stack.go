package temp

import (
	"fmt"
)

func mainStack() {
	n := []int{1, 2, 3, 4, 5, 6}

	popped := n[len(n)-1]
	n = n[:len(n)-1]

	fmt.Println(popped)
	fmt.Println(n)
}
