package lc3754

func sumAndMultiply(n int) int64 {
	num := n
	sum := int64(0)
	x := int64(0)
	mul := 1
	for num > 0 {
		digit := num % 10
		if digit != 0 {
			x += (int64(digit) * int64(mul))
			mul *= 10
			sum += int64(digit)
		}
		num = num / 10
	}
	return sum * x
}
