# Hash Map (Frequency Counting, Lookup Tables)

## When to Use

- **Constant-time lookup** for counting, grouping, or mapping
- **Two-sum pattern:** find complement in O(1)
- **Frequency counting:** character/element frequencies for comparison
- **Grouping:** bucket elements by a key (anagrams, duplicates)

Key signals: "count occurrences," "group by," "find pair with sum," "is anagram"

## Core Templates

### Frequency Counting
```go
freq := map[byte]int{}
for i := 0; i < len(s); i++ {
    freq[s[i]]++
}
```

### Two-Sum Pattern
```go
seen := map[int]int{} // value → index
for i, num := range nums {
    complement := target - num
    if j, ok := seen[complement]; ok {
        return []int{j, i}
    }
    seen[num] = i
}
```

### Grouping by Key
```go
groups := map[string][]string{}
for _, s := range strs {
    key := sortedString(s) // or frequency signature
    groups[key] = append(groups[key], s)
}
```

## Variants

### Frequency Array vs Hash Map
When the alphabet is fixed (e.g., lowercase letters), a `[26]int` array is faster than a map:
- Array: O(1) access, no hashing overhead
- Map: O(1) amortized, but higher constant factor

**Use array when:** alphabet is known and small (a-z, 0-9, ASCII)
**Use map when:** keys are arbitrary integers, strings, or structs

### Sorted String as Key (Anagrams)
Instead of frequency counting, sort each string to create the key:
```go
func sortedKey(s string) string {
    bytes := []byte(s)
    sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
    return string(bytes)
}
```
O(k log k) per string vs O(k) for frequency array, but simpler.

## Common Mistakes

1. **Using map when array suffices** — for fixed alphabets, `[26]int` is faster than `map[byte]int`.
2. **Storing all occurrences when only the last matters** — for two-sum, store the most recent index. For subarray problems, store the most recent prefix sum.
3. **Map key type issues in Go** — `int` keys work, but `float64` doesn't (not comparable). Use `int` for numeric keys.
4. **Forgetting map zero value** — in Go, `freq[ch]` returns 0 for missing keys. This is useful for counting but can cause false matches when checking existence (use `_, ok := freq[ch]` pattern).

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 1** — Two Sum | Complement lookup in O(1) |
| **LC 383** — Ransom Note | Frequency comparison with array |
| **LC 49** — Group Anagrams | Frequency array as key (or sorted string) |
| **LC 13** — Roman to Integer | Map for symbol values + subtraction rules |
| **LC 1331** — Rank Transform | Sort + map for rank assignment |