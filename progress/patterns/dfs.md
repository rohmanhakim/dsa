# DFS

**Status:** ⚠️ Regressed (2026-10-02) — sev-3 guided session on LC 22 broke the 9/9 streak. Re-cert bar: 3 consecutive sev≤1 across 2 distinct days. Re-cert progress: **1/3 sev≤1, day 1/2** (LC 1971 sev 1, 2026-10-08).
**Solved:** LC 104, LC 46, LC 17, LC 77, LC 2467, LC 200, LC 733, LC 1992, LC 39, LC 22, LC 1971

## Variants

### Tree DFS (parent-tracking) — LC 104, LC 2467
`(node, parent)` eliminates visited maps entirely for trees. Iterate neighbors, skip if `nb == parent`. Natural fit for depth, path, and property-checking on trees.

### Backtracking DFS (visited set / reuse) — LC 46, LC 17, LC 77, LC 39
Maintain visited set, iterate over unchecked elements, append/recurse/pop. Base case adds result when path is complete. `num` parameter enforces forward-only exploration for combinations.

**LC 39 variant (reuse same element):** Pass `i` (not `i+1`) to recursive call to allow picking the same candidate again. No visited set needed — the `idx` parameter both prevents duplicate combinations and enables reuse. Prune branches where `total + candidates[i] > target`.

**LC 22 variant (choice-gated generation, no visited set):** Build strings one char at a time with two branch gates derived from the prefix invariant `close <= open <= n`: recurse on `cur+"("` only if `openCount < n`, on `cur+")"` only if `closeCount < openCount`. Collect at `len == 2n` (both counts pinned to n) — the only node-killing `return` besides the base case. Critical distinction: *invalid state* (`close > open`) kills the whole node; *choice-limited* state (`open == n`) gates exactly ONE branch — using `return` for the latter silently drops every legal `)` placement. Go strings have no pop: pass `cur+"("`/`cur+")"` as call arguments, never mutate-then-recurse. Exploring `(` first yields lexicographic output for free.

### Grid DFS (2D visited array) — LC 200, LC 733, LC 1992
`if invalid { return }` guard at start of DFS with boundary, visited, and target checks, then recurse 4 directions directly. Nested loop over grid finds unvisited land cells. No need for map/coord formatting — grid + visited array is sufficient.

**LC 733 variant:** Color-swap itself acts as visited marker — once changed, a cell won't match `oldColor` again. Avoids O(m*n) visited allocation.

**LC 1992 variant:** Connected component with bounding box tracking — update min/max coordinates during traversal to find the rectangle's corners. Same pattern as LC 200 with added coordinate tracking.

### General Graph DFS (adjacency list + visited) — LC 1971
Undirected graph that is neither tree nor grid: build a bi-directional adjacency list from `edges`, recursive DFS with a visited set, early return on `curr == destination`. **Mark visited on arrival** — top of dfs, after the destination check — never conditionally inside the neighbor loop: visited tracks *curr's* state ("I have arrived, this node is spent"), not its neighbors' state. Marking only when an unvisited neighbor is found means a node whose neighbors were all already visited never gets marked → re-explored from other parents (cycle 0-1-3-2-0 with isolated dest → `dfs(2)` ran twice). The destination flag can never cause a false skip: the mark line is unreachable for destination (early return fires first). For 0..n-1 labels use `make([][]int, n)` / `make([]bool, n)`, not maps. n ≤ 2×10⁵ → recursion depth up to 2×10⁵ on path-shaped graphs; iterative DFS or BFS sidesteps the stack question.

## Decision Framework

"Is this a **tree** (use parent-tracking), a **permutation/combination** (use visited set), a **grid** (use 2D visited array + guard clause), or a **general graph** (adjacency list + visited marked on arrival)?"

## Lessons

- **LC 2467:** Two-phase DFS for multi-player tree traversal. Phase 1: DFS from Bob to find his unique path to root (parent-tracking). Phase 2: DFS from root exploring all leaf paths, computing income by comparing arrival times. Standard 3-line adjacency list with `append` to nil slices — avoid verbose conditional construction. Parent-tracking DFS is cleaner than global-state closure pattern for trees.
- **LC 104:** Initially attempted iterative traversal before switching to recursive. Call stack tracks depth naturally.
- **LC 46:** Slow pass (34:15) — backtracking with visited map for permutations. Key template: visited set + iterate unchecked + append/recurse/pop.
- **LC 17:** Clean implementation (17:51). Backtracking for letter combinations — digit-to-letter mapping with standard explore/choose/backtrack.
- **LC 77:** Slow pass (47:10) — combinations with `num` parameter for forward-only exploration. Clean but took time.
- **LC 200:** Grid DFS for counting connected components. Used `lands` map + `getNeighbors` helper — over-engineered. Standard template: nested loop over grid, DFS with boundary+visited+target guard at top, recurse 4 directions directly. Need to internalize this template as a third DFS variant alongside tree DFS and backtracking DFS.
- **LC 39:** Clean implementation (21:01). Combination Sum — backtracking DFS with element reuse. Key distinction from LC 46/17/77: pass `i` instead of `i+1` to allow unlimited reuse of the same candidate. No visited set needed since `idx` controls exploration scope. This is the third backtracking variant: (1) visited set for permutations (LC 46), (2) `i+1` for combinations (LC 77), (3) `i` for combinations with reuse (LC 39).
- **LC 22:** Regressed session (guided, sev 3, 01:45:45 wall). Own attempt used insertion generation + sort/Compact dedup (over-generates → dedup hacks; closure-shared `depth` corrupted branch termination). Guided to forward-build backtracking, then 3 more debug rounds: rebound `cur` local (Go strings no pop), `open >= n` used as node-killer instead of branch gate, B-gate used `close > open` instead of `close < open`. Core lesson: classify each condition as *invalid state* vs *choice-limited* before coding it. Related to weakness #7 (mechanics before invariant).
- **LC 1992:** Grid DFS for connected components with bounding box tracking. Same over-engineering as LC 200 — used `getNeighbors` helper that allocates a new slice on every call. Standard template recurses directly on 4 directions. Key insight missed: the rectangular property of farmland groups could be exploited for a greedy O(1) space solution (scan right and down from top-left corner to find bottom-right). Always check if the input has structural guarantees that enable a simpler approach.
- **LC 1971:** General graph DFS (does a path exist source→destination), slow pass 29:03, sev 1 — right pattern identified immediately; the flaw was **where** visited was marked. `visited[curr] = true` sat inside the neighbor loop under `!visited[ne]`, so a node whose neighbors were all already visited was never marked → re-explored from other parents (cycle 0-1-3-2-0, isolated dest → dfs(2) ran twice). Accepted anyway: this bug only wastes work, never flips the answer — **LC passing is not proof of a sound template.** Found in review, fixed post-acceptance by moving the mark to the top of dfs (on arrival, after the dest check, which also guarantees dest can never be flagged and skipped). Ties to weakness #7 — the invariant ("visited = I have arrived here") needed stating explicitly instead of inferred from mechanics.

## Speed
LC 104: 20:47, LC 46: 34:15, LC 17: 17:51, LC 77: 47:10, LC 2467: 01:26:56, LC 200: 29:41, LC 733: 10:24, LC 1992: 26:23, LC 39: 21:01, LC 22: guided (01:45:45, not counted), LC 1971: 29:03 (sev 1).
