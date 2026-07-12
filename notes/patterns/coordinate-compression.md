# Coordinate Compression (Sort + Hash Map)

## When to Use

- Problem asks for "rank," "position in sorted order," or "relative ordering"
- Value range is huge (±10⁹) but count is small (n ≤ 10⁵)
- Need array-indexed DP on non-contiguous or large-range values
- Graph nodes are values (not indices) and need compact IDs

## Core Idea

Map arbitrary values to a compact range of consecutive integers (1..k) while **preserving relative order**. The absolute magnitudes don't matter — only which value is bigger/smaller than which.

## Template (3 steps)

```
1. Sort unique values → establishes order
2. Build map: {value → rank} (1-indexed, skip duplicates)
3. Map back to original positions
```

## Example Trace

```
Input:  [40, 10, 20, 30]
Sorted: [10, 20, 30, 40]
Ranks:  {10→1, 20→2, 30→3, 40→4}
Output: [4, 1, 2, 3]
```

Duplicate handling:
```
Input:  [100, 100, 100]
Sorted: [100]
Ranks:  {100→1}
Output: [1, 1, 1]
```

## Complexity

- **Time:** O(n log n) for sorting + O(n) for map build and lookup
- **Space:** O(n) for the sorted copy and the map

## Variants

### Variant A: Rank assignment (LC 1331)
Direct application: sort → build rank map → map back. The rank map is the answer.

### Variant B: Compressed-index DP
When you need DP on values but the range is too large for an array:
1. Compress values to 0..k
2. Allocate DP array of size k+1
3. Run DP using compressed indices
Useful for: LIS variants, interval scheduling with arbitrary coordinates, graph problems where nodes are values.

## Recognition Signals

| Signal | Why coordinate compression |
|---|---|
| "Assign ranks based on value" | Classic rank assignment |
| Value range ±10⁹ but n ≤ 10⁵ | Can't allocate array on raw values, compress first |
| "k-th smallest/largest" problems | Compression as preprocessing |
| DP on x-coordinates where x is sparse | Compress to contiguous indices |
| Nodes are values, not 0..n-1 | Compress node IDs for array-based adjacency |

## Common Mistake

Reaching for arithmetic tricks (e.g., `maxVal - a`) to compute rank instead of the straightforward sort+map. The arithmetic approach breaks on duplicates and non-uniform distributions. Always use sort+map — it's simpler and correct.

## Relation to Other Patterns

- **Sorting** is the prerequisite — coordinate compression is sorting + a map on top
- **Hash Map** is the lookup structure — O(1) per element after O(n log n) sort
- Often combined with **Binary Search** when the compressed index is used for lookups in a sorted structure

## Problems

- LC 1331 — Rank Transform of an Array (direct application, Variant A)
- USACO 1063 — Rectangular Pasture
- CF gym102951D — Static Range Queries
- CSES 1619 — Restaurant Customers
- CF 1000C — Covered Points Count
- USACO 786 — Lifeguards
- USACO 787 — Rental Service-
- USACO 896 — Mountain View
- USACO 1064 — Stuck in a Rut
- USACO 645 — Splitting the Field
- CF 632C — The Smallest String Concatenation
- CF 1478C — Nezzar and Symmetric Array
- CF 1472E — Correct Placement
- USACO 1015 — Triangles
- USACO 834 — Out Of Sorts
- USACO 967 — Meetings
