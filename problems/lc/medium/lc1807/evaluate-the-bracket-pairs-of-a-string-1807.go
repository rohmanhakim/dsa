package lc1807

import (
	"slices"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {

	stack := []rune{}

	kMap := make(map[string]string)
	for _, k := range knowledge {
		kMap[k[0]] = k[1]
	}

	var sb strings.Builder
	capture := true
	for _, ch := range s {
		if ch == ')' {
			keySlice := []rune{}
			for len(stack) > 0 {
				if stack[len(stack)-1] != '(' {
					keySlice = append(keySlice, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					break
				}
			}
			slices.Reverse(keySlice)
			key := string(keySlice)

			if _, ok := kMap[key]; ok {
				sb.WriteString(kMap[key])
			} else {
				sb.WriteString("?")
			}
			capture = true
		} else {
			stack = append(stack, ch)
			if ch == '(' {
				capture = false
			}
			if capture {
				sb.WriteRune(ch)
			}
		}
	}
	return sb.String()
}
