# LeetCode's Official Editorial

### Approach: Enumeration

#### Intuition

We can enumerate three distinct indices i, j, and k, and use their corresponding digits as the hundreds, tens, and ones digits, respectively.

A valid three-digit even number must satisfy the following conditions:

- Its hundreds digit cannot be 0.
- Its ones digit must be even.
- Its three digits must come from distinct indices.

Because the array can contain duplicate digits, different index triples may form the same three-digit number. We can use a boolean array vis of length 1000 to record whether each three-digit number has already appeared. Whenever a valid, unseen number is formed, we mark it as visited and increment the answer.

#### Complexity Analysis

Let n be the length of the digits array.

- Time complexity: O(n3).

 We use three nested loops to enumerate the three indices i, j, and k, which takes O(n3) time.

- Space complexity: O(1).

 The boolean array vis has a fixed length of 1000, independent of the input size.

---
