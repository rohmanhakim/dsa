package lc3622

func checkDivisibility(n int) bool {
	num := n

	prod := 1
	sum := 0
	for num > 0 {
		digit := num % 10

		sum += digit
		prod *= digit
		num = num / 10
	}

	return (n%(prod+sum) == 0)
}
