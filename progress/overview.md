# Progress Overview

Compact index. Per-pattern details are in `progress/patterns/<name>.md`.

## Patterns Mastered

| Pattern | Streak | Distinct Days | Key Speed |
|---------|--------|---------------|-----------|
| DFS | 9/9 sev≤1 | 6+ | 10:24–86:56 |
| Binary Search | 8/8 sev≤1 | 5+ | 5:40–90:04 |
| Greedy (Intervals) | 5/5 sev≤1 | 5 | 7:28–74:17 |
| Hash Map | 4/4 sev≤1 | 4+ | 3:47–15:13 |

## Patterns Near Mastery

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| Stack | 4/4 sev≤1 | 4 total (0 guided) | 1 more sev≤1 pass across 3+ days (already has 4 distinct days) |
| Sliding Window | 4/4 sev≤1 | 12 total (3 guided) | 1 more sev≤1 pass (already has 3 distinct days) |
| Prefix Sum | 3/3 sev≤1 | 10 total (4 independent) | 2 more consecutive sev≤1 passes |
| Two Pointers | 2/2 sev≤1 | 7 total (2 guided) | 3 more consecutive sev≤1 passes |
| Monotonic Stack | 2/2 sev≤1 | 4 total (2 guided) | 3 more consecutive sev≤1 passes |

## Patterns In Progress

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| Math | 0/5 sev≤1 | 16 total (5 guided) | Last 3 sev≤1 then sev 3 broke streak. Needs 5 consecutive independent sev≤1. |
| Digit Manipulation | 0/5 sev≤1 | 7 total (2 guided) | Streak broken by guided session. Needs fresh 5 consecutive sev≤1. |
| Greedy (Observation) | 0/5 sev≤1 | 15 total (9 guided) | Streak broken by guided sessions. Needs independent practice. |
| BFS | 0/5 sev≤1 | 5 total (2 guided) | Streak broken by guided session (lc3568). Needs 5 independent passes. |
| Linked List | 1/1 sev≤1 | 1 session | 4 more |
| Simulation | 1/1 sev≤1 | 1 session | 4 more |
| 1D Dynamic Programming | 0/5 | 15 total (13 guided) | Needs independent sessions. Key: recurrence derivation. |
| 2D Dynamic Programming | 0/5 | 4 guided | Needs independent sessions. |
| Minimax DP | 0/5 | 4 guided | Needs independent sessions. |
| BST | 0/5 | 2 guided | Needs independent sessions. |
| Segment Tree | 0/5 | 1 guided | Needs independent sessions. |
| Binary Search on Answer | 0/5 | 1 guided | Needs independent sessions. |
| Binary Lifting | 0/5 | 1 guided | Needs independent sessions. |

## Top Weaknesses (compact)

1. **Deriving recurrence from invariant** — 9+ documented struggles (1D DP)
2. **Over-engineering / overthinking** — Building complex data structures when simpler solutions exist
3. **Brute-force feasibility check** — Not checking constraint bounds first
4. **Debug print left in code** — 5 occurrences (fmt.Println)
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
