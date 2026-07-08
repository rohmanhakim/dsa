# 4. Greedy Geometric Observation

## When to Use
When the optimal solution has a recognizable shape (triangle, mountain, valley) that can be described by a simple formula.

## How It Works
1. Observe the geometric shape of the constraint landscape
2. Derive a closed-form formula for the maximum/minimum in a region
3. Apply the formula across all regions and take the global optimum

## Example: LC 1840
Between two constrained points with heights `hL` and `hR` separated by distance `d`:
```
The optimal profile forms a triangle:
    peak = (hL + hR + d) / 2

Because: rise by 1 from left, fall by 1 from right
         meeting point is the peak
```

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Jump Game (LC 55)** | Greedy: track farthest reachable position |
| **Jump Game II (LC 45)** | Greedy BFS on jump ranges |
| **Maximum Subarray (LC 53)** | Kadane's greedy: extend or restart |
| **Container With Most Water (LC 11)** | Two pointers: move the shorter side |
