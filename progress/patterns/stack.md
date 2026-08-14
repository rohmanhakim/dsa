# Stack

**Status:** ✅ Mastered — 5/5 consecutive independent sessions with severity ≤ 1 (4+ distinct days).
**Solved:** LC 20, LC 150, LC 71, LC 155, LC 2211

Core idea: LIFO data structure for balanced matching, expression evaluation, and simulation.

## Applications

- **LC 20 (Valid Parentheses):** Push opening brackets onto stack; on closing bracket, check top for valid pair, pop if match, else fail. Early exit: odd-length strings are always invalid. Key insight: helper closure (`pair`) cleanly handles multiple bracket types.
- **LC 150 (Evaluate Reverse Polish Notation):** Push numbers onto stack; on operator, pop two operands, apply operator, push result. Key insight: **first pop is the second operand** (right side of operator) — so `res = b <op> a` where `a` is first pop. Go's integer division truncates toward zero natively, matching the problem requirement.
- **LC 71 (Simplify Path):** Split path by `/`, use stack to filter components: skip empty/`.`, pop on `..`, push everything else. `strings.Split` naturally handles consecutive slashes (produces empty strings filtered by `case ""`). Clean, well-proportioned solution — exactly the right amount of code.
- **LC 155 (Min Stack):** Two-stack approach: main stack `data` stores all values, auxiliary stack `minVals` tracks minimums. Push to `minVals` only when `value <= minVals[-1]`. Pop from `minVals` when popped value equals current min. The `<=` condition handles duplicate minimums correctly. Classic pattern for constant-time min queries.
- **LC 2211 (Count Collisions on a Road):** Stack-based simulation for car collision counting. Push R's onto stack; on L, pop top (R→count+2, S→count+1) and push S; on S, drain all R's from stack (each R→count+1) and push S. Post-loop cleanup handles trailing R's behind S's by peeling off the S layer and counting remaining R's.

## Key Lessons

- **LC 2211 key insight:** The post-loop logic is the trickiest part — an S at the stack top acts as a "barrier" that traps R's underneath it. **O(1) alternative missed:** The editorial's counting observation (leading R's + trailing L's never collide, everything else collides exactly once) eliminates the need for a stack entirely. **Pattern recognition:** when the problem involves sequential collision/interaction simulation, a stack naturally models the "pending" state. But always check if a structural observation can simplify to O(1) space.

## Speed
LC 20: (slow pass), LC 150: 20:56, LC 71: 12:59, LC 155: 20:43, LC 2211: 43:40 (slow — over target by ~14 min, likely on post-loop edge cases).