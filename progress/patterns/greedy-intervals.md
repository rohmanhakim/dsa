# Greedy (Intervals)

**Status:** ✅ Mastered — 5/5 consecutive independent sessions with severity ≤ 1 across 5 distinct days.
**Solved:** LC 435, LC 1288, LC 56, LC 452, LC 57

## Variants

### Sort by end → greedy selection (LC 435, LC 1288)
Sort by end time, greedily keep the interval with earliest end. Used for maximum non-overlapping intervals / minimum removals.

### Sort by start → merge/group overlapping (LC 56, LC 452, LC 57)
Sort by start time (or already sorted), single-pass. Each interval either extends the current merged range or starts a new one. Used for combining all overlapping intervals, minimum arrows/groups, and inserting new intervals.

## Decision Framework

"Do I need to **combine** overlapping intervals (sort by start) or **select** non-overlapping ones (sort by end)?"

## Lessons

- **LC 452:** Track the intersection window `[xs, xe]`. The partial overlap case (new balloon starts within window but ends past it) requires no update — the sort guarantees future balloons have `start >= current start`, so only the right endpoint matters for future overlap checks.
- **LC 57:** For Insert Interval (intervals already sorted by start), the canonical approach is a 3-phase linear scan: (1) add all intervals ending before newInterval starts, (2) merge all overlapping intervals by updating newInterval bounds, (3) add remaining intervals. Each phase has a single clear invariant — no need for extra state variables. Non-standard single-pass approach with `prev`/`merged` flags worked but was significantly harder to implement and debug (2+ hours).