package lc71

import (
	"strings"
)

func simplifyPath(path string) string {

	pathItems := strings.Split(path, "/")

	stack := []string{}
	var sb strings.Builder
	for _, pi := range pathItems {
		switch pi {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		case "":
		default:
			stack = append(stack, pi)
		}
	}

	if len(stack) == 0 {
		return "/"
	}
	for _, item := range stack {
		sb.WriteString("/")
		sb.WriteString(item)
	}

	return sb.String()
}
