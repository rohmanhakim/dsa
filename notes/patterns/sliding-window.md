# 7. Sliding Window

## When to Use
When you need to find a **subarray/substring** that satisfies a constraint (count, sum, unique characters, etc.). The key signal: "longest/shortest subarray/substring with property X."

## How It Works
1. Maintain a window `[left, right)` with two pointers
2. **Expand** `right` to include new elements
3. **Shrink** `left` when the window becomes invalid
4. Track the answer (max length, min length, etc.) at each step

## Key Insight: Superset Property
When a window is valid, all extensions of it are also valid (for "at most" constraints). This means you can count extensions rather than just tracking the window itself.

## Duplicate Detection: Hash Map vs `strings.Contains`
The standard approach for "longest substring without repeating characters" uses a **hash map** (or array) to track the last seen index of each character:

```
seen[char] = last index where char appeared
when duplicate found: left = max(left, seen[char] + 1)
```

**Why `max(left, seen[char]+1)`?** The duplicate's last occurrence might be *outside* the current window (stale entry). For example, in "abba":
- When processing the second 'a' at index 3, `seen['a'] = 0` but `left = 2` (already past index 0)
- `max(2, 0+1) = 2` — correct, don't move left backward

**Why not `strings.Contains`?** It's O(n) per call, making the overall solution O(n²). The hash map approach is O(n) total.

## Two Counting Perspectives
When counting valid windows (e.g., "number of substrings with at most K distinct characters"):
- **Count extensions right:** For each valid window `[left, right]`, all extensions `[left, right+k]` are also valid → count `n - right`
- **Count extensions left:** For each valid window, all extensions `[left-k, right]` are also valid → count `left`

## Example: LC 3 (Longest Substring Without Repeating Characters)
```
s = "abcabcbb"

right=1, curr='a', sub="" → sub="a", maxLength=1
right=2, curr='b', sub="a" → sub="ab", maxLength=2
right=3, curr='c', sub="ab" → sub="abc", maxLength=3
right=4, curr='a', sub="abc" → contains 'a' → left=1 → sub="bca", maxLength=3
right=5, curr='b', sub="bca" → contains 'b' → left=2 → sub="cab", maxLength=3
...
Result: 3 ✓
```

## Complexity
- **Time:** O(n) — each element visited at most twice (once by right, once by left)
- **Space:** O(min(n, alphabet size)) — hash map stores at most one entry per unique character

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Longest Substring Without Repeating Characters (LC 3)** | Hash map of last-seen indices, `left = max(left, seen[char]+1)` |
| **Longest Substring with At Most K Distinct Characters** | Hash map + shrink when count > K |
| **Minimum Window Substring (LC 76)** | Shrink from left when valid, track minimum |
| **Permutation in String (LC 567)** | Fixed-size window, count array for frequency |
| **Fruit Into Baskets (LC 904)** | Longest subarray with at most 2 distinct elements |
