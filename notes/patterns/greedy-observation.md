# Greedy (Observation)

## When to Use

The optimal solution comes from **observing the net effect** of an operation or structure, rather than simulating it step-by-step. Key signals:
- "At most one operation" problems
- "What is the minimum/maximum after one change?"
- The operation seems complex to simulate, but the net effect might be simple
- Ask: **"What is the net effect?"** before coding

## How It Works

1. **Model the operation** — what actually changes?
2. **Find the invariant** — what stays the same, what changes?
3. **Derive the formula** — express the gain/loss in terms of measurable quantities
4. **Optimize** — apply the formula across all candidates, take the global optimum

The key insight is that many operations have a **net effect** that is much simpler than the operation itself. The removed part is often restored (net zero), and the actual gain comes from a smaller, identifiable component.

## Example: LC 3499 (Maximize Active Section with Trade I)

**Operation:** Remove a block of 1s (surrounded by 0s), then convert the resulting merged 0-block to 1s.

**Naive approach (what the user tried):** Scan left/right around each 1-block to count surrounding regions — complex, 3 nested loops per transition point.

**Observation approach:** The removed 1-block is restored (net zero change). The gain = sum of the two adjacent 0-block lengths. So: `answer = count_ones + max(adjacent zero-block pair sum)`.

**Why it works:** When you remove a 1-block of length b and merge the adjacent 0-blocks of lengths a and c, the merged block has length a+b+c. Converting it to 1s gives a+b+c ones. But you removed b ones, so net gain = (a+b+c) - b = a+c.

This reduces the entire problem to: collect zero-block lengths, find max adjacent pair. One pass, O(n).

## Example: LC 55 / LC 45 (Jump Game)

**Operation:** Jump from position i to any position in [i+1, i+nums[i]].

**Naive approach:** Try all possible jump sequences (exponential).

**Observation approach:** Track the farthest reachable position. At each step, the farthest you can reach is `max(farthest, i + nums[i])`. If the current position exceeds farthest, you're stuck.

## Example: LC 53 (Maximum Subarray / Kadane's)

**Operation:** Choose a contiguous subarray.

**Naive approach:** Try all O(n²) subarrays.

**Observation approach:** At each position, the maximum subarray ending here is either the current element alone or the current element extending the previous maximum subarray. `dp[i] = max(nums[i], dp[i-1] + nums[i])`.

## Example: LC 11 (Container With Most Water)

**Operation:** Choose two lines to form a container.

**Naive approach:** Try all O(n²) pairs.

**Observation approach:** The area is `min(h[left], h[right]) * (right - left)`. Start with the widest container. Moving the shorter line inward can only increase area (because the width decreases but the height might increase). Moving the taller line inward can only decrease area.

## Example: LC 1840
Between two constrained points with heights `hL` and `hR` separated by distance `d`:
```
The optimal profile forms a triangle:
    peak = (hL + hR + d) / 2

Because: rise by 1 from left, fall by 1 from right
         meeting point is the peak
```

## Complexity

- **Time:** O(n) — typically a single pass after the observation
- **Space:** O(1) to O(n) — depends on what data you need to track

## Common Mistakes

1. **Simulating the operation step-by-step** when the net effect is a simple formula. Always ask "what is the net effect?" first.
2. **Building complex data structures** (like region structs) to model every possible configuration, when a single observation eliminates the need.
3. **Starting with implementation mechanics** ("what data moves where") instead of stating the invariant ("what property determines the answer?").

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 3499** — Maximize Active Section with Trade | Net effect: gain = sum of adjacent zero-block lengths |
| **LC 55** — Jump Game | Track farthest reachable position |
| **LC 45** — Jump Game II | Greedy BFS on jump ranges |
| **LC 53** — Maximum Subarray | Kadane's: extend or restart |
| **LC 11** — Container With Most Water | Two pointers: move the shorter side |
| **LC 1840** — Maximum Building Height | Geometric: peak = (hL + hR + d) / 2 |