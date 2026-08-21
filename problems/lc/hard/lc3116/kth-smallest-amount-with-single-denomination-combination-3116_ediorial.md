# LeetCode's Official Editorial

### Approach 1: Binary Answer + Inclusion-Exclusion Principle

#### Intuition

Given a set of coins, there are n (n≤15) types of coins. The i\-th type of coin has a denomination of coins\[i\] (1≤coins\[i\]≤25), and each type of coin has an unlimited supply. We need to find the k\-th smallest amount that can be formed using coins of the same denomination.

Since k can be very large, it is not possible to directly enumerate all amounts that can be formed using coins of the same denomination. However, given an amount x, we can use the **inclusion-exclusion principle** to calculate how many amounts are less than or equal to x.

Define count\[x\] as the number of integers in the interval \[1,x\] that are divisible by at least one coins\[i\]. Since count\[x\] is non-decreasing with respect to x, we can use binary search to find the smallest x such that count\[x\]≥k.

Next, we discuss how to calculate count\[x\]. Suppose there are only two denominations, a and b. There are ⌊x/a⌋ multiples of a and ⌊x/b⌋ multiples of b. However, numbers that are multiples of both a and b are counted twice, so we need to subtract ⌊x/lcm(a,b)⌋, where lcm(a,b) is the least common multiple of a and b.

When extending this to multiple denominations, we can apply the inclusion-exclusion principle:

- If the size of the subset is odd, add ⌊x/lcm⌋, where lcm is the least common multiple of all denominations in the subset.

- If the size of the subset is even, subtract ⌊x/lcm⌋, where lcm is the least common multiple of all denominations in the subset.

#### Details

- In the binary search, the left endpoint is k, and the right endpoint is k×mincoins. If we use a right-open interval, we need to add 1 to the maximum value. In this implementation, the binary search uses ⌊(l+r)/2⌋ as the midpoint and therefore never evaluates r directly.

- Since the maximum size of a subset is only 15, we can precompute the least common multiple of every subset and then traverse all subsets during the binary search to calculate count\[x\].

- When precomputing the least common multiple of every subset, the least common multiple may exceed the integer range in some languages or become greater than the right boundary r of the binary search. Such subsets will not contribute to count\[x\], so they can be skipped.

#### Complexity Analysis

Let n be the length of coins.

- Time complexity: O(2n×(nlog(max{coins})+log(k×min{coins})))

 During the preprocessing stage, calculating the least common multiple of each subset takes O(nlog(maxcoins)) time per subset, and there are O(2n) subsets. During the binary search, we perform O(log(k×mincoins)) iterations, and each iteration takes O(2n) time to calculate count\[x\].

- Space complexity: O(2n).

### Approach 2: Optimized Based on Approach 1

#### Intuition

**Optimization 1: Remove Redundant Denominations**

If coins contains denominations 3 and 9, every amount that can be formed using 9 can also be formed using 3. Therefore, if one denomination is a multiple of another, the larger denomination can be removed directly.

**Optimization 2: Reduce the Preprocessing Time for the Least Common Multiple of Subsets**

For convenience, we use the integer mask to represent a subset of coin denominations. If the i\-th bit in the binary representation of mask is 1, the subset contains coins\[i\]; otherwise, it does not. Let lcm(mask) denote the least common multiple of all denominations in the subset, and define lcm(0)\=1.

We enumerate all non-empty subsets mask in ascending order. Let the lowest set bit in its binary representation be at position i. After removing this bit, we obtain the subset \\textit{pre\_mask}, which is given by

\\textit{pre\_mask}=\\textit{mask} \\& (\\textit{mask}-1).

Since \\textit{pre\_mask}<\\textit{mask}, \\textit{lcm}(\\textit{pre\_mask}) has already been calculated when we process mask. The subset mask can be considered as adding the denomination coins\[i\] to \\textit{pre\_mask}, so we have

\\textit{lcm}(\\textit{mask}) = \\textit{lcm}\\bigl(\\textit{lcm}(\\textit{pre\_mask}),\\textit{coins}\[i\]\\bigr).

Therefore, the least common multiple of each subset can be recursively computed from a smaller subset in O(logmax(coins)) time.

#### Complexity Analysis

Let n be the length of coins.

- Time complexity: O(n2+2n×(log(maxcoins)+log(k×mincoins))).

 Removing redundant denominations takes O(n2) time. During the preprocessing stage, the least common multiple of each subset can be computed in O(log(maxcoins)) time, and there are O(2n) subsets. During the binary search, we perform O(log(k×mincoins)) iterations, and each iteration takes O(2n) time to calculate count\[x\].

- Space complexity: O(2n).

---
