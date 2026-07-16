# LeetCode's Official Editorial

### Approach: Simulation

#### Intuition

The problem requires us to construct the array prefixGcd. According to its definition, we first need to construct the array mx.

A brute-force approach would compute each mxi​ independently, resulting in a time complexity of O(n2), which is too slow. Observe that mx is simply the prefix maximum array of nums. Therefore, we can maintain the current prefix maximum while traversing the array once, allowing us to construct mx in linear time.

Once mx has been constructed, we can directly compute the array prefixGcd according to its definition.

Next, we sort the array prefixGcd. We then repeatedly pair the smallest remaining element with the largest remaining element and add their greatest common divisor to the answer. This process can be simulated using the two-pointer technique.

Note that if the length of the array is odd, one element will remain unpaired. According to the problem statement, this element should simply be ignored.

#### Complexity Analysis

Let n be the length of the array nums, and let U be the maximum value in nums.

- Time complexity: O(nlogn+nlogU).

 Sorting prefixGcd takes O(nlogn) time. Computing each greatest common divisor takes O(logU) time, and this operation is performed O(n) times.

- Space complexity: O(n).

 Used to store the arrays mx and prefixGcd.

---
