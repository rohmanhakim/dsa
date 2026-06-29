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
- **Over-engineering / Overthinking** — Tendency to build data structures (maps, arrays) when a simpler built-in function would suffice. Before coding, ask: "What's the minimum I need to compute?" Use `strings.Contains` directly when the problem doesn't require custom substring search. Also reinforced in LC 14 — used 4 state variables when 1 sufficed. Practice: count your variables before submitting, and ask if each one is truly necessary.

### Practice Plan
- LC 1143 - Longest Common Subsequence (2D DP, similar structure to Edit Distance)
- LC 5 - Longest Palindromic Substring (2D DP, different recurrence)
- LC 64 - Minimum Path Sum (2D DP, simpler)
- **String Matching Algorithms** — Learn Rabin-Karp (rolling hash), KMP, Trie basics, and Aho-Corasick. Practice when to choose which based on constraints.

### Constraint-Based Algorithm Selection (Reference)
| Constraint Size | Complexity Target | Algorithm Family |
|---|---|---|
| n ≤ 10⁴ | O(n²) | Brute force, nested loops |
| n ≤ 10⁵ | O(n log n) | Sorting, binary search, two pointers |
| n ≤ 10⁶ | O(n) | Sliding window, hash map, BFS/DFS |
| n ≤ 10⁷ | O(n) | Linear scan, monotonic stack |
| n ≤ 10⁸ | O(n log n) | Rare, usually requires optimization tricks |

### Problem Categories Solved
| Category | Count |
|----------|-------|
| 2D DP (Strings) | 1 |
| String Manipulation | 2 |

### Monthly Progress
- [June 2026](2026-06.md)