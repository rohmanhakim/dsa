# 7. Sliding Window

## When to Use
When you need to find a **subarray/substring** that satisfies a constraint (count, sum, unique characters, etc.). The key signal: "longest/shortest subarray/substring with property X."

## How It Works
1. Maintain a window `[left, right)` with two pointers
2. **Expand** `right` to include new elements
3. **Shrink** `left` to restore validity or optimize the answer
4. Track the answer (max length, min length, etc.) at each step

### Three Sub-Variants

| Variant | Goal | When to shrink | Inner loop |
|---|---|---|---|
| **Longest valid** | Maximize window size | Shrink when **invalid** | `while !valid { shrink }` |
| **Shortest valid** | Minimize window size | Shrink while **valid** | `while valid { update answer; shrink }` |
| **Fixed-size** | Match exact window size | No shrink loop | slide one step at a time |

**Longest valid (LC 3):** Expand right to grow the window. When the window becomes invalid (e.g., duplicate found), shrink left until valid again. Answer is updated during expansion.

**Shortest valid (LC 209):** Expand right to grow the window. When the window becomes valid (e.g., sum ≥ target), shrink left while still valid, updating the minimum at each step. Answer is updated during shrinking.

**Fixed-size (LC 438):** Fill an initial window of exact size `wlen`. Check, then slide one position at a time: remove leftmost element, add next element, check again. No inner while loop needed — both pointers advance by exactly 1 each step. Use frequency array + `matches` counter for O(1) per-step checking instead of O(alphabet) comparison.

### Standard Template (Shortest Valid)
```go
left := 0
sum := 0
minLen := math.MaxInt
for right := 0; right < n; right++ {
    sum += nums[right]
    for sum >= target {
        minLen = min(minLen, right-left+1)
        sum -= nums[left]
        left++
    }
}
// minLen == math.MaxInt → return 0
```

### Standard Template (Longest Valid)
```go
left := 0
count := 0
maxLen := 0
for right := 0; right < n; right++ {
    // add nums[right] to window state
    for !valid {
        // remove nums[left] from window state
        left++
    }
    maxLen = max(maxLen, right-left+1)
}
```

## Key Insight: Superset Property
When a window is valid, all extensions of it are also valid (for "at most" constraints). This means you can count extensions rather than just tracking the window itself.

**For shortest-valid problems:** The superset property means once you find a valid window, you should try to shrink it — a smaller valid window is always better. This is why the inner loop shrinks while valid.

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

## Common Mistakes
- **Non-standard initialization** (e.g., `left=0, right=1, sum=nums[0]`) adds unnecessary complexity: extra guards, boolean flags, slice allocations. Always use the standard template: both pointers at 0, sum at 0.
- **Confusing the two sub-variants** — using `while !valid` (longest) when you need `while valid` (shortest), or vice versa. Before coding, ask: "Am I maximizing or minimizing the window?"

## Classic Problems to Practice
| Problem | Variant | Key Idea |
|---|---|---|
| **Longest Substring Without Repeating Characters (LC 3)** | Longest valid | Hash map of last-seen indices, `left = max(left, seen[char]+1)` |
| **Longest Substring with At Most K Distinct Characters** | Longest valid | Hash map + shrink when count > K |
| **Minimum Size Subarray Sum (LC 209)** | Shortest valid | Minimize window with sum ≥ target, shrink while valid |
| **Minimum Window Substring (LC 76)** | Shortest valid | Shrink from left when valid, track minimum |
| **Permutation in String (LC 567)** | Fixed-size | Fixed-size window, count array for frequency |
| **Fruit Into Baskets (LC 904)** | Longest valid | Longest subarray with at most 2 distinct elements |
| **Longest Substring with At Least K Repeating Characters (LC 395)** | Iterate over unique count | Non-monotonic validity → iterate `currUnique` 1..maxUnique, run standard window for each |
| **Binary Subarrays with Sum (LC 930)** | Counting exact matches with prefix zeros | Each leading zero in a window with `sum == goal` is an extra valid starting point. Track `prefixZeros`, shrink past zeros even when `sum ≤ goal`. `count += prefixZero + 1` |

### Sub-Variant: Iterate Over Unique Count (LC 395)

When the validity condition is **non-monotonic** for a standard window (e.g., "all characters appear ≥ k times" — adding or removing a character can each make the window invalid), use a different strategy: **iterate over the number of unique characters** in the window.

**When to use:** The problem asks for the longest/shortest substring where *every* character meets a frequency threshold. A single sliding window can't maintain this because the condition isn't monotonic.

**How it works:**
1. Count total unique characters in `s` → `maxUnique`
2. For `currUnique` from 1 to `maxUnique`:
   - Run a standard sliding window that maintains **at most `currUnique` unique characters** (shrink when `len(freq) > currUnique`)
   - Track `countAtLeastK` = number of unique characters in the window with freq ≥ k
   - When `countAtLeastK == currUnique`, the window is valid → update answer

**Key tracking: `countAtLeastK` (threshold crossing counter)**
Instead of tracking exact minimum frequency (which is hard to maintain during shrink), track how many characters have *crossed* the k threshold:
- **Adding a character:** if `freq[ch]` goes from `k-1` to `k` → `countAtLeastK++`
- **Removing a character:** if `freq[ch]` goes from `k` to `k-1` → `countAtLeastK--`

This is O(1) per add/remove and doesn't require recomputing the minimum.

**Template:**
```go
maxUnique := len(uniqueChars(s))
best := 0
for currUnique := 1; currUnique <= maxUnique; currUnique++ {
    freq := map[byte]int{}
    countAtLeastK := 0
    left := 0
    for right := 0; right < n; right++ {
        // add s[right]
        freq[s[right]]++
        if freq[s[right]] == k { countAtLeastK++ }
        // shrink if too many unique
        for len(freq) > currUnique {
            if freq[s[left]] == k { countAtLeastK-- }
            freq[s[left]]--
            if freq[s[left]] == 0 { delete(freq, s[left]) }
            left++
        }
        // check validity
        if countAtLeastK == currUnique {
            best = max(best, right-left+1)
        }
    }
}
```

**Complexity:** O(maxUnique × n) = O(26 × n) = O(n) for lowercase letters.
