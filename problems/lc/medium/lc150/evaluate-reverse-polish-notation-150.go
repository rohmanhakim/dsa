package lc150

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []string{}
	for _, n := range tokens {
		if n == "+" || n == "-" || n == "*" || n == "/" {
			a, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			b, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			c := 0
			switch n {
			case "+":
				c = b + a
			case "-":
				c = b - a
			case "*":
				c = b * a
			case "/":
				c = b / a
			}
			stack = append(stack, strconv.Itoa(c))
		} else {
			stack = append(stack, n)
		}
	}

	res, _ := strconv.Atoi(stack[0])
	return res
}

func evalRPNClaudeSonnet(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, b+a)
			case "-":
				stack = append(stack, b-a)
			case "*":
				stack = append(stack, b*a)
			case "/":
				stack = append(stack, b/a)
			}
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func EvaluateReversePolishNotation() {
	// tokens := []string{"2", "1", "+", "3", "*"} // 9
	// tokens := []string{"4", "13", "5", "/", "+"} // 6
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"} // 22
	out := evalRPN(tokens)
	fmt.Println(out)
}
