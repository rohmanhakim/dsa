# LeetCode's Official Editorial

### Approach: Enumeration

#### Intuition

Let n be the length of the array nums. Since the data range is small, we can directly enumerate each index and calculate its instability score.

Enumerate indices from 0 to n−1 in order. For each i, traverse nums\[0..i\] to find the maximum value, and traverse nums\[i..n−1\] to find the minimum value. The difference between these two values is the instability score of the index. If this score is less than or equal to k, then i is a stable index. Since we are iterating from left to right, the first stable index we find will be the smallest one, so we can return it immediately.

If no index satisfies the condition, we return −1.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n2).

 When enumerating each index, it takes O(n) time to compute the prefix maximum and suffix minimum values.

- Space complexity: O(1).

---
