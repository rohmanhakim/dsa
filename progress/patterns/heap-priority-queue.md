# Heap (Priority Queue)

**Status:** In Progress — 1/5 sev≤1 streak (day 1)

## Description

A heap gives O(log n) push/pop and O(1) access to the min (or max) element. The key trick: **bound the heap to size k** — for kth-largest problems, use a *min*-heap holding the k largest elements seen, so the root IS the kth largest. Evict the root whenever size exceeds k. O((M+N)·log k) time, O(k) space — far better than sorting the whole stream.

## Sessions

| Date | Problem | Difficulty | Severity | Time | Status | Key Learning |
|------|---------|------------|----------|------|--------|-------------|
| 2026-10-10 | lc703 | easy | 3 (hint) | 0 | guided | Never call the interface's raw `Push`/`Pop` directly — go through the `heap` package |
| 2026-10-10 | abc234d | easy | 3 (hint) | 0 | guided | Invariant: heap must hold *exactly* the k largest of the prefix — missing eviction → root decays to prefix min |
| 2026-10-11 | lc215 | medium | 0 | 7:03 | independent | **First cold rep** — both prior lessons applied unprompted: size-k invariant + `heap.Pop(&h)` correct, inverted sizing recalled |

## Lessons Learned

- **Never call the interface's raw `Push`/`Pop` directly.** Your `Pop()` method is what `heap.Pop` calls *internally* after swapping the root to the end. Calling `h.Pop()` yourself only truncates the slice — it removes the **last array element** (arbitrary, not the min) and skips the sift-down, silently breaking the invariant. Always use `heap.Push(&h, x)` / `heap.Pop(&h)`.
- **State the invariant**: "the min-heap holds the k largest elements of the entire stream → root = kth largest." The bug showed up as every answer returning the global minimum (got `[3 3 3 3 3]` instead of `[4 5 5 8 8]`).
- **One push → at most one eviction.** A `for range h.Len()-k` loop is dead complexity; an `if h.Len() > k` suffices.
- **Min-heap for kth-largest, max-heap for kth-smallest** — inverted sizing is the whole idea, don't reach for a max-heap when asked for largest.
- **Two ways to maintain the size-k invariant (abc234d), same end state:**
  1. **push-then-evict**: `heap.Push` unconditionally, then `if h.Len() > k { heap.Pop(&h) }`. No comparison guard needed — if the new element is unworthy it just pops itself.
  2. **conditional replace-root**: `if x > h[0] { h[0] = x; heap.Fix(&h, 0) }`. Skipping when `x ≤ root` saves a log op; overwriting the root with a *larger* value only needs sift-down, which `Fix` does.
- **State the loop boundaries from the invariant, not by trial.** Off-by-one symptom: duplicate element pushed by two loops. Check: exactly `n−k+1` answers printed, each element pushed exactly once.
- **Root decay diagnosis:** if answers drift toward the global minimum, the heap contains the whole prefix — the size-k eviction is missing or misplaced (this was abc234d's core bug).

## Variants

1. **Kth largest in a stream** (lc703): min-heap of size k, root = answer
2. **Prefix k-th max** (abc234d): for each prefix i≥k, root of the size-k min-heap = k-th largest of the prefix — works because the prefix only grows
3. **Kth smallest in a stream**: max-heap of size k (or min-heap of size n-k+1)
4. **Top-k frequent** (lc347): heap ordered by frequency instead of value
5. **Merge k sorted lists** (lc23): heap of one pointer per list, size k
6. **Streaming median** (lc295): two heaps — max-heap lower half + min-heap upper half
7. **Dijkstra / uniform-cost search**: priority queue ordered by distance
8. **Static-array kth largest** (lc215): single pass over the whole array with the same size-k min-heap (push-then-evict) — no streaming, no sorting; O(n log k)

## Weaknesses

- **Debug prints during development** — `fmt.Println` of heap state on every op (weakness #4; 8th occurrence lc703, 9th abc234d first draft — both removed before finishing)
- **Missing/misplaced size-k eviction** — heap grows to hold the whole prefix, root becomes the prefix minimum, answers decay toward the global min (abc234d). Detection heuristic: "are answers drifting to 1?"
