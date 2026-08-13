# LeetCode's Official Editorial

### Approach 1: Segment Tree

#### Intuition

We can use a segment tree to maintain information about the string s. Each node of the segment tree stores the following information for its corresponding interval:

- pre: The length of the longest continuous sequence of identical characters at the beginning of the interval.

- suf: The length of the longest continuous sequence of identical characters at the end of the interval.

- maxLen: The length of the longest continuous sequence of identical characters within the interval.

- leftChar: The character at the leftmost position of the interval.

- rightChar: The character at the rightmost position of the interval.

For a leaf node, whose interval has length 1, we have pre\=suf\=maxLen\=1, and leftChar\=rightChar\=s\[i\].

For a non-leaf node, let its left and right child nodes be left and right, respectively. We can merge their information according to the following rules:

- leftChar is inherited from the leftChar of the left child.

- rightChar is inherited from the rightChar of the right child.

- pre initially takes the value of pre∗left. If the prefix of the left child covers the entire left interval, and rightChar∗left is equal to leftChar∗right, then the prefix can be extended into the right child. In this case, pre\=pre∗left+preright​.

- suf initially takes the value of suf∗right. If the suffix of the right child covers the entire right interval, and rightChar∗left is equal to leftChar∗right, then the suffix can be extended into the left child. In this case, suf\=suf∗right+sufleft​.

- maxLen initially takes the maximum of maxLen∗left and maxLen∗right. If rightChar∗left is equal to leftChar∗right, the suffix of the left child and the prefix of the right child can be concatenated. Therefore, we update maxLen with suf∗left+pre∗right.

For each query, we perform a **point update**, changing the character at the corresponding position to the new character. After each update, maxLen at the root of the segment tree represents the length of the longest substring consisting of a single repeated character in the current string.

The initial construction of the segment tree takes O(n) time, and each point update takes O(logn) time. Therefore, the total time complexity is O(n+klogn).

#### Complexity Analysis

Let n be the length of the string s, and let k be the number of queries.

- Time complexity: O((n+k)logn).

 Building the segment tree initially takes O(n) time. Each query performs a single-point update, which takes O(logn) time. Therefore, the total time complexity is O(n+klogn).

- Space complexity: O(n).

 The segment tree uses arrays of size O(n).

---

### Approach 2: Ordered Set Simulating Interval Merging

#### Intuition

We can use an ordered set, such as `set` in C++, `TreeMap` in Java, or `SortedList` in Python, to maintain all **maximal contiguous intervals** in the string s whose characters are identical. In addition, we use another ordered set to maintain the lengths of these intervals, allowing us to obtain the maximum interval length efficiently.

Initially, we traverse the string s and record every maximal contiguous interval, including its left and right endpoints. Each interval is represented by (L,R), indicating that all characters in the range \[L,R\] are identical.

For each query (pos,ch), we perform the following steps:

1. If s\[pos\] is already equal to ch, the string remains unchanged, so the current maximum interval length can be used directly.

2. Otherwise, find the interval \[L,R\] containing position pos. Remove this interval from the interval set and remove its length from the length set.

3. If L<pos, add the remaining left interval \[L,pos−1\] to both sets.

4. If pos<R, add the remaining right interval \[pos+1,R\] to both sets.

5. Position pos now becomes the character ch, initially forming the interval \[pos,pos\]. We then check the adjacent left and right intervals. If an adjacent interval exists and has the same character ch, we merge it with the new interval.

6. Finally, update s\[pos\] to ch and add the maximum interval length to the answer.

Because each update only affects the interval containing the modified position and its immediate neighbors, each query requires only a constant number of ordered-set operations.

#### Complexity Analysis

Let n be the length of the string s, and let k be the number of queries.

- Time complexity: O((n+k)logn).

 Constructing the initial intervals requires O(n) time. For each query, we perform a constant number of searches, deletions, and insertions in the ordered sets, each taking O(logn) time. Therefore, the total time complexity is O(n+klogn).

- Space complexity: O(n).

 There can be at most n intervals, so the ordered sets require O(n) space.
