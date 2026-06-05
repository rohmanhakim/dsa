package easy

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	pick := 6

	if num < pick {
		return 1
	} else if num > pick {
		return -1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left := 1
	right := n
	res := n

	for left <= right {
		mid := left + ((right - left) / 2)
		res = guess(mid)
		switch res {
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}

func GuessNumberHigherOrLower() {
	n := 10
	out := guessNumber(n)
	fmt.Println(out)
}
