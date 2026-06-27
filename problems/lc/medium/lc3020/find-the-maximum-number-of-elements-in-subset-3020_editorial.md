Approach: Hash Table + Enumeration
Intuition

The problem asks us to find the longest possible mountain array that can be formed from the given array. Such an array first increases and then decreases. During the increasing phase, each element is the square of the previous one, while during the decreasing phase, each element is the square root of the previous one.

Since the values grow extremely quickly, the maximum possible length of such a mountain array is very small within the given constraints. Let the maximum value in the array be M, and let the first element of the mountain array be `x`. Then the `(p+1)`-th element is `x^(2^p)`. When `x≥2`, we must have

```
x^(2^p) ≤ M ,
```

which implies

```
p ≤ log2 * ​logx * ​M
```

Therefore, the height of the mountain is bounded by `O(loglogM)`, and under the given constraints, `p` is at most about `4`.

This observation allows us to enumerate each distinct value in the array as the starting value `x` of a mountain array. We then repeatedly check whether `x^2`, `x^4`, `x^8`, and so on appear at least twice in the array, since one occurrence is needed on the increasing side and the other on the decreasing side.

Because a mountain array contains exactly one peak element, its length is always odd. For every value below the peak, we count pairs of identical elements. After constructing all possible pairs, if the peak value still exists in the array, we can place one occurrence at the peak and increase the total length by 1. Otherwise, one element from the last pair must be used as the peak, reducing the total length by 1.

We must also handle the special case where all elements in the mountain array are equal to 1.

Complexity Analysis

Let n be the length of `nums`, and let `M` be the maximum value in `nums`.

    Time complexity: `O(nloglogM)`.

    Building the hash table takes `O(n)` time. For each distinct starting value, we repeatedly square the current value, and the number of iterations is at most `O(loglogM)`.

    Space complexity: `O(n)`.

    The hash table stores the frequency of each element.
