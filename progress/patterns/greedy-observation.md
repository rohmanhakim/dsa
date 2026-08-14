# Greedy (Observation)

**Status:** ⏳ 3/5 consecutive independent sessions with severity ≤ 1 (3 distinct days).
**Sessions:** LC 3499 (ind, sev 1, 1:55:00), LC 3517 (ind, sev 0, 13:43), LC 3518 (guided, sev 3), LC 3016 (ind, sev 1, 39:01).

Model an operation's net effect as a formula, then optimize. Key insight: the net effect of a trade (or any complex operation) is often much simpler than the operation itself — the removed part is restored (net zero) and the actual gain comes from a smaller, identifiable component.

## Lessons

- **LC 3499:** Built complex region-scanning mechanism (3 nested loops per transition point + special cases) when the editorial's observation (gain = sum of adjacent zero-block lengths) reduces the problem to a simple one-pass scan. Time: 1:55:00 vs 30-min target. **Key lesson:** Before building complex scanning logic, spend 5 minutes tracing through one example and ask: "What is the net effect of this operation?" If the net effect can be expressed as a formula, use the formula — don't simulate step by step.
- **LC 3501 (guided, 5+ hours):** Same observation extended to multiple range queries. The core formula (cnt1 + bestGain) was already internalized from LC 3499 — difficulty was purely in range query implementation. 10+ edge case bugs: boundary truncation, binary search predicate direction, off-by-one in slice endpoints and RMQ formula, sparse table built per-query causing TLE. **Key lesson:** When extending a single-query solution to multiple queries, separate boundary-finding from computation — compute all boundary values (i, j, zzi, zzj) FIRST, then compute the three cases.