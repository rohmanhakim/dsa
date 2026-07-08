# LeetCode's Official Editorial

### Approach: Prefix Array

#### Intuition

We define `pow10[i]` as (10^i \\bmod \\textit{MOD}). By precomputing the `pow10` array once, we avoid recomputing these values for every test case.

As in the string-based solution for **3754\. Concatenate Non-Zero Digits and Multiply by Sum I**, we build prefix arrays to obtain the values of `x` and `sum` for any prefix of the input string.

We also maintain a third prefix array, `cnt`, where `cnt[i]` denotes the number of non-zero digits in the prefix ending before index `i`.

Using these three prefix arrays, we can compute the values of `x` and `sum` for any query range by taking differences between the corresponding prefix values.

Finally, we compute the answer for each query and return the resulting array.

#### Complexity Analysis

Let `MAX_N` be the maximum possible string length, n be the length of the input string, and m be the number of queries.

Preprocessing:

- Time complexity: O(MAX_N).
 
- Space complexity: O(MAX_N).
 

For each test case:

- Time complexity: O(n+m).
 
- Space complexity: O(n).
 
 The returned array is not included in the space complexity.
 
