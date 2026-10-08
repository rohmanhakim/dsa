# Sliding Window

**Status:** ⚠️ Streak broken 2026-09-23 (lc1658, guided sev 3; also lc1477 secondary sev 3). Was 4/4 near mastery — data gap from the DB reset (2026-09-08..09-26) was backfilled 2026-10-08. Needs fresh 5 consecutive sev≤1.
**Sessions:** 15 DB sessions (6 guided).

## Seven Sub-Variants

### Longest valid (expand when valid, shrink when invalid)
**LC 3, LC 2958** — longest subarray/substring satisfying a constraint. Hash map tracks frequencies; shrink from left when constraint violated. Key insight: only the newly added element can violate the constraint, so the inner check is O(1).

### Shortest valid (shrink when valid, track min)
**LC 209** — minimum subarray sum ≥ target.

### Fixed-size window (fill, then slide one at a time)
**LC 438, LC 567** — find all anagrams / permutation in string.

### Longest with deletions (maxFreq trick)
**LC 2831** — longest equal subarray after removing at most k. Window valid when `windowSize - maxFreq <= k`. Key insight: **maxFreq never decreases**.

### Iterate over unique count (non-monotonic validity)
**LC 395** — iterate `currUnique` from 1 to `maxUnique`. Track `countAtLeastK` instead of exact minimum frequency.

### Counting exact matches with prefix zeros
**LC 930** — each leading zero in a valid window is an extra starting point.

### Shortest valid with lexicographic comparison
**LC 2904** — shrink past leading zeros when `count == k`.

### Counting subarray membership (frequency across all windows)
**LC 3471** — count how many subarrays of size k contain each element. Maintain freq map across sliding windows. Key insight: element at position p belongs to `min(p+1, n-k+1, k)` subarrays. Also, for 1<k<n, only first/last elements can appear in exactly one subarray (editorial observation shortcut).

## Key Insight

When a window is valid, all extensions are also valid (superset property). Two counting perspectives: count extensions right (`n-right`) or count extensions left (`left`).

## Lessons

- **LC 3:** Used `strings.Contains` for duplicate detection (O(n²)). Hash map is O(n).
- **LC 209:** Non-standard init added unnecessary complexity. Adopt standard template.
- **LC 438:** The `matches` counter optimization tracks how many of the 26 positions match — O(1) per step.
- **LC 567:** Copy-pasted template from LC 438. Pattern recognition improving but speed needs work.
- **LC 2831:** The `maxFreq` reframing eliminates recomputation. **Pattern:** "longest subarray after removing at most k elements" → `windowSize - maxFreq <= k`.
- **LC 395:** Threshold-crossing counters (`countAtLeastK`) for non-monotonic validity.
- **LC 2904:** Shrink past leading zeros when count == k optimizes both length and lexicographic order. Re-attempt: 41:53 to 20:28, sev 1 to 0. The invariant that once ones reaches k it stays k (leading-zero removal preserves count) is the key insight that makes the final gate check correct.
- **LC 2958:** Clean "longest valid" implementation — inner `for freq[curr] > k` correctly checks only the newly added element. 5:24 solve time, well under 30-min target.
- **LC 3090:** Clean "longest valid" with `freq[ch] > 2` threshold. Key insight: for "at most K occurrences" constraint, only the newly added character can violate window validity — shrinking only decreases other characters' counts, so they stay valid. 4:05 solve time.
- **LC 3471:** Counting subarray membership. 22:49 solve time (over 15-min target). Code structure non-standard: re-increments elements from left→right-1 on each slide instead of decrement old / increment new. Works correctly but harder to reason about. Bonus observation: for 1<k<n, only first/last elements can appear in exactly one subarray. Weakness #2 (over-engineering) triggered — full sliding window unnecessary when a direct check suffices.
- **LC 219 (guided, sev 3, primary, 2026-10-08):** Contains Duplicate II — TLE on large inputs. Window template + freq map were correct, but the validity check scanned the **entire map** per step (zero-count keys linger → O(n) scan → O(n²) total). Fixed with O(1) check `freq[nums[right]] > 1` after the shrink loop. **Invariant:** after add+shrink, `freq[v]` = count of `v` in `[left..right]`; only the newly added element can newly cross 1. The O(1)-check lesson was already recorded 3× (LC 3, LC 2958, LC 3090) but **not retrieved under TLE pressure** — transfer failure, not a knowledge gap. New corollary: per-step cost of the validity check is part of the complexity budget, not just the outer loops.

## Backfill 2026-10-08 (DB-reset gap 2026-09-08..09-26)

- **LC 1658 (guided, sev 3, primary, 2026-09-23):** "Reduce x from both ends" → complement subarray sum (target `total - x`); answer `n - longest` window. Tried DP first (pattern misidentification, weakness #5). **The window we want is the LONGEST, not the shortest — the optimization direction flips with the complement.** This sev-3 primary session broke the 4/4 near-mastery streak.
- **LC 1477 (guided, sev 3, secondary, 2026-09-17):** sliding window finds candidate windows; 1D prefix-min DP (`dp[left-1]`) enforces non-overlap. Combination pattern: window + "best before the left edge". Under the strict multi-pattern rule this also counts against the streak.
