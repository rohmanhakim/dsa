# LeetCode's Official Editorial

### Approach: Hash Table

#### Intuition

Store the pairs from knowledge in a hash table dict for O(1) average-time lookup. As we traverse s, each character falls into one of four cases: an opening parenthesis, a closing parenthesis, a lowercase letter outside parentheses, or a lowercase letter inside parentheses.

Use key to store the key within the current pair of parentheses, and let a boolean flag addKey indicate whether we are currently inside parentheses. While traversing s:

- When the current character c is an opening parenthesis, set addKey\=true.
- When c is a closing parenthesis, key is complete. Append dict\[key\] to the result if the key exists; otherwise append `'?'`. Then clear key and set addKey\=false.
- When c is a lowercase letter, append it to key if addKey is true; otherwise append it directly to the result.

> **Note:** Alternatively, instead of collecting characters into a separate buffer, we can record the index of `'('` using a variable start (with start\=−1 indicating we are outside parentheses) and extract the key directly as a substring slice when reaching `')'`. This variation is demonstrated in the Python and Go implementations below.

#### Complexity Analysis

Let n be the length of s, and let m be the total length of all strings in knowledge.

- Time complexity: O(n+m).

 Building the hash table takes O(m) time. Traversing s, extracting keys, looking them up in the hash table, and appending characters and values to the result string takes O(n+m) time. Thus, the total time complexity is O(n+m).

- Space complexity: O(n+m).

 The hash table stores all key-value pairs from knowledge, which requires O(m) space. The output string (and key buffer) requires O(n+m) space.

---
