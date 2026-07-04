package lc1318

/*
Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.



Example 1:

```
initial:
0010 -> a
0110 -> b
----
0101 -> c

after 3x flipping:
0001 -> a
0100 -> b
----
0101 -> c
```

Input: a = 2, b = 6, c = 5
Output: 3
Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)

Example 2:

Input: a = 4, b = 2, c = 7
Output: 1

Example 3:

Input: a = 1, b = 2, c = 3
Output: 0


Constraints:

1 <= a <= 10^9
1 <= b <= 10^9
1 <= c <= 10^9
*/

import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	count := 0
	for i := range 32 {
		x := (a >> i) & 1
		y := (b >> i) & 1
		z := (c >> i) & 1

		if x|y != z {
			if z == 1 {
				count++
			} else {
				count += x + y
			}
		}
	}

	return count
}

func MinimumFlipstoMakeaORbEqualtoc() {
	a := 8
	b := 3
	c := 5
	out := minFlips(a, b, c)
	fmt.Println(out)
}
