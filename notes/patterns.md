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

## 5. Monotonic Stack (Next Greater/Smaller Element)

### When to Use
When you need to find, for each element, the **next greater** or **next smaller** element in the array. Classic signal: "for each day/bar/element, when does something [warmer/taller/smaller] come next?"

### How It Works
Maintain a stack that is monotonic (either decreasing or increasing). Process elements left to right:
1. For each new element, pop all stack elements that are "resolved" by this element.
2. Set their answer (e.g., distance/index).
3. Push the current element onto the stack.

### Key Decision: Decreasing vs Increasing?

| You want... | Stack holds... | Stack is... | Pop when... |
|---|---|---|---|
| Next **greater** | items waiting for bigger | **Decreasing** (top is smallest) | `curr > top` |
| Next **smaller** | items waiting for smaller | **Increasing** (top is largest) | `curr < top` |

**Mnemonic:** The stack stores a "monotonic reserve" of unresolved elements. An element gets resolved when something *opposite* to the stack's order arrives.

### Example: LC 739 (Decreasing Stack)
```
temperatures = [73, 74, 75, 71, 69, 72, 76, 73]

Stack (bottom→top)  |  Current  |  Action
[(73)]              |  74       |  74>73, pop 73, res[0]=1
[(74)]              |  75       |  75>74, pop 74, res[1]=1
[(75)]              |  71       |  71<75, push → [(75,71)]
[(75,71)]           |  69       |  69<71, push → [(75,71,69)]
[(75,71,69)]        |  72       |  72>69→pop, res[4]=1; 72>71→pop, res[3]=2
[(75,72)]           |  76       |  76>72→pop, res[5]=1; 76>75→pop, res[2]=4
[(76)]              |  73       |  73<76, push → [(76,73)]
Done → remaining elements get 0
Result: [1,1,4,2,1,1,0,0] ✓
```

### Implementation Styles

There are two ways to implement a monotonic stack, depending on what the problem asks:

| Style | Stack stores | Answer computed by | Extra data needed |
|---|---|---|---|
| **Span-based** | `(value, accumulated_span)` | Sum of popped spans + 1 | None |
| **Index-based** | `indices` into a separate array | `len(array) - barrier_index - 1` | Full history array |

**When to use which:**
- **Span-based** when you only need a count or distance (e.g., LC 901 Stock Span → "how many days back?", LC 739 Daily Temperatures → "how many days forward?")
- **Index-based** when you need to identify the actual barrier element (e.g., "what is the temperature of the next warmer day?")

**Span-based trace (LC 901):**
```
Stack stores (price, span). When popping, accumulate spans:

next(70) with stack [(100,1),(80,1),(60,1)]:
  70 ≥ 60 → pop (60,1), span = 1+1 = 2
  70 < 80 → stop
  push (70, 2)
  return 2
```

**Index-based trace (LC 901, LeetCode official):**
```
Stack stores indices. t[] stores all prices. Compute from indices:

next(70), stack=[0,1,2], t=[100,80,60]:
  70 ≥ t[2]=60 → pop index 2
  70 < t[1]=80 → stop
  t = [100,80,60,70], push index 3
  span = len(t) - stack_top - 1 = 4 - 1 - 1 = 2
```

### Design Principle: Invariant-First

**Before coding, state your invariant in one sentence.** Example: "My stack maintains a strictly decreasing sequence of prices from bottom to top."

Then the popping rule is the *only* thing that preserves that invariant:
- "If `current >= top`, the decreasing property is broken → pop to restore it."

This prevents the common mistake of solving by mechanics ("move data between structures") instead of by invariant ("maintain a property").

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Daily Temperatures (LC 739)** | Decreasing stack: next warmer day |
| **Next Greater Element I (LC 496)** | Decreasing stack on nums2, map for nums1 |
| **Next Greater Element II (LC 503)** | Circular array: process 2× to simulate wrap-around |
| **Largest Rectangle in Histogram (LC 84)** | Increasing stack: next smaller bar (left & right boundaries) |
| **Online Stock Span (LC 901)** | Decreasing stack: consecutive days with ≤ price |
| **Asteroid Collision (LC 735)** | Stack for simulating collision resolution |

### Complexity
- **Time:** O(n) — each element pushed and popped at most once
- **Space:** O(n) — stack size

---

## 6. Backward Two-Pointer Merge

### When to Use
You have two sorted sequences and need to merge them into a buffer that has **extra space at the end**. Forward insertion causes O(n) shifting per element; filling from the back eliminates shifting entirely.

### How It Works
1. Start pointers at the **end** of both sorted sequences (`p1 = m-1`, `p2 = n-1`)
2. Start the write pointer at the **end** of the result buffer (`p = m+n-1`)
3. Compare `nums1[p1]` vs `nums2[p2]`, place the **larger** value at `nums1[p]`, decrement the winning pointer and `p`
4. When one sequence is exhausted, copy remaining elements from the other (they're already in place if it's nums1)

### Key Insight
The empty slots at the end of nums1 are **unused space** — writing there never overwrites valid data. By always placing the largest remaining element at the current back position, you guarantee sorted order with zero shifting.

### Example: LC 88
```
nums1 = [1,2,3,0,0,0]  nums2 = [2,5,6]

p1=2(3), p2=2(6), p=5  → 6>3 → nums1[5]=6
p1=2(3), p2=1(5), p=4  → 5>3 → nums1[4]=5
p1=2(3), p2=0(2), p=3  → 3>2 → nums1[3]=3
p1=1(2), p2=0(2), p=2  → 2≥2 → nums1[2]=2
p1=0(1), p2=-1,   p=1  → nums2 done, copy nums1[0:1] → already in place
Result: [1,2,2,3,5,6] ✓
```

### Complexity
- **Time:** O(m+n) — single pass, each element placed exactly once
- **Space:** O(1) — in-place

### Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Merge Sorted Array (LC 88)** | Classic backward merge into buffer with trailing zeros |
| **Squares of a Sorted Array (LC 977)** | Largest squares at the ends → fill result from back |
| **Sort Transformed Array (LC 360)** | Quadratic transform preserves sorted property at extremes |

---

## Quick Reference: Pattern Recognition

| If the problem involves... | Consider... |
|---|---|
| Constraints from both directions | Two-pass propagation |
| Huge input with few events | Sparse point processing + formula |
| Edge cases that need special handling | Virtual/sentinel points |
| An optimal shape that's geometrically obvious | Greedy geometric formula |
| Next greater/smaller element for each item | Monotonic stack |
| Merging sorted data into a buffer with space at the end | Backward two-pointer merge |
