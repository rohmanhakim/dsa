package lc678

func checkValidString(s string) bool {

	low, high := 0, 0
	for _, ch := range s {

		switch ch {
		case ')':
			low = max(low-1, 0)
			high--
		case '(':
			low++
			high++
		default:
			low = max(low-1, 0)
			high++
		}
		if high < 0 {
			return false
		}
	}

	return low == 0
}
