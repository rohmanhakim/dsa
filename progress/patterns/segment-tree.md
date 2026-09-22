# Segment Tree

**Status:** ⏳ 0/5 sessions toward mastery (guided sessions don't count).
**Sessions:** 1 total (LC 3525 guided, severity 3).

## Core Concept

Binary tree for range queries and point updates. Each node stores interval metadata; merge combines children.

**When to use:** Need both point updates AND range queries. If only queries (no updates), use prefix sums.

## Node Structure (LC 3525 variant)

For problems tracking remainder distributions:
- `pre[x]`: count of subarrays starting at left endpoint with product ≡ x (mod k)
- `mul`: product of entire range (mod k)

## Merge Operation

Given left child `[l, m]` and right child `[m+1, r]`:

1. **Subarrays staying in left:** use `left.pre` directly
2. **Subarrays spanning into right:** shift `right.pre` by multiplying with `left.mul`
   ```
   combined_remainder = (left.mul * y) % k
   spanning_pre[combined] += right.pre[y]
   ```
3. **Merged result:**
   ```
   merged.pre[i] = left.pre[i] + spanning_pre[i]
   merged.mul = (left.mul * right.mul) % k
   ```

## Query Decomposition

Standard segment tree query:
1. Collect O(log n) canonical nodes that partition the range
2. Merge them left-to-right with identity node
3. Identity node: `mul = 1` (multiplicative identity), `pre = [0,...,0]`

## Edge Cases

- **Extra leaves:** When `n` (tree size) > `len(nums)`, initialize extra leaves to identity (`pre=zeros`, `mul=1`)
- **Identity node:** Always initialize `pre: make([]int, k)` (not nil)

## Complexity

- Build: O(nk)
- Update: O(k log n)
- Query: O(k log n)
- Space: O(nk)

## Lessons

- **LC 3525:** First segment tree problem. Key insight: k ≤ 5 allows tracking distributions. Merge operation is the core — understanding how spanning subarrays shift remainders by left.mul. Query decomposition requires collecting canonical nodes, not concatenating pre arrays.

## Weakness
- **New pattern** — needs more practice to build fluency with merge operations and query decomposition.
