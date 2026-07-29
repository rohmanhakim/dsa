# LeetCode's Official Editorial

### Approach: Combinatorial Mathematics + Trial and Error Method

#### Intuition

This problem is an extension of [3517\. Smallest Palindromic Rearrangement I](https://leetcode.com/problems/smallest-palindromic-rearrangement-i/), and it is recommended to solve that problem first.

Here, we are asked to find the k\-th lexicographically smallest palindromic permutation. Compared with the previous problem, the only difference is that we now need to construct an arbitrary lexicographical rank instead of only the smallest one.

From the previous problem, we know that the symmetry of a palindrome implies that the multisets of characters on the left and right halves are completely determined once one half is fixed. Therefore, we do not need to construct the entire palindrome directly. Instead, we only need to consider the permutations of the left half of the string. The proof of this property can be found in the editorial for the previous problem.

Thus, the problem is transformed into the following:

> Given a multiset of characters, find its k\-th lexicographically smallest permutation.

We solve this using a trial-and-error approach. At each position, we try every available character in lexicographical order.

Suppose the prefix constructed so far is prefix, and we are currently trying character c. Let L and R denote the lexicographical ranks of the smallest and largest permutations that begin with prefix∥c, respectively.

There are three cases:

- If R≥k, then the desired permutation lies within this interval. Since characters are tried in lexicographical order, the current character must be part of the answer. We append c to the prefix and continue constructing the next position.

- If R<k, then all permutations beginning with this prefix come before the desired one. We restore the character to the multiset, continue trying the next candidate character, and update L\=R+1.

- If every possible character has been tried and still R<k, then fewer than k permutations exist, so we return an empty string.

Initially, L\=1. The remaining question is how to compute R once the current prefix is fixed.

The Cantor expansion algorithm computes the lexicographical rank of a given permutation. Here, we essentially perform the inverse process. Since Cantor expansion is based on combinatorial counting, we use combinatorial mathematics to count how many permutations can be formed after fixing the current prefix.

Suppose the remaining number of positions is rem after fixing both the current prefix and the trial character c. Let the remaining multiset be W.

Assume the distinct characters in W are c1​,c2​,…,cσ​, with frequencies cnt\[ci​\], where

∑cnt\[ci​\]\=rem.

The number of distinct permutations of these remaining characters is

P\=cnt\[c1​\]!×cnt\[c2​\]!×⋯×cnt\[cσ​\]!rem!​,

or equivalently,

P\=∏c∈W​cnt\[c\]!rem!​.

Therefore, 
R\=L+P−1.

Unfortunately, directly computing factorials is not practical because the problem does not permit arbitrary-precision arithmetic or modular inverses. Instead, we rewrite the multinomial coefficient as a product of ordinary combinations.

1. Among the rem remaining positions, choose cnt\[c1​\] positions for character c1​. This contributes

Cremcnt\[c1​\]​.

2. After placing all copies of c1​, choose cnt\[c2​\] positions from the remaining positions. This contributes

Crem−cnt\[c1​\]cnt\[c2​\]​.

3. Continue in this manner until every character has been placed.

By the multiplication principle,

P\=∏i\=1σ​Crem−∑j\=1i−1​cnt\[cj​\]cnt\[ci​\]​.

We still cannot compute combinations using factorials directly. Instead, we evaluate each combination using its multiplicative formula:

Cnm​\=∏i\=1m​in−i+1​.

First, we exploit the symmetry of combinations by replacing m with 
m\=min(m,n−m),

thereby reducing the number of iterations.

During the computation, let resi​ denote the value after the i\-th iteration. Then

res∗i\=res∗i−1×in−i+1​.

Since the product of any i consecutive integers is divisible by i!, we can safely multiply first and divide 
afterward, ensuring that every intermediate result remains an integer.

Moreover, we never need an exact value once it exceeds k. Therefore, we terminate the computation immediately whenever resi​\>k, which avoids overflow while remaining sufficient for our comparison.

Using this combinatorial counting procedure together with the trial-and-error construction described above, we can determine each character of the left half and finally reconstruct the palindrome.

##### Time Complexity Analysis of the Permutation Count

First, consider the complexity of computing a single combination.

Each computation either terminates early or performs at most cnt\[ci​\] iterations. Suppose the loop executes x iterations, where x≤m. Using the multiplicative expansion,

(xn​)\=x(x−1)⋯1n(n−1)⋯(n−x+1)​\=∏j\=0x−1​x−jn−j​.

Since n≥2m≥2x, every term satisfies

x−jn−j​≥x−j2x−j​\=1+x−jx​≥2.

Therefore, the intermediate result at least doubles in every iteration. Because the computation stops as soon as the value exceeds k, the number of iterations is bounded by O(logk).

Hence, the complexity of computing a single combination is

O(min(cnt\[ci​\],logk)).

Next, consider computing the entire permutation count P.

Suppose the combination computations execute x1​,x2​,…,xp​ iterations, respectively. Their product grows by at least

2x1​×2x2​×⋯×2xp​\=2∑xj​.

Since the computation stops once the product exceeds k,

2∑xj​≤k,

which implies

∑xj​≤logk.

Therefore, the total work of all inner combination computations is O(logk).

On the other hand, the total number of iterations is also bounded by the number of remaining characters, namely 
∑xj​≤rem,

and rem is O(n).

Combining these two bounds, the total number of iterations performed by all combination computations is 
O(min(n,logk)).

Finally, computing P also requires scanning the entire character set once, contributing an additional O(σ) cost.

Therefore, the total time required to compute P is

O(σ+min(n,logk)).

#### Complexity Analysis

Let n be the length of the palindrome string s, and let σ denote the size of the character set, which is 26 in this problem.

- Time complexity: O(n⋅σ⋅(σ+min(n,logk))).

 For each position in the left half, we enumerate all candidate characters, which costs O(σ). For every candidate, computing the number of remaining permutations takes O(σ+min(n,logk)). Therefore, the total time complexity is 
 O(n⋅σ⋅(σ+min(n,logk))). 
 The analysis of the permutation-counting procedure is given above.

- Space complexity: O(1) or O(n).

 The frequency array bucket occupies O(σ) space, which is constant. The overall space complexity depends on whether the palindrome is constructed in place or by creating additional strings.

---