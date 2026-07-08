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

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Daily Temperatures (LC 739)** | Decreasing stack: next warmer day |
| **Next Greater Element I (LC 496)** | Decreasing stack on nums2, map for nums1 |
| **Next Greater Element II (LC 503)** | Circular array: process 2× to simulate wrap-around |
| **Largest Rectangle in Histogram (LC 84)** | Increasing stack: next smaller bar (left & right boundaries) |
| **Online Stock Span (LC 901)** | Decreasing stack: consecutive days with ≤ price |
| **Asteroid Collision (LC 735)** | Stack for simulating collision resolution |

## Complexity
- **Time:** O(n) — each element pushed and popped at most once
- **Space:** O(n) — stack size
