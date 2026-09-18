# LeetCode's Official Editorial

### Approach: Greedy

#### Intuition

The problem requires that "if a substring contains a character c, then every occurrence of c in s must be inside that substring," and we must minimize the final total length. Because non-overlapping valid substrings cannot share characters, there can be at most Σ substrings, where Σ is the size of the character set. Suppose we have found the shortest valid substring s\[la​,ra​\] containing character a. It may seem that s\[la​−1,ra​\] or s\[la​,ra​+1\] could also be valid, but because the total length must be minimized, we only need to consider the shortest substring containing each character.

The first step is therefore to preprocess the shortest substring corresponding to every character in the character set. Because the character set is small, we can process this part directly. First, traverse the string to determine the first occurrence li​ and the last occurrence ri​ of each character i. Because other characters may appear inside \[li​,ri​\], to satisfy the second condition of the problem, traverse every character in \[li​,ri​\] and expand li​ and ri​ to include their first and last occurrences until the boundaries stabilize. This guarantees that if a substring contains a character c, every occurrence of c in s is inside that substring.

After preprocessing, regard the starting and ending positions of each substring as an interval \[li​,ri​\]. The problem is transformed into the following: **on a one-dimensional number line \[0,n−1\], where n is the length of s, select as many mutually disjoint intervals as possible while minimizing their total length**. This is a classic greedy interval scheduling problem. Sort the intervals in ascending order by their right endpoint (primary key) and ascending order by length (secondary key, i.e., larger left endpoint first). Traverse them from left to right, greedily adding an interval whenever its left endpoint is strictly greater than the end position of the previously chosen interval. To see why this is correct, consider two overlapping intervals \[l1​,r1​\] and \[l2​,r2​\] with r1​<r2​. Selecting \[l1​,r1​\] leaves more of the number line available for subsequent intervals. Furthermore, if two intervals share the same right endpoint, choosing the shorter one leaves more space to the left and directly minimizes the total length.


#### Complexity Analysis

Let n be the length of the string, and let Σ be the size of its character set.

- Time complexity: O(nΣ+ΣlogΣ).

 Preprocessing the left and right endpoints of every interval takes O(nΣ) time. Greedy selection takes O(ΣlogΣ+Σ) time. Therefore, the total time complexity is O(nΣ+ΣlogΣ).

- Space complexity: O(Σ).

 We need O(Σ) space to record the left and right endpoints of the interval containing each character.

---
