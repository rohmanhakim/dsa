# Algorithmic Patterns Study Notes

1. [[two-pass-constraint-propagation|Two-Pass Constraint Propagation]] (Left→Right + Right→Left)
2. [[sparse-point-processing|Sparse Point Processing]] (Large Input Space)
3. [[virtual-or-sentinel-points|Virtual / Sentinel Points]]
4. [[greedy-geometric-observation|Greedy Geometric Observation]]
5. [[monotonic-stack|Monotonic Stack]] (Next Greater/Smaller Element)
6. [[backward-two-pointer-merge|Backward Two-Pointer Merge]]
7. [[sliding-window|Sliding Window]]

## Quick Reference: Pattern Recognition

| If the problem involves... | Consider... |
|---|---|
| Constraints from both directions | Two-pass propagation |
| Huge input with few events | Sparse point processing + formula |
| Edge cases that need special handling | Virtual/sentinel points |
| An optimal shape that's geometrically obvious | Greedy geometric formula |
| Next greater/smaller element for each item | Monotonic stack |
| Merging sorted data into a buffer with space at the end | Backward two-pointer merge |
| Subarray/substring with constraint (unique, count, sum) | Sliding window |
