## Problem Description
```
	You are given three integers n, l, and r.

	A ZigZag array of length n is defined as follows:

	Each element lies in the range [l, r].
	No two adjacent elements are equal.
	No three consecutive elements form a strictly increasing or strictly decreasing sequence.
	Return the total number of valid ZigZag arrays.

	Since the answer may be large, return it modulo 109 + 7.

	A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

	A sequence is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).



	Example 1:

	Input: n = 3, l = 4, r = 5

	Output: 2

	Explanation:

	There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:

	[4, 5, 4]
	[5, 4, 5]
	Example 2:

	Input: n = 3, l = 1, r = 3

	Output: 10

	Explanation:

	​​​​​​​There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

	[1, 2, 1], [1, 3, 1], [1, 3, 2]
	[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
	[3, 1, 2], [3, 1, 3], [3, 2, 3]
	All arrays meet the ZigZag conditions.



	Constraints:

	3 <= n <= 10^9
	1 <= l < r <= 75​​​​​​​
```

## Approach: Matrix Exponentiation

The problem statement is identical to [Problem I](number-of-zigzag-arrays-i-3699.md),
but the constraints are drastically different:

| | Problem I | Problem II |
|---|---|---|
| n | ≤ 2,000 | ≤ **10^9** |
| range (r-l+1) | ≤ 2,000 | ≤ **75** |

The bottom-up O(n × range) approach from Problem I would need ~75 billion
operations for n=10^9 — far too slow. However, the small range (≤ 75) means
the state vector is compact, and the transition is a **fixed linear
transformation** — perfect for **matrix exponentiation**.

### Step 1: General Principle

When a DP transition is the **same at every step** (position-independent),
we can represent it as a matrix `M` and compute `M^k` via fast exponentiation
in O(range³ × log k) instead of O(k × range²).

**Matrix size** = product of all state dimensions **except** the advancing one
(n, which becomes the exponent).

For our problem:
```
States:     pos × isUp × val
              ↓       ↓
            n       2    × range

Matrix:     2 × range    (excluding pos)
Exponent:   n - 2         (number of transitions from position 1 to n-1)
```

### Step 2: State Vector

The state vector at each position has **2 × range** elements:

```
s = [up[l], up[l+1], ..., up[r], down[l], down[l+1], ..., down[r]]
     ←──── range elements ────→   ←──── range elements ────→
```

Where:
- `up[v]` = number of valid arrays ending with value `v`, last step went **up**
- `down[v]` = same, but last step went **down**

### Step 3: Transition Matrix M

The transition rules (from Problem I) are:
```
newUp[v]   = sum of down[u] for all u < v
newDown[v] = sum of up[u] for all u > v
```

Since the zigzag rule forbids consecutive same-direction steps, the matrix has
a **block structure**:

```
            up[l..r]       down[l..r]
           ┌───────────┬───────────┐
up[l..r]   │  all 0s   │  ???      │   ← newUp depends on down
           ├───────────┼───────────┤
down[l..r] │  ???      │  all 0s   │   ← newDown depends on up
           └───────────┴───────────┘
```

The off-diagonal blocks contain 0s and 1s:
- **Top-right**: `M[i][range+j] = 1` when `j < i` (newUp[i] += down[j])
- **Bottom-left**: `M[range+i][j] = 1` when `j > i` (newDown[i] += up[j])

### Step 4: Concrete Example (l=4, r=5)

Range = 2, so M is **4×4**. State vector: `[up[4], up[5], down[4], down[5]]`

```
            up[4]  up[5]  down[4]  down[5]
newUp[4]  [  0      0       0        0    ]
newUp[5]  [  0      0       1        0    ]
newDown[4][  0      1       0        0    ]
newDown[5][  0      0       0        0    ]
```

**Verification** for n=3:
- Initial state at position 1: s = [0, 1, 1, 0]
- M¹ × s = [0, 1, 1, 0]
- Sum = 2 ✓

### Step 5: Fast Exponentiation (Repeated Squaring)

To compute M^(n-2) in O(log n) matrix multiplications:

```
result = identity matrix
base = M
exp = n - 2

while exp > 0:
    if exp is odd:
        result = result × base
    base = base × base
    exp = exp / 2
```

Each squaring doubles the exponent, so we reach any power in O(log n) steps.
For n=10^9, that's ~30 matrix multiplications.

### Step 6: Initial State Vector

At position 1 (bootstrapping from count = [1,1,...,1]):
```
s[i] = i                    for up[l+i]   (values less than l+i)
s[range+i] = range - 1 - i  for down[l+i] (values greater than l+i)
```

### Step 7: Final Answer

```
answer = sum of all elements in M^(n-2) × s
```

### The Code

```go
func zigZagArraysII(n int, l int, r int) int {
	const MOD = 1e9 + 7

	rangeLength := r - l + 1
	m := make([][]int, 2*rangeLength)
	for i := range len(m) {
		m[i] = make([]int, 2*rangeLength)
	}

	for i := 0; i < rangeLength; i++ {
		for j := 0; j < i; j++ {
			m[i][rangeLength+j] = 1
		}
		for j := i + 1; j < rangeLength; j++ {
			m[rangeLength+i][j] = 1
		}
	}

	mult := func(a, b [][]int) [][]int {
		size := len(a)
		result := make([][]int, size)
		for i := range result {
			result[i] = make([]int, size)
		}

		for i := 0; i < size; i++ {
			for k := 0; k < size; k++ {
				if a[i][k] == 0 {
					continue
				}
				for j := 0; j < size; j++ {
					result[i][j] = (result[i][j] + a[i][k]*b[k][j]) % MOD
				}
			}
		}
		return result
	}

	result := make([][]int, 2*rangeLength)
	for i := range len(result) {
		result[i] = make([]int, 2*rangeLength)
		result[i][i] = 1
	}

	exp := n - 2
	for exp > 0 {
		if exp%2 == 1 {
			result = mult(result, m)
		}
		m = mult(m, m)
		exp = exp / 2
	}

	s := make([]int, 2*rangeLength)
	for i := 0; i < rangeLength; i++ {
		s[i] = i
		s[rangeLength+i] = rangeLength - 1 - i
	}

	finalMatrix := make([]int, 2*rangeLength)
	for i := 0; i < 2*rangeLength; i++ {
		for j := 0; j < 2*rangeLength; j++ {
			finalMatrix[i] = (finalMatrix[i] + result[i][j]*s[j]) % MOD
		}
	}

	answer := 0
	for i := 0; i < 2*rangeLength; i++ {
		answer = (answer + finalMatrix[i]) % MOD
	}

	return answer
}
```

### Complexity

- **Time**: O(range³ × log n) — each matrix multiply is O(range³), with O(log n) multiplies
- **Space**: O(range²) — storing the matrices

### Comparison with Problem I Approaches

| | Top-Down DP | Bottom-Up DP | Matrix Exponentiation |
|---|---|---|---|
| Time | O(n × range²) | O(n × range) | O(range³ × log n) |
| Space | O(n × range) | O(range) | O(range²) |
| Best for | n ≤ 2000, range ≤ 2000 | n ≤ 2000, range ≤ 2000 | **n ≤ 10^9**, range ≤ 75 |
| TLE on large n | Yes | Yes | **No** |

Matrix exponentiation trades the linear dependency on `n` for a logarithmic
one, making it the only viable approach when `n` is extremely large.
