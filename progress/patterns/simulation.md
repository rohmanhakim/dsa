# Simulation

**Status:** In Progress

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