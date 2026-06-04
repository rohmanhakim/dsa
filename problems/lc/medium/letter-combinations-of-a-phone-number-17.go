package medium

import (
	"fmt"
)

func phoneDigitCartesian(res *[]string, pools [][]rune, path *[]rune) {
	if len(pools) == len(*path) {
		*res = append(*res, string(*path))
		return
	}

	for _, item := range pools[len(*path)] {
		*path = append(*path, item)
		phoneDigitCartesian(res, pools, path)
		a := *path
		*path = a[:len(a)-1]
	}
}

func letterCombinations(digits string) []string {
	keypad := make(map[rune][]rune)
	keypad['2'] = []rune{'a', 'b', 'c'}
	keypad['3'] = []rune{'d', 'e', 'f'}
	keypad['4'] = []rune{'g', 'h', 'i'}
	keypad['5'] = []rune{'j', 'k', 'l'}
	keypad['6'] = []rune{'m', 'n', 'o'}
	keypad['7'] = []rune{'p', 'q', 'r', 's'}
	keypad['8'] = []rune{'t', 'u', 'v'}
	keypad['9'] = []rune{'w', 'x', 'y', 'z'}

	pools := [][]rune{}
	path := []rune{}
	res := []string{}

	for _, d := range digits {
		pools = append(pools, keypad[d])
	}

	phoneDigitCartesian(&res, pools, &path)

	return res
}

func LetterCombinationsOfPhoneNumber() {
	digits := "23"
	out := letterCombinations(digits)
	fmt.Println(out)
}
