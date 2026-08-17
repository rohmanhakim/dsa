# Weaknesses to Watch

## Recurring Issues (ranked by impact)

### 1. Deriving Recurrence from Invariant (8+ struggles)
Currently guessing recurrence from example tables rather than deriving it from the invariant. Practice habit: define invariant → derive recurrence → verify with examples. Two-sentence framework: (1) `dp[i]` **means** what? (2) `dp[i]` **depends on** what?

**Reinforced in:** LC 198, LC 322, LC 139, LC 300, LC 518, LC 53, LC 1406, LC 1140. The two-question framework was the breakthrough for LC 322 — needs to be applied as the **first** step, not after trying combinations.

### 2. Over-Engineering / Overthinking
Tendency to build data structures (maps, arrays) when a simpler built-in function or observation would suffice. Before coding, ask: "What's the minimum I need to compute?" Count your variables before submitting, and ask if each one is truly necessary.

**Reinforced in:** LC 14, LC 1268, LC 2685, LC 1331, LC 3499, LC 669, LC 3345, LC 3702. LC 3702: built 2D DP table for a problem solvable in O(n) with a 3-case observation. Positive signals: self-corrected in LC 1268 and LC 3345 via difficulty heuristic ("easy problem → simpler solution exists").

### 3. Brute-Force Feasibility Check
Spent ~1 hour on brute-force TLE approaches before recognizing the correct pattern. Root cause: not checking upfront whether brute force can work within the constraints.

**Action:** Add to pre-coding checklist: "How many items am I iterating over? If > 10^7, I need a smarter approach." If high is 10^9, iterating every number is O(10^9) — instant TLE. **Reinforced in LC 3345 (positive twist):** brute force IS the solution when bounded enumeration is O(10·log n).

### 4. Debug Print Left in Code (4 occurrences)
`fmt.Println` / `fmt.Printf` left in code after solving. Occurred in LC 216, LC 66, LC 941, and LC 13. Before submission, **always scan every line** for `fmt.Print` and remove them.

### 5. Pattern Misidentification
- **LC 1358:** Initially tried prefix/suffix sum when sliding window was the right tool. Ask: "Is this a range query problem or a window property problem?"
- **LC 3336:** Initially tried prefix GCD (range query) when the problem required counting pairs of subsequences (combinatorial DP). Subsequences are not contiguous — prefix structures don't capture the combinatorial nature. Ask: "Am I querying ranges, or counting assignments?"
- **LC 1081:** Initially identified as DP when monotonic stack was correct.
- **LC 3702:** Initially tried DP (running XOR subsequence) when the problem is a pure observation/classification problem. Ask: "Is the operation associative/commutative? If so, order doesn't matter — can I reason about the total and exclusion instead of building up?"

### 6. Greedy Instinct
- **LC 2812:** Initially tried greedy pathfinding (pick highest-safety neighbor) which fails for "maximize the minimum" problems. Ask: "Is the locally optimal choice also globally optimal?"
- **LC 486:** Initially tried greedy (always pick larger end) which fails because it ignores opponent's optimal response.
- **LC 1536:** Initially tried greedy (minimize difference between left/right sums at each split). Fails because locally optimal split doesn't guarantee globally optimal score — full game tree must be explored via interval DP. "Stone Game" variants with "Alice divides, Bob chooses" structure are interval DP, not greedy. Ask: "Does my split at step k affect future steps?"

### 7. Solving by Mechanics vs Invariants
Started with implementation mechanics ("what data moves where") instead of stating the invariant first. Root cause of over-engineering: starting with mechanics instead of "what property does my structure maintain?"

**Reinforced in:** LC 901, LC 14, LC 1268, LC 581. Before coding any monotonic structure, write one sentence: "My stack maintains ___ from bottom to top."

### 8. Modular Arithmetic / Operator Precedence with Percent
Recurring bug in LC 3756 (4+ occurrences in one session). Go percent operator binds tighter than multiply and plus, so the percent MOD must wrap the entire expression. Also, 10e8 is a float64 literal in Go — always use 1_000_000_007 for integer MOD.

### 9. Library Function Shortcut in Interviews
Tendency to reach for built-in functions (e.g., `bits.OnesCount`) when the problem is designed to teach an algorithmic pattern. For interview prep, always try to find the DP/algorithmic pattern first, then mention the built-in as a follow-up.

### 10. Index Alignment in DP
Confusing table indices (1-indexed) with string indices (0-indexed). Remember: `dp[i][j]` uses `word1[i-1]` and `word2[j-1]`.

## Pattern-Specific Weaknesses

### Binary Search
- **Template confusion:** Not internalizing "closed vs half-open interval" framework. Action: practice templates back-to-back.

### BFS
- **`continue` vs `if !visited`:** The `if visited[neighbor] { continue }` pattern conflates edge processing with enqueue control. Adopt `if !visited[neighbor] { enqueue }` as default.

### Binary Lifting
- **Template recall:** Had 2 bugs: wrong start position and missing guard condition. Action: practice the 3-line loop template until muscle memory.

### Linked Lists
- **Identity vs value confusion:** Used hash map keyed on `(value, index)` to detect cycles — fails because values repeat. Only the pointer uniquely identifies a node.
- **Sentinel/dummy node idiom:** Initially tried to manage head separately. Always start with `dummy := &ListNode{}; tail := dummy`.

### Grid Problems
- **Grid visualization for 2D shift/rotation:** Struggled to visualize element movement. Key learning: flatten to 1D and use modulo index mapping.

### Graph Structure Claims
- **Accepting unverified proofs:** In LC 3534, accepted "diameter ≤ 2" without tracing a counterexample. Always construct the worst-case example and trace through it before committing to a structural claim.

### XOR Pattern Recognition
- Couldn't independently identify that XOR is the key for "pairs cancel with O(1) space" problems. Need to internalize: `x ⊕ x = 0` is the cancellation invariant.
- **LC 3702:** Jumped to DP instead of recognizing XOR's commutative/associative properties reduce the problem to totalXor classification. Additional XOR properties to internalize: (1) `a ⊕ 0 = a`, (2) `a ⊕ a = 0`, (3) if `totalXor = 0` and you remove `x ≠ 0`, remaining XOR = `x`.

### Array Operations
- **Ignoring shifting cost:** In LC 88, used sort.Search + shifting which appears O(n log n) but is O(n*m). Always analyze the full cost including data movement.

### Coordinate Compression
- **Sort + Hash Map pattern:** Couldn't independently recognize "assign ranks based on sorted order" maps to sort + hash map. Internalize: sort a copy → build `map[value] = rank` → map back.