package lc395

func longestSubstring(s string, k int) int {
	n := len(s)

	freq := make(map[rune]int)

	maxUnique := 0

	longest := 0

	for _, ch := range s {
		freq[ch]++
	}

	maxUnique = len(freq)

	for currUnique := 1; currUnique <= maxUnique; currUnique++ {
		freq = make(map[rune]int)

		countAtLeastK := 0

		left := 0
		right := 0

		for right < n && left < n {
			ch := rune(s[right])
			f := freq[ch]

			if f+1 == k {
				countAtLeastK++
			}

			freq[ch]++

			for len(freq) > currUnique {
				lch := rune(s[left])
				f := freq[lch]
				if f == k {
					countAtLeastK--
				}
				freq[lch]--

				if freq[rune(s[left])] == 0 {
					delete(freq, rune(s[left]))
				}

				left++
			}

			windowSize := right - left + 1
			if windowSize > longest && countAtLeastK == currUnique {
				longest = windowSize
			}

			right++
		}
	}

	return longest
}
