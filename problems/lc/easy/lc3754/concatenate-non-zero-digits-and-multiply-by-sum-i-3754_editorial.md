# LeetCode's Official Editorial

### Approach 1: Traverse from Left to Right

#### Intuition

Convert the input integer n into a string s, then traverse the digits from left to right while simulating the process described in the problem statement.

#### Complexity Analysis

- Time complexity: O(logn).
 
- Space complexity: O(logn).
 

---

### Approach 2: Traverse from Right to Left

#### Intuition

Instead of converting the integer into a string, we repeatedly divide n by 10 to traverse its digits from right to left. During the traversal, we update both x and sum.

Whenever we encounter a non-zero digit, we append it to x. Since the digits are processed from right to left, we maintain a variable pow10 that stores the current power of 10, allowing us to place each non-zero digit in its correct position within x.

#### Complexity Analysis

- Time complexity: O(logn).
 
- Space complexity: O(1).
