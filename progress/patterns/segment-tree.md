# Segment Tree

**Status:** ⏳ 0/5 sessions toward mastery (guided).
**Sessions:** LC 2213 (guided, severity 3).

Binary tree for range queries and point updates. Each node stores interval metadata; merge combines children. Key template for problems requiring O(log n) point updates + O(1) range queries on dynamic data.

## LC 2213 — Longest Substring of One Repeating Character

Segment tree storing (pre, suf, maxLen, leftChar, rightChar, iLen) per node. Merge rule: pre extends across boundary only if left child is fully uniform AND boundary chars match; suf extends symmetrically; maxLen = max(left.maxLen, right.maxLen, left.suf+right.pre if chars match).

Three bugs caught: (1) unconditionally adding boundary run to maxLen regardless of char match, (2) not checking interval coverage for pre/suf extension, (3) Go zero-value bool is false — padding leaves need explicit isEmpty=true. Initially attempted interval/run tracking with Go slice (editorial Approach 2) which was O(n) per query due to slice shifting.

## Key Lesson

The merge function is the entire algorithm. Before coding, trace through 2-3 merge examples on paper to verify pre/suf/maxLen propagation. For segment trees with padding (non-power-of-2 lengths), always explicitly mark padding nodes.