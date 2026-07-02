/*
Find all valid combinations of `k` numbers that sum up to `n` such that the following conditions are true:

- Only numbers `1` through `9` are used.
- Each number is used **at most once**.

Return *a list of all possible valid combinations*. The list must not contain the same combination twice, and the combinations may be returned in any order.

**Example 1:**

```
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
```

**Example 2:**

```
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
```

**Example 3:**

```
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
```

**Constraints:**

- `2 <= k <= 9`
- `1 <= n <= 60`
*/

package lc216

import (
	"fmt"
	"slices"
)

func combinationSum3(k int, n int) [][]int {

	sum := 0
	res := [][]int{}
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		if sum > n {
			return
		}

		if len(path) == k {
			if sum == n {
				res = append(res, slices.Clone(path))
			}
			return
		}

		for i := start; i < 9; i++ {
			path = append(path, i+1)
			sum = sum + (i + 1)
			dfs(i + 1)

			path = path[:len(path)-1]
			sum = sum - (i + 1)
		}
	}

	dfs(0)
	return res
}

func CombinationSumIII() {
	k := 3
	n := 9
	out := combinationSum3(k, n)
	fmt.Println(out)
}
