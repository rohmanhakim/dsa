# Progress Overview

## Current Skill Level

### Patterns Mastered
- **2D Dynamic Programming** — Understand DP table setup, base cases, and recurrence relations for string problems (e.g., Edit Distance). Comfortable with the idea of subproblems building on each other via a 2D table.
- **Trie / Prefix Tree** — Implemented standard 26-child trie. Understand the character-by-character navigation pattern shared across insert/search/startsWith. Differentiate between Search (checks `isEnd`) and StartsWith (just checks path existence). Applied to a second problem (LC 1268) with DFS traversal for collecting words — can now build and query tries independently. Starting to recognize over-engineering in DFS and self-correct.
- **Backtracking** — Learned the standard `append → recurse → pop` pattern via LC 17. Each level of recursion is a decision point, each branch is a choice. Base case saves the result. Applied independently to Letter Combinations of a Phone Number (both iterative cartesian product and backtracking). Applied to a second problem (LC 216) with sum constraint — same pattern, different base case. This pattern generalizes to permutations, subsets, N-Queens, Sudoku, etc.
- **Monotonic Stack** — Solved 2 problems (LC 739, LC 901). Understand the "next greater/smaller element" pattern. Key decision framework: "next greater" → decreasing stack (pop when `curr >= top`), "next smaller" → increasing stack (pop when `curr <= top`). Two implementation styles: *span-based* (store value+span, absorb popped spans) for count/distance answers, and *index-based* (store indices, compute from index arithmetic) for identifying actual barrier elements. Invariant-first design principle: state the stack invariant before coding. Recognized as the same pattern family as LC 84 (Largest Rectangle in Histogram). Ready for "next smaller" variant practice.
- **Bit Manipulation / DP** — Applied via LC 338 (Counting Bits). Recurrence: `dp[i] = dp[i>>1] + (i&1)` — count 1s in i/2 plus last bit. Also learned XOR as cancellation operator via LC 136 (Single Number): `x ⊕ x = 0`, `x ⊕ 0 = x`. XOR is the go-to for "pairs cancel" problems. First XOR problem — need more practice to internalize.
- **Array Manipulation / Carry Propagation** — Applied via LC 66 (Plus One). Initialize carry to the value being added (1), propagate through digits using `sum % 10` (digit) and `sum / 10` (carry). Eliminates special-casing the first iteration. Also learned immutability habit: `slices.Clone(digits)` to avoid mutating input, and read consistency (read from and write to the same slice).
### Patterns In Progress
- **Sliding Window (Two Pointers)** — First application via LC 1358. Understand the core idea: maintain a window, expand right, shrink left when valid. Key insight: when a window is valid, all extensions are also valid (superset property). Used `count[3]` array for O(1) validity checks instead of re-scanning. Two counting perspectives learned: count extensions right (`n-right`) or count extensions left (`left`). Need more practice to solidify pattern recognition.
- **Binary Search on Answer ("maximize the minimum")** — Applied twice with guidance: LC 2812 (BFS + binary search) and LC 3620 (topo sort + DP + binary search). Core idea: binary search on the answer + feasibility check. Monotonicity property: if threshold `mid` works, then `mid-1` also works. The feasibility check can use any reachability/search algorithm (BFS, DFS, DP). **Still requires guidance to recognize** — need to internalize: ask "am I maximizing the minimum?" before choosing an approach.
- **DAG Topo Sort + DP** — First application via LC 3620 (with guidance). Post-order DFS topo sort: push node after visiting all neighbors, pop from stack = topological order (no reverse needed). Forward DP: `dp[v] = min over predecessors u of (dp[u] + edge_cost(u,v))`. Combined with binary search for optimization problems on DAGs. **Needs independent practice.**
- **BFS / Connected Component Traversal** — Solved LC 2492 with BFS. Key template learned: `if !visited[neighbor] { enqueue }` instead of `if visited[neighbor] { continue }` — separates edge processing from enqueue control. Understand when BFS suffices (component-level properties like min/max edge weight) vs when backtracking is needed (enumerating distinct paths).

### Weaknesses to Watch
- **Library function shortcut in interviews** — Tendency to reach for built-in functions (e.g., `bits.OnesCount`) when the problem is designed to teach an algorithmic pattern. For interview prep, always try to find the DP/algorithmic pattern first, then mention the built-in as a follow-up. Distinguish between "efficient solution" and "solution that demonstrates understanding."
- **DP case completeness** — Tendency to forget the "match" case in DP recurrences (only handling mismatches). Always ask: "what happens when the condition IS met?"
- **Index alignment in DP** — Confusing table indices (1-indexed) with string indices (0-indexed). Remember: `dp[i][j]` uses `word1[i-1]` and `word2[j-1]`.
- **Deriving recurrence from invariant** — Currently guessing recurrence from the example tables rather than deriving it from the invariant. Practice habit: define invariant → derive recurrence → verify with examples.
- **Solving by mechanics vs invariants** — In LC 901 first attempt, built complex data structures (stack + highestRecords + temp array) by thinking about "what data moves where" instead of "what property does my stack maintain." The root cause of over-engineering: starting with implementation mechanics instead of stating the invariant first. Before coding any monotonic structure, write one sentence: "My stack maintains ___ from bottom to top." Then the popping rule is the only thing that preserves it. Reinforced in LC 901; first appeared as "over-engineering" in LC 14 and LC 1268.
- **Over-engineering / Overthinking** — Tendency to build data structures (maps, arrays) when a simpler built-in function would suffice. Before coding, ask: "What's the minimum I need to compute?" Use `strings.Contains` directly when the problem doesn't require custom substring search. Also reinforced in LC 14 — used 4 state variables when 1 sufficed. Practice: count your variables before submitting, and ask if each one is truly necessary. **Positive signal:** In LC 1268, self-corrected the over-engineered DFS without guidance — pattern recognition is improving.
- **Pattern misidentification** — In LC 1358, initially tried prefix/suffix sum when sliding window was the right tool. Need to ask: "Is this a range query problem or a window property problem?" before committing to an approach. **Positive signal:** After hints, successfully implemented sliding window with correct counting logic.
- **Greedy instinct** — In LC 2812, initially tried greedy pathfinding (pick highest-safety neighbor) which fails for "maximize the minimum" problems. Need to ask: "Is the locally optimal choice also globally optimal?" before committing to a greedy approach.
- **DP objective misidentification** — In LC 3620, initially computed min-cost path (shortest path DP) instead of recognizing the "maximize the minimum" binary search pattern. Need to ask: "What am I optimizing — total cost, or the weakest link?" before choosing a DP approach.
- **Consistent filtering across components** — In LC 3620, DFS filtered edges by threshold but DP didn't, causing subtle bugs. When a filter applies in one place, it must apply everywhere.
- **Map vs slice for performance** — In LC 3620, using `map[int][]edge` and `map[int]bool` for adjacency/visited caused TLE on large inputs. Slice-based data structures are significantly faster for dense integer-keyed access.
- **TLE debugging instinct** — When facing TLE, check if your algorithm visits too much work first (wrong complexity class) before changing data structures. In LC 2492, first instinct was to change the map key type, which didn't address the root cause (map vs array for visited). Ask: "Is my visited check O(1) per call?"
- **`continue` vs `if !visited` in BFS** — The `if visited[neighbor] { continue }` pattern conflates edge processing with enqueue control. Adopt `if !visited[neighbor] { enqueue }` as default — it separates the two concerns and avoids missing edge-level work for already-visited nodes.
- **XOR pattern recognition** — Couldn't independently identify that XOR is the key for "pairs cancel with O(1) space" problems. Had seen the problem before (map solution) but couldn't bridge to the XOR solution. Need to internalize: "When pairs should cancel and space is O(1) → XOR." Build mental connection: `x ⊕ x = 0` is the cancellation invariant.
- **Debug print left in (recurring)** — `fmt.Println` / `fmt.Printf` left in code after solving. Same issue in LC 216 and LC 66. Before submission, always scan for debug prints and remove them. This is a code hygiene habit that needs to become automatic.
- **Shallow copy confusion** — In Go, `res := digits` copies the slice header but shares the underlying array. Modifying `res[i]` also modifies `digits[i]`. Use `slices.Clone(digits)` or `make` + `copy` when you need to avoid mutating the input.

### Practice Plan
- LC 1143 - Longest Common Subsequence (2D DP, similar structure to Edit Distance)
- LC 5 - Longest Palindromic Substring (2D DP, different recurrence)
- LC 64 - Minimum Path Sum (2D DP, simpler)
- **String Matching Algorithms** — Learn Rabin-Karp (rolling hash), KMP, and Aho-Corasick. Practice when to choose which based on constraints. ~~Trie basics~~ ✓ Done (LC 208) ✓ Applied (LC 1268)

### Constraint-Based Algorithm Selection (Reference)
| Constraint Size | Complexity Target | Algorithm Family |
|---|---|---|
| n ≤ 10⁴ | O(n²) | Brute force, nested loops |
| n ≤ 10⁵ | O(n log n) | Sorting, binary search, two pointers |
| n ≤ 10⁶ | O(n) | Sliding window, hash map, BFS/DFS |
| n ≤ 10⁷ | O(n) | Linear scan, monotonic stack |
| n ≤ 10⁸ | O(n log n) | Rare, usually requires optimization tricks |

### Problem Categories Solved
| Category | Count |
|----------|-------|
| 2D DP (Strings) | 1 |
| String Manipulation | 2 |
| Trie / Prefix Tree | 2 |
| Sliding Window | 1 |
| Graph / Grid BFS | 1 |
| Graph / DAG (Binary Search + DP) | 1 |
| Graph / Connected Component (BFS) | 1 |
| Backtracking | 2 |
| Monotonic Stack | 2 |
| Bit Manipulation / DP | 3 |
| Array Manipulation / Carry Propagation | 1 |

### Monthly Progress
- [June 2026](2026-06.md)
- [July 2026](2026-07.md)