# Binary Search

**Status:** ✅ Mastered — 6/6 consecutive independent sessions with severity ≤ 1.
**Solved:** LC 35, LC 704, LC 34, LC 278, LC 153, LC 162

## Core Decision

What you do with `mid` determines the template:
- **Eliminate mid** (checked, ruled out) → `left <= right` + `right = mid - 1`; init: `right = len(nums) - 1` (closed)
- **Preserve mid** (could be the answer) → `left < right` + `right = mid`; init: `right = len(nums)` (half-open)

Applied half-open to range-finding (LC 34), "find first X" (LC 278), and rotated array search (LC 153). Applied half-open with left-neighbor comparison to peak finding (LC 162).

## Peak Finding Lesson (LC 162)

Used left-neighbor comparison with `mid == 0` special case — unnecessary complexity. **Right-neighbor comparison** (`nums[mid]` vs `nums[mid+1]`) is standard because `mid+1` is always valid in closed interval, avoiding boundary checks entirely. Convergence variant uses `left < right` + `right = mid` with `right = len(nums) - 1` init — a hybrid that doesn't fit either pure template.

## Decision Framework

"Am I **eliminating** a known-bad candidate, or **preserving** a potentially-good one?"

## Speed
LC 35: 14:59, LC 704: 7:04, LC 34: 28:49, LC 278: 5:40, LC 153: 17:48, LC 162: 52:19 (slow — edge-case overhead from suboptimal template choice).

## Weaknesses
- **Template confusion:** Spent significant time deciding between `left <= right` vs `left < right`, `right = mid` vs `right = mid - 1`. Root cause: not internalizing the "closed vs half-open interval" framework. **Action:** Practice LC 704 (Template 1), LC 34 (both templates), LC 278 (Template 2 variant) back-to-back to build muscle memory.