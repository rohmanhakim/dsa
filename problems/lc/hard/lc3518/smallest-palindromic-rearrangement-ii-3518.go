package lc3518

import (
	"sort"
	"strings"

	"github.com/rohmanhakim/leetcode/lib"
)

func smallestPalindrome(s string, k int) string {

	n := len(s)
	mid := n / 2
	sub := strings.Split(s[:mid], "")
	subLen := len(sub)
	res := []string{}

	freq := make(map[string]int)
	for _, ch := range sub {
		freq[ch]++
	}

	sorted := []string{}

	for key := range freq {
		sorted = append(sorted, key)
	}

	sort.Strings(sorted)

	countPermutations := func(remaining int) int {
		result := 1
		pos := remaining
		for _, c := range sorted {
			if freq[c] == 0 {
				continue
			}
			result *= lib.Comb(pos, freq[c], k)
			if result >= k {
				return k
			}
			pos -= freq[c]
		}
		return result
	}

	if k > countPermutations(subLen) || k < 1 {
		return ""
	}

	var sb strings.Builder

	for i := range sub {
		remLen := subLen - 1 - i

		for _, ch := range sorted {
			if freq[ch] > 0 {
				freq[ch]--

				perms := countPermutations(remLen)

				if k <= perms {
					res = append(res, ch)
					break
				} else {
					k -= perms
					freq[ch] += 1
				}
			}
		}
	}

	for _, ch := range res {
		sb.WriteString(ch)
	}

	if n%2 == 1 {
		sb.WriteByte(s[mid])
	}

	for i := len(res) - 1; i >= 0; i-- {
		sb.WriteString(res[i])
	}

	return sb.String()
}
