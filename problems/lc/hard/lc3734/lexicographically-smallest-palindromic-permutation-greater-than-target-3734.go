package lc3734

import (
	"slices"
)

func lexPalindromicPermutation(s string, target string) string {
	n := len(s)

	freq := make([]int, 26)
	leftHalfFreq := make([]int, 26)
	hasMid := false
	mid := byte(0)
	countOdd := 0

	if n%2 != 0 {
		hasMid = true
	}
	for _, ch := range s {
		freq[ch-'a']++
		leftHalfFreq[ch-'a'] = freq[ch-'a'] / 2
	}

	for i, v := range freq {
		if v%2 != 0 {
			mid = byte(i + 'a')
			countOdd++
		}
	}

	if countOdd > 1 {
		return ""
	}

	largestPalindrome := func(prefix string) []byte {
		freqCopy := slices.Clone(freq)
		res := make([]byte, len(s))

		for i, ch := range prefix {
			res[i] = byte(ch)
			res[len(s)-1-i] = byte(ch)
			freqCopy[ch-'a']--
		}

		w := len(prefix)

		for i := 25; i >= 0; i-- {
			for freqCopy[i] > 1 {
				res[w] = byte(i + 'a')
				res[len(s)-1-w] = byte(i + 'a')
				freqCopy[i] -= 2
				w++
			}
		}

		for i := 25; i >= 0; i-- {
			if freqCopy[i] > 0 {
				res[w] = byte(i + 'a')
				freqCopy[i]--
				w++
			}
		}
		return res
	}

	res := []byte{}
	for i := 0; i < n/2; i++ {
		found := false
		for c := 'a'; c <= 'z'; c++ {
			if leftHalfFreq[c-'a'] > 0 {
				res = append(res, byte(c))
				freq[c-'a']--
				lp := largestPalindrome(string(res))
				if string(lp) > target {
					leftHalfFreq[c-'a']--
					found = true
					break
				} else {
					freq[c-'a']++
					res = res[:len(res)-1]
				}
			}
		}
		if !found {
			return ""
		}
	}

	revRes := slices.Clone(res)
	slices.Reverse(revRes)
	resstring := string(res)
	if hasMid {
		resstring += string(mid)
	}
	resstring += string(revRes)

	if resstring <= target {
		return ""
	}

	return resstring
}
