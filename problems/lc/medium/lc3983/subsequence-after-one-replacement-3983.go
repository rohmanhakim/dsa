package lc3983

func canMakeSubsequence(s string, t string) bool {

	ns := len(s)
	nt := len(t)

	if ns > nt {
		return false
	}

	prefix := make([]int, ns+1)
	for i := range prefix {
		prefix[i] = -1
	}
	prefix[0] = -1
	for i := 1; i <= ns; i++ {
		prefix[i] = prefix[i-1]
		found := false
		for j := prefix[i-1] + 1; j < nt; j++ {
			if s[i-1] == t[j] {
				prefix[i] = j
				found = true
				break
			}
		}
		if !found {
			prefix[i] = -1
			break
		}
	}

	if prefix[ns] != -1 {
		return true
	}

	// Suffix: rightmost match for s[i..n-1] in t (greedy right-to-left)
	suffix := make([]int, ns+1)
	for i := range suffix {
		suffix[i] = -1
	}
	suffix[ns] = nt // sentinel: empty suffix starts past the end
	for i := ns - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1]
		found := false
		for j := suffix[i+1] - 1; j >= 0; j-- {
			if s[i] == t[j] {
				suffix[i] = j
				found = true
				break
			}
		}
		if !found {
			suffix[i] = -1
			break
		}
	}

	// For each replacement position i, check if there's room between
	// the prefix end and the suffix start.
	for i := 0; i < ns; i++ {
		if i > 0 && prefix[i] < 0 {
			continue
		}
		if prefix[i]+1 < suffix[i+1] {
			return true
		}
	}
	return false
}
