# 3. Virtual / Sentinel Points

## When to Use
When you need to avoid special-casing boundary conditions. Add an artificial "anchor" point that represents a known base case.

## How It Works
1. Append a virtual constraint/point with a known value (e.g., position 1 with height 0)
2. Sort it into the array
3. The loop now naturally handles the edge case without extra `if` statements

## Example: LC 1840
```go
// Building 1 always has height 0 — treat it as a constraint
restrictions = append(restrictions, []int{1, 0})
sort.Slice(restrictions, ...)
```
This eliminates the need to special-case the first building in the gap formula.

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Remove Nth Node From End (LC 19)** | Dummy head node simplifies edge removal |
| **Linked List Cycle II (LC 142)** | Sentinel values for Floyd's algorithm |
| **Largest Rectangle in Histogram (LC 84)** | Add 0 at end to flush remaining stack |
| **Trapping Rain Water (LC 42)** | Boundary elements act as natural sentinels |
