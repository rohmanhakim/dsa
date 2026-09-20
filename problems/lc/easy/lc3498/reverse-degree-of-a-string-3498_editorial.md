# LeetCode's Official Editorial

### Approach: Simulation

#### Intuition

According to the definition, the reverse degree is computed as follows:

- For each character c, multiply its position in the reversed alphabet by its 1-indexed position in the string.
- Add all of these products.

Traverse the string once. The reversed-alphabet position of a lowercase letter c is 26−(c−’a’). Multiplying this value by the character's 1-indexed position in the string and adding it to the running total produces the answer.

#### Complexity Analysis

Let n be the length of the string.

- Time complexity: O(n).

 We traverse the string once.

- Space complexity: O(1).

---
