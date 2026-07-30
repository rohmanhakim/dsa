# LeetCode's Official Editorial

### Approach 1: Greedy

#### Intuition

There are 8 keys, and the k\-th letter assigned to a key requires exactly k key presses. To minimize the total number of key presses, we should assign letters to positions requiring fewer presses first. In other words, we fill all positions requiring one key press, then all positions requiring two key presses, and so on.

Suppose the n letters in the string word are numbered from 0 to n−1. Then the i\-th letter is assigned to a position requiring

⌊8i​⌋+1

key presses. Therefore, the answer is simply the sum of these values over all letters.

#### Complexity Analysis

Let n be the length of the string word.

- Time complexity: O(n).

- Space complexity: O(1).

---

### Approach 2: Mathematics

#### Intuition

This approach is based on the same greedy strategy as Approach 1. Instead of summing the number of key presses one by one, we can derive the answer directly using a mathematical formula.

Let the length of the string word be n, and let

m\=⌊8n−1​⌋+1

denote the maximum number of key presses required for any letter.

There are m−1 complete groups, each containing 8 letters. The letters in the i\-th group require exactly i key presses, where 1≤i≤m−1. Therefore, these groups contribute

8⋅2m(m−1)​\=4m(m−1)

key presses in total.

The remaining n−8(m−1) letters each require exactly m key presses, contributing

(n−8(m−1))⋅m

key presses.

Hence, the total number of key presses is

4m(m−1)+(n−8(m−1))⋅m.

#### Complexity Analysis

- Time complexity: O(1).

- Space complexity: O(1).
