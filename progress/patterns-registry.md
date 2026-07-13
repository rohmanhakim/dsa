# Pattern Registry

Canonical pattern names for the `session_patterns` junction table. The mentor (Cline) must use these exact names when logging sessions. Each session can have multiple patterns — the `role` column distinguishes between primary (the skill deliberately practiced) and secondary (supporting patterns).

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
| `Two Pointers (Fast/Slow)` | Floyd's cycle detection: slow moves 1, fast moves 2 | Linked list cycle detection, finding middle, finding cycle start |
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
| `Binary Lifting` | Sparse table / binary jumping for path reachability queries | Shortest path on sorted array graphs, Jump Game II with queries, O(log n) per query |

### Data Structures

| Canonical Name | Description | When to Use |
|---|---|---|
| `Stack` | LIFO data structure for balanced matching, nested lookups | Valid parentheses, expression evaluation, min-stack |
| `Monotonic Stack` | Next greater/smaller element | Daily temperatures, stock span, histogram |
| `BST` | Binary Search Tree properties: in-order = sorted order, search, insert, validate | BST-specific problems: min/max diff, validate, search, kth smallest, LCA in BST |

### Other

| Canonical Name | Description | When to Use |
|---|---|---|
| `Hash Map` | Frequency counting, lookup tables, grouping | Two-sum pattern, anagram grouping, frequency analysis |
| `Trie` | Prefix tree | Word search, autocomplete, prefix matching |
| `Digit Manipulation` | Extracting/processing individual digits | Digit DP, number properties, base conversion |
| `Two-Pass Propagation` | Left→right + right→left passes | Constraint propagation from both directions |
| `Prefix Sum` | Cumulative sums for range queries | Range sum, subarray sum, difference array |

## Migration Note

The `patterns` table in the database is now the source of truth for pattern definitions. This registry is kept as a documentation reference. When adding a new pattern, insert it into the `patterns` table first, then update this registry.

### Adding a new pattern

1. **Discover a new pattern** (e.g., from an editorial or a new problem)
2. **Add it to the `patterns` table**:
   ```sql
   INSERT INTO patterns (name, description, category) 
   VALUES ('Rabin-Karp', 'Rolling hash for string matching', 'String Matching');
   ```
3. **Update this registry** to document the new pattern (for reference)
4. **Use it in the session** (via `session_patterns` junction table)
5. **Track mastery** for that pattern

---

*Last updated: 2026-07-13*

### Changelog

- **2026-07-13**: Added `BST` (Data Structures category) for Binary Search Tree property-based problems.
- **2026-07-10**: Added `Stack` (Data Structures category) for balanced parentheses / expression evaluation problems.
- **2026-07-10**: Added `Two Pointers (Fast/Slow)` (Two Pointers category) for Floyd's cycle detection on linked lists.
