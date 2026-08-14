# Minimax DP

**Status:** ⏳ 0/5 sessions toward mastery (guided).
**Sessions:** LC 486 (guided, sev 3), LC 1406 (guided, sev 3), LC 1140 (guided, sev 3), LC 1510 (guided, sev 3).

## Variants

### Score difference (2D) — LC 486
`dp[i][j]` = best score difference from subarray. Takes from either end → 2D DP.

### Score difference (1D) — LC 1406, LC 1140
`dp[i]` = best score difference from suffix. Takes from front only → 1D DP.

### Boolean win/lose (1D) — LC 1510
`dp[i]` = can current player force a win with `i` stones? Recurrence: `dp[i] = OR over moves of (NOT dp[i-move])`. No score, just win/lose. Simpler than score-difference variants — no subtraction trick needed.

## Decision Framework

"Is this a **score-based** game (use score difference + subtraction trick) or a **win/lose** game (use boolean + OR over NOT)?"

## Lessons

- **LC 486:** `dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])`. The subtraction captures the opponent playing optimally against you. Initially tried greedy (always pick larger end) which fails because it ignores opponent's optimal response. **Loop order:** when `dp[i][j]` depends on `dp[i+1][j]` and `dp[i][j-1]`, iterate `i` from right to left (or by length).
- **LC 1406:** `dp[i] = max(take1 - dp[i+1], take2 - dp[i+2], take3 - dp[i+3])`. Takes from front only → remaining game is always a suffix → 1D DP. Correctly identified minimax; once given the score difference framing, implementation was clean. **Key lesson:** when the game only takes from one end, the state collapses to 1D.
- **LC 1140:** `dp[i][m]` = best score difference from suffix `piles[i..n-1]` with current max-take `m`. Recurrence: `dp[i][m] = max over X in [1,2m] of (suffix[i] - suffix[i+X] - dp[i+X][max(m,X)])`. Take-all base case when `i+2m >= n` → `dp[i][m] = suffix[i]`. **Key lesson:** when the legal move set depends on a parameter (M) that changes turn-to-turn, that parameter is part of the game state and must be a DP dimension — M is *carried state*, not a constant. M pulls two levers at every state: caps options (X in [1,2m]) and transforms into the next state (`max(m,X)`). Suffix sum makes `take` O(1). Answer conversion: `(total + dp[0][1]) / 2`. Also hit a negative-accumulator bug: `best := 0` for a max over possibly-negative score differences returns a fake 0 when all candidates are negative — init to first candidate or `math.MinInt`.
- **LC 1510:** `dp[i] = OR over moves of (NOT dp[i-move])`. No score, just win/lose. Simpler than score-difference variants.