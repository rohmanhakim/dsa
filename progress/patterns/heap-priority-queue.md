# Heap (Priority Queue)

**Status:** In Progress — 0/5 sev≤1 streak

## Description

A heap gives O(log n) push/pop and O(1) access to the min (or max) element. The key trick: **bound the heap to size k** — for kth-largest problems, use a *min*-heap holding the k largest elements seen, so the root IS the kth largest. Evict the root whenever size exceeds k. O((M+N)·log k) time, O(k) space — far better than sorting the whole stream.

## Sessions

| Date | Problem | Difficulty | Severity | Time | Status | Key Learning |
|------|---------|------------|----------|------|--------|-------------|
| 2026-10-10 | lc703 | easy | 3 (hint) | 0 | guided | Never call the interface's raw `Push`/`Pop` directly — go through the `heap` package |

## Lessons Learned

- **Never call the interface's raw `Push`/`Pop` directly.** Your `Pop()` method is what `heap.Pop` calls *internally* after swapping the root to the end. Calling `h.Pop()` yourself only truncates the slice — it removes the **last array element** (arbitrary, not the min) and skips the sift-down, silently breaking the invariant. Always use `heap.Push(&h, x)` / `heap.Pop(&h)`.
- **State the invariant**: "the min-heap holds the k largest elements of the entire stream → root = kth largest." The bug showed up as every answer returning the global minimum (got `[3 3 3 3 3]` instead of `[4 5 5 8 8]`).
- **One push → at most one eviction.** A `for range h.Len()-k` loop is dead complexity; an `if h.Len() > k` suffices.
- **Min-heap for kth-largest, max-heap for kth-smallest** — inverted sizing is the whole idea, don't reach for a max-heap when asked for largest.

## Variants

1. **Kth largest in a stream** (lc703): min-heap of size k, root = answer
2. **Kth smallest in a stream**: max-heap of size k (or min-heap of size n-k+1)
3. **Top-k frequent** (lc347): heap ordered by frequency instead of value
4. **Merge k sorted lists** (lc23): heap of one pointer per list, size k
5. **Streaming median** (lc295): two heaps — max-heap lower half + min-heap upper half
6. **Dijkstra / uniform-cost search**: priority queue ordered by distance

## Weaknesses

- **Debug prints during development** — `fmt.Println` of heap state on every op (removed before finishing; weakness #4, 8th occurrence)
