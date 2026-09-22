# LeetCode's Official Editorial

### Approach: Segment Tree

#### Intuition

Readers are advised to first solve the prerequisite problem, [3524\. Find X Value of Array I](https://leetcode.com/link/?target=https%3A%2F%2Fleetcode.com%2Fproblems%2Ffind-x-value-of-array-i%2F).

In this problem, we need to answer queries queries\[i\]\=\[indexi​,valuei​,starti​,xi​\].

For each query, we perform the following steps:

- Persistently update the element: nums\[indexi​\]:=valuei​.

- Consider the array after removing the prefix nums\[0…starti​−1\] (or removing no elements if starti​\=0).

We need to find the xi​\-value of this resulting array.

The definition of the x\-value is the same as in the prerequisite problem: it is the number of ways to perform one operation such that the product of the remaining elements has remainder x when divided by k.

However, the definition of an operation in this problem differs from that in the prerequisite problem. In the prerequisite problem, we can remove any non-overlapping prefix and suffix in one operation, whereas in this problem, we can only remove one suffix.

Note that we need the xi​\-value of the array after the changes. The prefix nums\[0…starti​−1\] has already been removed before we perform the operation, after which we can perform one operation to remove a suffix. Therefore, for each query, the remaining part after the operation must be a **non-empty subarray** nums\[starti​…j\] of the updated array, where starti​≤j≤n−1.

Furthermore, there is a one-to-one correspondence between different subarrays nums\[starti​…j\] and valid operations. Specifically, to obtain the subarray nums\[starti​…j\], we remove the suffix nums\[j+1…n−1\], where n is the length of nums.

Therefore, each query can be broken down into two tasks:

1. Perform a point update on nums\[indexi​\].
2. Among all non-empty subarrays nums\[starti​…j\] (j≥starti​), count the number of subarrays whose product of elements has remainder xi​ when divided by k.

The dynamic programming solution to the prerequisite problem requires traversing the entire array nums. Therefore, running the dynamic programming solution anew for every query would result in an O(nq) time complexity, which will exceed the time limit (TLE).

Note that there are only k possible remainders of the product of a subarray's elements when divided by k, and the range of k in this problem is still very small. We consider maintaining the statistics corresponding to different remainders in each node of a segment tree and merging intervals by enumerating the remainders. For convenience, unless otherwise specified, the remainder below refers to the remainder obtained after dividing the product of the elements by k.

Specifically, for the interval represented by each node in the segment tree, we maintain an array of length k:

pre\[x\]: the number of **non-empty** subarrays starting at the left endpoint of this interval whose product of elements has remainder x.

Suppose the interval corresponding to the current node is \[l,r\], and its left and right child nodes correspond to the intervals \[l,m\] and \[m+1,r\], respectively.

We now discuss how to merge the information from the child nodes to obtain the information for the current node.

For pre, consider a subarray whose left endpoint is l. There are two cases:

**1\. It lies entirely within the interval corresponding to the left child node**

\[l,j\]l≤j≤m

In this case, the number of subarrays can be obtained directly from the pre of the left child node.

**2\. It contains the entire interval corresponding to the left child node and continues to extend rightward into a partial interval of the right child node**

\[l,j\]\=\[l,m\]+\[m+1,j\]m+1≤j≤r

In this case, the subarray is formed by concatenating two parts: the complete interval \[l,m\] represented by the left child node and the interval \[m+1,j\] starting at the left endpoint m+1 of the interval represented by the right child node.

Let the remainder of the product of the elements in the entire interval of the left child node be mulL​. If the remainder of the product of a prefix subarray of the right child node is x, then the remainder of the product after merging is

(mulL​×x)modk

Therefore, we only need to enumerate all prefix remainders of the right child node to complete the transition. The time complexity for merging the pre arrays to obtain a single node is O(k). In addition, this transition requires the remainder of the product of the elements in the entire interval, which we maintain separately in the segment tree.

Using the method above, the segment tree allows us to query the prefix remainder counts for any range.

For each query, we perform a range query on the segment tree for the interval \[starti​,n−1\]. This decomposes the interval into O(logn) canonical nodes, combines their pre arrays via interval merging, and returns the merged pre distribution. The answer to the query is then pre\[xi​\].

In the sample code below, the length of the pre array is set to k+1, and the last element represents the remainder mul of the product of the elements in the entire interval.

> Follow-up:
> 
> In this problem, for each query, we need to answer the number of non-empty subarrays whose left endpoint is starti​ and whose product of elements has remainder xi​. Here, the left endpoints of the subarrays included in the count are fixed.
> 
> If neither the left nor the right endpoint of the subarrays to be counted is fixed, can the code for this problem directly solve the problem? In other words, if we need to answer the number of all subarrays **inside** the subarray nums\[starti​…endi​\] whose product of elements has remainder xi​, how should we solve it?

#### Complexity Analysis

Let n be the length of nums and q be the number of queries.

- Time complexity: O((n+q)klogn).

 Building the segment tree takes O(nk) time. Each of the q queries performs a point update and a range query, each taking O(klogn) time. This yields an overall time complexity of O((n+q)klogn).

- Space complexity: O(nk).

 The segment tree contains O(n) nodes, each storing an array of length k+1.

---
