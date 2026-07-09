# 8. Prefix Sum

## When to Use

- **Range sum queries** on a static array: answer `sum(arr[l..r])` in O(1) after O(n) preprocessing.
- **Range emptiness/count queries**: answer "are there any flagged positions in range [u, v]?" in O(1).
- Trigger signal: **n and q up to 10⁵, O(nq) is too slow — think prefix sums.**

## Variant A: Range Sum (Cumulative Sum)

### How It Works
1. Build `prefix[i] = prefix[i-1] + arr[i]` (1-indexed with sentinel at index 0).
2. Query `[l, r]` (0-indexed, inclusive) = `prefix[r+1] - prefix[l]`.

### Example: LC 3756
```
arr = [2, 4, 6, 8], prefix = [0, 2, 6, 12, 20]

Query [1, 3] = prefix[4] - prefix[1] = 20 - 2 = 18 ✓ (4+6+8=18)
```

### Key Detail
1-indexed arrays with sentinel at index 0 avoid edge cases where `l = 0`. The sentinel `prefix[0] = 0` means no special-casing needed.

### Complexity
- **Time:** O(n) build, O(1) per query
- **Space:** O(n)

---

## Variant B: Range Emptiness (Break/Flag Counting)

### When to Use
You have a boolean property between consecutive elements (e.g., "is there a gap > threshold?") and need to answer: "does this property hold for ALL pairs in range [u, v]?"

### How It Works
1. Define a boolean flag array: `flag[i] = 1` if the property is violated between elements i and i+1, else 0.
2. Build prefix sum of flags: `prefix[i] = prefix[i-1] + flag[i-1]`.
3. Query `[u, v]` (u < v): **no violations** in range iff `prefix[v] - prefix[u] == 0`.

### Key Insight
Since `nums` is sorted, connected components in the graph are always contiguous subarrays. A gap `> maxDiff` between consecutive elements creates a wall — nothing to the left can reach anything to the right (because values only increase further). This transforms "path existence" into a range emptiness query.

### Example: LC 3532
```
nums = [2, 5, 6, 8], maxDiff = 2

Gaps: |5-2|=3 > 2 (BREAK), |6-5|=1 ≤ 2, |8-6|=2 ≤ 2
flag = [1, 0, 0]
prefix = [0, 1, 1, 1]

Query (1, 3): prefix[3] - prefix[1] = 1 - 1 = 0 → same component → true ✓
Query (0, 1): prefix[1] - prefix[0] = 1 - 0 = 1 → different components → false ✓
```

### Complexity
- **Time:** O(n) build, O(1) per query
- **Space:** O(n)

---

## When NOT to Use Prefix Sum

- When the query is on a **mutable** array (updates change values). Use segment tree or Fenwick tree instead.
- When the property isn't aggregatable with addition (e.g., min/max queries). Use sparse table or segment tree.

## Classic Problems to Practice

| Problem | Variant | Key Idea |
|---|---|---|
| **Range Sum Query - Immutable (LC 303)** | A (range sum) | Classic 1-indexed prefix sum |
| **LC 3756** | A (range sum) | Prefix sum for concatenated digit sums with modular arithmetic |
| **Path Existence Queries (LC 3532)** | B (range emptiness) | Prefix sum of breaks in sorted array → connected component queries |
| **Subarray Sum Equals K (LC 560)** | A variant | Prefix sum + hash map: count subarrays where `prefix[j] - prefix[i] = k` |
| **Maximum Size Subarray Sum Equals k (LC 325)** | A variant | Prefix sum + hash map for longest subarray |
| **Contiguous Array (LC 525)** | A variant | Treat 0 as -1, prefix sum + hash map for equal 0s and 1s |
| **Binary Subarrays With Sum (LC 930)** | A variant | Prefix sum + hash map with frequency count |