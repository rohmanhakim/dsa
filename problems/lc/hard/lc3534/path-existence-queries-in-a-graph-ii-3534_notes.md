# LC 3534 — Path Existence Queries in a Graph II: Study Notes

## Problem Summary

Given `n` nodes (labeled 0 to n-1) with values `nums[i]`, an undirected edge exists between nodes `i` and `j` if `|nums[i] - nums[j]| ≤ maxDiff`. For each query `[u, v]`, find the **minimum number of edges** (shortest path) from `u` to `v`, or `-1` if unreachable.

**Constraints:** n ≤ 10⁵, queries.length ≤ 10⁵ → need O(log n) or O(1) per query after preprocessing.

---

## Key Insight: Sort + Binary Lifting (Jump Game II)

**Step 1 — Sort nodes by value.** This transforms the problem from an arbitrary graph into a structure where each node's reachable range is a contiguous segment in sorted order.

**Step 2 — Observe:** In the sorted array, the shortest path from sorted position `a` to sorted position `b` (a < b) only moves **rightward**. Going left then right is never shorter, because if you're at position `p` and go left to `q < p`, the rightward reach from `q` is at most the rightward reach from `p` (since `nums[q] ≤ nums[p]` implies `nums[q]+maxDiff ≤ nums[p]+maxDiff`).

**Step 3 — Reduce to Jump Game II:** Each sorted position `i` can "jump" to any position within `maxDiff`. We want the minimum number of jumps from `a` to reach `≥ b`.

**Step 4 — Binary lifting** answers "what's the farthest I can reach in 2^k jumps?" in O(1) per query.

---

## Why "Diameter ≤ 2" Is Wrong (Common Pitfall)

It's tempting to claim: "any two nodes in the same connected component have distance ≤ 2, because there's always a middle node bridging them."

**Counterexample:** `nums = [0, 8, 14, 18]`, `maxDiff = 8`

Sorted: `[0, 8, 14, 18]`. All consecutive diffs ≤ 8, so all in one component.

Query: node with value 0 → node with value 18.

- Can any single intermediate node bridge? 
  - val 8: `|0-8|=8 ≤ 8` ✓, but `|8-18|=10 > 8` ✗
  - val 14: `|0-14|=14 > 8` ✗
- **No single node connects to both!** Shortest path: `0 → 8 → 14 → 18` (3 hops).

The proof fails because `nums[w] - nums[v-1]` can be up to `2 × maxDiff` (not ≤ maxDiff) when there's a gap between `v-1` and `v`.

---

## Algorithm (Step by Step)

### Preprocessing (O(n log n))

```
1. Sort nodes by value → sortedIdx[]
2. Map original index → sorted position → pos[]
3. For each sorted position i, compute reach[i]:
     reach[i] = farthest sorted position j where nums[j] - nums[i] ≤ maxDiff
     (binary search: find first j where nums[j] - nums[i] > maxDiff, then j-1)
4. Build sparse table:
     jump[0][i] = reach[i]
     jump[k][i] = jump[k-1][ jump[k-1][i] ]
5. Compute prefix breaks for connectivity:
     breaks[i] = 1 if nums[i+1] - nums[i] > maxDiff (consecutive sorted values)
     prefixBreaks[i] = cumulative sum of breaks
     Two positions are in the same component iff their prefixBreaks values are equal
```

### Per Query (O(log n))

```
if u == v → answer is 0
if prefixBreaks[pos[u]] ≠ prefixBreaks[pos[v]] → answer is -1 (different components)
otherwise:
  a = min(pos[u], pos[v])
  b = max(pos[u], pos[v])
  steps = 0
  for k from LOG-1 down to 0:
      if jump[k][a] < b:     ← GUARD: only jump if we haven't reached b yet
          a = jump[k][a]
          steps += 2^k
  answer = steps + 1         ← one final jump to reach ≥ b
```

---

## Walkthrough: Example 5

```
nums = [18, 14, 8, 18, 0], maxDiff = 8
query = [4, 0] (node 4 has val 0, node 0 has val 18)
```

**Sorted values:** `[0, 8, 14, 18, 18]` → sorted indices: `[4, 2, 1, 0, 3]`

**reach[i]** (farthest position reachable in one jump):
| pos | val | reach |
|-----|-----|-------|
| 0 | 0 | 1 (val 8, diff=8) |
| 1 | 8 | 2 (val 14, diff=6) |
| 2 | 14 | 4 (val 18, diff=4) |
| 3 | 18 | 4 (val 18, diff=0) |
| 4 | 18 | 4 |

**Sparse table:**
- `jump[0] = [1, 2, 4, 4, 4]`
- `jump[1] = [2, 4, 4, 4, 4]`  (jump[1][0] = jump[0][jump[0][0]] = jump[0][1] = 2)

**Query: pos 0 → pos 3** (sorted positions of nodes 4 and 0):
- k=1: `jump[1][0] = 2 < 3` → take it. `a=2, steps=2`
- k=0: `jump[0][2] = 4 ≥ 3` → skip
- Answer: `2 + 1 = 3` ✓

Path: node 4 (val 0) → node 2 (val 8) → node 1 (val 14) → node 0 (val 18)

---

## Binary Lifting Template (Reusable)

```go
LOG := 0
for (1 << LOG) <= n {
    LOG++
}

// Build reach[] via binary search, then:
jump := make([][]int, LOG)
jump[0] = reach
for k := 1; k < LOG; k++ {
    jump[k] = make([]int, n)
    for i := 0; i < n; i++ {
        jump[k][i] = jump[k-1][jump[k-1][i]]
    }
}

// Query: min jumps from sorted position a to reach ≥ b
steps := 0
for k := LOG - 1; k >= 0; k-- {
    if jump[k][a] < b {
        a = jump[k][a]
        steps += 1 << k
    }
}
answer = steps + 1
```

**Three things to remember:**
1. `a = min(pos[u], pos[v])`, `b = max(pos[u], pos[v])` — always go left-to-right in sorted order
2. Guard condition: `if jump[k][a] < b` — only take the jump if it doesn't reach the target yet
3. Return `steps + 1` — the loop gets you as close as possible; one more jump finishes

---

## Complexity

| Phase | Time | Space |
|-------|------|-------|
| Sort | O(n log n) | O(n) |
| reach[] (binary search) | O(n log n) | O(n) |
| Sparse table | O(n log n) | O(n log n) |
| Prefix breaks | O(n) | O(n) |
| Per query | O(log n) | — |
| **Total** | **O(n log n + q log n)** | **O(n log n)** |

---

## Lessons Learned

1. **Always verify graph diameter claims with a concrete counterexample.** The "diameter ≤ 2" proof sounded plausible but was wrong. Before committing to a structural property, build the worst-case example and trace through it.

2. **Binary lifting = Jump Game II with queries.** When the problem asks "minimum jumps to reach from A to B" on a structure where each position has a contiguous reachable range, think binary lifting / sparse table.

3. **The guard condition is everything.** `if jump[k][a] < b` prevents overshooting. Without it, you unconditionally take all jumps and get the wrong answer. This is analogous to the guard in binary search (`left <= right`).

4. **Only go rightward in sorted order.** In an undirected graph where edges connect nodes within `maxDiff`, the shortest path between two nodes only uses edges that go rightward in the sorted array. Going left then right is never shorter because the rightward reach from a position to the left is always ≤ the rightward reach from the current position.

5. **Prefix breaks for component connectivity.** After sorting, connected components are contiguous segments. Two nodes are in the same component iff there are no "breaks" (consecutive sorted values with diff > maxDiff) between them. Prefix sum of boolean break flags gives O(1) connectivity checks.