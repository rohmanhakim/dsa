# Progress Overview

## Current Skill Level

### Patterns Mastered
- **2D Dynamic Programming** — Understand DP table setup, base cases, and recurrence relations for string problems (e.g., Edit Distance). Comfortable with the idea of subproblems building on each other via a 2D table.

### Patterns In Progress
- None currently tracked

### Weaknesses to Watch
- **DP case completeness** — Tendency to forget the "match" case in DP recurrences (only handling mismatches). Always ask: "what happens when the condition IS met?"
- **Index alignment in DP** — Confusing table indices (1-indexed) with string indices (0-indexed). Remember: `dp[i][j]` uses `word1[i-1]` and `word2[j-1]`.
- **Deriving recurrence from invariant** — Currently guessing recurrence from example tables rather than deriving it from the invariant. Practice habit: define invariant → derive recurrence → verify with examples.

### Practice Plan
- LC 1143 - Longest Common Subsequence (2D DP, similar structure to Edit Distance)
- LC 5 - Longest Palindromic Substring (2D DP, different recurrence)
- LC 64 - Minimum Path Sum (2D DP, simpler)

### Problem Categories Solved
| Category | Count |
|----------|-------|
| 2D DP (Strings) | 1 |

### Monthly Progress
- [June 2026](2026-06.md)