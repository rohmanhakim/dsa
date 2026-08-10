# BFS (Breadth-First Search)

## When to Use

- **Shortest path in unweighted graphs** — BFS guarantees the first time you reach a node is via the shortest path
- **Level-order traversal** — processing nodes level by level (tree or graph)
- **Connected components** — finding all nodes reachable from a source
- **Multi-source BFS** — multiple starting points simultaneously (e.g., rotten oranges, multi-source flood fill)

Key signals: "shortest path," "minimum steps," "level order," "nearest," "all nodes at distance k"

## Core Template

```go
func bfs(start int, adj map[int][]int) {
    visited := make([]bool, len(adj))
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        for _, nb := range adj[node] {
            if !visited[nb] {
                visited[nb] = true
                queue = append(queue, nb)
            }
        }
    }
}
```

**Key pattern:** `if !visited[neighbor] { enqueue }` — separates edge processing from enqueue control. Avoid `if visited[neighbor] { continue }` which conflates the two.

## Variants

### Level-Order (with level tracking)

```go
level := 0
for len(queue) > 0 {
    size := len(queue)
    for i := 0; i < size; i++ {
        node := queue[0]
        queue = queue[1:]
        // process node at current level
        for _, nb := range adj[node] {
            if !visited[nb] {
                visited[nb] = true
                queue = append(queue, nb)
            }
        }
    }
    level++
}
```

**Used in:** LC 199 (binary tree right side view — track last node per level)

### Multi-Source BFS

Initialize queue with all source nodes simultaneously:

```go
queue := []int{}
for i := 0; i < n; i++ {
    if sources[i] {
        queue = append(queue, i)
        visited[i] = true
    }
}
// run standard BFS from all sources at once
```

**Used in:** LC 994 (rotting oranges), LC 2812 (BFS for suspicious methods)

## When BFS Suffices vs When DFS Is Needed

| Need | BFS | DFS |
|---|---|---|
| Shortest path (unweighted) | ✅ First visit = shortest | ❌ May find longer path first |
| Enumerate all paths | ❌ | ✅ Natural for backtracking |
| Level-by-level processing | ✅ Natural | ❌ Requires explicit level tracking |
| Memory for wide graphs | ❌ Queue can be large | ✅ Stack is usually smaller |
| Detect cycles | ✅ | ✅ |

## Common Mistakes

1. **Using BFS for weighted graphs** — BFS assumes all edges have equal weight. For weighted graphs, use Dijkstra's.
2. **`if visited[neighbor] { continue }` instead of `if !visited[neighbor] { enqueue }`** — the former skips edge processing for already-visited nodes. If you need to process edges regardless of visit status (e.g., counting edges), use the latter.
3. **Forgetting to mark visited on enqueue** — marking on dequeue causes nodes to be enqueued multiple times, leading to O(n²) in worst case.
4. **Using map for visited on dense graphs** — use slice `[]bool` for O(1) access. Maps have higher constant factor.

## Complexity

- **Time:** O(V + E) — each node and edge processed once
- **Space:** O(V) — queue + visited array

## Classic Problems to Practice

| Problem | Variant | Key Idea |
|---|---|---|
| **LC 199** — Binary Tree Right Side View | Level-order | Track last node added at each level |
| **LC 994** — Rotting Oranges | Multi-source | All rotten oranges start simultaneously; count levels |
| **LC 2685** — Count Complete Components | BFS + degree check | Component with k vertices is complete iff every vertex has degree k-1 |
| **LC 3310** — Remove Methods from Project | BFS reachability | Forward BFS from k finds suspicious set; check edges |
| **LC 2812** — Safest Path in a Grid | BFS + binary search | BFS to compute max-safety path for each threshold |