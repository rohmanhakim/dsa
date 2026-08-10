# Minimax DP (Game Theory)

## When to Use
Two-player game theory problems where both players play optimally. Key signals:
- "Alice and Bob take turns"
- "Both players play optimally"
- "Return the winner (Alice/Bob/Tie)"
- Options are constrained (take 1, 2, or 3; take from left or right)

## Core Concept: Score Difference

The trick is to define `dp[i]` as the **score difference** (current player − opponent) from game state `i`.

Why score difference? Because when it's your turn, you want to maximize your total score. But your opponent will also play optimally against you. By tracking the *difference*, we capture both goals in a single recurrence.

**The subtraction trick:** When you take some value and move to state `next`, your score difference is:
```
value_taken − dp[next]
```
The subtraction (`− dp[next]`) captures the fact that from `next` onward, it's the **opponent's** turn — their advantage is your disadvantage.

## Three Distinct Variants

### Variant A: Take from Either End (2D DP)

**Example: LC 486 — Predict the Winner**

At each turn, the current player takes from the **left** or **right** end. The remaining game is a subarray `nums[i..j]`.

```
dp[i][j] = max score difference from subarray nums[i..j]
         = max(nums[i] − dp[i+1][j],    // take left
               nums[j] − dp[i][j-1])    // take right
```

- **Loop order:** iterate `i` from right to left (or by length), since `dp[i][j]` depends on `dp[i+1][j]` and `dp[i][j-1]`
- **Base case:** `dp[i][i] = nums[i]` (only one element left)
- **Space:** O(n²)

### Variant B: Take from Front Only (1D DP)

**Example: LC 1406 — Stone Game III**

At each turn, the current player takes 1, 2, or 3 stones from the **front**. The remaining game is always a suffix `stones[i..n-1]`.

```
dp[i] = max score difference from suffix stones[i..n-1]
      = max(take1 − dp[i+1],
            take2 − dp[i+2],
            take3 − dp[i+3])
```

where `take_k = sum(stones[i..i+k-1])`.

- **Loop order:** right to left (dependencies at larger indices)
- **Base case:** `dp[n] = 0` (no stones left, score difference is 0)
- **Space:** O(n) — this is the key simplification vs 2D

**Key lesson:** When the game only takes from one end, the remaining state is always a suffix → 1D DP suffices.

### Variant C: Boolean Win/Lose (1D DP)

**Example: LC 1510 — Stone Game IV**

No score — just win or lose. `dp[i]` = can the current player force a win with `i` stones?

```
dp[i] = OR over all valid moves s (where s² ≤ i) of (NOT dp[i - s²])
```

- **Base case:** `dp[0] = false` (no stones → can't move → lose)
- **Loop order:** left to right (dp[i] depends on smaller indices dp[i-s²])
- **Space:** O(n)

**Key insight:** No subtraction trick needed. "I win if I can move to a state where you lose." Simpler than score-difference variants.

## Template (1D Variant — LC 1406)

```go
func stoneGameIII(stoneValue []int) string {
    n := len(stoneValue)
    dp := make([]int, n+3)  // +3 for safe indexing

    for i := n - 1; i >= 0; i-- {
        takes := []int{}
        take := 0
        for k := 0; k < 3 && i+k < n; k++ {
            take += stoneValue[i+k]
            takes = append(takes, take - dp[i+k+1])
        }
        dp[i] = slices.Max(takes)
    }

    switch {
    case dp[0] > 0:  return "Alice"
    case dp[0] < 0:  return "Bob"
    default:         return "Tie"
    }
}
```

## Common Mistakes

1. **Using accumulated score instead of score difference.** The recurrence `dp[i] = max(take − dp[next])` uses subtraction, not addition. Adding `take + dp[next]` ignores that the opponent will also play optimally.
2. **Wrong loop order.** Since `dp[i]` depends on `dp[i+1]`, `dp[i+2]`, `dp[i+3]`, you must iterate from right to left. Iterating left to right uses uninitialized values.
3. **Forgetting the base case.** When no stones remain, the score difference is 0 (not the accumulated score). Initialize `dp[n] = 0`.
4. **Confusing 2D vs 1D.** "Take from either end" → 2D (subarray). "Take from front only" → 1D (suffix). Check which end the game takes from before choosing dimensionality.

## Key Questions to Ask

1. **What does `dp[i]` (or `dp[i][j]`) mean?** — Best score difference (current − opponent) from state i (or [i,j]).
2. **What does it depend on?** — All states reachable by valid moves (take 1/2/3, or take left/right).
3. **Which direction to iterate?** — Always from dependencies first (right to left for suffixes, by increasing length for subarrays).

## Complexity

- **1D variant:** O(n) time, O(n) space (or O(1) with rolling 3 variables)
- **2D variant:** O(n²) time, O(n²) space

## Problems

| Problem | Variant | Key Idea |
|---|---|---|
| **LC 486 — Predict the Winner** | 2D (take from either end) | `dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])` |
| **LC 1406 — Stone Game III** | 1D (take from front) | `dp[i] = max(take_k - dp[i+k+1])` for k=1,2,3 |
| **LC 877 — Stone Game** | 2D (take from either end) | Same as LC 486; parity argument makes Alice always win |
| **LC 1563 — Stone Game V** | 1D variant with partition | Different constraint (split into two groups), similar minimax structure |
| **LC 1510 — Stone Game IV** | 1D (boolean win/lose) | `dp[i] = OR over s² ≤ i of (NOT dp[i-s²])` |
