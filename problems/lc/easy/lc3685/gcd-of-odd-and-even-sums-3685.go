package lc3685

import "fmt"

func gcdOfOddEvenSums(n int) int {

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	sumodd, sumeven := 0, 0

	count := 0
	for i := 1; count < n; i++ {
		if i%2 == 1 {
			sumodd += i
			count++
		}
	}
	count = 0

	for i := 2; count < n; i++ {
		if i%2 == 0 {
			sumeven += i
			count++
		}
	}

	return gcd(sumodd, sumeven)
}

func LC3685() {
	n := 4
	out := gcdOfOddEvenSums(n)
	fmt.Println(out)
}
