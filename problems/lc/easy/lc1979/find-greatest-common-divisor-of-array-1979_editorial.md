# LeetCode's Official Editorial

### Approach: Calculate as Required

#### Intuition

First, traverse the array nums to find its maximum and minimum values. The answer is then the greatest common divisor of these two values.

To compute the greatest common divisor, both the C++ and Python standard libraries provide built-in functions for calculating the greatest common divisor of two integers.

**Euclidean Algorithm**

A common method for computing the greatest common divisor gcd(a,b) of two integers is the **Euclidean algorithm**, also known as the **division algorithm**. Its key recurrence is:

gcd(a,b)\=gcd(b,amodb).

#### Complexity Analysis

Let n be the length of the array nums, and let M be the maximum value in nums.

- Time complexity: O(n+logM).

 Finding the maximum and minimum values requires O(n) time, and computing their greatest common divisor requires O(logM) time.

- Space complexity: O(1).
