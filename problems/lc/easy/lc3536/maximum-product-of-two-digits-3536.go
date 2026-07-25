package lc3536

func maxProduct(n int) int {
	largest := -1
	secondLargest := -1

	num := n

	for num > 0 {
		digit := num % 10
		if digit > largest {
			if largest > -1 {
				secondLargest = largest
			}
			largest = digit

		} else if digit > secondLargest {
			secondLargest = digit
		}
		num = num / 10
	}
	return largest * secondLargest
}
