# Prefix Sum

**Status:** ⏳ 3/5 consecutive independent sessions with severity ≤ 1 (3 distinct days).
**Sessions:** 10 total (4 independent: LC 3904, LC 3903, LC 238, LC 3499; 6 guided/secondary: LC 3756, LC 3532, LC 3534, LC 3312, LC 581, LC 1872).

Core idea: precompute cumulative values in O(n) to answer range queries in O(1).

## Variants

### Variant A — Range Sum
1-indexed arrays with sentinel at index 0, query on [l,r] uses `pref[r+1] - pref[l]`. For non-trivial functions, precompute powers of the base.

### Variant B — Range Emptiness
Prefix sum of boolean flags (breaks/gaps), query [u,v] checks `prefix[v] - prefix[u] == 0`. Useful when sorted array + connectivity queries imply contiguous components.

### Variant C — Cumulative Count + Binary Search
Prefix sum of exact counts, query finds smallest index where `prefix[g] > q` via binary search. Used when the answer is the k-th element in a sorted frequency distribution.

### Variant D — Prefix/Suffix Product
Prefix product array (left→right) + suffix product array (right→left), then combine: `result[i] = prefix[i-1] * suffix[i+1]`. Generalizes prefix sums to products. Used for "product of all elements except self" without division. Follow-up: O(1) space by building prefix into result array and using a running suffix variable.

### Variant E — Prefix/Suffix Aggregates (Max/Min)
Build prefix max array (left→right, track running max) and suffix min array (right→left, track running min). Combine at each index: `score[i] = prefixMax[i] - suffixMin[i]`. Useful when an index's "score" depends on properties of its left and right neighborhoods. Can reduce O(n²) brute-force to O(n). Space can be O(1) if only one direction's aggregate is stored as an array (the other tracked as a running variable during the scan).

## Key Recognition Signal

When constraints show n=10^5 and q=10^5, O(nq) is too slow — think prefix sums. But simple prefix/suffix patterns (like LC 238) are recognizable independently.

## Lessons

- **LC 238:** Clean implementation (7:24). Prefix/suffix product arrays are a direct generalization of prefix sums. Three-pass structure (build prefix, build suffix, combine) is clean and readable. The follow-up O(1) space solution combines the prefix pass into the result array and uses a single running suffix variable — good to know for interviews.
- **LC 3903:** Built prefixLargest + suffixSmallest arrays to solve "instability score" (max prefix − min suffix) in O(n), beating editorial's O(n²) brute force. Good instinct to precompute running aggregates from both ends instead of rescanning. Space could be reduced to O(n) by keeping only suffix array and tracking prefix max as a variable during scan, but O(n) is fine here.
- **LC 3904:** Follow-up to LC 3903 requiring O(n) algorithm. Same Variant E pattern — suffix min array + running prefix max. Solved independently in 5:50, clean implementation. Confirms strong command of this variant: precompute one direction as array, track the other as running variable during scan.

## Weakness
Still requires guidance to recognize in complex variants.