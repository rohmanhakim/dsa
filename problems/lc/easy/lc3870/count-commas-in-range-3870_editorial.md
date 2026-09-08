# LeetCode's Official Editorial

### Approach 1: Traversing Count

#### Intuition

This approach counts the number of qualifying numbers by directly traversing all integers from 1 to n. According to the problem statement, only numbers greater than 999 will contain a comma in the standard format within the range \[1,n\] (n≤105). Therefore, each time the loop variable is greater than 999, the result is incremented by 1, and the total count is returned at the end.

#### Complexity Analysis

- Time complexity: O(n).

- Space complexity: O(1).

### Approach 2: Direct Calculation

#### Intuition

This approach calculates the answer in O(1) time through mathematical computation.

According to the problem constraints, within the range \[1,n\] (n≤105), only numbers greater than or equal to 1000 will contain exactly one comma (e.g., 1,000 to 100,000).

Therefore, if n≥1000, the count of numbers containing a comma is n−999; if n<1000, the count is 0. We can directly take the maximum value between n−999 and 0 as the final result, avoiding the need for any loops.

#### Complexity Analysis

- Time complexity: O(1).

- Space complexity: O(1).

---
