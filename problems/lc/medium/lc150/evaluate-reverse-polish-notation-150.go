package lc150

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	isOperator := func(ch string) bool {
		return ch == "+" || ch == "-" || ch == "*" || ch == "/"
	}

	stack := []int{}
	for _, t := range tokens {
		if isOperator(t) {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res := 0
			switch t {
			case "+":
				res = b + a
			case "-":
				res = b - a
			case "*":
				res = b * a
			case "/":
				res = b / a
			}
			stack = append(stack, res)
		} else {
			val, _ := strconv.Atoi(t)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

func LC150() {
	// tokens := []string{"2", "1", "+", "3", "*"}
	tokens := []string{"4", "13", "5", "/", "+"}
	out := evalRPN(tokens)
	fmt.Println(out)
}
