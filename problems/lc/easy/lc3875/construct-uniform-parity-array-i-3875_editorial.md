# LeetCode's Official Editorial

### Approach: Mathematical Analysis

#### Intuition

This problem can be solved by simply analyzing the parity.

Let the length of the array nums1​ be n, and the task is to construct nums2​ such that all elements have the same parity (either all odd or all even).

Consider the following two cases:

1.  nums1​ is all odd or all even: In this case, simply let nums2​\[i\]\=nums1​\[i\] (i.e., use the first operation for all indices i), and the resulting nums2​ will be the same as the original array, naturally satisfying the condition.
    
2.  nums1​ contains both odd and even numbers: Since even minus odd results in an odd number, we can arbitrarily choose an odd number x from nums1​ and then construct nums2​ as follows:
    
    -   If nums1​\[i\] is odd, then nums2​\[i\]\=nums1​\[i\] (the first operation).
        
    -   If nums1​\[i\] is even, then nums2​\[i\]\=nums1​\[i\]−x (the second operation, choosing j such that nums1​\[j\]\=x).
        

In this way, all elements in nums2​ will be odd, satisfying the requirements of the problem.

In summary, for any input that meets the conditions of the problem, we can always construct a valid nums2​, so the answer is always true.

#### Complexity Analysis

-   Time complexity: O(1).
    
    Directly return the result, independent of the input size.
    
-   Space complexity: O(1).
    
    Use only constant extra space.
    

---

 