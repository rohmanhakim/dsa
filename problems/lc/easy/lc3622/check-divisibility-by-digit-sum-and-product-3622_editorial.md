# LeetCode's Official Editorial

### Approach: Simulation

#### Intuition

According to the problem statement, we can extract each digit of n from the least significant digit to the most significant digit and calculate the digit sum digitSum and digit product digitProduct.

Finally, we can determine whether n is divisible by the sum of its digit sum and digit product by checking whether

nmod(digitSum+digitProduct)\=0.

Note that extracting the digits modifies the value of n. Therefore, we need to save the original value of n in advance.

#### Complexity Analysis

Let n be the given number.

- Time complexity: O(logn).

 The number of digits in n is O(logn), and we process each digit exactly once.

- Space complexity: O(1).
