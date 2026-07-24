# LeetCode's Official Editorial

### Approach 1: Enumeration

#### Intuition

Observe that the XOR operation never increases the number of bits required to represent a value. Let m be the maximum element in nums, and let U be the smallest power of 2 that is greater than m. Then the XOR of any two or three elements must be less than U.

We can find m with a single traversal of the array, compute U, and use a boolean array of size U to record all possible XOR values.

First, enumerate all possible XOR values of two elements (including the case where the same element is chosen twice). Then, XOR each of these values with every element in nums to obtain all possible XOR values of three elements.

The algorithm proceeds as follows:

1. Enumerate all pairs satisfying i≤j, compute nums\[i\]⊕nums\[j\], and record the result in set S.
2. For each value x in S, traverse every element v in nums, compute x⊕v, and record the result in set T.
3. The number of distinct values in T is the answer.

Since the value range is bounded, we can replace the hash sets with boolean arrays to reduce the constant factor.

#### Complexity Analysis

Let n be the length of the array, and let m be the maximum element in the array.

- Time complexity: O(n2+nm).

 The first double loop enumerates all XOR values of two elements in O(n2) time. The second phase enumerates the XOR of each two-element XOR value with every array element, resulting in O(nm) time.

- Space complexity: O(m).

 Two boolean arrays of size O(m) are used.

---

### Approach 2: Enumeration (Optimization)

#### Intuition

Instead of explicitly enumerating all pairs, we can build the set of attainable XOR values incrementally, similar to dynamic programming.

Define three sets:

- one stores all XOR values obtainable from a single element (that is, the element itself).
- two stores all XOR values obtainable from two elements (repetition is allowed).
- three stores all XOR values obtainable from three elements (repetition is allowed).

**Stage 1:** Construct one and two.

Traverse the array nums. For each element v:

- Add v to one.
- For every value x currently in one, add x⊕v to two.

**Stage 2:** Construct three.

At this point, two already contains every XOR value obtainable from two elements. Traverse nums again, and for each element v and every value x in two, add x⊕v to three.

Finally, the number of distinct values in three is the answer.

As in the previous approach, since the value range is bounded, boolean arrays can be used instead of sets.


#### Complexity Analysis

Let n be the length of the array, and let m be the maximum element in the array.

- Time complexity: O(nm).

 For each array element, we traverse the entire value range once to perform the state transitions.

- Space complexity: O(m).

 Three boolean arrays of size O(m) are used.
