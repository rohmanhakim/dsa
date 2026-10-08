# Greedy (Intervals)

**Status:** 🔄 In Progress — REGRESSED (was 5/5 mastered). lc2472 (09-15) + lc1520 (09-19) guided sev 3; backfilled 2026-10-08. Re-cert: 3 sev≤1 across 2 days.
**Solved:** LC 435, LC 1288, LC 56, LC 452, LC 57, LC 1520, LC 2472, HR Hackerland Radio Transmitters

## Variants

### Sort by end → greedy selection (LC 435, LC 1288)
Sort by end time, greedily keep the interval with earliest end. Used for maximum non-overlapping intervals / minimum removals.

### Sort by start → merge/group overlapping (LC 56, LC 452, LC 57)
Sort by start time (or already sorted), single-pass. Each interval either extends the current merged range or starts a new one. Used for combining all overlapping intervals, minimum arrows/groups, and inserting new intervals.

### Closure + Interval Scheduling (LC 1520)
Preprocessing step: compute "closure" ranges for each character before greedy selection. Key insight: constraint "if substring contains c, must contain ALL occurrences of c" creates a chain reaction. Simple first/last tracking is insufficient — must iteratively scan positions and expand until stable. After closure computation, apply standard interval scheduling (sort by end, greedy pick).

## Decision Framework

"Do I need to **combine** overlapping intervals (sort by start) or **select** non-overlapping ones (sort by end)?"

## Lessons

- **LC 452:** Track the intersection window `[xs, xe]`. The partial overlap case (new balloon starts within window but ends past it) requires no update — the sort guarantees future balloons have `start >= current start`, so only the right endpoint matters for future overlap checks.
- **LC 57:** For Insert Interval (intervals already sorted by start), the canonical approach is a 3-phase linear scan: (1) add all intervals ending before newInterval starts, (2) merge all overlapping intervals by updating newInterval bounds, (3) add remaining intervals. Each phase has a single clear invariant — no need for extra state variables. Non-standard single-pass approach with `prev`/`merged` flags worked but was significantly harder to implement and debug (2+ hours).
- **LC 1520:** Chain reaction concept — including one character forces inclusion of all its occurrences, which may force inclusion of other characters. **Overlap ≠ containment**: just because ranges overlap doesn't mean one contains the other (e.g., b[1,1] overlaps a[0,4] in "abcaa" but b doesn't contain a). Simple first/last tracking fails because [first,last] alone cannot determine if a character appears inside a range — must scan actual positions. Preprocessing requires iterative expansion until stable.- **LC 2472 (Palindrome Substrings, guided sev 3, backfilled 2026-10-08):** Earliest-finish greedy sweep — at each position take the shortest valid palindrome (length k, then k+1) starting at the earliest feasible index; `start` pointer enforces non-overlap. **Lesson:** "maximize count of non-overlapping X" → earliest-finish greedy; the validity check (palindrome) is a filter, not the pattern. A 1D DP framing was tried first — pattern misidentification. **Regression:** this session + lc1520 (both guided sev 3) broke the 5/5 mastered streak on 2026-09-19. Re-cert: 3 sev≤1 across 2 days.
- **HR Hackerland Radio Transmitters (guided sev 3, 2026-10-08):** Coverage variant — minimum transmitters (each *on a house*) to cover all houses; dual of activity selection. Invariant: everything left of pointer `i` is covered; at the leftmost uncovered house, place the transmitter at the **rightmost house within k** (exchange argument: still covers the current house, extends coverage maximally right), then skip all houses ≤ spot+k. **Wrong model first:** grouped houses by consecutive gaps ≤ k then `ceil(span/(2k+1))` per group — counts *positions* not *houses*; fails on sparse groups (e.g. `[1,4,7,10,13,16] k=3`: formula says 3, correct 2). **sort.Search gotcha:** returns first-true index (= count of elements ≤ target); the transmitter index is result-1.
