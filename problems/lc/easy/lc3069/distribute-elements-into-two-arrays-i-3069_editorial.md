# LeetCode's Official Editorial

### Preface

The original problem statement uses array indices starting from 1. To maintain consistency with the indexing conventions of most programming languages, this solution converts the array index range to 0 through n−1. That is, the original nums\[1\] and nums\[2\] correspond to nums\[0\] and nums\[1\] in this solution, respectively. All subsequent solutions will consistently use the index range 0 through n−1.

### Approach 1: Two List Simulation

#### Intuition

We can simulate the process described in the problem statement. Initialize two lists, arr1​ and arr2​, and add nums\[0\] and nums\[1\] to them, respectively.

Then, starting from i\=2, traverse the array nums and compare the last elements of arr1​ and arr2​:

- If the last element of arr1​ is greater than the last element of arr2​, append nums\[i\] to arr1​.

- Otherwise, append nums\[i\] to arr2​.

After traversing the array, concatenate arr2​ to the end of arr1​ and return the resulting array.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n).

 We traverse the array once, with each comparison and insertion taking O(1) time. Concatenating the two lists also takes O(n) time.

- Space complexity: O(n).

 The two lists together contain all n elements of the array.

---

### Approach 2: Single Array with Two Pointers

#### Intuition

We can perform the allocation using only one array arr. Place nums\[0\] at the beginning of arr, corresponding to the first element of arr1​, and place nums\[1\] at the end of arr, corresponding to the first element of arr2​.

We use two pointers, idx and revIdx, to represent the positions of the last elements of arr1​ and arr2​, respectively.

We then iterate over the remaining elements of nums and compare arr\[idx\] with arr\[revIdx\]:

- If arr\[idx\]\>arr\[revIdx\], place the current element at the next position of idx, namely idx+1, and then move idx one position to the right.

- Otherwise, place the current element at the previous position of revIdx, namely revIdx−1, and then move revIdx one position to the left.

After the traversal, arr\[0..idx\] contains arr1​, while arr\[revIdx..n−1\] contains arr2​. Since the elements of arr2​ are stored from right to left, their order is reversed. Therefore, we reverse arr\[revIdx..n−1\] to restore the correct order.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n).

 We traverse the array in O(n) time and reverse part of the array in O(n) time. Therefore, the overall time complexity is O(n).

- Space complexity: O(1).

 Apart from the output array itself, we only use a constant number of additional variables, so the extra space complexity is O(1).

---
