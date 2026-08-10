# Binary Lifting (Sparse Table / Binary Jumping)

## When to Use

- Need to answer **path reachability queries** on a graph (can I get from node A to node B in ≤ X steps?)
- **Shortest path on sorted array graphs** — when edges go from each node to the farthest reachable node
- **Jump Game II with queries** — "what's the minimum jumps to reach position B from position A?"
- Key signal: O(log n) per query needed on a static graph structure

## Core Idea

Precompute a sparse table `jump[k][i]` = the node reached from `i` after 2^k jumps. Each query answers in O(log n) by decomposing the jump distance into binary bits.

## Template (3 steps)

### Step 1: Compute reachability (farthest position in one jump)

For each node `i`, find the farthest node reachable in a single jump:

```go
reach := make([]int, n)
for i := 0; i < n; i++ {
    // binary search or scan for farthest reachable
    reach[i] = farthestReachable(i)
}
```

### Step 2: Build sparse table

```go
jump := make([][]int, K) // K = ceil(log2(n))
jump[0] = reach
for k := 1; k < K; k++ {
    jump[k] = make([]int, n)
    for i := 0; i < n; i++ {
        jump[k][i] = jump[k-1][jump[k-1][i]]
    }
}
```

**Invariant:** `jump[k][i]` = node reached from `i` after 2^k jumps.

### Step 3: Per-query greedy (largest jump first)

```go
func query(a, b int) int {
    pos := a
    steps := 0
    for k := K - 1; k >= 0; k-- {
        if jump[k][pos] < b {  // guard: only jump if not overshooting
            pos = jump[k][pos]
            steps += 1 << k
        }
    }
    return steps + 1  // one more jump to reach b
}
```

**Key guard:** `if jump[k][pos] < b` — only take the jump if it doesn't overshoot the target. This prevents jumping past `b`.

## Common Mistakes

1. **Wrong guard condition** — using `jump[k][pos] <= b` instead of `< b` causes overshoot. The guard ensures we only jump when the destination is strictly before the target.
2. **Hardcoded start position** — the query start position `a` varies per query. Don't hardcode `pos = 0`.
3. **Forgetting the +1** — after the greedy loop, `pos` is at the last node before `b`. One more jump is needed to reach `b`.

## Complexity

- **Build:** O(n log n) — sparse table construction
- **Query:** O(log n) — binary decomposition of jump distance
- **Space:** O(n log n) — sparse table

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 3534** — Process Queries On A Graph | Sort nodes, compute reach via binary search, build sparse table, answer queries |
| **LC 142** — Linked List Cycle II | Floyd's is a form of binary jumping (2x speed) |
| **Jump Game II (LC 45)** | Greedy BFS is equivalent to binary lifting with reach[i] = i + nums[i] |