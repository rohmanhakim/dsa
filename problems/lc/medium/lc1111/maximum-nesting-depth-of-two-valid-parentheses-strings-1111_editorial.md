# LeetCode'sOfficial Editorial

### Approach 1: Bracket Matching Using a Stack

#### Intuition

To split the string into two subsequences that minimize the maximum nesting depth, we first need to understand how to determine the nesting depth of each parenthesis. We can determine it by simulating parenthesis matching with a stack:

Maintain a stack and traverse each character in the parentheses string from left to right:

- If the current character is `(`, push it onto the stack; the nesting depth of this `(` is the current size of the stack.

- If the current character is `)`, the nesting depth of this `)` is the current size of the stack, and then pop a `(` from the stack.

Below is the nesting depth at each character for the bracket sequence `(()(())())`:

| Parentheses | ( | ( | ) | ( | ( | ) | ) | ( | ) | ) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Index** | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
| **Nesting depth** | 1 | 2 | 2 | 2 | 3 | 3 | 2 | 2 | 2 | 1 |

Knowing how to calculate the nesting depth, the solution becomes clear: by distributing the parentheses so that roughly half of the nested depth belongs to sequence A and the other half belongs to sequence B, we minimize the maximum depth of both subsequences to ⌈max\_depth/2⌉. To achieve this balanced distribution, we can assign parentheses at odd nesting depths to one group and those at even nesting depths to the other (using depthmod2). For the example above, parentheses at nesting depths 1 and 3 (forming `(())`) are assigned to one group, while parentheses at nesting depth 2 (forming `()()()`) are assigned to the other.

Furthermore, since the stack only tracks `(`, we do not need an actual stack data structure; an integer variable representing the current stack depth is sufficient.

#### Complexity Analysis

Let n be the length of the string.

- Time complexity: O(n).

 We only need to traverse the input string once.

- Space complexity: O(1).

 Apart from the answer array, we only need a constant number of variables.

---

### Approach 2: Find the Pattern

#### Intuition

We still use the above example `(()(())())`, but here we divide the nesting depth of `(` and `)` into two lines:

| Parentheses | ( | ( | ) | ( | ( | ) | ) | ( | ) | ) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Index** | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
| **Depth of '('** | 1 | 2 | \- | 2 | 3 | \- | \- | 2 | \- | \- |
| **Depth of ')'** | \- | \- | 2 | \- | \- | 3 | 2 | \- | 2 | 1 |

Have you noticed any pattern?

- For an opening parenthesis `(`, the parity of the index is opposite to the parity of the nesting depth:
 - An opening parenthesis at an odd index has an even nesting depth (assigned to group `0`).
 - An opening parenthesis at an even index has an odd nesting depth (assigned to group `1`).
- For a closing parenthesis `)`, the parity of the index is the same as the parity of the nesting depth:
 - A closing parenthesis at an odd index has an odd nesting depth (assigned to group `1`).
 - A closing parenthesis at an even index has an even nesting depth (assigned to group `0`).

This allows us to determine the group of each parenthesis directly from its character type and its index, without tracking stack depth.

Readers interested in the mathematical proof can refer to the section below, or skip directly to the implementation.

#### Proof

**For an opening parenthesis `'('`:** 
Let its 0-based index be x and its nesting depth be y. If there are l opening parentheses and r closing parentheses preceding it in the string:

- According to the definition of nesting depth: 
 y\=l−r+1
- The 0-based index x is the total count of characters preceding it: 
 x\=l+r

Adding the two equations yields x+y\=2l+1, which is always odd. Therefore, the index x and the nesting depth y must have opposite parity.

**For a closing parenthesis `')'`:** 
Let its 0-based index be x and its nesting depth be y. If there are l opening parentheses and r closing parentheses preceding it:

- The nesting depth is: 
 y\=l−r
- The 0-based index is: 
 x\=l+r

Subtracting the two equations yields x−y\=2r, which is always even. Therefore, the index x and the nesting depth y must have the same parity.


#### Complexity Analysis

Let n be the length of the string.

- Time complexity: O(n).

 We only need to traverse the input string once.

- Space complexity: O(1).

 Apart from the answer array, we only need a constant number of variables.

---
