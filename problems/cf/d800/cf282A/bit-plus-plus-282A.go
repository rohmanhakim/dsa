package cf282A

import "fmt"

func bitPlusPlus(statements []string) int {
	x := 0
	for _, s := range statements {
		for _, c := range s {
			if c == '+' {
				x += 1
				break
			} else if c == '-' {
				x -= 1
				break
			}
		}
	}
	return x
}

func BitPlusPlus() {
	var n int
	var statements []string

	fmt.Scan(&n)
	for range n {
		var s string
		fmt.Scan(&s)
		statements = append(statements, s)
	}

	out := bitPlusPlus(statements)
	fmt.Println(out)
}
