# LeetCode's Official Editorial

### Approach 1: Binary Search + Segment Tree

#### Intuition

Readers are advised to first solve the prerequisite problem [3499\. Maximize Active Section with Trade I](https://leetcode.com/problems/maximize-active-section-with-trade-i/).

In [3499\. Maximize Active Section with Trade I](https://leetcode.com/problems/maximize-active-section-with-trade-i/), we proved that after performing one operation, the maximum number of active sections in the string is

cnt1​+bestGain

where

- cnt1​ is the number of characters 1 in the string s.
- bestGain is the maximum sum of the lengths of two adjacent consecutive blocks of 0s.

In this problem, however, we need to answer multiple queries.

For each query \[l,r\], the operation is restricted to the substring s\[l..r\]. Note that the required answer is still the maximum number of active sections in the **entire string**, not just inside the substring. Therefore, for each query, we only need to compute the corresponding bestGain inside s\[l..r\].

A brute-force solution would enumerate every consecutive block of 0s inside the substring for every query, resulting in a time complexity of

O(nq),

which is too slow.

Therefore, we need a faster way to answer each query.

---

## Preprocessing Consecutive Blocks of 0s

First, extract the lengths of all consecutive blocks of 0s in the original string s.

Let 
zeroBlocks\=\[z0​,z1​,…,zm−1​\]

where

- m is the number of consecutive blocks of 0s.
- zk​ is the length of the k\-th consecutive block of 0s.

For a query \[l,r\], let 
subZeroBlocks

denote the array of consecutive 0 block lengths inside the substring s\[l..r\].

The following examples illustrate the possible forms of subZeroBlocks.

```
Index  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
S      0  0  1  1  0  1  0  0  1  0  0  0  0  1  1  1

zeroBlocks = [2, 1, 2, 4]

queries[i] = {1,12}:
    - subZeroBlocks = {1,1,2,4}

queries[i] = {4,6}
    - subZeroBlocks = {1,1}

queries[i] = {1,11}
    - subZeroBlocks = {1,1,2,3}

queries[i] = {4,7}
    - subZeroBlocks = {1,2}

queries[i] = {5,5}
    - subZeroBlocks = {}

- Partial 0-blocks at the substring boundaries can change both ends of subZeroBlocks
- subZeroBlocks may have < 2 entries
```

Observe that, except for the first and last elements, subZeroBlocks is always a **contiguous subarray** of zeroBlocks.

The reason is simple: a substring may cut through only the first and/or last block. Therefore, **only the lengths of the first and last blocks can change**.

Thus,

subZeroBlocks\=\[zi′​,zi+1​,…,zj−1​,zj′​\].

Here,

- zi′​ is the actual length of the first consecutive block of 0s inside s\[l..r\], where zi′​≤zi​.
- zj′​ is the actual length of the last consecutive block of 0s inside s\[l..r\], where zj′​≤zj​.

Furthermore,

- zi′​<zi​ if and only if s\[l\]\=′0′ and l is not the left endpoint of that block.
- zj′​<zj​ if and only if s\[r\]\=′0′ and r is not the right endpoint of that block.

##### Computing bestGain

From now on, for brevity, we refer to consecutive blocks of 0s simply as **blocks**.

Since 
bestGain\=max(zk​+zk+1​),

the answer for a query is simply the maximum of the following three cases.

#### Case 1

Use the first two blocks of subZeroBlocks: 
val1​\=zi′​+zi+1​.

#### Case 2

Use the last two blocks:

val2​\=zj−1​+zj′​.

#### Case 3

Use two adjacent blocks that lie completely inside the substring: 
val3​\=maxi+1≤k≤j−2​(zk​+zk+1​).

Therefore, 
bestGain\=max(val1​,val2​,val3​).

##### Locating Blocks Efficiently

To compute these values efficiently, we need to quickly identify the corresponding blocks.

Preprocess two arrays:

- blockLeft\[k\]: the left endpoint of the k\-th block.
- blockRight\[k\]: the right endpoint of the k\-th block.

Since the blocks do not overlap, both arrays are strictly increasing, allowing us to locate the relevant blocks using binary search.

---

#### Processing a Query

For a query \[l,r\]: 
First, binary search blockRight to find the first block satisfying 
blockRight\[i\]≥l.

Suppose this block is 
\[Li​,Ri​\].

Then the actual length of the first block inside the substring is 
zi′​\=Ri​−max(Li​,l)+1.

Similarly, binary search blockLeft to find the last block satisfying 
blockLeft\[j\]≤r.

Suppose this block is 
\[Lj​,Rj​\].

Then 
zj′​\=min(Rj​,r)−Lj​+1.

To compute Case 3, define

tmpSumk​\=zk​+zk+1​.

Then 
val3​\=maxi+1≤k≤j−2​tmpSumk​.

This becomes a standard range maximum query on the array tmpSum, which can be answered using a segment tree.

---

#### Edge Cases

Several special cases should be handled.

- Binary search may fail, meaning i\>m−1 or j<0. In this case, the substring contains no consecutive block of 0s, so bestGain\=0.
- If i≥j, the substring contains at most one consecutive block of 0s, so there are no adjacent blocks and bestGain\=0.
- If the substring contains exactly two consecutive blocks of 0s, then there are no complete interior blocks, making Case 3 irrelevant. In this case,

bestGain\=firstLen+lastLen.

- If the original string contains fewer than two consecutive blocks of 0s, no operation can increase the answer, so every query simply returns cnt1​.

#### Complexity Analysis

Let n be the length of the string s.

- Time Complexity: O(n+qlogn), including the preprocessing part such as building the segment tree, with a time complexity of O(n). For each query:

- The time complexity of two binary searches is O(logn);

- The time complexity of one segment tree range maximum query is O(logn).

Therefore, the total time complexity is:

O(n+qlogn)

- Space Complexity: O(n).

---

### Approach 2: Binary Search + Sparse Table

#### Intuition

In Approach 1, we use a segment tree to answer range maximum queries on the array `tmpSum`.

Note that in this problem, the array tmpSum does not change after preprocessing, so we can also use a Sparse Table to solve the static RMQ (Range Maximum Query) problem.

#### Complexity Analysis

- Time complexity: O((n+q)logn), where n is the length of the string s. The preprocessing time complexity of the Sparse Table is O(nlogn). For each query:

 - The time complexity of two binary searches is O(logn);

 - The time complexity of a single Sparse Table query is O(1).

- Space complexity: O(nlogn). The Sparse Table requires O(nlogn) additional space.

### Approach 3: Mo's Algorithm (Not the expected complexity for this problem, optional reading)

#### Intuition

Mo's algorithm is a classic offline algorithm for answering multiple static range queries using a sliding window. By carefully ordering the queries, it reduces the total cost of moving the sliding window, achieving a time complexity on the order of the square root of the input size.

Since this problem involves multiple interval queries without any updates, it provides a good introduction to Mo's algorithm. However, note that this solution does **not** achieve the expected time complexity for this problem and is therefore **not guaranteed to pass all test cases**. It is included for readers who are interested in learning this technique.

The core idea behind Mo's algorithm is **square root decomposition**. We partition the string s into blocks. For simplicity, let the block length be

B\=n​.

A brute-force solution processes a query in O(n) time, where n is the length of the substring. Therefore, for queries whose interval length is at most B, we simply compute the answer by brute force, resulting in a total complexity of

O(qn​).

We only need to apply Mo's algorithm to queries whose interval length exceeds B. Since such queries are longer than one block, their left and right endpoints cannot lie in the same block.

We group these queries according to the block containing their **left endpoint**. Within each group, the queries are sorted in ascending order of their **right endpoint**.

> Since queries within the same group may be processed in any order, we can simply sort all remaining queries using the block ID of the left endpoint as the primary key and the right endpoint as the secondary key. As a result, queries belonging to the same group become consecutive, and their right endpoints are automatically sorted in ascending order.

Next, we process the groups one by one. For now, assume that every sliding window update and every query can be answered in O(1) time.

Before processing each group, we initialize the sliding window and the auxiliary data structures required to answer the queries. Throughout the discussion below, we represent the sliding window using **open intervals**.

Suppose we are processing the queries whose left endpoints belong to the k\-th block (where k is zero-indexed). We initialize the sliding window as

(L,R)\=(kB−1,(k+1)B).

That is, L is initialized to the right endpoint of the k\-th block, while R is initialized to the left endpoint of the next block.

Since the interval is open, neither endpoint belongs to the window. Therefore, after initialization, the sliding window is empty.

At the same time, we initialize an array subZeroBlocks to record the lengths of consecutive blocks of 0s inside the current window.

We now process the queries in each group from left to right.

---

#### Step 1: Expand the Right Endpoint

For a query \[l,r\], we first expand the right endpoint until R\>r.

During this process, new characters continuously enter the window from the right.

- If s\[R\]\=′1′, then subZeroBlocks remains unchanged.
- Otherwise, we update subZeroBlocks.
- If R is the first position of a consecutive block of 0s in the original string, append a new element of value 1 to the end of subZeroBlocks, indicating that a new block of length 1 has entered the window.
- Otherwise, increase the last element of subZeroBlocks by 1, since the rightmost block has been extended.

---

#### Step 2: Expand the Left Endpoint

Next, expand the left endpoint until L<l.

During this process, new characters continuously enter the window from the left.

- If s\[L\]\=′1′, then subZeroBlocks remains unchanged.
- Otherwise, update subZeroBlocks.
- If L is the first position of a consecutive block of 0s in the original string, prepend a new element of value 1 to the front of subZeroBlocks, indicating that a new block has entered the window.
- Otherwise, increase the first element of subZeroBlocks by 1, since the leftmost block has been extended.

Using these rules, we can maintain subZeroBlocks dynamically as the sliding window expands.

Once both endpoints have been adjusted, the current window exactly corresponds to the query interval \[l,r\]. We can therefore compute bestGain from the current subZeroBlocks and answer the query.

---

#### Step 3: Roll Back the Left Endpoint

After answering the current query, we restore the left endpoint to its initial position so that the next query in the same group can be processed.

As the window shrinks from the left, characters continuously leave the window. We update subZeroBlocks using the reverse of the operations described in Step 2.

After restoring the left endpoint to the right boundary of the current block, we return to Step 1 to process the next query.

#### Complexity Analysis

- Time complexity: O(qlogq+nn​+qn​), which simplifies to O(nn​) when q and n are of the same order.

- Space complexity: O(n).

Below, assume that both sliding window updates and query answering can be performed in O(1) time. We first analyze the complexity of the algorithm under this assumption.

Since the queries within each group are sorted by non-decreasing right endpoint, the right endpoint R only moves monotonically to the right.

We partition the string into blocks of length

B\=n​,

so the total number of blocks is approximately

blockCount\=Bn​\=n​.

Within each group, the right endpoint moves from its initial position to the end of the string at most once. Therefore, over all groups, the total number of movements of the right endpoint is

O(n⋅blockCount)\=O(nn​).

Now consider the left endpoint.

For every query, the left endpoint first expands until it reaches the query boundary and is then restored to its initial position. Since both the query's left endpoint and the initial position belong to the same block, the total movement of the left endpoint for a single query is at most twice the block length.

Therefore, the total complexity contributed by the left endpoint is

O(2qB)\=O(qn​).

As mentioned earlier, queries whose interval length does not exceed B are handled using brute force, contributing another

O(qn​)

time.

Finally, sorting all queries requires

O(qlogq).

Combining all components, the overall time complexity is

O(qlogq+nn​+qn​).

When q and n are of the same order, this simplifies to

O(nn​).

This is the central idea behind Mo's algorithm: by carefully reordering the queries, the total movement of the sliding window is minimized, making it possible to answer a large number of static interval queries efficiently.

---

In the previous discussion, we assumed that both sliding window updates and query answering can be performed in O(1) time. We now explain how to achieve this.

Since elements need to be inserted and removed at both ends of subZeroBlocks, a deque is a natural choice for maintaining it, allowing each update to be completed in O(1) time.

However, maintaining only subZeroBlocks is not sufficient. To answer a query, we would still need to scan the entire deque to compute bestGain, which is no longer an O(1) operation.

Therefore, we need to maintain additional information that allows the answer to be updated incrementally.

#### Maintaining bestGain

Recall that bestGain is defined as the maximum sum of two adjacent elements in subZeroBlocks. Instead of recomputing it from scratch for every query, we maintain its value dynamically.

Before processing each group of queries, in addition to initializing an empty deque subZeroBlocks, we also initialize

bestGain\=0.

Updating bestGain during window expansion is straightforward. Whenever one end of subZeroBlocks changes, only the newly formed adjacent pair near that end can affect the answer.

Suppose the current deque is

subZeroBlocks\=\[z0​,z1​,…,zm−1​\].

When appending or modifying an element on the right, we simply update

bestGain\=max(bestGain,zm−2​+zm−1​).

Similarly, when inserting or modifying an element on the left, we update

bestGain\=max(bestGain,z0​+z1​).

As a result, every query can be answered in O(1) time using the maintained value of bestGain.

---

#### The Rollback Problem

The difficulty arises when restoring the left endpoint after answering a query.

During rollback, one of the following may happen:

- The leftmost element z0​ is removed from subZeroBlocks.
- The value of z0​ decreases.

Suppose the current maximum is

bestGain\=z0​+z1​.

After modifying or removing z0​, what should the new value of bestGain be?

Unfortunately, this cannot be determined in O(1) time. We only know the current maximum, but we do not know the **second-largest** adjacent sum. Once the pair contributing the maximum disappears, we have no way to immediately determine the new maximum.

One possible solution is to maintain all adjacent sums in an ordered set or a lazy-deletion heap. Then, whenever an adjacent pair becomes invalid, the next largest value can be obtained efficiently.

However, every update would then require an additional

O(logn)

time, increasing the complexity of every sliding window update. Can we avoid this?

---

#### Rollback Optimization

Notice how the algorithm is organized.

For every query, we first expand the left endpoint until it reaches the query boundary, answer the query, and then restore the left endpoint to its original position before processing the next query.

The key observation is that **no query is answered during the rollback process**.

Therefore, after rollback, we only need the maintained information to be exactly the same as it was before expanding the left endpoint.

In particular, the value of bestGain after rollback must be identical to its value before expansion, because the sliding window itself has been restored to the same state.

Thus, before expanding the left endpoint, we simply save the current value of

bestGain.

After rollback, we restore it directly.

This completely avoids the need to maintain the "second-largest" adjacent sum, eliminating any additional data structures and keeping both sliding window updates and query answering at O(1) time.

This idea is commonly known as **Rollback Mo's Algorithm**. Interested readers may continue exploring this technique by solving [3636\. Threshold Majority Queries](https://leetcode.com/problems/threshold-majority-queries/description/i/).

---

#### Skipping Entire Blocks

We can further optimize the sliding window movement through preprocessing.

Specifically, preprocess two arrays:

- left\[i\]: the length of the consecutive block ending at position i that has the same character as s\[i\].
- right\[i\]: the length of the consecutive block starting at position i that has the same character as s\[i\].

For example,

s\=0011100

gives

left\=\[1,2,1,2,3,1,2\]

and

right\=\[2,1,3,2,1,2,1\].

Both arrays can be computed in linear time.

With these arrays, we no longer need to move the sliding window one character at a time. Instead, we can jump directly over entire blocks of identical characters.

For example, suppose the current right endpoint is R.

- If s\[R\]\=′1′, then the entire consecutive block of 1s contributes nothing to the answer and can be skipped.
- Otherwise, the length of the current consecutive block of 0s is

sz\=min(right\[R\],,r−R+1),

where the minimum is taken because the block may extend beyond the query boundary.

We can then move the right endpoint directly by

R←R+sz,

instead of advancing one position at a time.

The same optimization applies to the left endpoint. If the current left endpoint is L, then the length of the current consecutive block is

sz\=min(left\[L\],,L−l+1),

and we update

L←L−sz.

In this way, the sliding window skips entire blocks instead of individual characters.

This optimization makes the algorithm significantly faster in practice, especially on random data or strings containing long consecutive blocks.

---

#### Simplifying Rollback

The rollback process can be simplified even further.

Observe that expanding the left endpoint affects subZeroBlocks in only two ways:

- If the left endpoint falls inside an existing consecutive block of 0s, only the first element of subZeroBlocks changes.
- Otherwise, one or more new blocks are inserted at the front.

Therefore, instead of undoing every operation one by one during rollback, we only need to record two pieces of information before expanding the left endpoint:

- the original value of the first element of subZeroBlocks;
- the number of new elements inserted at the front.

During rollback, we simply:

- remove the recorded number of newly inserted elements from the front;
- restore the original value of the first element.

This restores subZeroBlocks to exactly the same state it had before the expansion, while keeping the rollback operation simple and efficient.

