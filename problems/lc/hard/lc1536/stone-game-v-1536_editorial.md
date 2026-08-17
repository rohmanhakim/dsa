# LeetCode's Official Editorial

### Approach 1: Dynamic Programming

#### Intuition

We use f\[l\]\[r\] to represent the maximum score Alice can achieve when she is given the continuous segment of stones from position l to r in the array stoneValue. Since Alice needs to split this segment into two parts, we can enumerate the split position i, where the left part consists of the stones from position l to i, and the right part consists of the stones from position i+1 to r. Let sum(l,r) denote the total value of the stones from position l to r. According to the problem statement, there are three cases:

- If sum(l,i)<sum(i+1,r), Bob discards the right part, so Alice continues with the left part. The state transition is:

 f\[l\]\[r\]\=f\[l\]\[i\]+sum(l,i)

- If sum(l,i)\>sum(i+1,r), Bob discards the left part, so Alice continues with the right part. The state transition is:

 f\[l\]\[r\]\=f\[i+1\]\[r\]+sum(i+1,r)

- If sum(l,i)\=sum(i+1,r), Bob lets Alice choose which part to discard. Therefore, Alice chooses the better of the two parts, giving the transition:

 f\[l\]\[r\]\=maxf\[l\]\[i\],f\[i+1\]\[r\]+sum(l,i)

We can precompute sum(l,r) using prefix sums. Alternatively, when enumerating the split position i, we can maintain sum(l,i) directly. Then, the sum of the right part can be calculated as:

sum(i+1,r)\=sum(l,r)−sum(l,i)

which can be obtained in O(1) time.

When only one stone remains, Alice cannot make another split, so her score is 0. This gives the boundary condition:

f\[l\]\[l\]\=0

The final answer is f\[0\]\[n−1\], where n is the length of the array stoneValue.

**Note:** A straightforward implementation with three nested loops over l, r, and i may exceed the time limit in C++. Therefore, we use a top-down dynamic programming approach with memoization. This allows us to skip states that are never reached and pass within the time limit.

#### Complexity Analysis

Let n be the length of the array stoneValue.

- Time complexity: O(n3).

- Space complexity: O(n2).

 The O(n2) space is required to store all DP states.

---

### Approach 2: Dynamic Programming Optimization

#### Intuition

This approach is more advanced, but it does not require any sophisticated optimization techniques.

Consider a particular state transition involving l, r, and i. If

sum(l,i)<sum(i+1,r),

then the transition is:

f\[l\]\[r\]\=f\[l\]\[i\]+sum(l,i)

When considering r+1, as long as

sum(l,i)<sum(i+1,r+1),

the same transition remains valid:

f\[l\]\[r+1\]\=f\[l\]\[i\]+sum(l,i)

Therefore, the value f\[l\]\[i\]+sum(l,i) may be used repeatedly for different values of r. This repeated computation is the main bottleneck of Approach 1.

To avoid this, we maintain two auxiliary arrays, maxl and maxr:

\\left{ \\begin{aligned} \\textit{maxl}\[l\]\[r\] &= \\max\_{i=l}^{r} \\left{f\[l\]\[i\] + \\textit{sum}(l, i)\\right} \[4pt\] \\textit{maxr}\[l\]\[r\] &= \\max\_{i=l}^{r} \\left{f\[i\]\[r\] + \\textit{sum}(i, r)\\right} \\end{aligned} \\right.

In this way, for any fixed l and r, there exists an index i0​∈\[l−1,r) such that:

\\left{ \\begin{aligned} &\\textit{sum}(l, i\_0) \\leq \\textit{sum}(i\_0+1, r), \\ &\\textit{sum}(l, i\_0+1) > \\textit{sum}(i\_0+2, r). \\end{aligned} \\right.

Here, we define sum(x,y)\=0 when x\>y.

When i≤i0​, the corresponding transitions can be combined into:

f\[l\]\[r\]\=maxl\[l\]\[i0​\]

When i\>i0​, the corresponding transitions can be combined into:

f\[l\]\[r\]\=maxr\[i0​+2\]\[r\]

Furthermore, if

sum(l,i0​)\=sum(i0​+1,r),

there is an additional transition:

f\[l\]\[r\]\=maxr\[i0​+1\]\[r\]

Therefore, if we know maxl, maxr, and i0​, we can calculate f\[l\]\[r\] in O(1) time.

The remaining question is how to efficiently calculate these values.

**Details**

From the definitions of maxl and maxr, we can calculate them incrementally while enumerating l and r. Specifically:

\\left{ \\begin{aligned} \\textit{maxl}\[l\]\[r\] &= \\max{\\textit{maxl}\[l\]\[r-1\], f\[l\]\[r\] + \\textit{sum}(l, r)} \[4pt\] \\textit{maxr}\[l\]\[r\] &= \\max{\\textit{maxr}\[l+1\]\[r\], f\[l\]\[r\] + \\textit{sum}(l, r)} \\end{aligned} \\right.

The boundary conditions are:

maxl\[l\]\[l\]\=#maxr\[l\]\[l\]stoneValue\[l\]

For i0​, let il,r​ denote its value for a particular pair (l,r). Since all elements in stoneValue are positive, we have:

il,r​≤il,r+1​

Therefore, when l is fixed, il,r​ is monotonically increasing as r increases.

This allows us to enumerate l in decreasing order and, for each fixed l, enumerate r in increasing order while maintaining il,r​.

- Initially, il,l​\=l−1.

- Once il,r​ is known, we can find il,r+1​ by continuously increasing il,r​ while

 sum(l,il,r​)≤sum(il,r​+1,r+1)

 remains true. Once this condition no longer holds, the value of il,r+1​ is il,r​.

With this monotonicity, each value of il,r​ only moves forward during the traversal. Therefore, the total time spent maintaining il,r​ is amortized O(1) per state.

As a result, maxl\[l\]\[r\] and maxr\[l\]\[r\] can be updated in O(1) time, the amortized time for maintaining il,r​ is O(1), and each state f\[l\]\[r\] can also be computed in O(1) time.

#### Complexity Analysis

Let n be the length of the array stoneValue.

- Time complexity: O(n2).

- Space complexity: O(n2).

 The O(n2) space is required to store the DP states and the auxiliary arrays maxl and maxr.
