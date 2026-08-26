# LeetCode's Official Editorial

### Approach: Enumeration + Hash Table

#### Intuition

We need to find the smallest positive integer that is a multiple of k and does not appear in the array nums.

First, we put all elements of nums into a hash set seen to support O(1) average-time lookups. Then, starting from k, we enumerate its positive multiples k,2k,3k,… until we find the first multiple that does not appear in seen. This value is the answer.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n).

 Constructing the hash set takes O(n) time. We may need to enumerate at most n+1 multiples of k, and each lookup takes O(1) average time.

- Space complexity: O(n).

 The hash set stores at most n distinct elements.
