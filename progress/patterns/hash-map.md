# Hash Map

**Status:** ✅ Mastered — 5/5 consecutive independent sessions with severity ≤ 1 across 5 distinct days.
**Sessions:** LC 383, LC 1331, LC 13, LC 49, LC 3731, LC 290, LC 1807 (secondary)

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

## Decision Framework

"Do I need to **count/compare frequencies**, **establish a bijection**, **look up ranks**, or **group by canonical form**?" All four are Hash Map applications but with different map structures.

## Lessons

- **LC 13:** Two-map approach for symbol-to-value mapping with subtraction rules (e.g., IV = 4, not 1+5). Keep the maps simple — one for values, one for subtraction cases.
- **LC 49:** Character frequency array as key is more efficient than sorting each string. 26-int array is the canonical hash map key for anagram problems.
- **LC 3731:** Sort + map for rank assignment (coordinate compression). Clean implementation.
- **LC 290:** Bidirectional bijection with two hash maps. Early length check guard. Check the unmapped direction only when the forward map has no entry.
- **LC 1807 (secondary):** Lookup table for bracket-key → value replacement (`kMap[key]`, comma-ok check → `"?"` fallback). Map build O(m), lookups O(1); the bracket parsing (stack/flag) was the primary pattern.