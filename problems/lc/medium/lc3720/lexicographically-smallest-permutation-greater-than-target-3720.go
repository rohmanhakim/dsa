package lc3720

func lexGreaterPermutation(s string, target string) string {
	freq := make([]int, 26)

	for _, ch := range s {
		freq[ch-'a']++
	}

	getLargest := func() []byte {
		ret := []byte{}
		for i := 25; i >= 0; i-- {
			for j := freq[i]; j > 0; j-- {
				ret = append(ret, byte(i+'a'))
			}
		}
		return ret
	}

	res := []rune{}

	for i, ch := range target {
		if freq[ch-'a'] > 0 {
			freq[ch-'a']--
			if string(getLargest()) > target[i+1:] {
				res = append(res, ch)
				continue
			} else {
				freq[ch-'a']++
			}
		}

		for c := ch + 1; c <= 'z'; c++ {
			if freq[c-'a'] > 0 {
				freq[c-'a']--
				res = append(res, c)

				for j := 0; j < 26; j++ {
					for freq[j] > 0 {
						freq[j]--
						res = append(res, rune(j+'a'))
					}
				}
				return string(res)
			}
		}

		return ""
	}

	return ""
}
