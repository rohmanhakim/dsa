# Binary Search on Answer

**Status:** 🔄 In Progress — 1/5 sev≤1 (day 1). 2 total sessions (1 independent, 1 guided).
**Solved:** LC 3116 (guided, sev 3), LC 875 (independent, sev 1, 34:28 slow)

## Core Idea

When the problem asks "find the kth smallest/largest X" or "minimize the maximum / maximize the minimum", and you can define a **monotonic predicate** `count(x) ≥ k` (or `feasible(x) = true/false`), binary search on the answer space.

1. **Define the search range:** `[lo, hi]` where `lo` is the smallest possible answer and `hi` is a safe upper bound.
2. **Define count(x):** Given a candidate `x`, count how many valid values are ≤ x (or check feasibility).
3. **Binary search:** Find the smallest `x` where `count(x) ≥ k`.

## Template (half-open style with `left < right`)

```go
left, right := lo, hi
for left < right {
    mid := left + (right-left)/2
    if count(mid) < k {
        left = mid + 1   // mid is too small, eliminate it
    } else {
        right = mid       // mid could be the answer, preserve it
    }
}
return left  // or right (they're equal)
```

**Critical:** When using `left < right` + `left = mid + 1`, the else branch MUST be `right = mid` (not `mid - 1`). Setting `right = mid - 1` throws away the answer when `count(mid) ≥ k`.

**Alternative valid form — closed interval + minK (LC 875):** `left <= right` with `right = mid - 1` / `left = mid + 1`, recording `minK = mid` *before* shrinking on feasible. Safe because a smaller feasible value overwrites `minK`. **Both forms find min-feasible — pick the one you're fluent in; don't re-derive under the clock.**

## Variants

| Variant | Description | Example |
|---------|-------------|---------|
| kth smallest | count(x) = number of valid values ≤ x; find smallest x where count(x) ≥ k | LC 3116 (multiples of coins via inclusion-exclusion) |
| minimize the maximum | feasibility(x) = can we partition/assign with max ≤ x? | LC 410 (split array) |
| minimize rate under budget | feasibility(x) = does rate x satisfy the global budget (time/hours)? | LC 875 (min k with Σ⌈pᵢ/k⌉ ≤ h) |
| maximize the minimum | feasibility(x) = can we place/assign with min ≥ x? | LC 774 (minimize max distance to gas station) |

## Lessons

- **LC 3116 (guided, sev 3):** Binary search on answer + inclusion-exclusion for counting multiples. The `count` function and inclusion-exclusion were correct, but the binary search boundary was wrong: used `right = mid - 1` instead of `right = mid` with `left < right` template. **Root cause:** template confusion — mixing up the "eliminate mid" vs "preserve mid" branches. When `count(mid) ≥ k`, mid could be the answer, so `right = mid` preserves it.
- **LC 875 (independent, sev 1, 34:28):** Brute force → monotone feasibility → answer-space search `[1, max(piles)]` derived cleanly; closed+minK template paired correctly (lc3116 lesson held). Cost: 4:28 choosing the binary search *form* — recall recipe needed, not re-derivation. Also: debug `fmt.Println` left in code (weakness #4) and pseudocode artifact went untraced (else branch moved `right` instead of `left`).
- **Key insight:** Instead of enumerating k values (impossible for k up to 2×10^9), binary search on x and use a counting function to check if x is large enough.

## Weaknesses

- **Binary search template confusion (reinforced):** Used `right = mid - 1` when `right = mid` was needed. This is the same template confusion documented in the Binary Search pattern. For `left < right` + `left = mid + 1`: `right = mid` preserves the answer. For `left <= right`: `right = mid - 1` eliminates mid. Don't mix templates.

## LeetCode Problems
- 875. Koko Eating Banana
- 1011. Capacity To Ship Packages Within D Days (exact same pattern)
- 2187. Minimum Time to Complete Trips (binary search on time)
- 410. Split Array Largest Sum (binary search on maximum subarray sum)
- Aggressive Cows (binary search on minimum distance)