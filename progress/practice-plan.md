# Practice Plan

## Immediate Priorities (closest to mastery)

1. **Hash Map re-certification (REGRESSED 2026-10-07)** — sev-3 on hr-icecream-parlor (hint needed on two-sum recognition) broke the 4/4 streak. Bar: **3 consecutive sev≤1 across 2 distinct days**. Progress: **1/3 sessions, day 1/2** — hr-missing-numbers sev 0 (frequency-difference variant, independent, same day as regression). Two-sum recognition still unproven: LC 1 (Two Sum — classic complement map), LC 560 (Subarray Sum Equals K — prefix sum + map), LC 454 (4Sum II — map of pair sums). One session per day — no same-day retries.
2. **DFS re-certification (REGRESSED 2026-10-02)** — sev-3 on lc22 broke the 9/9 streak. Bar: **3 consecutive sev≤1 across 2 distinct days**. Progress: **1/3 sessions, day 1/2** — LC 1971 sev 1 (general graph DFS variant, 2026-10-08; visited-marking fix post-acceptance). Remaining 2 must span ≥1 more distinct day; prefer a fresh *backtracking* variant (LC 78 Subsets, LC 131 Palindrome Partitioning, LC 51 N-Queens, LC 79 Word Search) since that is where the regression happened. One session per day — no same-day retries (contaminated sampling).
3. **1D Dynamic Programming** — 3 guided sessions, 0 independent. Highest session count, most practice. Prioritize independent problems: LC 70 (Climbing Stairs), LC 198 (House Robber), LC 322 (Coin Change).
4. **2D Dynamic Programming** — 2 guided sessions. Try LC 64 (Minimum Path Sum) or LC 62 (Unique Paths) independently.
5. **Greedy (Intervals)** — 2 guided sessions. Try LC 435 (Non-overlapping Intervals) independently.

## Problem Queue

- LC 1 - Two Sum (Hash Map re-cert, classic complement lookup)
- LC 560 - Subarray Sum Equals K (prefix sum + map, re-cert different flavor)
- LC 1143 - Longest Common Subsequence (2D DP, similar structure to Edit Distance)
- LC 5 - Longest Palindromic Substring (2D DP, different recurrence)
- LC 64 - Minimum Path Sum (2D DP, simpler)
- LC 142 - Linked List Cycle II — Floyd's Phase 2: finding the cycle start node (extends LC 141 pattern)
- LC 70 - Climbing Stairs (1D DP, simplest entry point)
- LC 198 - House Robber (1D DP, classic)
- LC 347 - Top K Frequent Elements (heap next rep — frequency-ordered heap, new sub-variant after lc215 sev 0)
- LC 1514 - Path with Maximum Probability (Dijkstra next rep — multiplicative/max combine operator, fresh sub-variant after lc743 sev 0)

## String Matching Algorithms

Learn Rabin-Karp (rolling hash), KMP, and Aho-Corasick. Practice when to choose which based on constraints. ~~Trie basics~~ ✓ Done (LC 208) ✓ Applied (LC 1268)
