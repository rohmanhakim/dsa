package lc394

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []rune{}

	// - loop over input, check each character
	// - if character is not ] push it to the stack
	// - if character is ] begin loop & popping the stack
	//	 - put letter to separate stack
	//   - if stumbled upon [ while popping start capturing number
	//   - put the number to its own stack
	//   - if stumbled upon non-number or the stack only has one item:
	//   	- loop & pop number stack and costruct integer from it
	//		- loop & pop letter stack and construct substring from it
	//		- if stack has only one item left, pop it first
	//      - up to that integer, begin a loop:
	//        - append the substring to the main stack
	//   - break the loop of popping the main stack
	// - convert the stack into string

	for _, c := range s {
		if c == ']' {
			letter := []rune{}
			num := []rune{}
			captureNum := false
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if !captureNum {
					if top >= 'a' && top <= 'z' {
						letter = append(letter, top)
					} else if top == '[' {
						captureNum = true
					}
				} else {
					isNumber := top >= '0' && top <= '9'
					if isNumber {
						num = append(num, top)
					}
					if !isNumber || len(stack) == 1 {
						times := constructInteger(num)
						sub := constructSubstring(letter)
						if len(stack) == 1 {
							stack = stack[:len(stack)-1]
						}
						for range times {
							for _, l := range sub {
								stack = append(stack, l)
							}
						}
						break
					}
				}
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}

	return string(stack)
}

func constructSubstring(letter []rune) []rune {
	out := []rune{}
	for len(letter) > 0 {
		out = append(out, letter[len(letter)-1])
		letter = letter[:len(letter)-1]
	}
	return out
}

func constructInteger(num []rune) int {
	temp := []rune{}
	for len(num) > 0 {
		temp = append(temp, num[len(num)-1])
		num = num[:len(num)-1]
	}
	out, _ := strconv.Atoi(string(temp))
	return out
}

func decodeStringClaudeCode(s string) string {
	countStack := []int{}
	strStack := []string{}
	current := ""
	k := 0

	for _, c := range s {
		if c >= '0' && c <= '9' {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			countStack = append(countStack, k)
			strStack = append(strStack, current)
			current = ""
			k = 0
		} else if c == ']' {
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			current = prev + strings.Repeat(current, count)
		} else {
			current += string(c)
		}
	}

	return current
}

func DecodeString() {
	// s := "3[a]2[bc]" // aaabcbc
	s := "3[a2[bc]]" // abcbcabcbcabcbc
	out := decodeString(s)
	fmt.Println(out)
}
