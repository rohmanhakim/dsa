# Algorithmic Patterns Study Notes

## 1. Two-Pass Constraint Propagation (Left→Right + Right→Left)

### When to Use
When constraints propagate in both directions along a linear structure (array, line of buildings, etc.), and you need to satisfy all constraints simultaneously.

### How It Works
1. **Left→Right pass:** Each element considers constraints from its left neighbor
2. **Right→Left pass:** Each element considers constraints from its right neighbor
3. After both passes, each element holds the value that satisfies constraints from **both** directions

### Example: LC 1840 - Maximum Building Height
```
Restrictions: position 3 → height ≤ 1, position 5 → height ≤ 100

Left→Right:  position 5 gets min(100, 1+2) = 3  (constrained by position 3)
Right→Left:  position 3 stays min(1, 3+2) = 1   (already tight)
```
Without the right-to-left pass, position 5 would incorrectly use height 4 (only constrained from building 1), leading to wrong gap calculations.

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Trapping Rain Water (LC 42)** | leftMax[i] + rightMax[i] - height[i] |
| **Candy (LC 135)** | Left pass for rising, right pass for falling |
| **Product of Array Except Self (LC 238)** | Prefix product (L→R) + suffix product (R→L) |
| **Gas Station (LC 134)** | Track deficit from both directions |

---

## 2. Sparse Point Processing (Large Input Space)

### When to Use
When `n` is huge (e.g., 10^9) but the number of "interesting" points (constraints, events) is small (e.g., 10^5). You **cannot** iterate over all `n` elements.

### How It Works
1. Only store and process the constraint/event points
2. For gaps between consecutive points, derive the answer using a **formula** (not iteration)
3. The formula captures what happens in the gap based on the two endpoint values

### Example: LC 1840
```
n = 10^9, restrictions at positions 2, 5, 7
→ Only 3 points to store
→ Gap formula: maxInGap = (hL + hR + d) / 2
→ No need to build an array of 10^9 elements
```

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Merge Intervals (LC 56)** | Process interval endpoints, not all integers |
| **Non-overlapping Intervals (LC 435)** | Greedy on endpoints |
| **Insert Interval (LC 57)** | Only iterate over overlapping intervals |
| **Range Sum Query - Immutable (LC 303)** | Prefix sum avoids re-iterating subarrays |

---

## 3. Virtual / Sentinel Points

### When to Use
When you need to avoid special-casing boundary conditions. Add an artificial "anchor" point that represents a known base case.

### How It Works
1. Append a virtual constraint/point with a known value (e.g., position 1 with height 0)
2. Sort it into the array
3. The loop now naturally handles the edge case without extra `if` statements

### Example: LC 1840
```go
// Building 1 always has height 0 — treat it as a constraint
restrictions = append(restrictions, []int{1, 0})
sort.Slice(restrictions, ...)
```
This eliminates the need to special-case the first building in the gap formula.

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Remove Nth Node From End (LC 19)** | Dummy head node simplifies edge removal |
| **Linked List Cycle II (LC 142)** | Sentinel values for Floyd's algorithm |
| **Largest Rectangle in Histogram (LC 84)** | Add 0 at end to flush remaining stack |
| **Trapping Rain Water (LC 42)** | Boundary elements act as natural sentinels |

---

## 4. Greedy Geometric Observation

### When to Use
When the optimal solution has a recognizable shape (triangle, mountain, valley) that can be described by a simple formula.

### How It Works
1. Observe the geometric shape of the constraint landscape
2. Derive a closed-form formula for the maximum/minimum in a region
3. Apply the formula across all regions and take the global optimum

### Example: LC 1840
Between two constrained points with heights `hL` and `hR` separated by distance `d`:
```
The optimal profile forms a triangle:
    peak = (hL + hR + d) / 2

Because: rise by 1 from left, fall by 1 from right
         meeting point is the peak
```

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Jump Game (LC 55)** | Greedy: track farthest reachable position |
| **Jump Game II (LC 45)** | Greedy BFS on jump ranges |
| **Maximum Subarray (LC 53)** | Kadane's greedy: extend or restart |
| **Container With Most Water (LC 11)** | Two pointers: move the shorter side |

---

## Quick Reference: Pattern Recognition

| If the problem involves... | Consider... |
|---|---|
| Constraints from both directions | Two-pass propagation |
| Huge input with few events | Sparse point processing + formula |
| Edge cases that need special handling | Virtual/sentinel points |
| An optimal shape that's geometrically obvious | Greedy geometric formula |