# LEetCode's Official Editorial

### Approach: Classification Discussion

#### Intuition

The XOR operation has the following three properties:

> 1. XORing any number with 0 results in the original number, i.e., a⊕0\=a.
> 2. XORing any number with itself results in 0, i.e., a⊕a\=0.
> 3. The XOR operation satisfies the commutative and associative laws, i.e., a⊕b\=b⊕a and a⊕(b⊕c)\=(a⊕b)⊕c.
> 
> Here, ⊕ denotes the XOR operator.

Let n be the length of the array nums. Clearly, the answer cannot exceed n.

Let the bitwise XOR of all elements in the array be totalXor.

**Case 1:** totalXor\=0

In this case, the entire array nums itself is a non-empty subsequence whose bitwise XOR is non-zero. Therefore, the answer is n.

**Case 2:** totalXor\=0 and all elements in nums are 0

In this case, the bitwise XOR of any non-empty subsequence is 0. Therefore, there is no non-empty subsequence whose bitwise XOR is non-zero, so the answer is 0.

**Case 3:** totalXor\=0 and there exists at least one element x\=0 in nums

Since totalXor\=0 and x\=0, there must be at least one other element in nums.

Let the bitwise XOR of the remaining n−1 elements be subXor. Then:

x⊕subXor\=totalXor\=0

XORing both sides with x, we get:

x⊕subXor⊕x\=0⊕x

Using the properties of the XOR operation, we obtain:

subXor\=x\=0

Therefore, after removing the element x, the remaining n−1 elements form a subsequence whose bitwise XOR is non-zero. Hence, the answer is n−1.

#### Complexity Analysis

Let n be the length of the array nums.

- Time complexity: O(n).

- Space complexity: O(1).

---