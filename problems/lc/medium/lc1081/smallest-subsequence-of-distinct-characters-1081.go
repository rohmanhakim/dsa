package lc1081

func smallestSubsequence(s string) string {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	alphabet := [26]bool{}
	stack := []rune{}
	for _, ch := range s {
		if alphabet[ch-'a'] == false {
			for len(stack) > 0 && stack[len(stack)-1] > ch {
				top := stack[len(stack)-1]
				if freq[top] == 0 {
					break
				}
				alphabet[top-'a'] = false
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, ch)
			alphabet[ch-'a'] = true
		}
		freq[ch]--
	}
	return string(stack)
}
