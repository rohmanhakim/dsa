# Hash Map

**Status:** ⏳ Re-certification after regression — 1/3 sev≤1 (1/2 days). REGRESSED 2026-10-07 (sev 3, hr-icecream-parlor: hint needed on two-sum recognition). Re-cert same day: hr-missing-numbers sev 0 (frequency-difference variant, independent, 794s). Two-sum recognition remains the open weakness. Prior "✅ Mastered — 5/5" header was inconsistent with DB (actual primary streak was 4/4: lc1807 logged secondary).
**Sessions:** LC 383, LC 1331, LC 13, LC 49, LC 3731, LC 290, LC 1807 (secondary), HR Ice Cream Parlor (guided, sev 3), HR Missing Numbers (independent, sev 0), LC 219 (secondary, sliding window)

Core use: constant-time lookup for grouping, mapping symbols/characters to values.

## Variants

### Frequency counting (LC 383)
Count character occurrences, compare frequency maps. Simple and clean — no over-engineering needed.

### Bidirectional mapping / bijection (LC 13, LC 290)
Two hash maps for one-to-one correspondence between two domains. LC 290 (Word Pattern): `char2word` + `word2char` for bijection check in O(n). LC 13 (Roman to Integer): value map + subtraction rules map.

### Rank assignment (LC 1331, LC 3731)
Sort + map for rank/coordinate compression. Hash Map supports the lookup, but sorting is the core pattern.

### Canonical key for grouping (LC 49)
Character frequency array (26 counts) as hash map key for anagram grouping — O(n·k) canonical approach.

### Two-sum complement lookup (HR Ice Cream Parlor)
Find two distinct elements summing to a target on an **unsorted** array, return original 1-based indices. Single pass: for each cost, complement = m − cost; look up complement in map of cost → earlier index; if found return (idx+1, i+1); else insert cost → i. O(n) time, O(n) space, no sorting. **Check-before-insert ordering is the invariant** that makes duplicates safe: the map holds only earlier indices, so complement == cost matches a different earlier occurrence — never the element itself. Contrast with Two Pointers (LC 167): same problem family, but sorted input enables the pointer version; unsorted input + index preservation demands the map.

## Decision Framework

"Do I need to **count/compare frequencies**, **establish a bijection**, **look up ranks**, or **group by canonical form**?" All four are Hash Map applications but with different map structures.

## Lessons

- **LC 13:** Two-map approach for symbol-to-value mapping with subtraction rules (e.g., IV = 4, not 1+5). Keep the maps simple — one for values, one for subtraction cases.
- **LC 49:** Character frequency array as key is more efficient than sorting each string. 26-int array is the canonical hash map key for anagram problems.
- **LC 3731:** Sort + map for rank assignment (coordinate compression). Clean implementation.
- **LC 290:** Bidirectional bijection with two hash maps. Early length check guard. Check the unmapped direction only when the forward map has no entry.
- **LC 1807 (secondary):** Lookup table for bracket-key → value replacement (`kMap[key]`, comma-ok check → `"?"` fallback). Map build O(m), lookups O(1); the bracket parsing (stack/flag) was the primary pattern.
- **HR Ice Cream Parlor (sev 3, guided):** Sort + binary search over the price-value domain with per-iteration clone + delete for duplicates — over-engineered AND incorrect (`sort.Slice` comparator used `<=`, violating strict-weak-ordering). The complement map eliminated both failure modes at once. Recognition signal: "pair sums to target + unsorted + return original indices" → `map[value] = earlier index`. Regression cause: needed the hint — two-sum recognition is not yet automatic.
- **HR Missing Numbers (sev 0, re-cert 1/3):** Frequency-difference variant: `missing(x) ⟺ count_brr(x) > count_arr(x)` — one map of arr-counts + decrement sweep + collect keys with negative counts. Latent edge case caught in review: numbers absent from `arr` entirely with multiplicity ≥2 in `brr` were appended once per occurrence (slice-append in else-branch), violating "include once"; fixed post-acceptance with set-based dedup (`map[int32]bool`) so both branches are idempotent, plus regression tests. Hash map version avoids the editorial's max−min ≤ 10000 array assumption; paid O(r log r) sort instead. Pattern recognized instantly — contrast same-day hr-icecream-parlor.