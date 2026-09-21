# LeetCode's Official Editorial

### Approach: Dynamic Programming

#### Intuition

We use nums\[i,j\] to denote the subarray \[nums\[i\],nums\[i+1\],…,nums\[j\]\]. In particular, when i\>j, the subarray is defined to be empty.

According to the problem statement, in one operation, we can remove any non-overlapping prefix and suffix of nums while keeping nums non-empty. The remaining elements form a **non-empty subarray** nums\[i,j\].

Furthermore, there is a **one-to-one correspondence** between different subarrays nums\[i,j\] and different valid operations. Specifically, to obtain nums\[i,j\], we can only perform the following operation:

- Remove the prefix nums\[0,i−1\] and the suffix nums\[j+1,n−1\].

Therefore, the problem asks us to count the number of subarrays in nums whose product has a remainder of x when divided by k.

We can solve this problem using dynamic programming.

Define dp\[i\]\[r\] as the number of non-empty subarrays ending at index i whose product of elements has remainder r when divided by k. Initially, all states are 0.

When considering index i, all non-empty subarrays ending at i come from two sources:

- A subarray of length 1 containing only the element nums\[i\];
- A new subarray obtained by appending the element nums\[i\] to the end of any non-empty subarray ending at i−1.

For the first case, the product of the subarray modulo k is

r\=nums\[i\]modk,

so we have

dp\[i\]\[r\]+\=1.

For the second case, suppose the product of a subarray ending at i−1 has remainder r modulo k. After appending the element nums\[i\], the product of the new subarray modulo k becomes

(r×nums\[i\])modk.

Therefore, we obtain the state transition equation:

dp\[i\]\[(r×nums\[i\])modk\]+\=dp\[i−1\]\[r\],0≤r<k.

After completing the state transitions for all indices i, we sum dp\[i\]\[x\] across all i to obtain the answer for each remainder x:

result\[x\]\=∑i\=0n−1​dp\[i\]\[x\].

The space complexity of the dynamic programming method above is O(nk), where n is the length of nums. Since the state transition depends only on the state from the previous layer, we can use a rolling array to optimize the space complexity further to O(k), as shown in the code implementation below.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(nk).

 For each of the n elements in nums, we iterate through all k possible remainders to perform the state transition.

- Space complexity: O(k).

 The rolling array dp (and ndp) requires O(k) auxiliary space.

---