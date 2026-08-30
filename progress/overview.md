# Progress Overview

Compact index. Per-pattern details are in `progress/patterns/<name>.md`.

## Patterns Mastered

| Pattern | Streak | Distinct Days | Key Speed |
|---------|--------|---------------|-----------|
| Greedy (Intervals) | 5/5 sev≤1 | 5 | Varies |
| Binary Search | 8/8 sev≤1 | 5+ | 5:40–90:04 |
| DFS | 9/9 sev≤1 | 6+ | 10:24–86:56 |
| Stack | 5/5 sev≤1 | 4+ | 12:59–43:40 |
| Hash Map | 5/5 sev≤1 | 5 | Varies |

## Patterns In Progress

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| Sliding Window | 3/5 within time | 12 total | 6 consecutive sev<=1 but 2 of last 5 over time; improving |
| Math | 0/5 sev≤1 | 13 total (4 guided) | Streak broken (LC 1927 guided). Needs 5 independent passes |
| Greedy (Observation) | 1/5 sev≤1 | 11 total (6 guided) | Streak restarted (LC 2091 ind sev 0). Needs 4 more consecutive independent sev≤1. |
| BFS | 3/5 sev≤1 | 5 total | 2 more |
| Two Pointers | 2/5 sev≤1 | 4 total | 3 more |
| Digit Manipulation | 2/5 sev≤1 | 5 total | 3 more |
| Monotonic Stack | 2/5 sev≤1 | 4 total | 3 more |
| Prefix Sum | 1/5 sev≤1 | 5 total | 4 more independent |
| Binary Lifting | 1/5 | 1 guided | Needs independent practice |
| 1D Dynamic Programming | 0/5 | 11 total (10 guided) | Needs independent sessions |
| 2D Dynamic Programming | 0/5 | 2 guided | Needs independent sessions |
| Minimax DP | 0/5 | 5 guided | Needs independent sessions |
| BST | 0/5 | 2 guided | Needs independent sessions |
| Segment Tree | 0/5 | 1 guided | Needs independent sessions |
| Binary Search on Answer | 0/5 | 1 guided | Needs independent sessions |
| DAG Topo Sort + DP | 0/5 | 1 guided | Needs independent sessions |
| Backward Two-Pointer Merge | 0/5 | 1 guided | Needs independent sessions |
| Two Pointers (Fast/Slow) | 0/5 | 1 guided | Needs independent sessions |
| Simulation | 1/5 sev≤1 | 1 total | 4 more |

## Top Weaknesses (compact)

1. **Deriving recurrence from invariant** — 8+ documented struggles (1D DP)
2. **Over-engineering / overthinking** — Building complex data structures when simpler solutions exist
3. **Brute-force feasibility check** — Not checking constraint bounds first
4. **Debug print left in code** — 5 occurrences (fmt.Println)
5. **Pattern misidentification** — Prefix/range-query vs combinatorial DP
6. **Greedy instinct** — Picking greedy when minimax/binary-search is needed
7. **Solving by mechanics vs invariants** — Starting with implementation instead of stating invariant

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