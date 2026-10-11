# Progress Overview

Compact index. Per-pattern details are in `progress/patterns/<name>.md`.

## Patterns Mastered

| Pattern | Streak | Distinct Days | Key Speed |
|---------|--------|---------------|-----------|
| Stack | 10/10 sev≤1 | 10 | 8:05–43:40 |

## Patterns Near Mastery

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| Two Pointers | 1/5 sev≤1 | 10 primary (5 guided) | Stale "2/2" corrected 2026-10-08: strict model had streak at 0 (lc3983 sev 3 anchor); LC 27 (Read/Write variant, sev 1) restarts it. Bar: 5 sev≤1 across 3 distinct days. |
| Monotonic Stack | 2/2 sev≤1 | 4 total (2 guided) | 3 more consecutive sev≤1 passes |

## Patterns In Progress

| Pattern | Streak | Sessions | Needs |
|---------|--------|----------|-------|
| DFS | 1/3 sev≤1 (re-cert, day 1/2) | 15 total (3 guided) | Re-cert after REGRESSED 2026-10-02 (sev 3, lc22): lc1971 sev 1 (2026-10-08) = 1/3, day 1/2. Needs 2 more sev≤1 across ≥1 more day. |
| Binary Search | 1/3 sev≤1 (re-cert, day 1/2) | 13 total (4 guided) | Re-cert: lc350 sev 0 (2026-10-08) = 1/3, day 1/2. Needs 2 more sev≤1 across ≥1 more day. |
| Greedy (Intervals) | 0/5 sev≤1 (re-cert) | 8 total (3 guided) | REGRESSED 2026-09-19 (sev 3, lc1520 + lc2472; backfilled 2026-10-08). Re-cert: 3 sev≤1 across 2 days. HR hackerland-radio-transmitters (guided sev 3): wrong-model formula attempt; rightmost-in-range greedy learned. |
| Math | 0/5 sev≤1 | 19 total (6 guided) | sev≤1 pair (lc3870, lc3871) wiped by guided lc1401 (sev 3) 6 days later. Needs 5 consecutive independent sev≤1. |
| Digit Manipulation | 2/5 sev≤1 | 8 total (2 guided) | Rising: lc3550 + lc3622 sev≤1; chain stopped by lc3348 (sev 3). 3 more needed. |
| Greedy (Observation) | 0/5 sev≤1 | 17 total (11 guided) | Streak broken by guided sessions. Needs independent practice. LC678 (guided, sev 3): wildcard range invariant learned. |
| BFS | 1/5 sev≤1 | 6 total (2 guided) | Streak restarted 2026-10-09 (lc463 sev 1 slow pass, 22:44 > 15-min easy target). Needs 4 more sev≤1 across ≥2 more days. |
| Hash Map | 1/3 sev≤1 (re-cert) | 10 total (3 guided) | REGRESSED 2026-10-07 (sev 3, hr-icecream-parlor). Re-cert 1/3, day 1/2: hr-missing-numbers sev 0 (freq-count variant). Two-sum recognition still open. |
| Linked List | 1/1 sev≤1 | 1 session | 4 more |
| Simulation | 1/5 sev≤1 | 7 total (2 guided) | Streak restarted 2026-10-08 (lc4056 sev 0). lc1096 (guided sev 3) previously anchored at 0 — grammar parsing w/ precedence (concatenation > union) still the reference lesson there. |
| Sliding Window | 0/5 sev≤1 | 15 total (6 guided) | Streak broken 2026-09-23 (sev 3, lc1658; also lc1477 secondary sev 3). lc219 sev 3 (2026-10-08): O(1) window-check lesson (LC 3/2958/3090) not retrieved — TLE from per-step full-map scan. Was 4/4 near mastery — needs fresh 5. |
| Prefix Sum | 1/5 sev≤1 (fresh) | 12 total (7 guided) | Streak restarted 2026-10-07 (hr-sherlock-and-array, sev 1 slow pass; latent last-index bug noted). |
| 1D Dynamic Programming | 0/5 | 19 total (17 guided) | Needs independent sessions. Key: recurrence derivation. |
| 2D Dynamic Programming | 0/5 | 7 total (7 guided) | Needs independent sessions. |
| Minimax DP | 0/5 | 4 guided | Needs independent sessions. |
| BST | 0/5 | 2 guided | Needs independent sessions. |
| Segment Tree | 0/5 | 2 total (2 guided) | Needs independent sessions. |
| Heap (Priority Queue) | 0/5 sev≤1 | 2 sessions (2 guided) | 2× guided sev 3: lc703 (raw `h.Pop()` vs `heap.Pop(&h)`), abc234d (missed size-k eviction → root decayed to prefix min). Right pattern both times; invariant not yet second nature — **next heap rep must be cold/independent** (e.g. lc215, lc347). |
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
