package lc392

func isSubsequence(s string, t string) bool {
	left := 0
	right := 0
	ns := len(s)
	nt := len(t)

	for left < ns && right < nt {
		if s[left] == t[right] {
			left++
		}
		right++
	}

	return left == ns
}
