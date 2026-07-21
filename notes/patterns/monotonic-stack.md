# 5. Monotonic Stack (Next Greater/Smaller Element)

## When to Use
When you need to find, for each element, the **next greater** or **next smaller** element in the array. Classic signal: "for each day/bar/element, when does something [warmer/taller/smaller] come next?"

## How It Works
Maintain a stack that is monotonic (either decreasing or increasing). Process elements left to right:
1. For each new element, pop all stack elements that are "resolved" by this element.
2. Set their answer (e.g., distance/index).
3. Push the current element onto the stack.

## Key Decision: Decreasing vs Increasing?

| You want... | Stack holds... | Stack is... | Pop when... |
|---|---|---|---|
| Next **greater** | items waiting for bigger | **Decreasing** (top is smallest) | `curr > top` |
| Next **smaller** | items waiting for smaller | **Increasing** (top is largest) | `curr < top` |

**Mnemonic:** The stack stores a "monotonic reserve" of unresolved elements. An element gets resolved when something *opposite* to the stack's order arrives.

## Example: LC 739 (Decreasing Stack)
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

## Implementation Styles

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

## Design Principle: Invariant-First

**Before coding, state your invariant in one sentence.** Example: "My stack maintains a strictly decreasing sequence of prices from bottom to top."

Then the popping rule is the *only* thing that preserves that invariant:
- "If `current >= top`, the decreasing property is broken → pop to restore it."

This prevents the common mistake of solving by mechanics ("move data between structures") instead of by invariant ("maintain a property").

## Sub-Variant: Boundary Detection (Two-Pass)

### When to Use
When you need to find the **leftmost and rightmost out-of-place elements** in an array — i.e., the boundaries of a subarray that, if sorted, would make the entire array sorted. Classic signal: "shortest unsorted continuous subarray."

### How It Works
Use **two passes** with monotonic stacks, each maintaining a single invariant:

1. **Left pass (left→right, decreasing stack):** Find the leftmost element that is larger than something to its right.
   - Stack maintains a decreasing sequence (bottom to top).
   - Pop when `nums[i] < stack[-1]` — the stack top is out of place.
   - Track `left = min(left, popped_index)`.

2. **Right pass (right→left, increasing stack):** Find the rightmost element that is smaller than something to its left.
   - Stack maintains an increasing sequence (bottom to top).
   - Pop when `nums[i] > stack[-1]` — the stack top is out of place.
   - Track `right = max(right, popped_index)`.

### Key Insight
Each pass has a **single clear invariant** (decreasing for left, increasing for right). This eliminates the need to track multiple state variables in a single pass. The two-pass approach is cleaner than trying to compute both bounds in one pass.

### Example: LC 581 (Shortest Unsorted Continuous Subarray)

```
nums = [2, 6, 4, 8, 10, 9, 15]

Left pass (decreasing stack, left→right):
  i=0: push [2@0]
  i=1: 6 >= 2, push [2@0, 6@1]
  i=2: 4 < 6 → pop 6@1, left = min(∞, 1) = 1. push [2@0, 4@2]
  i=3: 8 >= 4, push [2@0, 4@2, 8@3]
  i=4: 10 >= 8, push [2@0, 4@2, 8@3, 10@4]
  i=5: 9 < 10 → pop 10@4, left stays 1. push [2@0, 4@2, 8@3, 9@5]
  i=6: 15 >= 9, push [2@0, 4@2, 8@3, 9@5, 15@6]
  left = 1

Right pass (increasing stack, right→left):
  i=6: push [15@6]
  i=5: 9 < 15, push [15@6, 9@5]
  i=4: 10 > 9 → pop 9@5, right = max(-1, 5) = 5. push [15@6, 10@4]
  i=3: 8 < 10, push [15@6, 10@4, 8@3]
  i=2: 4 < 8, push [15@6, 10@4, 8@3, 4@2]
  i=1: 6 > 4 → pop 4@2, right = max(5, 2) = 5. push [15@6, 10@4, 8@3, 6@1]
  i=0: 2 < 6, push [15@6, 10@4, 8@3, 6@1, 2@0]
  right = 5

Answer: right - left + 1 = 5 - 1 + 1 = 5 ✓
```

### Common Mistake
Trying to compute both bounds in a **single pass** with multiple state variables (`tops`, `first`, `right`). This leads to incorrect bound computation because the right boundary should be `max(right, i)` on every pop (the current index where the violation is detected), not just the popped element's index. Two-pass eliminates this complexity.

### Complexity
- **Time:** O(n) — each pass is O(n), each element pushed and popped at most once per pass
- **Space:** O(n) — stack size per pass

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Daily Temperatures (LC 739)** | Decreasing stack: next warmer day |
| **Next Greater Element I (LC 496)** | Decreasing stack on nums2, map for nums1 |
| **Next Greater Element II (LC 503)** | Circular array: process 2× to simulate wrap-around |
| **Largest Rectangle in Histogram (LC 84)** | Increasing stack: next smaller bar (left & right boundaries) |
| **Online Stock Span (LC 901)** | Decreasing stack: consecutive days with ≤ price |
| **Asteroid Collision (LC 735)** | Stack for simulating collision resolution |
| **Shortest Unsorted Continuous Subarray (LC 581)** | Two-pass boundary detection: left pass (decreasing) + right pass (increasing) |

## Complexity
- **Time:** O(n) — each element pushed and popped at most once
- **Space:** O(n) — stack size
