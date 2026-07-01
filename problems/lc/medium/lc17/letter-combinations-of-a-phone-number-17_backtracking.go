/*
Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order**.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

```
1        2(abc)  3(def)
4(ghi)   5(jkl)  6(mno)
7(pqrs)  8(tuv)  9(wxyz)
*        0       #
```

**Example 1:**

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**

```
Input: digits = "2"
Output: ["a","b","c"]
```

**Constraints:**

- `1 <= digits.length <= 4`
- `digits[i]` is a digit in the range `['2', '9']`.
*/

package lc17

func letterCombinationsBacktracking(digits string) []string {
	phone := []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	path := []byte{}
	indices := make([]int, len(digits))
	for i, ch := range digits {
		indices[i] = int(ch - '0')
	}

	var dfs func(pos int)
	dfs = func(pos int) {
		// digit++
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}

		chars := phone[indices[pos]-1]

		for i := 0; i < len(chars); i++ {
			path = append(path, chars[i])
			dfs(pos + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return res
}
