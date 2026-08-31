# Linked List

**Status:** In Progress — 1/5 sev≤1 streak

## Description

Linked list traversal, prev/curr pointer tracking, node manipulation. Foundation for fast/slow pointers, in-place modification, and cycle detection.

## Sessions

| Date | Problem | Difficulty | Severity | Time | Status | Key Learning |
|------|---------|------------|----------|------|--------|-------------|
| 2026-08-31 | lc2058 | medium | 0 (pass) | 15:36 | independent | Critical points need both prev & next; O(1) space by tracking only first + prev index |

## Lessons Learned

- **Critical points require both neighbors**: skip head and tail (they lack prev or next)
- **MinDist = min gap between consecutive critical points**: only need the most recent index
- **MaxDist = last − first critical point**: only need to remember the first one
- **O(1) space achievable**: instead of storing all indices, track `firstCriticalIndex` and `prevCriticalIndex`

## Variants

1. **Local maxima/minima detection** (lc2058): one-pass prev/curr/next comparison
2. **Cycle detection**: Floyd's fast/slow pointer (see Two Pointers (Fast/Slow))
3. **In-place modification**: rearrange or delete nodes without extra space

## Weaknesses

- None identified yet (first session)
