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

## Sub-Variant: Closed-Form Answer Based on Input Size (LC 3513)

### When to Use
- The answer depends **only on n** (the input size), not on the specific values in the array.
- Strong signal: the input is a **permutation** of [1, n] or [0, n-1] — every value from 1 to n appears exactly once.
- The answer is bounded by some function of n (e.g., the smallest power of 2 greater than n).

### How It Works
1. **Trace small examples** — compute the answer for n=1, 2, 3, 4, 5, 6, 7, 8 manually.
2. **Identify the pattern** — look for a relationship between n and the answer.
3. **Derive the formula** — express the answer as a function of n.
4. **Verify** — check that the formula holds for edge cases (n < 3, n = 2^k, etc.).

### Key Question to Ask
Before coding, ask: **"Does the answer depend on the specific values, or just on n?"**

If the input is a permutation of [1, n], then the actual values don't matter — the answer is determined solely by n. This means there must be a closed-form formula, and any brute-force enumeration is unnecessary.

### Example: LC 3513 (Number of Unique XOR Triplets I)
- **Input:** nums is a permutation of [1, n]
- **Question:** How many unique XOR values can be formed by XORing 3 elements?
- **Trace small examples:**
  - n=1 → answer 1 (only 1^1^1=1)
  - n=2 → answer 2 (values {1, 2})
  - n=3 → answer 4 (values {0, 1, 2, 3})
  - n=4 → answer 8 (values {0, 1, ..., 7})
  - n=6 → answer 8 (values {0, 1, ..., 7})
- **Pattern:** For n ≥ 3, the answer is the **smallest power of 2 strictly greater than n**.
  - n=3 → 2^2=4, n=6 → 2^3=8, n=358 → 2^9=512.
- **Why it works:** When n ≥ 3, you have {1, 2, 3} which XORs to 0. For x in [1, n], 1^1^x = x. For x in [n+1, 2^(k+1)-1], use 2^k plus two elements to construct x. Every integer in [0, 2^(k+1)-1] is achievable.
- **Result:** O(log n) time, O(1) space — no enumeration needed.

### Example: LC 3685 (GCD of Odd and Even Sums)
- **Input:** n (count of odd/even numbers)
- **Question:** GCD(sum of first n odd numbers, sum of first n even numbers) = ?
- **Trace:** sumOdd = n², sumEven = n(n+1), GCD(n², n(n+1)) = n.
- **Result:** Answer is always n. O(1).

### Common Mistake
Jumping to brute-force enumeration (O(n³)) without first checking if the answer has a closed-form. The permutation property is a strong signal — when every value from 1 to n appears, the answer depends only on n.

### Pre-Coding Checklist Addition
Before coding, add these questions:
1. "Is there a formula for this sum/sequence?"
2. "Can GCD/LCM/GCD properties simplify this?"
3. "Is the answer constant or trivially derived?"
4. **"Does the answer depend on the specific values, or just on n?"** ← If yes, trace small examples and look for pattern.

## Pre-Coding Checklist

1. "Is there a formula for this sum/sequence?"
2. "Can GCD/LCM/GCD properties simplify this?"
3. "Is the answer constant or trivially derived?"
4. "Does the answer depend on the specific values, or just on n?"

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
| **Number of Unique XOR Triplets I (LC 3513)** | Closed-form answer: smallest power of 2 > n (for n ≥ 3) |
