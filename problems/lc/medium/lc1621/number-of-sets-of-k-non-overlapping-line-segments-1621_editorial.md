# LeetCode's Official Editorial

### Approach 1: Dynamic Programming

#### Intuition

We need to select exactly k non-overlapping line segments from n points.

Let dp\[i\]\[j\] denote the number of valid ways to construct i line segments using the points in \[0,j\].

First, consider the boundary conditions. When i\=0, there is exactly one way (selecting no segments), so dp\[0\]\[j\]\=1 for all j∈\[0,n−1\]. When i\>0 and j\=0, no valid segment can be formed using only a single point, so dp\[i\]\[0\]\=0.

When i\>0 and j\>0, there are two cases:

1. The right endpoint of the last line segment is not j. This is equivalent to selecting i line segments using points in \[0,j−1\], contributing dp\[i\]\[j−1\] ways.
2. The right endpoint of the last line segment is j. If the last line segment is \[p,j\] for some left endpoint p∈\[0,j), the first i−1 line segments must be constructed using points in \[0,p\]. Summing over all possible choices of p contributes p\=0∑j−1​dp\[i−1\]\[p\] ways.

Therefore, the transition is

dp\[i\]\[j\]\=dp\[i\]\[j−1\]+p\=0∑j−1​dp\[i−1\]\[p\].

To compute p\=0∑j−1​dp\[i−1\]\[p\] in O(1) time, we can maintain running prefix sums of dp\[i−1\]. Furthermore, since the transitions for i line segments only depend on the values for i−1 line segments, we can optimize the space to O(n) using a rolling array.

#### Complexity Analysis

Let n be the number of points and k be the target number of line segments.

- Time complexity: O(nk).

 There are O(nk) states, and each state takes O(1) time to compute.

- Space complexity: O(n).

 The rolling array dp and the prefix sums array prefixSums both require O(n) space.

### Approach 2: Combinatorics

#### Intuition

The problem is equivalent to counting the tuples (l1​,…,lk​,r1​,…,rk​) that satisfy

0 ≤ l1 ​< r1 ​≤ l2 ​< r2 ​≤ ⋯ ≤ lk ​< rk ​< n.

From a combinatorial perspective, if we can make all line-segment endpoints distinct, the problem becomes selecting 2k numbers from a certain set of numbers.

Let li′​\=li​+(i−1) and ri′​\=ri​+(i−1). This moves the endpoints of the second line segment one unit to the right, the endpoints of the third line segment two units to the right, and so on. After this transformation, the new l′ and r′ values correspond one-to-one with the original l and r values, and the endpoints no longer coincide. They satisfy

0 ≤ l1′ ​< r1′ ​< l2′ ​<r2′ ​< ⋯ < lk′ ​< rk′ ​< n + k − 1.

The problem is now equivalent to selecting 2k numbers from the n+k−1 numbers in \[0,n+k−1), so the answer is (2kn+k−1​).

#### Complexity Analysis

Let M\=109+7 be the modulus.

- Time complexity: O(k+logM).

 Computing the binomial coefficient takes O(k+logM) time.

- Space complexity: O(1).