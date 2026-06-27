package lc1137

import "fmt"

func tribonacciRecurse(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	return tribonacciRecurse(n-1) + tribonacciRecurse(n-2) + tribonacciRecurse(n-3)
}

func tribonacci(n int) int {

	t := [3]int{0, 1, 1}

	if n < 3 {
		return t[n]
	}

	out := 0
	for i := 0; i <= n-3; i++ {
		out = t[0] + t[1] + t[2]
		t[0] = t[1]
		t[1] = t[2]
		t[2] = out
	}

	return out
}

func tribonacciClaudeSonnet(n int) int {
	t := [3]int{0, 1, 1}
	if n < 3 {
		return t[n]
	}
	for i := 3; i <= n; i++ {
		t[0], t[1], t[2] = t[1], t[2], t[0]+t[1]+t[2]
	}
	return t[2]
}

func NthTribonacciNumber() {
	n := 2
	out := tribonacci(n)
	fmt.Println(out)
}
