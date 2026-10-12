# Heap (Priority Queue)

**Status:** In Progress — 0/5 sev≤1 (streak reset by lc373 guided sev 3 on 2026-10-11; prior best 1/5 from lc215; lc1631 guided sev 3 on 2026-10-12 — still no live streak)

## Description

A heap gives O(log n) push/pop and O(1) access to the min (or max) element. The key trick: **bound the heap to size k** — for kth-largest problems, use a *min*-heap holding the k largest elements seen, so the root IS the kth largest. Evict the root whenever size exceeds k. O((M+N)·log k) time, O(k) space — far better than sorting the whole stream.

## Sessions

| Date | Problem | Difficulty | Severity | Time | Status | Key Learning |
|------|---------|------------|----------|------|--------|-------------|
| 2026-10-10 | lc703 | easy | 3 (hint) | 0 | guided | Never call the interface's raw `Push`/`Pop` directly — go through the `heap` package |
| 2026-10-10 | abc234d | easy | 3 (hint) | 0 | guided | Invariant: heap must hold *exactly* the k largest of the prefix — missing eviction → root decays to prefix min |
| 2026-10-11 | lc215 | medium | 0 | 7:03 | independent | **First cold rep** — both prior lessons applied unprompted: size-k invariant + `heap.Pop(&h)` correct, inverted sizing recalled |
| 2026-10-11 | lc373 | medium | 3 (hint) | 0 | guided | TLE from exhaustive m·n candidate generation despite correct size-k filter; learned sorted-rows merge (#5) and the lazy-generation ladder to #7 |
| 2026-10-12 | lc1631 | medium | 3 (hint) | 0 | guided | Minimax Dijkstra (variant #7): 3 wrong cost models before the invariant was stated — sum → `|h−effort|` → last-edge-only; fix = candidate `max(edge, effortSoFar)` |

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
- **Selection vs generation (lc373).** A size-bounded heap fixes *filtering*, not *producing*. When k ≪ m·n (10⁴ vs 10¹⁰), the scan that generates candidates is the TLE — replace enumerate-then-filter with generate-on-demand. Check the constraint product *before* designing (weakness #3).
- **Sorted grid → sorted rows → heap-merge (lc373, variant #5).** Rows of `nums1[i]+nums2[j]` are sorted: seed the heap with the first column `(i,0)`, pop k times, push `(i,j+1)` on each pop. No visited set — each row advances linearly. Heap stays O(m); seed the *shorter* array as rows → O(min(m,n)).
- **`Less` must be strict.** `<=` on ties makes `Less(i,j)` and `Less(j,i)` both true → unspecified sort output. Write `if a.sum != b.sum { return a.sum < b.sum }; return a.a < b.a`.
- **Dijkstra's combine operator is the problem's cost operator (lc1631).** Textbook Dijkstra *adds* weights; minimax paths take `max(edgeWeight, effortSoFar)`. Writing `nd := curr.dist + ...` by muscle memory computes the wrong quantity entirely (sum-of-heights gave 11 instead of 2). The invariant sentence "`effort[r][c]` = min over paths of max edge weight" *dictates* the combine — state it before coding.
- **Edge weight comes from two *cells*, never from a cell and the path total.** `|h[nr][nc] − curr.effort|` mixes units (cell height vs accumulated effort) and lets effort *decrease* along a path — impossible for a max of non-negatives. Monotonicity (extending a path never lowers effort) is exactly what makes "first pop of the target is final" valid; if values can drop mid-path, the cost model is wrong.
- **Coincidence passes (lc1631 ex1–3).** Storing only the last-edge weight passed the first three examples because the minimizing path's last edge happened to equal its max; a single-row 7-edge case exposed it. Passing given examples ≠ correct invariant — find the case where the coincidence breaks.

## Variants

1. **Kth largest in a stream** (lc703): min-heap of size k, root = answer
2. **Prefix k-th max** (abc234d): for each prefix i≥k, root of the size-k min-heap = k-th largest of the prefix — works because the prefix only grows
3. **Kth smallest in a stream**: max-heap of size k (or min-heap of size n-k+1)
4. **Top-k frequent** (lc347): heap ordered by frequency instead of value
5. **Merge k sorted lists** (lc23): heap of one pointer per list, size k
6. **Streaming median** (lc295): two heaps — max-heap lower half + min-heap upper half
7. **Dijkstra / uniform-cost search**: priority queue ordered by distance — the combine operator must match the problem's cost: `+` for additive shortest path, `max(edgeWeight, effortSoFar)` for minimax (lc1631 path-with-minimum-effort)
8. **Static-array kth largest** (lc215): single pass over the whole array with the same size-k min-heap (push-then-evict) — no streaming, no sorting; O(n log k)
9. **k smallest pairs from two sorted arrays** (lc373): heap-merge of sorted rows — seed first column, advance col on pop, heap O(m); refinement = lazy Dijkstra frontier (#7) with visited set, heap O(k), O(k log k)

## Weaknesses

- **Debug prints during development** — `fmt.Println` of heap state on every op (weakness #4; 8th occurrence lc703, 9th abc234d first draft — both removed before finishing)
- **Missing/misplaced size-k eviction** — heap grows to hold the whole prefix, root becomes the prefix minimum, answers decay toward the global min (abc234d). Detection heuristic: "are answers drifting to 1?"
- **Constraint check skipped before designing the scan** — lc373 diagonal wavefront ignored m·n = 10¹⁰ (weakness #3). Also: non-strict `<=` inside `Less` on sum ties (lc373, flagged twice).
