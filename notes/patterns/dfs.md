# DFS (Depth-First Search)

## When to Use

- **Tree problems:** Path enumeration, property checking, height/depth computation
- **Permutation/combination generation:** Enumerate all arrangements of elements
- **Grid problems:** Connected components, flood fill, path finding
- **Graph problems:** Cycle detection, topological sort (via post-order), connectivity

Key signals: "all paths," "enumerate," "connected components," "island counting," "backtrack"

## Three Distinct Variants

### Variant A: Tree DFS (Parent-Tracking)

Use `(node, parent)` to avoid revisiting the parent without a visited set:

```go
var dfs func(curr, parent *TreeNode)
dfs = func(curr, parent *TreeNode) {
    if curr == nil { return }
    for _, nb := range adj[curr] {
        if nb == parent { continue }
        dfs(nb, curr)
    }
}
dfs(root, nil)
```

**Why parent-tracking works for trees:** Trees have no cycles — the only "visited" node is the parent. This eliminates the visited map entirely.

**Used in:** LC 104 (max depth), LC 2467 (two-phase tree traversal)

### Variant B: Backtracking DFS (Visited Set)

For permutations/combinations, maintain a visited set and explore all unvisited options:

```go
func backtrack(path []int, visited map[int]bool) {
    if len(path) == n {
        result = append(result, slices.Clone(path))
        return
    }
    for i := 0; i < n; i++ {
        if visited[nums[i]] { continue }
        visited[nums[i]] = true
        path = append(path, nums[i])
        backtrack(path, visited)
        path = path[:len(path)-1]
        visited[nums[i]] = false
    }
}
```

**Key template:** visited set + iterate unchecked + append/recurse/pop

**For combinations:** use a `start` or `num` parameter to enforce forward-only exploration (avoids permutations of the same combination).

**Used in:** LC 46 (permutations), LC 17 (letter combinations), LC 77 (combinations)

### Variant C: Grid DFS (2D Visited Array)

Use a guard clause at the top of DFS to handle boundary, visited, and target checks:

```go
func dfs(grid [][]int, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) { return }
    if visited[i][j] || grid[i][j] != target { return }
    visited[i][j] = true
    // recurse 4 directions directly
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j-1)
}
```

**Outer loop** finds unvisited cells matching the target (each such cell starts a new component).

**No helper functions needed** — grid + visited array is sufficient. Avoid `getNeighbors()` helpers that allocate slices on every call.

**Used in:** LC 200 (num islands), LC 733 (flood fill), LC 1992 (farmland groups)

**LC 733 variant:** Color-swap acts as visited marker — change `oldColor` to `newColor` at each cell. Once changed, cell won't match `oldColor` again. No separate visited array needed.

**LC 1992 variant:** Connected component with bounding box tracking — update min/max coordinates during traversal. Key insight missed: rectangular property of farmland groups enables O(1) space greedy (scan right and down from top-left).

## Key Decision Framework

| Input type | Use | Why |
|---|---|---|
| Tree | Parent-tracking | No cycles, only need to skip parent |
| Permutation/combination | Visited set | Must enumerate all arrangements |
| Grid / connected components | 2D visited array + guard | Grid structure enables boundary checks |
| Graph with cycles | Visited set + edge tracking | Need to distinguish "in progress" from "visited" |

## Common Mistakes

1. **Over-engineering grid DFS** — using `getNeighbors()` helper that allocates a new slice per call. Standard template recurses directly on 4 directions.
2. **Not checking input properties first** — LC 1992's farmland groups are guaranteed rectangular, enabling O(1) space greedy. Always ask: "Does the input have structural guarantees?"
3. **DFS for counting paths when DP suffices** — DFS enumerates all paths (exponential). If you only need path properties (not enumeration), use DP.
4. **Wrong visited state for graphs** — need to distinguish "fully processed" from "currently in recursion stack" for cycle detection in directed graphs.

## Complexity

- **Time:** O(V + E) for graphs, O(m × n) for grids, O(n!) for permutations
- **Space:** O(V) for recursion stack (worst case O(n) for skewed trees, O(m × n) for grids)

## Classic Problems to Practice

| Problem | Variant | Key Idea |
|---|---|---|
| **LC 104** — Maximum Depth of Binary Tree | Tree DFS | Recursive depth tracking |
| **LC 46** — Permutations | Backtracking | Visited set + explore/choose/backtrack |
| **LC 17** — Letter Combinations of Phone Number | Backtracking | Digit-to-letter mapping with standard template |
| **LC 77** — Combinations | Backtracking | `num` parameter for forward-only exploration |
| **LC 200** — Number of Islands | Grid DFS | Nested loop + guard clause + recurse 4 directions |
| **LC 733** — Flood Fill | Grid DFS | Color-swap as visited marker |
| **LC 1992** — Find All Groups of Farmland | Grid DFS | Bounding box tracking during traversal |
| **LC 2467** — Most Profitable Path in a Tree | Tree DFS (two-phase) | Phase 1: Bob's path. Phase 2: Root-to-leaf income |