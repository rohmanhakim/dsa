package lc49

func groupAnagrams(strs []string) [][]string {
	freqs := make(map[string][]string)

	for _, str := range strs {
		freq := make([]byte, 26)
		for _, ch := range str {
			freq[ch-'a']++
		}
		freqs[string(freq)] = append(freqs[string(freq)], str)
	}

	res := [][]string{}

	for _, v := range freqs {
		res = append(res, v)
	}

	return res
}
