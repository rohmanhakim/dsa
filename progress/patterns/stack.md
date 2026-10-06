# Stack

**Status:** ✅ Mastered — 9/9 consecutive independent sessions with severity ≤ 1 (9 distinct days).

**Sessions:** LC 20 (×2), LC 150, LC 2211, LC 1190, LC 1614, LC 1807, LC 856, LC 921, LC 1111 (secondary), LC 32 (secondary), LC 678 (secondary)

Core idea: LIFO data structure for balanced matching, expression evaluation, and simulation.

## Applications

- **LC 20 (Valid Parentheses):** Push opening brackets onto stack; on closing bracket, check top for valid pair, pop if match, else fail. **Re-attempt (2026-10-01):** single `popAndCheckTop(topShould)` closure DRYs the three closing-bracket branches; final `len(stack) == 0` check subsumes the odd-length early exit — the invariant does the work, no special case needed. Note: `default:` push assumes constraints guarantee only bracket chars.
- **LC 150 (Evaluate Reverse Polish Notation):** Push numbers onto stack; on operator, pop two operands, apply operator, push result. Key insight: **first pop is the second operand** (right side of operator) — so `res = b <op> a` where `a` is first pop.
- **LC 2211 (Count Collisions on a Road):** Stack-based simulation for car collision counting. Push R's onto stack; on L, pop top (R→count+2, S→count+1) and push S; on S, drain all R's from stack. **O(1) alternative missed:** editorial's counting observation eliminates stack entirely.
- **LC 1190 (Reverse Substrings Between Each Pair of Parentheses):** Push chars onto stack; on `)`, pop and reverse until `(`, then push reversed chars back. Nested parentheses handled naturally by LIFO — innermost reversed first. O(n²) time due to repeated reversals, O(n) space.
- **LC 1807 (Evaluate the Bracket Pairs of a String):** Push chars onto stack; on `)`, pop until `(`, reverse → key, then hash-map lookup (`kMap[key]` or `"?"`). No nesting guaranteed, so pop-until-`(` isolates the key exactly. A `capture` flag writes chars outside brackets directly to the builder. O(n+m) time, O(n+m) space.
- **LC 921 (Minimum Add to Make Parentheses Valid):** No stack needed — two counters in one pass: `unclosedLeft` (unmatched `(`, decremented by a `)` when > 0) and `unclosedRight` (orphaned `)` with no opener). Answer = sum of both leftovers. Exactly the editorial's open-bracket counter; O(n) time, O(1) space.

## Key Lessons

- **LC 2211 key insight:** Post-loop logic is trickiest — an S at stack top acts as "barrier" trapping R's. **O(1) alternative missed:** structural observation simplifies to counting.
- **LC 1190 key insight:** Stack naturally handles nested parentheses because LIFO ensures innermost pair is processed first. The "pop until `(`, reverse, push back" pattern is clean and generalizable.
- **LC 1614 key insight:** For single bracket type + guaranteed valid input, use a counter (O(1) space). Stack is reserved for multiple bracket types or validity checking. Don't over-engineer.
- **LC 1807 key insight:** Pop-until-`(` + reverse extracts the key (same mechanic as LC 1190 — direct transfer). **But read the guarantees first:** "no nested brackets" + "every `(` matched" means a `capture` bool + key buffer (the editorial approach) needs **no stack at all**. Reaching for a stack by reflex on a bracket problem is over-engineering (weakness #2).
- **LC 1111 (secondary):** depth **counter** (not stack) suffices for single bracket type + valid input — extends the LC 1614 lesson. The struggle was the layer-split observation, not stack mechanics → tracked under Greedy (Observation).
- **LC 32 (secondary):** stack held **indices**, not chars — a closing `)` needs the matching `(` position to compute pair length and to stitch the block before the pair. Mechanics correct; the session's struggle was DP state derivation → tracked under 1D Dynamic Programming.
- **LC 678 (secondary):** LIFO mechanics were sound; the failure was the wildcard invariant — the single-stack search irreversibly consumed `*`s while looking for `(`, and the end-check never paired leftover `(` with a newer `*`. Tracked under Greedy (Observation). Logged secondary (not primary) so the mastered streak is not nuked by a wrong-tool instinct — judgment recorded 2026-10-04.
- **LC 921 (primary):** went straight to the counter (not a stack) on a bracket problem — the LC 1614 / 1111 / 1807 "single bracket type ⇒ counter" lesson applied **without prompting**, and the two-counter split (leftover openers vs orphaned closers) matched the editorial exactly in 8:48 (sev 0). Weakness #2 (stack-by-reflex over-engineering) did not trigger. Nit: `unclosedRight` name reads as a mirror of `unclosedLeft` but means "needs an added `(`" — ambiguous names slow re-reads.
- **LC 856 (primary):** pop-until-`(` mechanic transferred cleanly from LC 1190/1807 — stack evaluator correct in 22:12 (sev 0). **Missed O(1) observation:** each `()` contributes 2^(# enclosing pairs), so a depth counter (`score += 1 << (depth-1)` at each `()` close) solves it in O(1) space; user re-derived and implemented it post-session (verified vs recursive reference on all balanced strings ≤ len 10). **Lesson:** verify the anchor example before generalizing — a mis-computed `((()))`=8 (actually 4) produced a 2^3 formula that was off by one.

## Speed
LC 20: 9:43 (slow pass) → 10:52 clean re-pass, LC 150: 20:56, LC 2211: 43:40 (slow pass), LC 1190: 8:05, LC 1614: 23:10 (slow pass), LC 1807: 13:25, LC 856: 22:12, LC 921: 8:48.
