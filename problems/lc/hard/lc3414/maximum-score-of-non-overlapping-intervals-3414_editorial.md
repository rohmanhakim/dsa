# LeetCode's Official Editorial

### Approach: Dynamic Programming + Binary Search

#### Intuition

This problem is closely related to [1751\. Maximum Number of Events That Can Be Attended II](https://leetcode.com/link/?target=https%3A%2F%2Fleetcode.com%2Fproblems%2Fmaximum-number-of-events-that-can-be-attended-ii%2Fdescription%2F). The key differences are that the maximum number of intervals is fixed at 4 (equivalent to k\=4 in Problem 1751), and among all valid selections that achieve the maximum score, we must return the one with the lexicographically smallest list of indices.

The score for selecting interval i is weighti​. We need to select at most 4 **non-overlapping** intervals such that their total score is **maximized** and their index list is **lexicographically smallest**. Because intervals cannot overlap and we can choose at most 4 of them, this problem is a variation of **weighted interval scheduling** (or 0/1 knapsack with non-overlapping constraints), which can be solved using dynamic programming.

First, pair each interval with its original index i, then sort all intervals in intervals by their right endpoints in ascending order. Define dp\[i\]\[j\] as the maximum total score obtained by selecting **at most** j non-overlapping intervals among the first i intervals, where i∈\[0,n\] and j∈\[0,4\]. Concurrently, let indices\[i\]\[j\] store the corresponding sorted list of original indices, allowing us to resolve ties via lexicographical comparison when two choices produce equal scores. We obtain the following transitions:

- If interval i is not selected, the maximum score obtained by selecting j **non-overlapping** intervals among the first i intervals equals the maximum score obtained by selecting j **non-overlapping** intervals among the first i−1 intervals. Thus,

 dp\[i\]\[j\]\=dp\[i−1\]\[j\]

- If interval i is selected, any preceding intervals must end strictly before li​ to avoid overlapping. Because the intervals are sorted by their right endpoints in ascending order, all candidate intervals ending strictly before li​ form a prefix of the sorted array. Let p be the count of such intervals. From these, we can select at most j−1 non-overlapping intervals. Thus:

 dp\[i\]\[j\]\=dp\[p\]\[j−1\]+weighti​

- Combining the two cases, we obtain

 dp\[i\]\[j\]\=max(dp\[i−1\]\[j\], dp\[p\]\[j−1\]+weighti​)

Since all intervals are sorted by their right endpoints, given the left endpoint li​, binary search can find the number of intervals p whose right endpoints are smaller than li​. We enumerate (i,j) in increasing order. While computing the state dp\[i\]\[j\], we also record the currently selected interval indices in indices\[i\]\[j\]. Specifically:

- Maintain a sorted index list for each state.
- When the scores of both choices are equal, compare the two index lists lexicographically and select the smaller one.

After the enumeration is complete, the maximum score is dp\[n\]\[4\], and the optimal index array is indices\[n\]\[4\].

#### Complexity Analysis

Let n be the length of the array intervals, and let k\=4.

- Time complexity: O(nlogn+nk2).

 Sorting the intervals takes O(nlogn) time. For each interval, finding the non-overlapping prefix via binary search takes O(logn) time, totaling O(nlogn) across all intervals. There are O(nk) dynamic programming states, and maintaining each index list takes O(k) time. Therefore, the total time complexity is O(nlogn+nk2).

- Space complexity: O(nk2).

 An array of length O(n) is required to store the sorted intervals. There are O(nk) dynamic programming states, and storing the index list for each state requires O(k) space. Therefore, the total space complexity is O(nk2).

---
