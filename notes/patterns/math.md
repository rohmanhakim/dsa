# Math (Arithmetic Series, Number Theory)

## When to Use

- Problem involves **sums of consecutive numbers** (odd, even, natural numbers, squares)
- Problem asks for **GCD/GCD-related** and there's a mathematical property that simplifies it
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
| **Sum of Multiples (LC 1085)** | Arithmetic series with common difference |