# LeetCode's Official Editorial

### Approach: Enumeration

#### Intuition

We simply enumerate integers starting from n. Within at most 10 attempts, we are guaranteed to encounter an integer whose last digit is 0. Since the product of its digits is then 0, it is divisible by any positive integer t. Therefore, a valid answer is guaranteed to be found within the next 10 integers.

#### Complexity Analysis

Let n be the integer given in the problem.

- Time complexity: O(10logn).

- Space complexity: O(1).

---
