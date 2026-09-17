# LeetCode's Official Editorial

### Approach 1: Prefix Sum + Dynamic Programming + Hash Table

#### Intuition

We need to find two non-overlapping subarrays whose sums both equal $target$ and whose total length is minimized.

First, we can use prefix sums to quickly determine which subarrays have a sum equal to $target$. Let $s[i]$ denote the prefix sum, that is, $s[i]=k=0∑i​arr[k]$. For an index $i$, if we can find an index $j$ such that $s[i]−target=s[j]$, then $arr[j+1…i]$ is a subarray whose sum equals $target$.

To find such an index $j$, we can use a hash table mapping each prefix sum to the index where it occurred, initialized with ${0:−1}$. Because all elements in $arr$ are positive ($arr[k]≥1$), prefix sums are strictly increasing, so each prefix sum is unique.

Next, we need the minimum length of a valid subarray ending at or before index $j$ (in $arr[0…j]$). We can maintain this using dynamic programming. Let $dp[j]$ be this minimum length. When we find a valid subarray $arr[j+1…i]$, if a valid first subarray exists in $arr[0…j]$, the total length of the two subarrays is $dp[j]+(i−j)$.

As we iterate, the minimum length of a valid subarray found so far never increases, so we can simply update $dp[i]$ as the running minimum length.

In the implementation, previous values are no longer needed, so instead of allocating another array, we can use $arr$ itself as the $dp$ array.

#### Complexity Analysis

Let $n$ be the length of the array $arr$.

- Time complexity: $O(n)$.
	We only need to traverse the array once, and hash table insertions and lookups take $O(1)$ time on average.
- Space complexity: $O(n)$.
	The hash table requires $O(n)$ space.

### Approach 2: Sliding Window + Dynamic Programming

#### Intuition

Since all elements in $arr$ are positive, we can use a sliding window instead of prefix sums to find subarrays whose sums equal $target$. Specifically, continuously move the right boundary to the right to expand the window. When the sum inside the window becomes greater than $target$, move the left boundary to the right to shrink the window.

Because shrinking the sliding window requires reading elements from $arr$, unlike in Approach 1, we can no longer overwrite $arr$ in place and must allocate a separate $dp$ array.


#### Complexity Analysis

Let $n$ be the length of the array $arr$.

- Time complexity: $O(n)$.
	Sliding window and dynamic programming are performed during one traversal of $arr$.
- Space complexity: $O(n)$.
	The array $dp$ requires $O(n)$ space.