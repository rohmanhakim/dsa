# Math (Arithmetic Series, Number Theory)

## When to Use

- Problem involves **sums of consecutive numbers** (odd, even, natural numbers, squares)
- Problem asks for **GCD/GCD-related** and there's a mathematical property that simplifies it
- Problem asks to **count pairs by GCD value** — inclusion-exclusion on multiples
- Recognizing a formula makes the solution O(1) instead of O(n)

## Core Idea

Many problems with "sum of first n [something]" have a closed-form formula. Always check before iterating.

## Common Formulas

| Series | Formula |
|--------|---------|
| 1 + 2 + ... + n | n(n+1)/2 |
| 1 + 3 + 5 + ... + (2n-1) | n² |
| 2 + 4 + 6 + ... + 2n | n(n+1) |
| 1² + 2² + ... + n² | n(n+1)(2n+1)/6 |

## Key Properties

- **GCD(n, n+1) = 1** — consecutive integers are always coprime
- **GCD(a, b) = GCD(b, a%b)** — Euclidean algorithm (O(log(min(a,b))))
- **GCD(a, b) = GCD(k*a, k*b) / k** — scaling property

## Recognition Signals

| Signal | Why Math? |
|---|---|
| "Sum of first n [odd/even/natural] numbers" | Arithmetic series formula exists |
| "GCD of two expressions" | GCD properties may simplify |
| "Return a formula result" | No need to iterate |
| Constraints allow O(1) but you used O(n) | Formula missed |

## Example: LC 3685 (GCD of Odd and Even Sums)

Problem: compute GCD(sumOdd, sumEven) where sumOdd = first n odd numbers, sumEven = first n even numbers.

- sumOdd = 1 + 3 + 5 + ... + (2n-1) = n²
- sumEven = 2 + 4 + 6 + ... + 2n = n(n+1)
- GCD(n², n(n+1)) = n × GCD(n, n+1) = n × 1 = n

Answer is always **n**. O(1) solution.

## Sub-Variant: Inclusion-Exclusion for GCD Pairs (LC 3312)

### When to Use
- Problem asks to count pairs `(i, j)` where `gcd(nums[i], nums[j]) = g` for various `g`
- n is large (10⁵) but value range is small (max = 5×10⁴) — enumerate values, not pairs
- Signals: "count pairs with GCD = X", "k-th smallest GCD", sorted GCD queries

### How It Works (4 steps)

**Step 1: Count by value**
```
cnt[v] = frequency of value v in nums
```

**Step 2: Count multiples (harmonic series)**
For each `g` from 1 to max, count how many elements are divisible by g:
```
divCount[g] = cnt[g] + cnt[2g] + cnt[3g] + ...
```
Total work: O(m log m) — harmonic series.

**Step 3: "At least g" → "exactly g" via inclusion-exclusion**
- Pairs with GCD **≥ g** = C(divCount[g], 2) = divCount[g] × (divCount[g]-1) / 2
- Pairs with GCD **exactly g** = pairs with GCD ≥ g − (pairs with GCD exactly 2g) − (pairs with GCD exactly 3g) − ...
- Iterate g from max down to 1, subtract multiples: `exact[g] = C(divCount[g], 2) − exact[2g] − exact[3g] − ...`

**Step 4: Prefix sum + binary search for queries**
Build `prefix[g] = exact[1] + exact[2] + ... + exact[g]`. For query `k`, binary search for smallest `g` where `prefix[g] > k`.

### Key Insight
The "at least g" counting via multiples is the standard technique. When value range [1, m] is small but n is large, always enumerate values, not pairs.

### Complexity
- **Time:** O(m log m + q log m) — harmonic series + binary search per query
- **Space:** O(m)

### Common Mistake
Starting the inclusion-exclusion inner loop at `j = i` instead of `j = 2*i`. When `j = i`, `exact[i]` is 0 (not yet computed), so it's harmless but wasteful. More importantly, iterating all `j > i` instead of only multiples of `i` gives wrong results.

## Common Mistake

Iterating to compute the sum when a formula exists. Before coding, ask: "Is there a closed-form for this sequence?"

## Pre-Coding Checklist

1. "Is there a formula for this sum/sequence?"
2. "Can GCD/LCM/GCD properties simplify this?"
3. "Is the answer constant or trivially derived?"

## Complexity

- **Time:** O(1) with formula (vs O(n) without)
- **Space:** O(1)

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **GCD of Odd and Even Sums (LC 3685)** | sumOdd = n², sumEven = n(n+1), GCD = n |
| **Sum of Squares (LC 509)** | Fibonacci – not arithmetic but uses formula |
| **GCD of Array (LC 1979)** | Euclidean algorithm |
| **Sorted GCD Pair Queries (LC 3312)** | Inclusion-exclusion on multiples + prefix sum + binary search |
| **Counting Coprime Pairs (CSES 2417)** | Inclusion-exclusion
| **Sum of Multiples (LC 1085)** | Arithmetic series with common difference |
