# 2. Sparse Point Processing (Large Input Space)

## When to Use
When `n` is huge (e.g., 10^9) but the number of "interesting" points (constraints, events) is small (e.g., 10^5). You **cannot** iterate over all `n` elements.

## How It Works
1. Only store and process the constraint/event points
2. For gaps between consecutive points, derive the answer using a **formula** (not iteration)
3. The formula captures what happens in the gap based on the two endpoint values

## Example: LC 1840
```
n = 10^9, restrictions at positions 2, 5, 7
→ Only 3 points to store
→ Gap formula: maxInGap = (hL + hR + d) / 2
→ No need to build an array of 10^9 elements
```

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Merge Intervals (LC 56)** | Process interval endpoints, not all integers |
| **Non-overlapping Intervals (LC 435)** | Greedy on endpoints |
| **Insert Interval (LC 57)** | Only iterate over overlapping intervals |
| **Range Sum Query - Immutable (LC 303)** | Prefix sum avoids re-iterating subarrays |
