package lc1291

import (
	"fmt"
	"strconv"
)

func sequentialDigits(low int, high int) []int {

	res := []int{}

	/*-------------- approach 1 (TLE)
	for i := low; i <= high; i++ {

		num := i

		prev := -1
		for num > 0 {
			lsd := num % 10
			if prev != -1 && prev-lsd != 1 {
				break
			}

			prev = lsd
			num = num / 10
		}
		if num == 0 {
			res = append(res, i)
		}

	}
	*/

	/*-------------- approach 2 (TLE)//
	for i := low; i <= high; i++ {
		numString := strconv.Itoa(i)
		left := 0
		right := len(numString) - 1

		seq := false
		lastRight := -1
		lastLeft := -1
		diff := len(numString) - 1
		for left < right {
			if int(numString[right]-numString[left]) != diff || (lastLeft != -1 && int(numString[left])-lastLeft != 1) || (lastRight != -1 && lastRight-int(numString[right]) != 1) {
				break
			}

			lastLeft = int(numString[left])
			lastRight = int(numString[right])
			diff -= 2
			left++
			right--
			if right == left && len(numString)%2 == 1 {
				midIdx := (len(numString) / 2)
				mid := numString[midIdx]
				prev := numString[midIdx-1]
				next := numString[midIdx+1]
				if next-mid == 1 && mid-prev == 1 {
					seq = true
				}
			} else if left-right == 1 {
				seq = true
			}
		}

		if seq {
			res = append(res, i)
		}

	}
	*/

	nLow := len(strconv.Itoa(low))
	nHigh := len(strconv.Itoa(high))

	numArr := []byte{}
	num := 0
	curr := '1'
	currLen := nLow
	for curr <= '9' && num < high && currLen <= nHigh {
		numArr = append(numArr, byte(curr))
		curr++
		num, _ = strconv.Atoi(string(numArr))

		if len(numArr) >= currLen && num <= high {
			if num >= low {
				res = append(res, num)
			}

			numArr = numArr[1:]
		}

		if curr == '9'+1 && currLen < nHigh {
			currLen++
			curr = '1'
			numArr = []byte{}
		}
	}

	return res
}

func LC1291() {
	low := 89
	high := 234
	// low := 10
	// high := 1000000000
	out := sequentialDigits(low, high)
	fmt.Println(out)
}
