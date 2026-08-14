# Binary Lifting

**Status:** ⏳ 1/5 sessions toward mastery (guided).
**Sessions:** LC 3534 (guided, severity 3).

Sparse table for answering "shortest path" queries on sorted array graphs in O(log n). Key template:

1. Sort nodes by value, compute `reach[i]` (farthest position in one jump) via binary search
2. Build sparse table: `jump[0] = reach`, `jump[k][i] = jump[k-1][jump[k-1][i]]`
3. Per query: greedy from largest k downward, only take jump if `jump[k][a] < b` (guard prevents overshoot)

## Key Lesson

Initially fell for a false "diameter ≤ 2" proof that broke on counterexample `[0, 8, 14, 18]` with maxDiff=8 (distance = 3, no single intermediate node bridges). Always verify graph diameter claims with a concrete path-trace before committing to an algorithm.

## Weakness
- **Binary lifting template recall:** Had 2 bugs in the binary lifting loop: wrong start position (hardcoded 0 instead of query-dependent `pos[u]`), and missing guard condition (`if jump[k][a] < b`) causing unconditional overshoot. **Action:** Practice the 3-line binary lifting loop template until it's muscle memory: init → for k descending → guard → take jump.