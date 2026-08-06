package lc3345

/*
func smallestNumber(n int, t int) int {
	prod := 1
	num := n
	digits := []int{}
	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		prod = prod * digit
		num = num / 10
	}

	candidates := []int{}
	for i := 0; i < len(digits); i++ {
		carry := 0
		newProd := prod

		newDigit := digits[i]
		for j := digits[i]; j <= 9 && newProd%t != 0; j++ {
			newProd = (prod / digits[i]) * j
			newDigit = j
		}

		if newProd%t != 0 {
			carry++
		}

		fmt.Println(newDigit, carry)
		mult := 1
		cand := 0
		for j := range len(digits) {
			if j == i {
				if carry == 0 {
					cand += mult * newDigit
				} else {
					cand += mult * 0
				}
			} else {
				cand += mult * digits[j]
			}
			mult *= 10
		}
		if carry > 0 {
			cand += carry * mult
		}
		candidates = append(candidates, cand)
	}

	smallest := slices.Min(candidates)
	return smallest
}
*/

func smallestNumber(n int, t int) int {
	num := n

	prod := 1
	for num > 0 {
		digit := num % 10
		prod = prod * digit
		num = num / 10
	}

	num = n

	for prod%t != 0 {
		num++
		newNum := num
		newProd := 1
		for newNum > 0 {
			digit := newNum % 10
			newProd = newProd * digit
			newNum = newNum / 10
		}
		prod = newProd
	}
	return num
}
