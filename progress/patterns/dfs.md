# DFS

**Status:** ✅ Mastered — 8/8 consecutive independent sessions with severity ≤ 1 across 5+ distinct days.
**Solved:** LC 104, LC 46, LC 17, LC 77, LC 2467, LC 200, LC 733, LC 1992

## Variants

### Tree DFS (parent-tracking) — LC 104, LC 2467
`(node, parent)` eliminates visited maps entirely for trees. Iterate neighbors, skip if `nb == parent`. Natural fit for depth, path, and property-checking on trees.

### Backtracking DFS (visited set) — LC 46, LC 17, LC 77
Maintain visited set, iterate over unchecked elements, append/recurse/pop. Base case adds result when path is complete. `num` parameter enforces forward-only exploration for combinations.

### Grid DFS (2D visited array) — LC 200, LC 733, LC 1992
`if invalid { return }` guard at start of DFS with boundary, visited, and target checks, then recurse 4 directions directly. Nested loop over grid finds unvisited land cells. No need for map/coord formatting — grid + visited array is sufficient.

**LC 733 variant:** Color-swap itself acts as visited marker — once changed, a cell won't match `oldColor` again. Avoids O(m*n) visited allocation.

**LC 1992 variant:** Connected component with bounding box tracking — update min/max coordinates during traversal to find the rectangle's corners. Same pattern as LC 200 with added coordinate tracking.

## Decision Framework

"Is this a **tree** (use parent-tracking), a **permutation/combination** (use visited set), or a **grid** (use 2D visited array + guard clause)?"

## Lessons

- **LC 2467:** Two-phase DFS for multi-player tree traversal. Phase 1: DFS from Bob to find his unique path to root (parent-tracking). Phase 2: DFS from root exploring all leaf paths, computing income by comparing arrival times. Standard 3-line adjacency list with `append` to nil slices — avoid verbose conditional construction. Parent-tracking DFS is cleaner than global-state closure pattern for trees.
- **LC 104:** Initially attempted iterative traversal before switching to recursive. Call stack tracks depth naturally.
- **LC 46:** Slow pass (34:15) — backtracking with visited map for permutations. Key template: visited set + iterate unchecked + append/recurse/pop.
- **LC 17:** Clean implementation (17:51). Backtracking for letter combinations — digit-to-letter mapping with standard explore/choose/backtrack.
- **LC 77:** Slow pass (47:10) — combinations with `num` parameter for forward-only exploration. Clean but took time.
- **LC 200:** Grid DFS for counting connected components. Used `lands` map + `getNeighbors` helper — over-engineered. Standard template: nested loop over grid, DFS with boundary+visited+target guard at top, recurse 4 directions directly. Need to internalize this template as a third DFS variant alongside tree DFS and backtracking DFS.
- **LC 1992:** Grid DFS for connected components with bounding box tracking. Same over-engineering as LC 200 — used `getNeighbors` helper that allocates a new slice on every call. Standard template recurses directly on 4 directions. Key insight missed: the rectangular property of farmland groups could be exploited for a greedy O(1) space solution (scan right and down from top-left corner to find bottom-right). Always check if the input has structural guarantees that enable a simpler approach.

## Speed
LC 104: 20:47, LC 46: 34:15, LC 17: 17:51, LC 77: 47:10, LC 2467: 01:26:56, LC 200: 29:41, LC 733: 10:24, LC 1992: 26:23.