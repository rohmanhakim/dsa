package lc151

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	ss := []string{}
	for w := range strings.SplitSeq(s, " ") {
		if len(strings.TrimSpace(w)) != 0 {
			ss = append(ss, w)
		}
	}
	slices.Reverse(ss)
	return strings.Join(ss, " ")
}

func reverseWordsFollowUp(s string) string {
	sb := []byte(s)

	slices.Reverse(sb)
	w := 0
	r := 0

	for r < len(sb) {
		for r < len(sb) && sb[r] == ' ' {
			r++
		}

		if r == len(sb) {
			break
		}

		if w > 0 {
			sb[w] = ' '
			w++
		}

		wordStart := w
		for r < len(sb) && sb[r] != ' ' {
			sb[w] = sb[r]
			w++
			r++
		}

		left := wordStart
		right := w - 1
		for left < right {
			sb[left], sb[right] = sb[right], sb[left]
			left++
			right--
		}
	}

	return string(sb[:w])
}
