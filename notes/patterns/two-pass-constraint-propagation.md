# 1. Two-Pass Constraint Propagation (Left→Right + Right→Left)

## When to Use
When constraints propagate in both directions along a linear structure (array, line of buildings, etc.), and you need to satisfy all constraints simultaneously.

## How It Works
1. **Left→Right pass:** Each element considers constraints from its left neighbor
2. **Right→Left pass:** Each element considers constraints from its right neighbor
3. After both passes, each element holds the value that satisfies constraints from **both** directions

## Example: LC 1840 - Maximum Building Height
```
Restrictions: position 3 → height ≤ 1, position 5 → height ≤ 100

Left→Right:  position 5 gets min(100, 1+2) = 3  (constrained by position 3)
Right→Left:  position 3 stays min(1, 3+2) = 1   (already tight)
```
Without the right-to-left pass, position 5 would incorrectly use height 4 (only constrained from building 1), leading to wrong gap calculations.

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Trapping Rain Water (LC 42)** | leftMax[i] + rightMax[i] - height[i] |
| **Candy (LC 135)** | Left pass for rising, right pass for falling |
| **Product of Array Except Self (LC 238)** | Prefix product (L→R) + suffix product (R→L) |
| **Gas Station (LC 134)** | Track deficit from both directions |
