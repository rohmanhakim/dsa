# LeetCode's Official Editorial

### Approach 1: Sorting + Enumeration

#### Intuition

We first sort the array. Then, for each pair of adjacent elements nums\[i\] and nums\[i+1\], every integer in the range \[nums\[i\]+1,nums\[i+1\]) is missing from the array. We simply enumerate all such values and add them to the answer.


#### Complexity Analysis

Let n be the length of the array nums, and let D be the difference between the maximum and minimum elements in nums.

- Time complexity: O(nlogn+D).

 Sorting the array takes O(nlogn) time. After sorting, enumerating all missing values between adjacent elements takes O(D) time in the worst case. Therefore, the overall time complexity is O(nlogn+D).

- Space complexity: O(1).

 Only a few extra variables are used. The returned array is not included in the space complexity.

---

### Approach 2: Hash Set + Enumeration

#### Intuition

We can store all elements of the array in a hash set. Then, enumerate every integer between the minimum and maximum values in the array. If a value does not exist in the hash set, it is a missing element and should be added to the answer.


#### Complexity Analysis

Let n be the length of the array nums, and let D be the difference between the maximum and minimum elements in nums.

- Time complexity: O(D+n).

 Building the hash set and finding the minimum and maximum elements each require traversing the array, which takes O(n) time overall. Enumerating every integer between the minimum and maximum values takes O(D) time. Therefore, the total time complexity is O(n+D).

- Space complexity: O(n).

 The hash set stores all elements of the array, requiring O(n) extra space.

---
