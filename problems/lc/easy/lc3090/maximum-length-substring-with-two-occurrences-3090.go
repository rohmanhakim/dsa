package lc3090

func maximumLengthSubstring(s string) int {
	n := len(s)
	freq := make([]int, 26)

	left := 0

	maxLen := 0
	for right := range n {
		freq[s[right]-'a']++

		for freq[s[right]-'a'] > 2 {
			freq[s[left]-'a']--
			left++
		}

		wLen := right - left + 1

		maxLen = max(wLen, maxLen)
	}
	return maxLen
}
