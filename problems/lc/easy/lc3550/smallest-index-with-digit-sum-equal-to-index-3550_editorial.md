# LeetCode's Official Editorial

### Approach: Traversal

#### Intuition

We need to find the smallest index i in nums such that the sum of the digits of nums\[i\] is equal to i.

We can iterate through nums from left to right, calculate the digit sum of nums\[i\], and check whether it equals i. Because we traverse indices in increasing order, the first index that satisfies this condition is guaranteed to be the smallest, so we can return it immediately. If the traversal finishes without finding any matching index, we return −1.

#### Complexity Analysis

Let n be the length of nums and m be the largest number in nums.

- Time complexity: O(nlogm).

 Calculating the digit sum by repeatedly removing the last digit takes O(logm) time.

- Space complexity: O(1).

---
