# Weaknesses to Watch

## Recurring Issues (ranked by impact)

### 1. Deriving Recurrence from Invariant (10+ struggles)
Currently guessing recurrence from example tables rather than deriving it from the invariant. Practice habit: define invariant → derive recurrence → verify with examples. Two-sentence framework: (1) `dp[i]` **means** what? (2) `dp[i]` **depends on** what?

**Reinforced in:** LC 198, LC 322, LC 139, LC 300, LC 518, LC 53, LC 1406, LC 1140, LC 940, LC 32. The two-question framework was the breakthrough for LC 322 — needs to be applied as the **first** step, not after trying combinations. LC 32: guessed `+2`/`max` mechanics from example tables for 2 hours before the state definition was fixed.

### 2. Over-Engineering / Overthinking
Tendency to build data structures (maps, arrays) when a simpler built-in function or observation would suffice. Before coding, ask: "What's the minimum I need to compute?" Count your variables before submitting, and ask if each one is truly necessary.

**Reinforced in:** LC 14, LC 1268, LC 2685, LC 1331, LC 3499, LC 669, LC 3345, LC 3702, LC 3718, LC 940. LC 3702: built 2D DP table for a problem solvable in O(n) with a 3-case observation. LC 3718: used binary search for membership checks on unsorted array when hash set is simpler. LC 1872: overcomplicated stone game mechanics with suffix sums and nested loops instead of seeking the simpler prefix-sum reformulation. LC 940: built 2D interval DP (`dp[i][j]`) with complex dedup loop when 1D running total suffices. LC 1807: reached for a stack on a bracket problem whose guarantees ("no nested brackets", "every `(` matched") allowed a `capture` bool + key buffer with no stack; also pushed outside-bracket chars onto the stack (only correct because the pop stops at `(`). LC 856: stack evaluator with `(marker, value)` struct + `'-'` sentinel where a plain depth counter suffices (`score += 1 << (depth-1)` at each `()`); post-session self-derived the O(1)-space alternative — the observation was reachable during the session. Positive signals: self-corrected in LC 1268 and LC 3345 via difficulty heuristic ("easy problem → simpler solution exists"). HR Ice Cream Parlor (hr-icecream-parlor): sort + struct-idx + per-iteration clone + delete where one complement map suffices — reached for sorting before asking "what do I actually need to look up?" LC 27: queue impulse on an easy compaction problem; self-corrected via "easy problem → simpler solution exists" heuristic (3rd documented self-correction: LC 1268, LC 3345, LC 27). LC 463: full BFS flood-fill (queue, dirs, visited-state, input mutation) on an easy problem whose guarantees (one island, no lakes) allow a single `4 - landNeighbors` scan — heuristic not invoked before coding (22:44 vs 15-min target); equivalent single-scan alternative written post-review.

### 3. Brute-Force Feasibility Check
Spent ~1 hour on brute-force TLE approaches before recognizing the correct pattern. Root cause: not checking upfront whether brute force can work within the constraints.

**Action:** Add to pre-coding checklist: "How many items am I iterating over? If > 10^7, I need a smarter approach." If high is 10^9, iterating every number is O(10^9) — instant TLE. **Reinforced in LC 3345 (positive twist):** brute force IS the solution when bounded enumeration is O(10·log n). **Reinforced in LC 3720:** started with DFS generating all permutations for n=300 (O(n!)) — constraint check would have immediately ruled this out (n=300 → O(n²) target). **Reinforced in HR Ice Cream Parlor:** committed to O(m·n) (value-domain loop + per-iteration clone) without checking n/m bounds — the constraint table row "n ≤ 10⁶ → O(n) hash map" applied. **LC 219:** correct O(n)-window logic but the validity check was O(map) per step → O(n²) TLE. Cost each *step's* check against constraints, not just the outer loop: n ≤ 10⁵ needs O(n) total, so any per-step helper must be O(1) amortized.

### 4. Debug Print Left in Code (5 occurrences)
`fmt.Println` / `fmt.Printf` left in code after solving. Occurred in LC 216, LC 66, LC 941, LC 13, LC 3734, and LC 678. Before submission, **always scan every line** for `fmt.Print` and remove them.

### 5. Pattern Misidentification
- **LC 1358:** Initially tried prefix/suffix sum when sliding window was the right tool. Ask: "Is this a range query problem or a window property problem?"
- **LC 3336:** Initially tried prefix GCD (range query) when the problem required counting pairs of subsequences (combinatorial DP). Subsequences are not contiguous — prefix structures don't capture the combinatorial nature. Ask: "Am I querying ranges, or counting assignments?"
- **LC 1081:** Initially identified as DP when monotonic stack was correct.
- **LC 3702:** Initially tried DP (running XOR subsequence) when the problem is a pure observation/classification problem. Ask: "Is the operation associative/commutative? If so, order doesn't matter — can I reason about the total and exclusion instead of building up?"
- **LC 1927:** Initially tried to simulate the game with prefix/suffix sums and complex branching logic. The problem is a game theory problem where optimal play reduces to a mathematical invariant. Ask: "What is the condition under which the second player can always match the first player's moves?" For game theory problems, always derive the invariant before trying to simulate.
- **LC 3568:** Initially tried TSP with precomputed pairwise BFS distances + bitmask DP. Failed because the energy constraint means the "cost" from A to B depends on remaining energy — can't precompute distances independently. Ask: "Does the cost between two nodes depend on a variable state?" If yes, incorporate that state into the BFS.
- **LC 1658:** Initially tried DP for "minimum operations to reduce x to zero" when sliding window was correct. Ask: "Is this a subarray sum problem in disguise?"
- **LC 1096:** Initially tried stack-based approach without clear plan for brace expansion. The problem is recursive grammar parsing with precedence rules. Ask: "Is this an expression parsing problem? What are the operators and their precedence?"

### 6. Greedy Instinct
- **LC 2812:** Initially tried greedy pathfinding (pick highest-safety neighbor) which fails for "maximize the minimum" problems. Ask: "Is the locally optimal choice also globally optimal?"
- **LC 486:** Initially tried greedy (always pick larger end) which fails because it ignores opponent's optimal response.
- **LC 1536:** Initially tried greedy (minimize difference between left/right sums at each split). Fails because locally optimal split doesn't guarantee globally optimal score — full game tree must be explored via interval DP. "Stone Game" variants with "Alice divides, Bob chooses" structure are interval DP, not greedy. Ask: "Does my split at step k affect future steps?"

### 7. Solving by Mechanics vs Invariants
Started with implementation mechanics ("what data moves where") instead of stating the invariant first. **Reinforced in:** LC 678 (committed to one wildcard reading before stating the range invariant). Root cause of over-engineering: starting with mechanics instead of "what property does my structure maintain?" **LC 219:** generic "is the window valid?" helper (full-map scan) instead of asking "what changed this step?" — only `nums[right]` changed, so only its count can newly cross 1. The one-sentence answer to that question *is* the invariant. **LC 1631:** built the whole Dijkstra PQ machinery without defining what `dist`/`effort` stores — three successive wrong cost models (sum of heights; `|h[nbr] − effort|` mixing cell & path units; last-edge-only, which passed examples 1–3 by coincidence until a single-row case exposed it). Fixed only after writing the one-sentence invariant. Habit: state "`effort[r][c]` = min over paths of max edge weight" *before* opening the editor. **CF 1473e (d2400, sev 3):** same weakness in the *relaxation loop* rather than the cost model — 3 debug rounds where the guard tested one dist cell (`dist[curr.node][1][1]`) while the write targeted another (`dist[dest][ns][na]`); also destination flags replaced instead of OR-chained (state regressed, "at most once" silently lost), and base cost omitted. Symptom range: all answers = sentinel → infinite zero-cost ping-pong. Fix discipline (applied once prompted, internalization pending): write the invariant sentence "`dist[v][s][a]` = min cost to reach `v` having spent subtract iff `s=1`, add iff `a=1`", then make every loop line *serve* it — one strict condition `newCost < dist[dest][ns][na]` gates both write and push. 4th documented occurrence.

**Reinforced in:** LC 901, LC 14, LC 1268, LC 581, LC 1872, LC 22. LC 1872: jumped straight into simulating stone merges instead of stating the invariant ("remaining stones always form a suffix, leftmost value = some prefix sum"). Before coding any monotonic structure, write one sentence: "My stack maintains ___ from bottom to top." Before coding any game simulation, write one sentence: "The game state can be described as ___." LC 22: started coding an insertion generator + `depth` counter without ever stating what makes a prefix well-formed — cost 4 debug rounds and a guided session (broke DFS mastery). Before coding any backtracking generator, write one sentence: "A partial candidate is valid iff ___" and derive every guard from it. HR Hackerland Radio Transmitters: committed to a position-tiling formula (`ceil(span/(2k+1))` per gap-group) without stating the placement invariant or checking a sparse counterexample — the formula counted positions, not houses; correct model needed "place at rightmost in-range house, skip covered". Habit: before committing to any closed-form, trace one sparse case by hand.

### 8. Modular Arithmetic / Operator Precedence with Percent
Recurring bug in LC 3756 (4+ occurrences in one session). Go percent operator binds tighter than multiply and plus, so the percent MOD must wrap the entire expression. Also, 10e8 is a float64 literal in Go — always use 1_000_000_007 for integer MOD.

### 9. Library Function Shortcut in Interviews
Tendency to reach for built-in functions (e.g., `bits.OnesCount`) when the problem is designed to teach an algorithmic pattern. For interview prep, always try to find the DP/algorithmic pattern first, then mention the built-in as a follow-up.

### 10. Confusing DP State Definition
Distinguishing "subarrays ending at i" from "total subarrays so far" — led to double-counting in LC 3524. When rolling DP, clarify: does dp[r] represent the count at the current step only, or accumulated across all steps? If both are needed, use two arrays (one for ending-at-i, one for running total).

**Reinforced in:** LC 3524, LC 32. LC 32: `dp[i]` as a running max made the recurrence impossible to express (couldn't stitch `() + (())`); flipping to "longest valid **ending exactly at i**" forced the correct recurrence `dp[i] = (i-j) + dp[j]`. Rule: before writing any recurrence, write one sentence defining whether `dp[i]` is **local (ending here)** or **cumulative (so far)** — and pick local unless the problem asks for a total.

### 11. Index Alignment in DP
Confusing table indices (1-indexed) with string indices (0-indexed). Remember: `dp[i][j]` uses `word1[i-1]` and `word2[j-1]`. **Also in 2D string DP:** the two strings have independent index spaces — `s[i-1]` vs `t[j-1]`, never mix them (e.g., writing `t[i-j]`). Always think of `i` and `j` as independent cursors over `s` and `t`.

**Reinforced in:** LC 72, LC 115 (described match as `t[i-j]` instead of `t[j-1]`).

### 12. Boundary Index Completeness
Special-casing one boundary index without checking its mirror. **HR Sherlock and Array:** combine loop `i < n-1` never tested the last element; `i==0` special-cased (`suffix[i+1]==0`) but `i==n-1` (left sum vs empty right = 0) forgotten — `[0,0,5]` returns NO instead of YES; HR tests lacked the case so it was accepted. When special-casing an endpoint, always ask: "what does the condition look like at the *other* endpoint, and does my loop visit it?" Prefer uniform formulations (running-variable `rightSum = total - leftSum - arr[i]`) that handle all indices identically.

### 13. Output-Emission Structure in Counter Solutions
When the invariant/counter is found quickly but the solution still stalls, the gap is often *where to emit output*, not the approach. **LC 1021:** had `balance` immediately, but assumed a counter solution has ONE global append site — spent ~20 min (of 26:10 on an easy) hunting for it. Unlock: derive the emit condition *per branch* from the invariant — each branch asks "is this char interior (depth ≥ 1)?" — `(` appended when `balance > 1` after increment, `)` appended when `balance > 0` after decrement. Rule: when a counter tracks state, write one emit condition per branch ("emit when ___") instead of searching for a single emit point.

**Reinforced in:** LC 1021.

## Pattern-Specific Weaknesses

### Backtracking (LC 22)
- **Node-killer vs branch-gate:** Repeatedly used a whole-node `return` for *choice-limited* conditions (`open == n`) that should gate exactly one branch — silently dropped every legal continuation through the other branch. Rule: a condition that says "this ONE move is unavailable" wraps that branch in `if`; only *invalid state* (`close > open`) may kill the node.
- **Path state leaking out of the path:** Closure-shared `depth` counter (attempt 1) and mutate-then-recurse on `cur` (attempt 2 — Go strings have no pop, so `cur += "("` bled into branch B). Rule: all path state lives in parameters; build children as `cur+"("` / `cur+")"` at the call site.
- **Dedup as a crutch:** `sort` + `slices.Compact` to clean an over-generating insertion scheme. If a generator needs dedup, its choice space is wrong — gate choices at the source instead.

### Binary Search
- **Template confusion:** Not internalizing "closed vs half-open interval" framework. Action: practice templates back-to-back.
- **sort.Search result semantics:** returns the count of elements ≤ target / first index where predicate turns true — the "last valid index" is result-1 (HR Hackerland Radio Transmitters: transmitter placed at `sorted[i+idx]` — first house *outside* range — instead of `sorted[i+idx-1]`). When using sort.Search for "rightmost element satisfying P", the index you want is always `result-1`.

### BFS
- **`continue` vs `if !visited`:** The `if visited[neighbor] { continue }` pattern conflates edge processing with enqueue control. Adopt `if !visited[neighbor] { enqueue }` as default.
- **State-space BFS when resource constraint exists:** When the problem has a resource (energy, fuel) that affects feasibility, TSP / precomputed-distances DP fails. Must incorporate the resource into the BFS state. Use a pruning table (e.g. `bestResource[x][y][mask]`) and only enqueue if strictly better than previously seen.

### Coordinate Convention
- **Key ordering inconsistency:** Stored map keys as `{row, col}` but looked up as `{col, row}` — all items mapped to the same index. Establish and document a convention (e.g., always `{row, col}` or always `{x, y}`) and use it consistently in storage and lookup.

### Binary Lifting
- **Template recall:** Had 2 bugs: wrong start position and missing guard condition. Action: practice the 3-line loop template until muscle memory.

### Linked Lists
- **Identity vs value confusion:** Used hash map keyed on `(value, index)` to detect cycles — fails because values repeat. Only the pointer uniquely identifies a node.
- **Sentinel/dummy node idiom:** Initially tried to manage head separately. Always start with `dummy := &ListNode{}; tail := dummy`.

### Grid Problems
- **Grid visualization for 2D shift/rotation:** Struggled to visualize element movement. Key learning: flatten to 1D and use modulo index mapping.

### Graph Structure Claims
- **Accepting unverified proofs:** In LC 3534, accepted "diameter ≤ 2" without tracing a counterexample. Always construct the worst-case example and trace through it before committing to a structural claim. LC 856: mis-computed `((()))`=8 (it is 4) and fitted a 2^3 formula to the bad number — verify the anchor example by hand before generalizing from it.

### Graph DFS Visited-Marking (LC 1971)
- **Conditional marking inside the neighbor loop:** `visited[curr] = true` was placed inside `for ne := range neighbors` guarded by `if !visited[ne]` — a node whose neighbors were all already visited never got marked and was re-explored from other parents (cycle 0-1-3-2-0 with isolated dest → dfs(2) ran twice). LC still accepted because the bug only wastes work, never flips the answer — **passing tests is not proof of a sound template.** **Invariant:** visited tracks *curr's* state — "I have arrived here, this node is spent" — so mark on arrival at the top of dfs, after the destination check (which also guarantees dest itself is never flagged and skipped). Mechanically-vs-invariant issue: see weakness #7.

### XOR Pattern Recognition
- Couldn't independently identify that XOR is the key for "pairs cancel with O(1) space" problems. Need to internalize: `x ⊕ x = 0` is the cancellation invariant.
- **LC 3702:** Jumped to DP instead of recognizing XOR's commutative/associative properties reduce the problem to totalXor classification. Additional XOR properties to internalize: (1) `a ⊕ 0 = a`, (2) `a ⊕ a = 0`, (3) if `totalXor = 0` and you remove `x ≠ 0`, remaining XOR = `x`.

### Array Operations
- **Ignoring shifting cost:** In LC 88, used sort.Search + shifting which appears O(n log n) but is O(n*m). Always analyze the full cost including data movement.

### Read/Write Pointer Invariant (LC 27)
- **Branch placement on the write pointer:** first attempt mutated `write` in the skip branch (`write = read` when `nums[read] == val`) — breaks on consecutive vals (`[3,3,2]`: the kept element lands one index late). Pointer roles are the invariant: `read` advances every iteration; `write` advances only when an element is kept; the skip branch touches nothing but `read`. Habit: trace one case per branch — here, a consecutive-duplicates case — before submitting (reinforces the HR-missing-numbers trace habit).

### Coordinate Compression
- **Sort + Hash Map pattern:** Couldn't independently recognize "assign ranks based on sorted order" maps to sort + hash map. Internalize: sort a copy → build `map[value] = rank` → map back.

### Two-Sum Recognition (HR Ice Cream Parlor)
- **Sorting before asking "what lookup do I need":** On "pair sums to target, unsorted, return original indices," reached for sort + binary search over the price-value domain instead of a complement map. Recognition signal: pair-sum + unsorted + original indices → `map[value] = earlier index`, check complement before inserting current. Also: `sort.Slice` comparator must be strict weak ordering (`<`, never `<=`) — `<=` is a correctness bug on duplicates.

### Frequency-Map Invariant (HR Missing Numbers)
- **Edge-case trace gap:** the `if key in map { decrement } else { append }` split handled numbers present in `arr`, but the else-branch appended once per occurrence for numbers fully absent (arr=[100], brr=[100 200 200] → [200 200] instead of [200]), violating "include once". Caught in mentor review, fixed post-acceptance with set-based dedup + regression tests. Habit: trace one case per branch of every if/else before submitting.
- **Positive:** frequency-difference recognition was instant (794s, within the 900s easy bar) the same day as the two-sum regression — the frequency variant is solid; two-sum recognition is the remaining gap.

### Mixed Buffered/Unbuffered I/O (CF 1473e)
- **Mixing I/O channels silently loses data:** header read via `bufio.Reader`+`Fscan`, edges read via raw `fmt.Scanf` — the reader had already slurped stdin into its buffer, so `Scanf` hit EOF and read 0 edges. No panic, no error surfaced, every answer printed as the sentinel. Each API is correct alone; the *mix* is fatal because the buffer is invisible state.
- **Detection heuristic:** *every* answer equals the sentinel/MaxInt → suspect input reading before suspecting the algorithm (the input-equivalent of "answers drifting to 1" for the size-k eviction bug).
- **Rule:** one buffered pair (`Fscan` reader / `Fprintln` writer) for the entire program, start to finish. Local end-to-end harness caught this before a wasted CF attempt (positive: the harness is doing its job).
- **Submission-boundary habit:** after local tests pass, confirm the file you submit is byte-identical to the file you tested (diff the CF source panel against the local file); timestamp mismatch between test-run and submit is the same "verified ≠ shipped" family.