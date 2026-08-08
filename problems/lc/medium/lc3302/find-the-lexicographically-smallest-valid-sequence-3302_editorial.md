# LeetCode's Official Editorial

### Approach: Prefix and Suffix Decomposition + Greedy

#### Intuition

This approach combines prefix-suffix decomposition with a greedy strategy to construct the lexicographically smallest valid sequence.

Suffix preprocessing

Let last\[j\] denote the rightmost index in word1 where word2\[j\] can be matched while still allowing the suffix word2\[j…m−1\] to be matched with a suffix of word1. We traverse both word1 and word2 from right to left, greedily finding the latest possible matching position for each character of word2 and storing it in the last array.

This preprocessing allows us to determine, at any point during the construction of the answer, whether the remaining characters of word2 can still be matched using the remaining suffix of word1 without requiring any additional modifications.

Greedy construction

Next, we traverse word1 from left to right to construct the lexicographically smallest valid sequence of indices. Since earlier indices always produce a lexicographically smaller sequence, we greedily select the earliest possible index whenever it is safe to do so.

If word1\[i\] equals word2\[j\], we simply match the two characters and add index i to the answer.

Otherwise, if the characters do not match, we may still use the one allowed modification, provided it has not already been used. Before doing so, we check the last array to ensure that after treating the current position as a match, the remaining suffix of word2 can still be matched with the remaining suffix of word1.

Since using the modification at the earliest possible position yields a lexicographically smaller sequence, we greedily use the modification as soon as this condition is satisfied.

After the traversal, if all m characters of word2 have been matched, we return the constructed sequence. Otherwise, no valid sequence exists, and we return an empty array.


#### Complexity Analysis

Let n be the length of word1 and m be the length of word2.

- Time complexity: O(n+m).

- Space complexity: O(m).

---
