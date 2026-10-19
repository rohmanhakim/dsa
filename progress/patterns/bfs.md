# BFS

**Status:** ⏳ 2/5 consecutive independent sessions with severity ≤ 1 (day 2/3: LC 463 sev 1 + LC 994 sev 1; needs 3 more across ≥1 more day).
**Sessions:** LC 2492 (never logged), LC 2685 (guided, sev 3), LC 199 (independent, sev 0, 16:39), LC 3310 (independent, sev 1, 45:27), LC 200 (independent, sev 0, 19:36 — BFS re-solve after DFS), LC 3568 (guided, sev 3 — state-space BFS with energy + bitmask), LC 463 (independent, sev 1, 22:44 — BFS correct, but single-scan solution exists), LC 994 (independent, sev 1, 39:43 — multi-source BFS right first try, slow on level bookkeeping).

Graph reachability from a source node. Key template: `if !visited[neighbor] { enqueue }` instead of `if visited[neighbor] { continue }` — separates edge processing from enqueue control.

## When BFS Suffices vs Backtracking

- **BFS:** Component-level properties (min/max edge weight, shortest path, reachability)
- **Backtracking:** Enumerating distinct paths

## Lessons

- **LC 3310:** After finding the suspicious set via forward BFS, the removal condition ("no method outside the group invokes any method within it") maps directly to iterating edges and checking `if !visited[a] && visited[b]`. Over-engineered by building an undirected BFS and comparing reachability counts instead. Always map the problem statement's condition to a concrete check before designing data structures.
- **LC 2685:** For checking if a connected component is complete, use the degree property — a component with `k` vertices is complete iff every vertex has degree `k-1`. This is O(k) per component instead of O(k²) pair enumeration.
- **LC 199:** BFS naturally gives right side view — track the last node added to the queue at each level.
- **LC 200:** Re-solving a known problem with a different traversal pattern (BFS after DFS) builds fluency. The `mark-on-enqueue` pattern prevents duplicate queue entries — essential for BFS correctness in graph traversal.
- **LC 3568:** When a problem has a resource constraint (energy) that affects movement feasibility, classic TSP / precomputed-distances DP fails — the "cost" from A to B depends on resource state. Must incorporate the resource into the BFS state itself: `(x, y, mask, remainingEnergy)`. The `bestEnergy[x][y][mask]` pruning table tracks max arrival energy per state. Key pitfalls: (1) map key ordering — stored `{row,col}` but looked up `{col,row}`, causing all L cells to map to the same index; (2) pruning comparison must compare like-with-like (arrival energy vs arrival energy); (3) fullMask check must come before energy guard so collecting the last litter with 0 energy still returns the answer.
- **LC 463:** Problem guarantees (*exactly one island*, *no lakes*) can eliminate traversal entirely — one pass with `perim += 4 - landNeighbors` per land cell, O(1) extra space (`_alternative.go`). Shared edges cancel because both sides decrement, so no visited state is needed. Check constraints for such guarantees BEFORE reaching for queue/visited machinery: the "easy problem → simpler solution exists" heuristic applies to choosing the algorithm, not just the data structure. BFS implementation itself was clean (mark-on-enqueue, visited-neighbor decrement correct); the cost was algorithm choice (22:44 vs 15-min easy target).
- **LC 994 (Rotting Oranges):** Multi-source BFS identified on first try — seeding ALL initially-rotten cells into the queue upfront turns per-minute re-scanning into a single O(m*n) BFS where max-distance from any source = answer. Mark-on-enqueue held (`grid[nr][nc] = 2` at enqueue). The cost was level/minute off-by-one debugging (39:43 vs 30-min medium target).

## Weakness
- **`continue` vs `if !visited` in BFS:** The `if visited[neighbor] { continue }` pattern conflates edge processing with enqueue control. Adopt `if !visited[neighbor] { enqueue }` as default — it separates the two concerns.
- **Level/minute off-by-one (LC 994):** State the layer↔minute invariant BEFORE coding, not during debugging: *oranges enqueued during layer k rotted at minute k; level after layer k = k−1*. A `level := -1` + increment-after-layer loop is correct only because it runs one trailing layer after the last fresh orange rots (those distance-ans cells enqueue nothing, bumping level to the answer); stopping the moment the fresh counter hits 0 returns answer−1.