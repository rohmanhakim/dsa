# LeetCode's Official Editorial

### Approach: Dynamic Programming

#### Intuition

First, the properties of parentheses strings and grid paths yield the following observations:

- In an n×m grid, any path from (0,0) to (n−1,m−1) moving only right or down has a fixed length of (n−1)+(m−1)+1\=n+m−1. Since a valid parentheses string must have an even length, if n+m−1 is odd, no valid path can exist.
- A valid parentheses string must start with ( and end with ). Therefore, grid\[0\]\[0\] must be ( and grid\[n−1\]\[m−1\] must be ).
- While traversing a parentheses string from left to right, let l1​ be the number of ( characters encountered and l2​ be the number of ) characters encountered. At every point, we must have l1​≥l2​ (the count of unmatched opening parentheses is non-negative). Furthermore, upon reaching the end, we must have l1​\=l2​ (the count of unmatched opening parentheses is exactly 0).

We can solve the problem using dynamic programming. Let dp\[i\]\[j\] store the set of all possible numbers of unmatched opening parentheses among all paths from (0,0) to cell (i,j). Consider transitions from the valid predecessor cells (i−1,j) and (i,j−1):

- If grid\[i\]\[j\]\=(, increase every value in the set by 1.
- If grid\[i\]\[j\]\=), decrease every value in the set by 1.
- If decreasing makes a value negative, it represents an invalid prefix state and is discarded from the set.

The transition equation is:

dp\[i\]\[j\]\={x+Δi,j​​x∈(dp\[i−1\]\[j\]∪dp\[i\]\[j−1\]),x+Δi,j​≥0}

where:

Δi,j​\={1,−1,​if grid\[i\]\[j\]\=(,if grid\[i\]\[j\]\=).​

For initialization, since grid\[0\]\[0\] must be (, we initialize dp\[0\]\[0\]\={1}.

Finally, a valid path exists if and only if 0∈dp\[n−1\]\[m−1\].

**State Transition Implementation Details**

Implementing these transitions directly requires a container to store reachable balance values:

- **Hash Set**: A set can store valid values for each cell. Since each step modifies all values and prunes invalid ones, constructing a new set per cell takes O(n+m) time.
- **Boolean Array**: Since each step changes the balance by ±1, the balance cannot exceed n+m−1. We can maintain a boolean array where index k is true if balance k is reachable. Each transition also takes O(n+m) time.
- **Bitset / Big Integer**: We can represent the set as a bitmask where the k\-th bit is set if balance k is reachable. Adding 1 corresponds to a left bit shift (`<< 1`), while subtracting 1 corresponds to a right bit shift (`>> 1`). The natural truncation during a right shift automatically discards negative balances (as bit 0 shifts out). This reduces each transition to O(Wn+m​) operations, where W is the machine word size (typically 64).

#### Complexity Analysis

Let W be the machine word size.

- Time complexity: O(Wnm(n+m)​) or O(nm(n+m)).

 Suppose the input grid is an n×m matrix. There are n⋅m cells, and each cell has at most n+m states. The time complexity of these transitions depends on the implementation; see the State Transition Implementation Details section above.

- Space complexity: O(nm(n+m)).

 This is the space cost for the dynamic programming array.

---
