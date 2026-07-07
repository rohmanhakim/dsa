# Pattern Registry

Canonical pattern names for the `sessions.pattern` column. The mentor (Cline) must use these exact names when logging sessions. Each session has **one primary pattern** — the skill deliberately practiced. Secondary patterns go in `approach_notes`.

## Naming Convention

- Format: `Technique` or `Technique (Variant)`
- No slashes or compound names
- Parenthetical variants only for distinct sub-techniques worth tracking separately

## Patterns

### Two Pointers

| Canonical Name | Description | When to Use |
|---|---|---|
| `Two Pointers` | General two-pointer technique on arrays/strings | Opposite-end convergence, partitioning, pair-finding |
| `Two Pointers (Backward Merge)` | Fill from end to avoid O(n) shifting | Merging sorted data into buffer with trailing empty space |
| `Two Pointers (Read/Write)` | Separate read and write pointers | In-place array compaction, deduplication |
| `Sliding Window` | Expand right, shrink left to maintain window property | Subarray/substring with constraint (count, sum, unique) |

### Binary Search

| Canonical Name | Description | When to Use |
|---|---|---|
| `Binary Search` | Standard search on sorted data or monotonic predicate | Find element, find insertion point, template 1/2 |
| `Binary Search on Answer` | Binary search on answer space + feasibility check | "Maximize the minimum" / "Minimize the maximum" |

### Greedy

| Canonical Name | Description | When to Use |
|---|---|---|
| `Greedy (Intervals)` | Interval scheduling, activity selection, merge intervals | Sort by endpoint, non-overlapping, coverage |

### Dynamic Programming

| Canonical Name | Description | When to Use |
|---|---|---|
| `1D Dynamic Programming` | Single-dimensional DP table | Sequence DP, house robber, climbing stairs |
| `2D Dynamic Programming` | Two-dimensional DP table | Edit distance, LCS, grid paths, knapsack |

### Graph

| Canonical Name | Description | When to Use |
|---|---|---|
| `BFS` | Breadth-first search, connected components | Shortest path (unweighted), level-order, component properties |
| `DFS` | Depth-first search, backtracking | Path enumeration, cycle detection, connected components |
| `Topological Sort` | DAG ordering | Task scheduling, dependency resolution, DP on DAG |

### Other

| Canonical Name | Description | When to Use |
|---|---|---|
| `Hash Map` | Frequency counting, lookup tables, grouping | Two-sum pattern, anagram grouping, frequency analysis |
| `Monotonic Stack` | Next greater/smaller element | Daily temperatures, stock span, histogram |
| `Trie` | Prefix tree | Word search, autocomplete, prefix matching |
| `Digit Manipulation` | Extracting/processing individual digits | Digit DP, number properties, base conversion |
| `Two-Pass Propagation` | Left→right + right→left passes | Constraint propagation from both directions |
| `Prefix Sum` | Cumulative sums for range queries | Range sum, subarray sum, difference array |

## Migration Note

When migrating from a junction table (if needed in the future), this registry becomes the `patterns` table with `name` and `description` columns. The `session_patterns` junction table would reference `patterns.name`.

---

*Last updated: 2026-07-07*