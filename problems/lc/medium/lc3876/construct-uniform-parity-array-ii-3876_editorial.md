# LeetCode's Official Editorial

### Approach: Classification Discussion

#### Intuition

This problem is similar to "[3875\. Construct Uniform Parity Array I](https://leetcode.com/link/?target=https%3A%2F%2Fleetcode.com%2Fproblems%2Fconstruct-uniform-parity-array-i%2F)", but with an additional constraint: when using the second operation, it must satisfy nums1​\[i\]−nums1​\[j\]≥1, meaning the minuend must be strictly greater than the subtrahend.

Since only subtracting an odd number changes the parity (even minus odd gets odd, odd minus odd gets even), we only need to focus on the minimum odd number in the array.

Let the length of nums1​ be n, and the minimum value in the array be mn. Since the minimum value cannot be changed by subtracting a smaller number, the parity of mn determines the target parity of nums2​.

We classify nums1​ into discussions:

1. nums1​ is all even: In this case, we can directly set each nums2​\[i\]\=nums1​\[i\] (using the first operation), and the resulting nums2​ will all be even, satisfying the condition.

2. nums1​ is all odd: Similarly, we can directly set each nums2​\[i\]\=nums1​\[i\], and the resulting nums2​ will all be odd, satisfying the condition.

3. nums1​ contains both odd and even numbers:

 - If mn is odd: We can construct the array to be all odd. For odd positions, use the first operation directly; for even positions, let the even number subtract mn (a smaller odd number), since even minus odd gets odd, and mn is the minimum value, any even number subtracting mn will result in a number ≥1, satisfying the problem's requirements.

 - If mn is even: It is impossible to construct a valid nums2​. Because:

 - If the goal is to be all even: The odd number must subtract an odd number to become even. But the smallest odd number cannot find a smaller odd number to subtract, so the parity cannot be changed.

 - If the goal is to be all odd: The even number mn must subtract an odd number to become odd. But mn is the minimum value, and there is no smaller number to subtract.

 Both cases are infeasible, so return false.

In summary, we only need to traverse the array once, find the minimum value mn, and record whether there is an odd number:

- If mn is odd, return true directly.

- If mn is even, return true only if there are no odd numbers in the array, otherwise return false.

#### Complexity Analysis

Let n be the length of the array nums1​.

- Time complexity: O(n).

 We only need to traverse the array once to find the minimum value and to determine if there is an odd number.

- Space complexity: O(1).

 Use only constant extra space.

---
