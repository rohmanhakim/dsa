# LeetCode's Official Editorial

### Approach: Bitwise Comparison

#### Intuition

We iterate through each digit of the given number n, maintaining the largest digit in first and the second largest digit in second. After processing all digits, the product of these two values is the maximum possible product, which is the answer.

#### Complexity Analysis

Let n be the given number.

- Time complexity: O(logn).

- Space complexity: O(1).

