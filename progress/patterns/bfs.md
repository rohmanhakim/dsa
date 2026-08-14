# BFS

**Status:** ⏳ 2/5 consecutive independent sessions with severity ≤ 1 (2 distinct days).
**Sessions:** LC 2492 (never logged), LC 2685 (guided, sev 3), LC 199 (independent, sev 0, 16:39), LC 3310 (independent, sev 1, 45:27).

Graph reachability from a source node. Key template: `if !visited[neighbor] { enqueue }` instead of `if visited[neighbor] { continue }` — separates edge processing from enqueue control.

## When BFS Suffices vs Backtracking

- **BFS:** Component-level properties (min/max edge weight, shortest path, reachability)
- **Backtracking:** Enumerating distinct paths

## Lessons

- **LC 3310:** After finding the suspicious set via forward BFS, the removal condition ("no method outside the group invokes any method within it") maps directly to iterating edges and checking `if !visited[a] && visited[b]`. Over-engineered by building an undirected BFS and comparing reachability counts instead. Always map the problem statement's condition to a concrete check before designing data structures.
- **LC 2685:** For checking if a connected component is complete, use the degree property — a component with `k` vertices is complete iff every vertex has degree `k-1`. This is O(k) per component instead of O(k²) pair enumeration.
- **LC 199:** BFS naturally gives right side view — track the last node added to the queue at each level.

## Weakness
- **`continue` vs `if !visited` in BFS:** The `if visited[neighbor] { continue }` pattern conflates edge processing with enqueue control. Adopt `if !visited[neighbor] { enqueue }` as default — it separates the two concerns.