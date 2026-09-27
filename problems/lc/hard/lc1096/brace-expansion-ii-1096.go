package lc1096

import (
	"slices"
	"strings"
)

func braceExpansionII(expression string) []string {

	var expr func(s string) []string
	var term func(s string) []string
	var item func(s string) []string

	var cartesian func(idx int, input [][]string, path []string, res *[]string)
	cartesian = func(idx int, input [][]string, path []string, res *[]string) {
		if len(path) == len(input) {
			var concatenated strings.Builder
			for _, s := range path {
				concatenated.WriteString(s)
			}
			*res = append(*res, concatenated.String())
			return
		}
		for i := 0; idx < len(input) && i < len(input[idx]); i++ {
			path = append(path, input[idx][i])
			cartesian(idx+1, input, path, res)
			path = path[:len(path)-1]
		}
	}

	expr = func(s string) []string {
		inputs := []string{}
		level := 0
		lastPos := -1
		for i, ch := range s {
			if ch == '{' {
				level++
			}
			if ch == '}' {
				level--
			}
			if ch == ',' {
				if level == 0 {
					inputs = append(inputs, s[lastPos+1:i])
					lastPos = i
				}
			}
			if i == len(s)-1 {
				if level == 0 {
					inputs = append(inputs, s[lastPos+1:])
				}
			}
		}
		outputs := []string{}
		for _, inp := range inputs {
			out := term(inp)
			outputs = append(outputs, out...)
		}
		slices.Sort(outputs)
		return slices.Compact(outputs)
	}

	term = func(s string) []string {
		inputs := []string{}
		level := 0
		lastPos := 0
		for i, ch := range s {
			switch ch {
			case '{':
				level++
			case '}':
				level--
				if level == 0 {
					inputs = append(inputs, s[lastPos:i+1])
					lastPos = i + 1
				}
			default:
				if level == 0 {
					inputs = append(inputs, s[lastPos:i+1])
					lastPos = i + 1
				}
			}
		}
		pool := [][]string{}
		for _, inp := range inputs {
			out := item(inp)
			pool = append(pool, out)
		}
		path := []string{}
		outputs := []string{}
		cartesian(0, pool, path, &outputs)
		return outputs
	}

	item = func(s string) []string {
		outputs := []string{}
		if len(s) > 1 && s[0] == '{' && s[len(s)-1] == '}' {
			outputs = expr(s[1 : len(s)-1])
		} else {
			if len(s) > 1 {
				outputs = expr(s)
			} else {
				outputs = append(outputs, s)
			}
		}
		return outputs
	}

	out := expr(expression)

	return out
}
