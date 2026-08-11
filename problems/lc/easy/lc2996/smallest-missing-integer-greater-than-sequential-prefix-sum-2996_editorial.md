# LeetCode's Official Editorial

### Approach 1: Simulation + Prefix Sum

#### Intuition

Starting from 0, we traverse nums in order to find the longest ordered prefix according to the definition in the problem, while calculating its sum at the same time. We stop when we encounter the first element that does not satisfy the definition of an ordered prefix.

Then, starting from this prefix sum, we incrementally check whether the current value exists in nums. The first value that does not appear in nums is the answer.

To efficiently determine whether a value exists in nums, we use the elements of nums to construct a hash set.

#### Complexity Analysis

Let n be the length of nums.

- Time complexity: O(n).

 Finding the longest ordered prefix takes O(n) time, and finding the first value that does not appear in nums takes O(n) time in the worst case.

- Space complexity: O(n).

 The hash set requires O(n) auxiliary space.

---

### Approach 2: Simulation + Series Sum Formula

#### Intuition

The idea is essentially the same as in Approach 1. However, the ordered prefix forms an arithmetic sequence with a common difference of 1. Therefore, instead of calculating its sum by traversing the prefix, we can directly use the formula for the sum of an arithmetic sequence.

First, we determine the length of the longest ordered prefix. We can then calculate its sum using the first and last elements of the prefix. Finally, we incrementally check whether the resulting value exists in nums, and the first value that does not appear is the answer.

#### Complexity Analysis

Let n be the length of nums.

- Time complexity: O(n).

 Finding the longest ordered prefix takes O(n) time, and finding the first value that does not appear in nums takes O(n) time in the worst case.

- Space complexity: O(n).

 The hash set requires O(n) auxiliary space.
