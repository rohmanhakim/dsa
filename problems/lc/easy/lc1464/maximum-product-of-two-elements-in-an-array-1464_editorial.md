# LeetCode's Official Editorial

## Solution

---

### Approach 1: Brute Force

**Intuition**

To start, we will simply check every pair of indices `(i, j)` and calculate `(nums[i] - 1) * (nums[j] - 1)`. We will take the maximum value as the answer.

Note that a pair of indices `(i, j)` will have the same result as `(j, i)`. Thus, to be more efficient, we will start iterating `j` from `i + 1`. This way, we don't check any duplicate pairs.

**Algorithm**

1. Initialize the answer `ans = 0`.
2. Iterate `i` over the indices of `nums`:
 - Iterate `j` over the indices of `nums`, starting from `i + 1`:
 - Calculate `(nums[i] - 1) * (nums[j] - 1)` and update `ans` if it is larger.
3. Return `ans`.

**Complexity Analysis**

Given n as the length of `nums`,

- Time complexity: O(n2)

 We have a nested for loop over the indices of `nums`. For `i = 0`, we will iterate `j` over n indices. For `i = 1`, we will iterate `j` over n−1 indices. For `i = 2`, we will iterate `j` over n−2 indices, and so on.

 In total, we iterate `j` over 1+2+3+...+n indices. This is the partial sum of [this series](https://leetcode.com/link/?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2F1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF%23Partial_sums), which is equal to 2n⋅(n+1)​\=O(n2).

- Space complexity: O(1)

 We aren't using any extra space.

---