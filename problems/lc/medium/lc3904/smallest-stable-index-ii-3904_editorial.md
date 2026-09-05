# LeetCode's Official Editorial

### Approach: Prefix Maximum + Suffix Minimum

#### Intuition

This problem is a follow-up to ["3903. Smallest Stable Index I"](https://leetcode.com/link/?target=https%3A%2F%2Fleetcode.com%2Fproblems%2Fsmallest-stable-index-i%2F), requiring an O(n) algorithm, where n is the length of the array nums.

We preprocess the array minValue, where minValue\[i\] represents the minimum value in nums\[i..n−1\]. Traverse from the end to the beginning:

- minValue\[n−1\]\=nums\[n−1\];

- For the case i\=n−1, minValue\[i\]\=min(minValue\[i+1\],nums\[i\]).

Then, traverse from the beginning to the end, maintaining a variable maxValue representing the maximum value in nums\[0..i\]. For each index i, the instability score is maxValue−minValue\[i\]. If this score is less than or equal to k, then i is a stable index. Since we are iterating from left to right, the first stable index we find will be the smallest one, so we can return it immediately. If no stable index is found, we return −1.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n).

 The minimum suffix preprocessing and the forward traversal each require O(n) time.

- Space complexity: O(n).

 The space required for the array minValue.
