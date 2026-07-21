# Algorithmic Patterns Study Notes

1. [[two-pass-constraint-propagation|Two-Pass Constraint Propagation]] (Left→Right + Right→Left)
2. [[sparse-point-processing|Sparse Point Processing]] (Large Input Space)
3. [[virtual-or-sentinel-points|Virtual / Sentinel Points]]
4. [[greedy-observation|Greedy (Observation)]]
5. [[monotonic-stack|Monotonic Stack]] (Next Greater/Smaller Element)
6. [[backward-two-pointer-merge|Backward Two-Pointer Merge]]
7. [[sliding-window|Sliding Window]]
8. [[coordinate-compression|Coordinate Compression]] (Sort + Hash Map)
9. [[bst|Binary Search Tree (BST)]] (Sorted Tree Properties)
10. [[binary-search|Binary Search]] (Sorted Array / Monotonic Predicate)
11. [[math|Math]] (Arithmetic Series, Number Theory)

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
