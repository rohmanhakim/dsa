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

import (
	"fmt"
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	pool := [][]rune{}
	curr := 'a'
	for i := range 9 {
		pool = append(pool, []rune{})

		if i == 0 {
			continue
		}
		for range 3 {
			pool[i] = append(pool[i], curr)
			curr++
		}
		if i == 6 || i == 8 {
			pool[i] = append(pool[i], curr)
			curr++
		}
	}

	prevDigit, _ := strconv.Atoi(string(digits[0]))

	prevChars := pool[prevDigit-1]
	result := []string{}
	for _, ch := range prevChars {
		result = append(result, string(ch))
	}

	for i := 1; i < len(digits); i++ {
		currDigit, _ := strconv.Atoi(string(digits[i]))
		currChars := pool[currDigit-1]
		product := []string{}
		n := len(result)
		for x := 0; x < n; x++ {
			for y := 0; y < len(currChars); y++ {
				var sb strings.Builder

				sb.WriteString(result[x])
				sb.WriteRune(currChars[y])
				product = append(product, sb.String())
			}
		}
		result = product
	}
	return result
}

func LetterCombinationsOfAPhoneNumber() {
	digits := "2"
	out := letterCombinationsBacktracking(digits)
	fmt.Println(out)
}
