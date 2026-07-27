package lc567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq1 := [26]int{}
	freq2 := [26]int{}

	for _, ch := range s1 {
		freq1[ch-'a']++
	}

	matches := 0
	for i := range 26 {
		if freq1[i] == freq2[i] {
			matches++
		}
	}

	for i := 0; i < len(s1); i++ {
		ch := s2[i] - 'a'
		if freq2[ch] == freq1[ch] {
			matches--
		}
		freq2[ch]++
		if freq2[ch] == freq1[ch] {
			matches++
		}
	}

	if matches == 26 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		rm := s2[i-len(s1)] - 'a'
		if freq1[rm] == freq2[rm] {
			matches--
		}
		freq2[rm]--
		if freq1[rm] == freq2[rm] {
			matches++
		}

		add := s2[i] - 'a'
		if freq1[add] == freq2[add] {
			matches--
		}
		freq2[add]++
		if freq1[add] == freq2[add] {
			matches++
		}

		if matches == 26 {
			return true
		}
	}

	return false
}
