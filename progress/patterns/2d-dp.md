# 2D Dynamic Programming

**Status:** ⏳ 0/5 sessions toward mastery (guided).
**Sessions:** LC 3336 (guided, severity 3).

## LC 3336 — Count Pairs of Disjoint Subsequences with Equal GCD

Key reframe: each element has 3 choices (seq1, seq2, or unused). State = running GCDs of both subsequences `(g1, g2)`. Since GCD is associative, only the aggregate matters — no need to track which specific elements. Constraints bound the state space: `nums[i] ≤ 200` → GCD values ∈ `[0, 200]` → `O(n × M²)` time, `O(M²)` space with rolling arrays.

Initially tried prefix GCD (range query) which is wrong for combinatorial counting problems.

## Weakness
- **Prefix/range-query misidentification for combinatorial problems:** When the problem says "count pairs of subsequences," reframe as element-wise choices first. Subsequences are not contiguous — prefix structures don't capture the combinatorial nature.