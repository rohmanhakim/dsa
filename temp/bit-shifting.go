package temp

import "fmt"

func mainBitShifting() {
	a := 10
	b := 3

	x := a << 1
	fmt.Println(x)

	b <<= 1
	fmt.Println(b)
}
