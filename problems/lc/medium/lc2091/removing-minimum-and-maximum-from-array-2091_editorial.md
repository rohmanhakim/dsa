# LeetCode's Official Editorial

### Approach: Classification Discussion

#### Intuition

We first traverse the array to find the indices of the minimum and maximum values. For convenience, let l be the smaller of the two indices and r be the larger one. Assuming the array length is n, these two indices divide the remaining elements of nums into three mutually disjoint subarrays, which are nums\[0..l−1\], nums\[l+1..r−1\], and nums\[r+1..n−1\], where \[a..b\] denotes a closed interval.

Since we can only remove elements from the beginning or end of the array, after removing the minimum and maximum values, the remaining subarray must be a **subarray** of one of the three blocks mentioned above. Therefore, to minimize the number of deletions, the remaining subarray must be **exactly one of these three blocks**. The corresponding number of deletions is the difference between the length of the original array n and the length of the remaining subarray.

Specifically:

- If the remaining subarray is nums\[0..l−1\], the number of deletions is n−l.

- If the remaining subarray is nums\[l+1..r−1\], the number of deletions is l+1+n−r.

- If the remaining subarray is nums\[r+1..n−1\], the number of deletions is r+1.

The minimum of these three values is the smallest possible number of deletions, which is the answer.

#### Complexity Analysis

Let n be the length of nums.

- Time complexity: O(n).

 We traverse the array once to find the indices of the minimum and maximum values.

- Space complexity: O(1).
