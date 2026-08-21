# Binary Search on Answer

**Status:** 🔄 In Progress — 0/5 sev≤1. 1 guided session.
**Solved:** LC 3116 (guided, sev 3)

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

## Variants

| Variant | Description | Example |
|---------|-------------|---------|
| kth smallest | count(x) = number of valid values ≤ x; find smallest x where count(x) ≥ k | LC 3116 (multiples of coins via inclusion-exclusion) |
| minimize the maximum | feasibility(x) = can we partition/assign with max ≤ x? | LC 410 (split array) |
| maximize the minimum | feasibility(x) = can we place/assign with min ≥ x? | LC 774 (minimize max distance to gas station) |

## Lessons

- **LC 3116 (guided, sev 3):** Binary search on answer + inclusion-exclusion for counting multiples. The `count` function and inclusion-exclusion were correct, but the binary search boundary was wrong: used `right = mid - 1` instead of `right = mid` with `left < right` template. **Root cause:** template confusion — mixing up the "eliminate mid" vs "preserve mid" branches. When `count(mid) ≥ k`, mid could be the answer, so `right = mid` preserves it.
- **Key insight:** Instead of enumerating k values (impossible for k up to 2×10^9), binary search on x and use a counting function to check if x is large enough.

## Weaknesses

- **Binary search template confusion (reinforced):** Used `right = mid - 1` when `right = mid` was needed. This is the same template confusion documented in the Binary Search pattern. For `left < right` + `left = mid + 1`: `right = mid` preserves the answer. For `left <= right`: `right = mid - 1` eliminates mid. Don't mix templates.