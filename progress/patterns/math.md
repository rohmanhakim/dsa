# Math

**Status:** In Progress — 2/5 consecutive independent sessions with severity ≤ 1. Needs 3 more consecutive independent sev≤1 passes across at least 3 distinct days.
**Sessions:** LC 3685 (ind, sev 0, 4:15), LC 3867 (ind, sev 0, 9:57), LC 3312 (guided, sev 3), LC 1979 (ind, sev 0, 2:08), LC 3513 (guided, sev 3), LC 628 (ind, sev 1, 56:54), LC 1464 (ind, sev 0, 2:06), LC 3518 (guided, sev 3), LC 3014 (ind, sev 1, 18:52), LC 3345 (ind, sev 1, 1:09:17), LC 2 (ind, sev 0, 17:58), LC 2996 (ind, sev 1, 80:48), LC 1927 (guided, sev 3), LC 3875 (ind, sev 0, 6:55), LC 3876 (ind, sev 1, 34:47).

Arithmetic series formulas, GCD properties, prefix max + GCD combination, inclusion-exclusion for GCD pair counting, Euclidean algorithm for GCD of extremes, closed-form answer based on input size, sort + candidate enumeration for max product, k-th lexicographic permutation construction via product of combinations, bounded enumeration for digit-product divisibility, carry-based digit addition for linked list number representation, sequential prefix sum + hash set for missing integer search, game theory invariant derivation (optimal play reduces to a mathematical equation), and parity analysis with minimum-value constraint (minimum parity determines feasibility when operations are constrained).

## Lessons

- **LC 3685:** Arithmetic series formulas + GCD properties. Key formulas: sum of first n odd numbers = n², sum of first n even numbers = n(n+1). Key property: gcd(n, n+1) = 1 (consecutive integers are coprime). Before iterating to compute a sum, check if a closed-form formula exists.
- **LC 3867:** Prefix max array (single pass O(n)) + GCD computation. Two-pointer pairing after sorting. Key pattern: when the problem involves prefix aggregates combined with GCD, compute prefix max in one pass, then apply GCD.
- **LC 3312:** Inclusion-exclusion for GCD pair counting. When GCD value range [1, m] is small but n is large, enumerate GCD values instead of pairs. Key sub-variant: (1) count elements divisible by g via harmonic series O(m log m), (2) C(cnt,2) = pairs with GCD ≥ g, (3) top-down subtraction to get exact counts, (4) prefix sum + binary search for query answering. **Key insight:** the "at least g" counting via multiples is the standard technique for GCD pair problems.
- **LC 1979:** Euclidean algorithm for GCD (`a, b = b, a%b` until b=0) combined with min/max scan is the textbook O(n) solution for "GCD of min and max" problems.
- **LC 3513:** Initially brute-forced O(n³) DFS — TLE. Failed to recognize that when nums is a permutation of [1, n], the answer depends only on n, not on the specific values. Answer is simply the smallest power of 2 greater than n (for n ≥ 3). **Key insight:** always ask "does the answer depend on the specific values, or just on n?"
- **LC 628:** After sorting, the max product of three numbers is either **three largest** or **two most negative × largest**. Only 2 candidates needed — user had 4 with one redundant. Key insight: negative × negative = positive.
- **LC 1464:** Sort array, take two largest elements, subtract 1 from each, multiply. O(n log n) works but O(n) single pass is optimal. Key insight: for simple array problems, always ask "Can I find the answer in one pass?"
- **LC 3518:** K-th lexicographic permutation using product of binomial coefficients. **Critical insight:** factorial-based counting overflows for n > 20. The product-of-combinations approach with early termination at k avoids overflow entirely. **Palindrome symmetry:** only the left half needs to be constructed; the right half is the mirror.
- **LC 2:** Carry-based digit addition with simultaneous linked list traversal. Max possible sum 9+9+1=19, so carry is always 0 or 1. Cleaner alternative: always create `current.Next` before advancing (dummy-head pattern) eliminates mid-loop nil guards.
- **LC 2996:** Two-phase simulation: (1) find longest sequential prefix and compute sum, (2) find smallest missing integer ≥ sum using hash set + linear scan. Key learning: for "find smallest missing integer" problems, hash set + linear scan is the standard approach. **Weakness:** Time significantly over target (80:48 vs 15 min for LC Easy).
- **LC 3116:** Inclusion-exclusion for counting multiples (used as secondary pattern with Binary Search on Answer). For each non-empty subset of denominations, compute LCM and add/subtract floor(x/LCM) based on subset size parity. **Key insight:** when counting "numbers divisible by at least one of [a,b,c,...]", inclusion-exclusion over all 2^n subsets is the standard technique. Skip subsets whose LCM exceeds x (contributes 0). Bitmask iteration over subsets is the mechanism; inclusion-exclusion is the math.
- **LC 1927:** Game theory problem where optimal play reduces to a mathematical invariant. Key formula: if total `?` marks are odd, Alice wins; if even, Bob wins iff `2*(n0-n1) = 9*(q1-q0)`. **Key insight:** instead of simulating the game, derive the condition under which the second player can always match the first player's moves. The `?` marks on both sides cancel out (Alice puts d, Bob puts d), and pairs on one side contribute exactly 9 (Alice puts d, Bob puts 9-d). This leads to a clean equation rather than a simulation. **Weakness:** initially tried to simulate the game with prefix/suffix sums and complex branching logic — pattern misidentification.
- **LC 3875:** Parity observation — for any array of distinct integers, a uniform-parity array is always constructible. If all same parity, use identity. If mixed, pick any odd element x and subtract it from even elements (even − odd = odd). **Key insight:** before coding, ask "is the answer always true/false?" For parity problems, the two operations (keep or subtract another element) give enough flexibility that the answer is always true.
- **LC 3876:** Parity with minimum-value constraint — when subtraction requires result ≥ 1, the **minimum value's parity** determines feasibility. Min odd → all-odd target achievable (even elements subtract min odd). Min even with mixed parity → impossible (min even can't change, smallest odd has no smaller odd to subtract). **Key insight:** for parity problems with constrained operations, analyze what the minimum element can/cannot do. **Weakness:** sorted entire array (O(n log n)) when only finding minimum was needed (O(n) single pass).

## Sub-techniques

### Inclusion-Exclusion

**When to use:** Counting objects satisfying *at least one* of several conditions (union counting). The formula alternates between adding and subtracting intersections.

**Signature:** "How many X are divisible by at least one of [a, b, c]?" or "How many X satisfy condition A or B or C?"

**Mechanism:** Iterate over all non-empty subsets (bitmask 1 to 2^n - 1). For each subset:
- Compute the combined constraint (e.g., LCM for divisibility)
- Add floor(x / constraint) if subset size is odd
- Subtract if subset size is even

**Seen in:** LC 3312 (GCD pair counting, primary), LC 3116 (multiples counting, secondary). If this appears as primary in 3+ distinct contexts, consider splitting into its own pattern.
