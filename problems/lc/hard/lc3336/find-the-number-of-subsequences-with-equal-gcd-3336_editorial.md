# LeetCode's Official Editorial

### Approach: Dynamic Programming

#### Intuition

Let n be the length of the array nums, and let m be the maximum element in nums. For any two positive integers a and b, their greatest common divisor satisfies

gcd(a,b)≤min(a,b),

where min(a,b) denotes the smaller of the two values. Therefore, the GCD of either subsequence can never exceed m.

Let dp\[i\]\[j\]\[k\] denote the number of ways to process the first i elements, where the current GCD of the first subsequence seq1​ is j, and the current GCD of the second subsequence seq2​ is k.

Initially, before processing any elements, both subsequences are empty. We define the GCD of an empty subsequence as 0, so

dp\[0\]\[0\]\[0\]\=1.

Now consider the i\-th element nums\[i\] (using 1\-based indexing). For every state dp\[i−1\]\[j\]\[k\], there are three possible choices:

- **Add nums\[i\] to seq1​.**

The GCD of seq1​ becomes gcd(j,nums\[i\]), while the GCD of seq2​ remains k. Therefore, dp\[i\]\[gcd(j,nums\[i\])\]\[k\] receives dp\[i−1\]\[j\]\[k\] additional ways.

- **Add nums\[i\] to seq2​.**

Similarly, the GCD of seq2​ becomes gcd(k,nums\[i\]), while the GCD of seq1​ remains unchanged. Thus, dp\[i\]\[j\]\[gcd(k,nums\[i\])\] receives dp\[i−1\]\[j\]\[k\] additional ways.

- **Do not add nums\[i\] to either subsequence.**

Both GCDs remain unchanged, so dp\[i\]\[j\]\[k\] receives dp\[i−1\]\[j\]\[k\] additional ways.

These three choices are mutually exclusive and cover all possible assignments of the current element. Hence, the recurrence can be written as

dp\[i\]\[j\]\[k\]\=​ dp\[i−1\]\[j\]\[k\] ​+j′∑​dp\[i−1\]\[j′\]\[k\](where j\=gcd(j′,nums\[i\])) ​+k′∑​dp\[i−1\]\[j\]\[k′\](where k\=gcd(k′,nums\[i\])).​

In practice, it is more convenient to iterate over every state dp\[i−1\]\[j\]\[k\] and distribute its value to the three destination states:

{dp\[i\]\[j\]\[k\]+\=dp\[i−1\]\[j\]\[k\], dp\[i\]\[gcd(j,nums\[i\])\]\[k\]+\=dp\[i−1\]\[j\]\[k\], dp\[i\]\[j\]\[gcd(k,nums\[i\])\]+\=dp\[i−1\]\[j\]\[k\].​

Since the answer is required modulo 109+7, all additions are performed modulo 109+7.

After processing all n elements, we need to count the states where the two subsequences have the same positive GCD. Since the GCD of an empty subsequence is defined as 0, states with GCD 0 are excluded. Therefore, the final answer is

ans\=∑i\=1m​dp\[n\]\[i\]\[i\].

Return the result modulo 109+7.

**Space Optimization**

Notice that dp\[i\] depends only on dp\[i−1\]. Therefore, we can optimize the first dimension using a rolling array.

Let dp\[j\]\[k\] denote the states after processing the current prefix of the array. For each new element num, create a temporary array ndp\[j\]\[k\] to store the updated states. After all transitions are complete, replace dp with ndp. This reduces the space complexity from O(nm2) to O(m2).

#### Complexity Analysis

Let n be the length of the array nums, and let m be the maximum element in nums.

- Time complexity: O(nm2logm).

 Finding the maximum element takes O(n) time. The dynamic programming table contains O(nm2) states, and each transition computes two GCD values, each taking O(logm) time. Therefore, the overall time complexity is O(nm2logm).

- Space complexity: O(m2).

 After applying the rolling array optimization, only two O(m2) two-dimensional arrays are maintained.
