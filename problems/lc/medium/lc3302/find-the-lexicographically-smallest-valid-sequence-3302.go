package lc3302

func validSequence(word1 string, word2 string) []int {
	n1 := len(word1)
	n2 := len(word2)
	suffix := make([]int, n1+1)
	j := n2 - 1
	for i := n1 - 1; i >= 0; i-- {
		if j >= 0 && word1[i] == word2[j] {
			j--
			suffix[i] = suffix[i+1] + 1
		} else {
			suffix[i] = suffix[i+1]
		}
	}
	j = 0
	usedChance := false
	res := []int{}
	for i := 0; i < n1 && j < n2; i++ {
		if word1[i] == word2[j] {
			res = append(res, i)
			j++
		} else if !usedChance && suffix[i+1] >= n2-j-1 {
			res = append(res, i)
			usedChance = true
			j++
		}
	}
	if len(res) < n2 {
		return []int{}
	}
	return res
}
