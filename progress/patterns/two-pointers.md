# Two Pointers

**Status:** ⏳ 1/5 consecutive sessions with severity ≤ 1 (1 distinct day). Streak fresh — anchored at 0 by LC 3983 (guided, sev 3, 2026-08-10); restarted by LC 27.
**Sessions:** LC 392 (independent, sev 0, 5:27), LC 21 (guided, sev 3), LC 167 (independent, sev 0, 12:35), LC 3983 (guided, sev 3), LC 27 (independent, sev 1, 19:35 — Read/Write variant).

## Variants

### Opposite-end scan (LC 392, LC 167)
Left pointer at start, right pointer at end. Move the pointer that's too large inward. Works when sorted order guarantees that moving one pointer eliminates a provably non-optimal candidate. LC 392 used this for subsequence matching; LC 167 used it for two-sum on sorted array.

### Same-direction merge (LC 21)
Two pointers advancing through two sorted linked lists, attaching the smaller element. Dummy node eliminates head-tracking. Key lesson: dummy → tail → attach smaller → advance chosen → return dummy.Next.

### Parallel greedy states (LC 3983)
Maintain two independent match pointers (`j0` without exception, `j1` with exception) in a single pass. `j1 = max(j1, j0+1)` ensures the "with exception" state is always ≥ without-exception + 1. O(m) time, O(1) space. **Pattern:** when the problem asks "can I do X with one exception," maintain two parallel greedy states. The `max` line captures "use exception now" as a single operation.

### Read/write compaction (LC 26, LC 27)
Separate read and write pointers for in-place array compaction. `read` scans every element; `write` marks where the next kept element goes and advances ONLY when an element is kept. Invariant: `nums[0:write]` = kept elements so far, and `write + skipped = read` — so the return value is simply `write`, no second counter needed. LC 26 (guided, sev 3, 2026-07-06) → LC 27 (independent, sev 1, 2026-10-08): recognition improved, speed still over the easy bar. LC 27 used the *stable* variant (order preserved) — more general than required, since the problem allows reordering.

## Lessons

- **LC 392:** Follow-up asks for map + binary search for repeated queries. Key insight: must store ALL occurrences per character for binary search; `prev` tracks position in `t`, not array index.
- **LC 167:** Sorted property enables O(1) space (no hash map needed). Each pointer movement provably eliminates candidates that cannot be part of the solution.
- **LC 3983:** Did not recognize "subsequence with one replacement" as a parallel-state problem. Built prefix+suffix precomputation approach instead (O(n) space, multiple off-by-one bugs). The official two-pointer solution is 10 lines vs our 60+. **Key recognition signal:** when you can rephrase the problem as "subsequence with one allowed mismatch," think parallel greedy states.
- **LC 27:** First attempt mutated `write` in the skip branch (`write = read`) — breaks on consecutive vals (`[3,3,2]`: kept element lands one index late). Pointer role division IS the invariant: skip branch touches only `read`. Mid-session queue impulse self-corrected via "this is easy problem" difficulty heuristic (3rd documented win: LC 1268, LC 3345).