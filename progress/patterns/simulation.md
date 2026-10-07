# Simulation

**Status:** 🔄 In Progress — 0/5 streak (6 total, 2 guided). lc1096 (2026-09-27, guided sev 3) anchors the streak at 0. Backfilled 2026-10-08.

## Pattern Description

Direct simulation of problem rules: translate the problem description to code 1:1 when the problem defines explicit step-by-step operations. No optimization or clever trick needed — just faithful implementation.

## When to Use

- Problem defines explicit step-by-step operations to perform
- Constraints are small enough (typically n ≤ 1000) that O(n) or O(n²) simulation is fine
- No known DSA pattern applies — the problem *is* the algorithm

## Sessions

| Date | Problem | Difficulty | Severity | Time | Notes |
|------|---------|-----------|----------|------|-------|
| 2026-08-20 | lc3069 - Distribute Elements Into Two Arrays I | easy | 0 (pass) | 6:52 | Direct 1:1 translation of problem rules; clean implementation |

## Lessons Learned

- (none yet — first session on this pattern)

## Variants

- **Array simulation**: distribute/merge elements based on comparisons (lc3069)
- **String simulation**: character-by-character processing with state
- **Grid simulation**: cell-by-cell updates with neighbor rules (e.g., Game of Life)
## Backfill 2026-10-08 (DB-reset gap 2026-09-08..09-26)

- **LC 835 Image Overlap (independent, sev 0, 28:46):** n ≤ 30 → brute-force translation over all (±i, ±j) offsets + bitwise overlay is the intended O(n⁴) solution. Minor: the 4 sign-combos redo identical work when i or j = 0.
- **LC 836 Rectangle Overlap (independent, sev 1, 23:11):** AABB overlap = axis-wise non-separation (`ax1 < bx2 && ax2 > bx1 && ay1 < by2 && ay2 > by1`); strict inequalities make edge/corner touching non-overlapping. Slow pass on an easy — deriving the condition from the "positive area" definition took too long.
- **LC 3498 Reverse Degree (independent, sev 0, 3:21):** the statement is already a formula — translate it 1:1, no data structures.
- **LC 1096 Brace Expansion II (guided, sev 3, 2026-09-27):** recursive-descent grammar parsing — `expr` (top-level comma split tracked by brace depth → union), `term` (concatenation via cartesian product of items), `item` (letter or braced expr). Initially tried stack-based parsing with no clear plan. Key: concatenation binds tighter than union. This guided session anchors the Simulation streak at 0.
