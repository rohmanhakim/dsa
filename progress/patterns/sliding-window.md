# Sliding Window

**Status:** ⏳ 4/5 consecutive independent sessions with severity ≤ 1 across 4 distinct days (need 1 more for mastery).
**Sessions:** 10 DB sessions (LC 1358 guided but never logged to DB).

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

## Key Insight

When a window is valid, all extensions are also valid (superset property). Two counting perspectives: count extensions right (`n-right`) or count extensions left (`left`).

## Lessons

- **LC 3:** Used `strings.Contains` for duplicate detection (O(n²)). Hash map is O(n).
- **LC 209:** Non-standard init added unnecessary complexity. Adopt standard template.
- **LC 438:** The `matches` counter optimization tracks how many of the 26 positions match — O(1) per step.
- **LC 567:** Copy-pasted template from LC 438. Pattern recognition improving but speed needs work.
- **LC 2831:** The `maxFreq` reframing eliminates recomputation. **Pattern:** "longest subarray after removing at most k elements" → `windowSize - maxFreq <= k`.
- **LC 395:** Threshold-crossing counters (`countAtLeastK`) for non-monotonic validity.
- **LC 2904:** Shrink past leading zeros when `count == k` optimizes both length and lexicographic order.
- **LC 2958:** Clean "longest valid" implementation — inner `for freq[curr] > k` correctly checks only the newly added element. 5:24 solve time, well under 30-min target.
- **LC 3090:** Clean "longest valid" with `freq[ch] > 2` threshold. Key insight: for "at most K occurrences" constraint, only the newly added character can violate window validity — shrinking only decreases other characters' counts, so they stay valid. 4:05 solve time.