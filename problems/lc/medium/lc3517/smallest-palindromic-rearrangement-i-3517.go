package lc3517

import (
	"sort"
	"strings"
)

func smallestPalindrome(s string) string {
	mid := len(s) / 2
	sub := s[:mid]

	split := strings.Split(sub, "")

	sort.Strings(split)
	var sb strings.Builder

	for _, ch := range split {
		sb.WriteString(ch)
	}

	if len(s)%2 == 1 {
		sb.WriteByte(s[mid])
	}

	for i := len(split) - 1; i >= 0; i-- {
		sb.WriteString(split[i])
	}
	return sb.String()
}
