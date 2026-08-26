# Binary Search

**Status:** ✅ Mastered — 8/8 consecutive independent sessions with severity ≤ 1.
**Solved:** LC 35, LC 704, LC 34, LC 278, LC 153, LC 162, LC 33, LC 3718

## Core Decision

What you do with `mid` determines the template:
- **Eliminate mid** (checked, ruled out) → `left <= right` + `right = mid - 1`; init: `right = len(nums) - 1` (closed)
- **Preserve mid** (could be the answer) → `left < right` + `right = mid`; init: `right = len(nums)` (half-open)

Applied half-open to range-finding (LC 34), "find first X" (LC 278), and rotated array search (LC 153). Applied half-open with left-neighbor comparison to peak finding (LC 162).

## Peak Finding Lesson (LC 162)

Used left-neighbor comparison with `mid == 0` special case — unnecessary complexity. **Right-neighbor comparison** (`nums[mid]` vs `nums[mid+1]`) is standard because `mid+1` is always valid in closed interval, avoiding boundary checks entirely. Convergence variant uses `left < right` + `right = mid` with `right = len(nums) - 1` init — a hybrid that doesn't fit either pure template.

## Decision Framework

"Am I **eliminating** a known-bad candidate, or **preserving** a potentially-good one?"

## Rotated Array Two-Phase vs Single-Pass (LC 33)

LC33 (Search in Rotated Sorted Array) initially attempted single-pass "detect which half is rotated" approach but kept getting WA on edge cases. The correct single-pass rule: if `nums[left] <= nums[mid]`, left half is sorted; check if target is in `[nums[left], nums[mid])` to decide left/right; else right half is sorted. But this case analysis is error-prone under time pressure.

**Reliable fallback:** Two-phase approach — (1) find pivot via `sort.Search` (or custom binary search for min element, same as LC153), (2) binary search the correct half. Both phases are O(log n), and each is conceptually simpler than the single-pass approach. Trade: slightly more code for much less cognitive overhead.

**When to use which:** Single-pass is more elegant but requires mastery of 3-4 case conditions. Two-phase is safer when under time pressure or when the rotation boundary logic is causing bugs. Practice both to build fluency.

## Speed
LC 35: 14:59, LC 704: 7:04, LC 34: 28:49, LC 278: 5:40, LC 153: 17:48, LC 162: 52:19 (slow — edge-case overhead from suboptimal template choice), LC 33: 90:04 (slow — struggled with single-pass approach before settling on two-phase), LC 3718: 10:55

## When NOT to Use Binary Search

Binary search is powerful but not always the right tool. **LC 3718:** Used binary search for membership checks on a small unsorted array (n ≤ 100). Hash set is simpler and O(n) vs O(n log n). Ask before coding: "Is the data already sorted? If not, is there a set/map that does this lookup in O(1)?" Binary search earns its keep when: (1) data is already sorted, (2) you need the position/index, or (3) you're searching an answer space (binary search on answer). For simple "does X exist?" on unsorted data, use a map/set.

## Weaknesses
- **Template confusion:** Spent significant time deciding between `left <= right` vs `left < right`, `right = mid` vs `right = mid - 1`. Root cause: not internalizing the "closed vs half-open interval" framework. **Action:** Practice LC 704 (Template 1), LC 34 (both templates), LC 278 (Template 2 variant) back-to-back to build muscle memory.
- **Single-pass rotated detection:** Cannot reliably implement the "detect which half is sorted, then decide" approach for rotated arrays. The case analysis (comparing `nums[mid]` with `nums[left]` and `nums[right]`) is error-prone. **Action:** Revisit when tackling LC 154 (find min with duplicates), LC 81 (search with duplicates). For now, two-phase approach is a valid alternative.
