# LeetCode's Official Editorial

### Approach 1: Greedy

#### Intuition

The goal is to maximize the number of `1`s in the string `s`. Let's first analyze the effect of performing a single operation.

According to the problem statement, we first add `1`s to both ends of `s` to obtain a new string `t`:

t\=1+s+1

Without loss of generality, suppose `t` contains the following pattern:

...1a000​​b111​​c00​​1...

where

- the middle block `111` is a contiguous block of `1`s of length b, surrounded by `0`s;
- the left and right blocks are contiguous blocks of `0`s with lengths a and c, respectively.

First, we change the middle block of `1`s into `0`s:

...1a000​​b000​​c00​​1...

The three consecutive zero blocks immediately merge into one larger block:

...1a+b+c00000000​​1...

Next, we change this entire block of `0`s into `1`s:

...1a+b+c11111111​​1...

Observe what happens during this process:

- The original block of `1`s of length b is first removed and then restored, so it contributes no net change.
- The two neighboring zero blocks of lengths a and c are converted into `1`s.

Therefore, the increase in the number of `1`s is 
(a+b+c)−b\=a+c.

In other words, **the gain from one operation equals the sum of the lengths of two adjacent zero blocks separated by a block of `1`s**. The length of the middle block of `1`s does not matter.

Therefore, we only need to record the lengths of all contiguous zero blocks in `t`. Suppose we obtain the following array:

\[z0​,z1​,…,zm−1​\]

We enumerate every pair of adjacent zero blocks. The final answer is

cnt1​+max0≤i<m−1​(zi​+zi+1​)

where cnt1​ is the number of `1`s in the original string `s`.

If there are fewer than two zero blocks, no valid operation can be performed, so the answer is simply cnt1​.

Notice that the auxiliary `1`s added to both ends do not affect the lengths of the zero blocks. Therefore, we can simply traverse the original string `s`.

### Notes

> **After changing a block of `1`s into `0`s in the first step, could the resulting zero block fail to be surrounded by `1`s, making the second step impossible?**

No.

If that happened, the string `t` would begin or end with `0`. However, the first and last characters of `t` are the auxiliary `1`s that we added, and these characters are never modified. Therefore, the merged zero block is always surrounded by `1`s, so the second operation can always be performed.

> **After changing a block of `1`s into `0`s, could it be better to choose a different zero block in the second step instead of the newly formed one?**

No.

Suppose

- the removed block of `1`s has length b;
- in the second step, we instead choose another zero block of length zi​.

The net gain would then be

zi​−b

However, every zero block is adjacent to another zero block through some block of `1`s, so it can always participate in a valid operation that yields

zi​+zi−1​orzi​+zi+1​,

Since

zi​+zi±1​\>zi​−b

choosing the newly formed merged zero block is always at least as good, and therefore optimal.


#### Complexity Analysis

Let n be the length of s, and let k be the number of contiguous zero blocks.

- Time complexity: O(n).

- Space complexity: O(k).

 The extra space is used to store the lengths of the zeroBlocks.

---

### Approach 2: Space Optimization

#### Intuition

In Approach 1, we store the lengths of all zero blocks in an array zeroBlocks.

However, we only need the maximum sum of two adjacent zero blocks, so storing the entire array is unnecessary. Instead, we maintain three variables:

- prev: the length of the previous zero block;
- cur: the length of the current zero block;
- bestGain: the maximum gain obtained so far.

Whenever we finish scanning a zero block, we compute

bestGain\=max(bestGain,prev+cur)

and then update

prev\=cur

Since each character is processed exactly once, the time complexity remains O(n), while the extra space complexity is reduced from O(k) to O(1).


#### Complexity Analysis

Let n be the length of s.

- Time complexity: O(n).

- Space complexity: O(1).
