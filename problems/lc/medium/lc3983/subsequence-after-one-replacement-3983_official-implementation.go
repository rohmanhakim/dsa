package lc3983

/*
func canMakeSubsequence(s string, t string) bool {
	n := len(s)
	j0, j1 := 0, 0
	for _, ch := range t {
		if s[j1] == byte(ch) {
			j1++
		}
		j1 = max(j1, j0+1)
		if s[j0] == byte(ch) {
			j0++
		}
		if j1 == n {
			return true
		}
	}
	return false
}
*/
