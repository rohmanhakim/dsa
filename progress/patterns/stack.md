# Stack

**Status:** ✅ Mastered — 5/5 consecutive independent sessions with severity ≤ 1 (5 distinct days).

**Sessions:** LC 20, LC 150, LC 2211, LC 1190, LC 1614, LC 1111 (secondary)

Core idea: LIFO data structure for balanced matching, expression evaluation, and simulation.

## Applications

- **LC 20 (Valid Parentheses):** Push opening brackets onto stack; on closing bracket, check top for valid pair, pop if match, else fail. Early exit: odd-length strings are always invalid. Key insight: helper closure (`pair`) cleanly handles multiple bracket types.
- **LC 150 (Evaluate Reverse Polish Notation):** Push numbers onto stack; on operator, pop two operands, apply operator, push result. Key insight: **first pop is the second operand** (right side of operator) — so `res = b <op> a` where `a` is first pop.
- **LC 2211 (Count Collisions on a Road):** Stack-based simulation for car collision counting. Push R's onto stack; on L, pop top (R→count+2, S→count+1) and push S; on S, drain all R's from stack. **O(1) alternative missed:** editorial's counting observation eliminates stack entirely.
- **LC 1190 (Reverse Substrings Between Each Pair of Parentheses):** Push chars onto stack; on `)`, pop and reverse until `(`, then push reversed chars back. Nested parentheses handled naturally by LIFO — innermost reversed first. O(n²) time due to repeated reversals, O(n) space.

## Key Lessons

- **LC 2211 key insight:** Post-loop logic is trickiest — an S at stack top acts as "barrier" trapping R's. **O(1) alternative missed:** structural observation simplifies to counting.
- **LC 1190 key insight:** Stack naturally handles nested parentheses because LIFO ensures innermost pair is processed first. The "pop until `(`, reverse, push back" pattern is clean and generalizable.
- **LC 1614 key insight:** For single bracket type + guaranteed valid input, use a counter (O(1) space). Stack is reserved for multiple bracket types or validity checking. Don't over-engineer.
- **LC 1111 (secondary):** depth **counter** (not stack) suffices for single bracket type + valid input — extends the LC 1614 lesson. The struggle was the layer-split observation, not stack mechanics → tracked under Greedy (Observation).

## Speed
LC 20: 9:43 (slow pass), LC 150: 20:56, LC 2211: 43:40 (slow pass), LC 1190: 8:05, LC 1614: 23:10 (slow pass).
