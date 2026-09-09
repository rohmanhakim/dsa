# LeetCode's Official Editorial

### Approach: Place Value Contribution

#### Intuition

In standard number formatting, a comma is inserted every three digits from the right (for numbers with at least four digits). Instead of counting the number of commas in each integer individually, we can look at it from another angle and calculate the "contribution" of each comma to the final answer:

- Numbers greater than or equal to 103 will contain at least 1 comma, contributing a total of n−103+1 commas.

- Numbers greater than or equal to 106 will contain at least 2 commas (i.e., an additional 1 comma per number from the previous step), contributing an additional n−106+1.

- In general, each power of 1000 (p) up to n contributes an additional n−p+1 commas. We can start at p\=1000, add n−p+1 while p≤n, and multiply p by 1000 in each iteration.

#### Complexity Analysis

- Time complexity: O(logn).

 The loop runs O(log1000​n) times because `p` is multiplied by 1000 in each iteration (at most 5 iterations for n≤1015).

- Space complexity: O(1).