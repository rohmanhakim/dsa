package lc3751

import "fmt"

func totalWaviness(num1 int, num2 int) int {

	total := 0
	for n := num1; n <= num2; n++ {
		number := n
		digits := []int{}
		for number > 0 {
			digits = append(digits, number%10)
			number = number / 10
		}
		waviness := 0
		if len(digits) >= 3 {
			for i := 1; i < len(digits)-1; i++ {
				peak := digits[i] > digits[i-1] && digits[i] > digits[i+1]
				valley := digits[i] < digits[i-1] && digits[i] < digits[i+1]
				if peak || valley {
					waviness++
				}
			}
		}
		total += waviness
	}

	return total
}

func TotalWavinessOfNumbersInRangeI() {
	num1 := 120
	num2 := 130 // 3

	// num1 := 198
	// num2 := 202 // 3

	// num1 := 4848
	// num2 := 4848 // 2
	out := totalWaviness(num1, num2)
	fmt.Println(out)
}
