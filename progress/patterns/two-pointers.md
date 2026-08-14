# Two Pointers

**Status:** ⏳ 2/5 consecutive independent sessions with severity ≤ 1 (2 distinct days).
**Sessions:** LC 392 (independent, sev 0, 5:27), LC 21 (guided, sev 3), LC 167 (independent, sev 0, 12:35), LC 3983 (guided, sev 3).

## Variants

### Opposite-end scan (LC 392, LC 167)
Left pointer at start, right pointer at end. Move the pointer that's too large inward. Works when sorted order guarantees that moving one pointer eliminates a provably non-optimal candidate. LC 392 used this for subsequence matching; LC 167 used it for two-sum on sorted array.

### Same-direction merge (LC 21)
Two pointers advancing through two sorted linked lists, attaching the smaller element. Dummy node eliminates head-tracking. Key lesson: dummy → tail → attach smaller → advance chosen → return dummy.Next.

### Parallel greedy states (LC 3983)
Maintain two independent match pointers (`j0` without exception, `j1` with exception) in a single pass. `j1 = max(j1, j0+1)` ensures the "with exception" state is always ≥ without-exception + 1. O(m) time, O(1) space. **Pattern:** when the problem asks "can I do X with one exception," maintain two parallel greedy states. The `max` line captures "use exception now" as a single operation.

## Lessons

- **LC 392:** Follow-up asks for map + binary search for repeated queries. Key insight: must store ALL occurrences per character for binary search; `prev` tracks position in `t`, not array index.
- **LC 167:** Sorted property enables O(1) space (no hash map needed). Each pointer movement provably eliminates candidates that cannot be part of the solution.
- **LC 3983:** Did not recognize "subsequence with one replacement" as a parallel-state problem. Built prefix+suffix precomputation approach instead (O(n) space, multiple off-by-one bugs). The official two-pointer solution is 10 lines vs our 60+. **Key recognition signal:** when you can rephrase the problem as "subsequence with one allowed mismatch," think parallel greedy states.