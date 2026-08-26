# LeetCode's Official Editorial

### Approach 1: Enumeration

#### Intuition

The task is to find the shortest and lexicographically smallest substring of the binary string s that contains exactly k ones.

Let n be the length of s. Since the length of the given string is relatively small, with n≤102, we can solve this problem using an algorithm with a time complexity of O(n3).

Suppose the length of the shortest valid substring is m. We can enumerate all substrings of length m in s, check whether each substring contains exactly k ones, and keep track of the lexicographically smallest valid substring. The possible values of m range from k to n.

Since we enumerate the lengths in increasing order, the first length for which a valid substring exists is the minimum possible length.

#### Implementation

#### Complexity Analysis

Let n be the length of the string s.

- Time complexity: O(n3).

 There are O(n) possible substring lengths, and for each length, we enumerate O(n) substrings. Checking the number of ones and extracting each substring both take O(n) time, resulting in a total time complexity of O(n3).

- Space complexity: O(n) or O(1).

 At any time, we store a substring of length at most n and the current answer, both of which require O(n) space.

---

### Approach 2: Sliding Window

#### Intuition

We can maintain a sliding window containing exactly k ones. As we expand the window from right to left, we shrink it whenever the number of ones exceeds k or when the leftmost character is 0. This allows us to find the shortest valid substring ending at each position.

More specifically, after adding `s[right]` to the window, we repeatedly move `left` forward while either the window contains more than k ones or the leftmost character is `0`. As a result, whenever the window contains exactly k ones, it is the shortest valid substring ending at `right`.

We then compare this substring with the current answer. A substring is better if it has a shorter length, or if the lengths are equal and it is lexicographically smaller.

#### Complexity Analysis

Let n be the length of the string s.

- Time complexity: O(n2).

 The sliding window itself takes O(n) time, since both `left` and `right` move from left to right at most once. However, extracting a substring takes O(n) time in the worst case, and this operation can be performed O(n) times. Therefore, the total time complexity is O(n2).

- Space complexity: O(n) or O(1).

 The current substring and the answer can each require O(n) space.
