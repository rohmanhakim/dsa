# Algorithmic Patterns Study Notes

1. [[two-pass-constraint-propagation|Two-Pass Constraint Propagation]] (Left→Right + Right→Left)
2. [[sparse-point-processing|Sparse Point Processing]] (Large Input Space)
3. [[virtual-or-sentinel-points|Virtual / Sentinel Points]]
4. [[greedy-observation|Greedy (Observation)]]
5. [[monotonic-stack|Monotonic Stack]] (Next Greater/Smaller Element)
6. [[two-pointers|Two Pointers]] (Opposite-End, Same-Direction, Parallel States)
7. [[backward-two-pointer-merge|Backward Two-Pointer Merge]]
8. [[sliding-window|Sliding Window]]
9. [[coordinate-compression|Coordinate Compression]] (Sort + Hash Map)
10. [[bst|Binary Search Tree (BST)]] (Sorted Tree Properties)
11. [[binary-search|Binary Search]] (Sorted Array / Monotonic Predicate)
12. [[math|Math]] (Arithmetic Series, Number Theory)
13. [[minimax-dp|Minimax DP]] (Game Theory)
14. [[floyd-cycle-detection|Floyd's Cycle Detection]] (Fast/Slow Pointers)
15. [[dfs|DFS]] (Depth-First Search)
16. [[bfs|BFS]] (Breadth-First Search)
17. [[hash-map|Hash Map]] (Frequency Counting, Lookup Tables)
18. [[trie|Trie]] (Prefix Tree)
19. [[binary-lifting|Binary Lifting]] (Sparse Table / Binary Jumping)
20. [[segment-tree|Segment Tree]] (Range Queries + Point Updates)

## Quick Reference: Pattern Recognition

| If the problem involves... | Consider... |
|---|---|
| Constraints from both directions | Two-pass propagation |
| Huge input with few events | Sparse point processing + formula |
| Edge cases that need special handling | Virtual/sentinel points |
| "At most one operation" / modeling net effect | Greedy (Observation) |
| Next greater/smaller element for each item | Monotonic stack |
| Merging sorted data into a buffer with space at the end | Backward two-pointer merge |
| Subarray/substring with constraint (unique, count, sum) | Sliding window |
| "Assign ranks" / large value range with small n / compress values to indices | Coordinate compression |
| Input is BST / tree with sorted property / need in-order traversal | BST properties |
| Sorted array / monotonic predicate / O(log n) required | Binary search |
| Sums of consecutive numbers / GCD-related / formula exists | Math |
| "Alice and Bob take turns" / "both play optimally" / game winner | Minimax DP (score difference + subtraction trick) |
| Point updates + range queries on dynamic array | Segment tree (merge function is the algorithm) |
