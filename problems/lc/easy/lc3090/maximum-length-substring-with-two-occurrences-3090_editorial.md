# LeetCode's Official Editorial

### Approach 1: Enumerate the Left Endpoint

#### Intuition

We can enumerate all substrings and check whether any character appears more than 2 times in each substring. Among the valid substrings, we find the one with the maximum length.

Specifically, we can enumerate the left endpoint left of the substring and continuously expand the right endpoint right. During the expansion, we use a count array to record the number of times each character appears in the current substring:

- If adding s\[right\] does not cause the count of that character to exceed 2, the current substring remains valid.

- If the count of a character becomes 3, the current substring is no longer valid.

Since further expanding the substring to the right cannot decrease the number of occurrences of this character, we can stop enumerating the current left endpoint.

#### Complexity Analysis

Let n be the length of s.

- Time complexity: O(n2).

- Space complexity: O(1).

---

### Approach 2: Sliding Window

#### Intuition

In Approach 1, we perform many redundant calculations for substrings with different left endpoints. For example, after processing the substring s\[left…right\], there is no need to recalculate the entire substring when moving the left endpoint to left+1. We only need to:

1. Remove the character at the left endpoint of the window.

2. Update the count of that character.

3. Continue expanding the window to the right.

Therefore, we can use a sliding window to maintain an interval \[left,right\] that always satisfies the condition that each character appears at most 2 times.

The right endpoint right moves continuously to the right, adding new characters to the window. If the newly added character appears more than twice in the window, we move the left endpoint left to the right until the count of that character becomes 2 again.

#### Complexity Analysis

Let n be the length of s.

- Time complexity: O(n).

 Each character is added to the sliding window at most once and removed from the sliding window at most once. Therefore, the overall time complexity is O(n).

- Space complexity: O(1).

 We use a fixed-size count array of size 26, so the extra space is O(1).

---
