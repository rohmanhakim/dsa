# 1D Dynamic Programming

**Status:** ⏳ 0/5 sessions toward mastery (guided sessions don't count).
**Sessions:** 13 total (LC 198 independent slow pass, LC 322 guided, LC 139 guided ×2, LC 300 guided, LC 213 guided, LC 91 guided, LC 3514 guided, LC 518 guided, LC 53 guided, LC 416 guided, LC 940 guided, LC 32 guided).

## Key Framework for Faster State Identification

(1) What does `dp[i]` **mean**? (2) What answers from before does `dp[i]` **depend on**? Should take ≤2 minutes to fill in.

## Core Recurrences

- `dp[i] = max(dp[i-1], dp[i-2] + nums[i])` — LC 198
- `dp[i] = min(dp[i-c]+1)` for each coin c — LC 322
- `dp[i] = any(s[i-len(w):i] == w && dp[i-len(w)])` — LC 139 word break
- `dp[i] = 1 + max(dp[j])` for j < i where nums[j] < nums[i] — LC 300 LIS
- `dp[i] = (s[i]!='0' ? dp[i-1] : 0) + (10<=s[i-1..i]<=26 ? dp[i-2] : 0)` — LC 91 decode ways
- `dp[i] += dp[i-c]` for each coin c — LC 518 counting combinations
- `dp[i] = max(nums[i], dp[i-1]+nums[i])` with answer = max over all dp[i] — LC 53 Kadane's

Two sliding variables replace the full DP table for LC 198.

## Lessons

- **LC 198:** Spent ~42 minutes experimenting with parity-based bucket tracking before arriving at a correct but non-standard formulation.
- **LC 322:** Spent 1.5 hours unable to identify states. Initially tried greedy (fails for arbitrary denominations), then multiple DP formulations. The two-question framework ("What does dp[i] mean?" / "What does dp[i] depend on?") was the breakthrough — needs to be applied as the **first** step, not after trying combinations.
- **LC 139 (re-attempt):** First attempt used `strings.LastIndex` — works but unnecessary O(n) scan. Re-attempt started with overcomplicated approach: tracked word indices in `dp[i]`, performed string replacements. After guidance, rewrote as clean boolean DP in 21 lines. `dp[i]` only needs to be a boolean (is it possible?), not a list of which words were used. Second guided session — 9th documented struggle with deriving recurrence.
- **LC 300:** Identified state (dp[i] = LIS ending at index i) correctly but couldn't derive the recurrence. Over-engineered the inner loop with unnecessary variables. Fourth documented struggle with deriving recurrence. After guidance: `dp[i] = max(dp[i], dp[j]+1)`.
- **LC 213:** Attempted to handle circular constraint within a single DP pass by tracking `robFirstHouse` flag. Failed because the 2-element DP state cannot capture the constraint. **Circular constraint pattern:** split into two linear subproblems: exclude-last and exclude-first. Reuse the linear solution (LC 198) on both.
- **LC 91:** Had correct intuition but botched the `0` case — set `dp[i] = 1` (constant) instead of `dp[i-2]` when encountering a `0` that pairs with the previous digit. The `0` can ONLY be decoded as part of a two-digit pair, so it contributes `dp[i-2]` ways. Rewrote using 1-indexed dp with sentinels.
- **LC 3514:** Initially brute-forced O(C(k,3)) DFS — TLE. After guidance, implemented incremental set building: boolean arrays `one`, `two`, `three` where each layer XORs every element with all achievable values from the previous layer. O(n × U) where U = 2048. **Key insight:** when the value range is bounded (nums[i] ≤ 1500 → XOR < 2048), build achievable values layer by layer instead of enumerating all combinations.
- **LC 518:** Stuck on the recurrence for counting coin combinations. **Critical insight:** coins must be the outer loop and amounts the inner loop to count combinations (not permutations). If you iterate amounts outer and coins inner, you count permutations (1+2 and 2+1 counted separately). 5th documented struggle with deriving recurrence.
- **LC 53:** Spent ~2 hours stuck. Mixed two state interpretations: "max subarray ending at i" vs "max subarray anywhere in nums[0..i]". The former gives a clean 2-candidate recurrence; the latter requires tracking two pieces of state. **Key lesson:** when choosing state definition, prefer the one that gives the simplest recurrence. 6th documented struggle.
- **LC 416:** Initially pursued element-difference tracking — a valid-sounding but wrong framing. After guidance, correctly implemented 1D DP subset sum: `dp[j] = dp[j] || dp[j-num]` with backwards iteration. **Key lesson:** when the problem asks to split into two equal halves, reframe immediately as "find any subset summing to total/2." This is the boolean 0/1 knapsack. 8th documented struggle.
- **LC 940:** Started with **2D interval DP** (`dp[i][j]` for substring `s[i..j]`) with complex inner dedup loop — wrong dimensionality entirely. After guidance, arrived at clean 1D running total: `dp = 2*dp - last[c]`. Struggled with: (1) identifying state as running total not interval, (2) understanding `last[c]` stores a **value** (snapshot) not an **index**, (3) MOD subtraction guard. **Key insight for counting problems:** when asked "count distinct X", the state is typically a cumulative total, not a range decomposition. The elegant formula `dp = 2*dp - last[c]` works because `-(-1)` gives `+1` for new characters automatically. 9th documented struggle with deriving recurrence.
- **LC 32:** Started with `dp[i]` as a **running max** ("so far") + char stack — recurrence degenerated to carry-forward/`max(dp[i-1], 2)` and couldn't stitch `() + (())` (got 4 instead of 6 on `()(())`). After hint: flipped to `dp[i]` = longest valid **ending exactly at i**, index stack, `dp[i] = (i - j) + dp[j]`. **Key lesson:** the state definition *forces* the recurrence — "ending at i" makes the stitch the only possible move. **10th documented struggle with deriving recurrence; also reinforces weakness #10 (state definition).** Secondary pattern: Stack (index stack mechanics were correct — positions not chars).

## Weakness
- **Deriving recurrence from invariant:** Currently guessing recurrence from the example tables rather than deriving it from the invariant. Practice habit: define invariant → derive recurrence → verify with examples. 10+ documented struggles (LC 198, LC 322, LC 139, LC 300, LC 518, LC 53, LC 1406, LC 1140, LC 940, LC 32).