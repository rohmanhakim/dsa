package lc17

func letterCombinations(digits string) []string {
	chars := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	path := []rune{}
	out := []string{}

	var dfs func(d int)
	dfs = func(d int) {

		if len(path) == len(digits) {
			out = append(out, string(path))
			return
		}

		char := chars[digits[d]-'0']
		for _, ch := range char {
			path = append(path, ch)
			dfs(d + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return out
}
