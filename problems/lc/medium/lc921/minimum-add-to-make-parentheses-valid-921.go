package lc921

func minAddToMakeValid(s string) int {

	unclosedLeft := 0
	unclosedRight := 0

	for _, ch := range s {
		if ch == '(' {
			unclosedLeft++
		} else {
			if unclosedLeft <= 0 {
				unclosedRight++
			} else {
				unclosedLeft--
			}
		}
	}

	return unclosedLeft + unclosedRight
}
