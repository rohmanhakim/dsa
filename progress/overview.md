# Progress Overview

Compact index. Per-pattern details are in `progress/patterns/<name>.md`.

## Patterns Mastered

_No patterns mastered yet. DB reset 2026-09-12. All sessions since reset are guided._

## Patterns In Progress

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| 1D Dynamic Programming | 0/5 | 3 guided (lc2472, lc1477, lc3524) | Needs independent sessions. Key: rolling DP for subarray counting, recurrence derivation. |
| 2D Dynamic Programming | 0/5 | 2 guided (lc3414, lc1621) | Needs independent sessions. |
| Greedy (Intervals) | 0/5 | 2 guided (lc2472, lc1520) | Needs independent sessions. |
| Binary Search | 0/5 | 1 guided (lc3414, secondary) | Needs independent practice. |
| Math | 0/5 | 1 guided (lc1401) | Needs independent practice. |
| Prefix Sum | 0/5 | 1 guided (lc1621, secondary) | Needs independent practice. |
| Sliding Window | 0/5 | 2 guided (lc1477 secondary, lc1658 primary) | Needs independent practice. Pattern recognition for "reduce to zero" problems. |
| Segment Tree | 0/5 | 1 guided (lc3525) | Needs independent sessions. Key: merge operation shifts right child's pre by left.mul; query collects O(log n) canonical nodes. |

## Top Weaknesses (compact)

1. **Deriving recurrence from invariant** — documented struggles (1D DP)
2. **Over-engineering / overthinking** — Building complex data structures when simpler solutions exist
3. **Brute-force feasibility check** — Not checking constraint bounds first
4. **Debug print left in code** — recurring (fmt.Println)
5. **Pattern misidentification** — Prefix/range-query vs combinatorial DP
6. **Greedy instinct** — Picking greedy when minimax/binary-search is needed
7. **Solving by mechanics vs invariants** — Starting with implementation instead of stating invariant
8. **Binary search indexing** — Confusion between returning index vs count (off-by-one errors)
9. **Tie-breaking logic** — Not considering all cases when selecting final answer
10. **Confusing dp state definition** — Not distinguishing "ending at i" vs "total so far" (lc3524)

→ Full weaknesses: `progress/weaknesses.md`

## Practice Plan

→ `progress/practice-plan.md`

## Constraint-Based Algorithm Selection

| Constraint | Target | Algorithm Family |
|-----------|--------|-----------------|
| n ≤ 10⁴ | O(n²) | Brute force, nested loops |
| n ≤ 10⁵ | O(n log n) | Sorting, binary search, two pointers |
| n ≤ 10⁶ | O(n) | Sliding window, hash map, BFS/DFS |
| n ≤ 10⁷ | O(n) | Linear scan, monotonic stack |
| n ≤ 10⁸ | O(n log n) | Rare, usually requires optimization tricks |
