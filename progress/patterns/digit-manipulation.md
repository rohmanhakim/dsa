# Digit Manipulation

**Status:** ⏳ 2/4 consecutive independent sessions with severity ≤ 1 (2 distinct days).
**Sessions:** LC 3754 (ind, sev 0), LC 1291 (ind, sev 2), LC 3536 (ind, sev 0, 6:14), LC 3348 (guided, sev 3).

## Lessons

- **LC 3348 (Smallest Divisible Digit Product II):** Rem array approach: `rem[i]` = remaining factor the suffix (positions i..n-1) must contribute, computed via `rem[i+1] = rem[i] / gcd(rem[i], num[i])`. Right-to-left scan: at each position, try bumping digit up, then greedily fill suffix (right-to-left, largest digit first) to consume factors on less-significant positions. Required 4 rounds of guided fixes: (1) suffix reversal for correct string order, (2) 1-padding when tNow resolves early to preserve answer length, (3) explicit `hasZero` flag (zero at last position makes `gcd(x,0)=x` → `rem[n]=1` falsely), (4) longer-number fallback (decompose t into digits, sort ascending, pad with leading 1s). Key insight: right-to-left greedy with largest-first ensures leftmost suffix positions get smallest digits, minimizing the overall number. **Complexity:** O(n + D·log₂t) time, O(n) space.
- **LC 3536 (Maximum Value):** Simple digit sum extraction. Clean, fast. Speed: under target.
- **LC 1291 (Digit Divides Value):** Wrong pattern initially (severity 2). Eventually identified digit enumeration approach. Key insight: when the constraint space is huge (10^9) but the answer set is tiny, generate the answer set directly.
- **LC 3754:** First digit manipulation session. Independent, clean.

## Decision Framework

"Do I need to **enumerate/transform digits** (Digit Manipulation) or **compute a property from digits** (Math)?" Digit Manipulation is primary when the problem requires constructing or modifying numbers based on digit-level constraints. Math is primary when the solution derives from a formula or number-theoretic property.

## Weaknesses
- **Greedy suffix fill: reversal and length preservation** — In LC 3348, greedy suffix fill produces digits in discovery order (rightmost position first), but string construction needs left-to-right order. Forgot to reverse. Also, when `tNow` resolves before all suffix positions are filled, remaining positions must be padded with 1s to preserve answer length. **Action:** After any greedy fill that processes positions in reverse order, always reverse the result before string construction.
- **Proxy variable for "no zeros" check** — Used `start == n-1` as proxy for "no zeros found," but when the zero is at the last position, `start` equals `n-1` just like the default. Use an explicit `hasZero` boolean flag instead.