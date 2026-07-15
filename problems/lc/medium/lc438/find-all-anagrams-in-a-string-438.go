package lc438

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	freqp := [26]int{}
	freqs := [26]int{}

	for i := 0; i < len(p); i++ {
		freqp[p[i]-'a']++
	}

	// matches = number of character positions where freqs[i] == freqp[i]
	// when matches == 26, we have an anagram (all frequencies match)
	matches := 0
	for i := 0; i < 26; i++ {
		if freqs[i] == freqp[i] {
			matches++
		}
	}

	res := []int{}

	// fill initial window [0, len(p))
	for i := 0; i < len(p); i++ {
		ch := s[i] - 'a'
		if freqs[ch] == freqp[ch] {
			matches-- // was matching, about to change
		}
		freqs[ch]++
		if freqs[ch] == freqp[ch] {
			matches++ // now matches again
		}
	}
	if matches == 26 {
		res = append(res, 0)
	}

	// slide window: add s[i], remove s[i-len(p)]
	for i := len(p); i < len(s); i++ {
		// remove leftmost character
		rm := s[i-len(p)] - 'a'
		if freqs[rm] == freqp[rm] {
			matches--
		}
		freqs[rm]--
		if freqs[rm] == freqp[rm] {
			matches++
		}

		// add new character on the right
		add := s[i] - 'a'
		if freqs[add] == freqp[add] {
			matches--
		}
		freqs[add]++
		if freqs[add] == freqp[add] {
			matches++
		}

		if matches == 26 {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}
