# Progress Overview

Compact index. Per-pattern details are in `progress/patterns/<name>.md`.

## Patterns Mastered

| Pattern | Streak | Distinct Days | Key Speed |
|---------|--------|---------------|-----------|
| Stack | 9/9 sev≤1 | 9 | 8:05–43:40 |

## Patterns Near Mastery

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| Two Pointers | 2/2 sev≤1 | 7 total (2 guided) | 3 more consecutive sev≤1 passes |
| Monotonic Stack | 2/2 sev≤1 | 4 total (2 guided) | 3 more consecutive sev≤1 passes |

## Patterns In Progress

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| DFS | 0/5 sev≤1 | 14 total (3 guided) | REGRESSED 2026-10-02 (sev 3, lc22). Re-cert: 3 sev≤1 across 2 days. |
| Binary Search | 0/5 sev≤1 (re-cert) | 12 total (4 guided) | REGRESSED 2026-09-12 (sev 3 secondary, lc3414; strict multi-pattern rule, backfilled 2026-10-08). Re-cert: 3 sev≤1 across 2 days. |
| Greedy (Intervals) | 0/5 sev≤1 (re-cert) | 7 total (2 guided) | REGRESSED 2026-09-19 (sev 3, lc1520 + lc2472; backfilled 2026-10-08). Re-cert: 3 sev≤1 across 2 days. |
| Math | 0/5 sev≤1 | 19 total (6 guided) | sev≤1 pair (lc3870, lc3871) wiped by guided lc1401 (sev 3) 6 days later. Needs 5 consecutive independent sev≤1. |
| Digit Manipulation | 2/5 sev≤1 | 8 total (2 guided) | Rising: lc3550 + lc3622 sev≤1; chain stopped by lc3348 (sev 3). 3 more needed. |
| Greedy (Observation) | 0/5 sev≤1 | 17 total (11 guided) | Streak broken by guided sessions. Needs independent practice. LC678 (guided, sev 3): wildcard range invariant learned. |
| BFS | 0/5 sev≤1 | 5 total (2 guided) | Streak broken by guided session (lc3568). Needs 5 independent passes. |
| Hash Map | 0/3 sev≤1 (re-cert) | 8 total (2 guided) | REGRESSED 2026-10-07 (sev 3, hr-icecream-parlor). Re-cert: 3 sev≤1 across 2 days. Was 4/4 — one shy of the bar. |
| Linked List | 1/1 sev≤1 | 1 session | 4 more |
| Simulation | 0/5 sev≤1 | 6 total (2 guided) | lc1096 (guided, sev 3) anchors streak at 0. Key: recursive grammar parsing with precedence (concatenation > union). |
| Sliding Window | 0/5 sev≤1 | 14 total (5 guided) | Streak broken 2026-09-23 (sev 3, lc1658; also lc1477 secondary sev 3). Was 4/4 near mastery — needs fresh 5. |
| Prefix Sum | 0/5 sev≤1 | 11 total (7 guided) | Streak broken by lc1621 (sev 3 secondary, strict multi-pattern rule). Was 3/3 — needs fresh 5. |
| 1D Dynamic Programming | 0/5 | 19 total (17 guided) | Needs independent sessions. Key: recurrence derivation. |
| 2D Dynamic Programming | 0/5 | 7 total (7 guided) | Needs independent sessions. |
| Minimax DP | 0/5 | 4 guided | Needs independent sessions. |
| BST | 0/5 | 2 guided | Needs independent sessions. |
| Segment Tree | 0/5 | 2 total (2 guided) | Needs independent sessions. |
| Binary Search on Answer | 0/5 | 1 guided | Needs independent sessions. |
| Binary Lifting | 0/5 | 1 guided | Needs independent sessions. |

## Top Weaknesses (compact)

1. **Deriving recurrence from invariant** — 10+ documented struggles (1D DP)
2. **Over-engineering / overthinking** — Building complex data structures when simpler solutions exist (lc678: committed to one wildcard reading instead of tracking the alive set)
3. **Brute-force feasibility check** — Not checking constraint bounds first
4. **Debug print left in code** — 7 occurrences (fmt.Println)
5. **Pattern misidentification** — Prefix/range-query vs combinatorial DP
6. **Greedy instinct** — Picking greedy when minimax/binary-search is needed
7. **Solving by mechanics vs invariants** — Starting with implementation instead of stating invariant (reinforced lc22, lc678)
8. **Binary search indexing** — Confusion between returning index vs count (off-by-one errors)
9. **Tie-breaking logic** — Not considering all cases when selecting final answer
10. **Confusing dp state definition** — Not distinguishing "ending at i" vs "total so far" (lc3524, lc32)

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
