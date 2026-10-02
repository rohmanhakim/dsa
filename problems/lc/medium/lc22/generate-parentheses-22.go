package lc22

func generateParenthesis(n int) []string {

	var explore func(cur string, closeCount, openCount int)

	res := []string{}
	explore = func(cur string, closeCount, openCount int) {
		if len(cur) == 2*n {
			res = append(res, cur)
			return
		}

		if openCount < n {
			explore(cur+"(", closeCount, openCount+1)
		}

		if closeCount < openCount {
			explore(cur+")", closeCount+1, openCount)
		}

	}

	explore("", 0, 0)
	return res
}
