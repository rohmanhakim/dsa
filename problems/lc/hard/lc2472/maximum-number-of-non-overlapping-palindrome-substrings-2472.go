package lc2472

func maxPalindromes(s string, k int) int {
	n := len(s)

	valid := func(sub string) bool {
		r := []rune(sub)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			if r[i] != r[j] {
				return false
			}
		}
		return true
	}

	start := 0
	count := 0
	for i := 0; i < n; i++ {
		if i+1 >= k && valid(s[i-k+1:i+1]) && i-k+1 >= start {
			start = i + 1
			count++
		} else if i >= k && valid(s[i-k:i+1]) && i-k >= start {
			start = i + 1
			count++
		}
	}
	return count
}
