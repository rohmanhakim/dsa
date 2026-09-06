# 2D Dynamic Programming

**Status:** ⏳ 0/5 sessions toward mastery (guided).
**Sessions:** LC 3336 (guided, severity 3), LC 486 (guided, severity 3), LC 1536 (guided, severity 3), LC 115 (guided, severity 3).

## LC 3336 — Count Pairs of Disjoint Subsequences with Equal GCD

Key reframe: each element has 3 choices (seq1, seq2, or unused). State = running GCDs of both subsequences `(g1, g2)`. Since GCD is associative, only the aggregate matters — no need to track which specific elements. Constraints bound the state space: `nums[i] ≤ 200` → GCD values ∈ `[0, 200]` → `O(n × M²)` time, `O(M²)` space with rolling arrays.

Initially tried prefix GCD (range query) which is wrong for combinatorial counting problems.

## LC 1536 — Stone Game V

Interval DP on subarray boundaries. State `dp(l, r)` = max score for `stones[l..r]`. For each split point `i`, compute `leftSum` and `rightSum` using prefix sums. Three transition cases:
- left < right → keep left: `leftSum + dp(l, i)`
- left > right → keep right: `rightSum + dp(i+1, r)`
- equal → Alice chooses: `max(dp(l, i), dp(i+1, r)) + leftSum`

Base case: single stone → 0. Memoized with 2D table. **O(n³)** time, **O(n²)** space.

Initially tried greedy (minimize difference at each split) — fails because locally optimal split doesn't guarantee globally optimal score. Classic "Stone Game" variant → interval DP.

## LC 115 — Distinct Subsequences

Standard 2D DP for counting distinct subsequences. State `dp[i][j]` = number of ways to form `t[0..j-1]` as a subsequence of `s[0..i-1]`.

- **Base case:** `dp[i][0] = 1` (empty t matched by choosing nothing from s)
- **Transition:** `dp[i][j] = dp[i-1][j]` (skip `s[i-1]`); if `s[i-1] == t[j-1]`, add `dp[i-1][j-1]` (extend all prior completions of `t[0..j-2]`)
- **O(m·n)** time, **O(m·n)** space

Key insight: when characters match, the contribution is `dp[i-1][j-1]` (the count of all ways to form the shorter prefix), not literally +1. This multiplicative branching is what makes subsequence counting work.

## Weakness
- **Prefix/range-query misidentification for combinatorial problems:** When the problem says "count pairs of subsequences," reframe as element-wise choices first. Subsequences are not contiguous — prefix structures don't capture the combinatorial nature.
- **Greedy instinct on game problems:** Stone Game / game theory problems with "Alice divides, Bob chooses" structure are interval DP, not greedy. Greedy fails because future game states depend on current split — must explore all possibilities.
- **Index confusion in 2D string DP:** `s[i-1]` vs `t[j-1]` — the two strings have independent index spaces. Writing `t[i-j]` instead of `t[j-1]` is a common off-by-one bug source. Always think of i and j as independent cursors over s and t.
