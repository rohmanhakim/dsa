## The Problem Description

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
[5, 4, 5]​​​​​​​

Example 2:

Input: n = 3, l = 1, r = 3

Output: 10

Explanation:

There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

[1, 2, 1], [1, 3, 1], [1, 3, 2]
[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
[3, 1, 2], [3, 1, 3], [3, 2, 3]
All arrays meet the ZigZag conditions.

Constraints:

3 <= n <= 2000
1 <= l < r <= 2000
```

## Approach 1: Recursive Top-Down

### Step 1: Understanding the ZigZag Constraint

The three rules can be restated more simply:

1. `a[i] != a[i-1]` — no equal adjacent elements
2. NOT `a[i-2] < a[i-1] < a[i]` — no three strictly increasing
3. NOT `a[i-2] > a[i-1] > a[i]` — no three strictly decreasing

Rules 2 and 3 together force the **direction to alternate** at every step. If the
last step went **up** (`a[i-1] > a[i-2]`), the next step **must** go **down**
(`a[i] < a[i-1]`). If the last step went down, the next must go up.

This is the key insight: the "zigzag" is literally an alternating up-down pattern.

### Step 2: Why Multi-Dimensional DP?

To decide what value to place at position `i`, we need more than just the position.
We need to know:

- **`pos`** — which position are we filling?
- **`val`** — what value is at the current position? (constrains what can come next)
- **`isUp`** — what was the direction of the last step? (determines whether the next
  step must go up or down)

This makes it a **3-dimensional DP** problem.

### Step 3: Modeling the State from the Final Position

We chose to model the recursion starting from the **last position** (`n-1`) and
working **backwards** to position 0. This is a natural top-down approach: define
the problem at the end, and recurse to reduce it.

**State**: `dp(pos, isUp, val)` = number of valid ways to fill positions
`[0, 1, ..., pos-1]` given that:
- Position `pos` has value `val`
- The step from position `pos-1` to `pos` went in direction `isUp`
  - `isUp = true` → value increased (val > value at pos-1)
  - `isUp = false` → value decreased (val < value at pos-1)

**Base case**: When `pos == 0`, there are no more positions to fill to the left →
return 1.

### Step 4: State Transition

From state `(pos, isUp, val)`, we need to determine what value `prevVal` could be
at position `pos-1`:

- If `isUp == true`: the step from `pos-1` to `pos` went up, meaning `val > prevVal`.
  So `prevVal` must be in `[l, val-1]`.
- If `isUp == false`: the step from `pos-1` to `pos` went down, meaning `val < prevVal`.
  So `prevVal` must be in `[val+1, r]`.

For each valid `prevVal`, we recurse to `dp(pos-1, !isUp, prevVal)` — flipping the
direction because the zigzag constraint requires alternating directions.

**Special case at `pos == 1`**: When we recurse to position 0, there's no zigzag
constraint (you need 3 consecutive elements for zigzag to apply). The only
requirement is that position 0's value is in `[l, r]` — and the direction constraint
already guarantees `prevVal != val`. So we just count 1 for each valid `prevVal`.

### Step 5: Initial Trigger

The answer is the sum of all possible states at position `n-1`:

```
answer = Σ dp(n-1, true, i)  +  Σ dp(n-1, false, i)
         for i in [l, r]           for i in [l, r]
```

This covers both possibilities: the last step went up, or the last step went down.

### Step 6: Bugs We Hit During Development

**Bug 1 — Missing direction enforcement (double counting)**:
Initially, the inner loop iterated over all `prevVal != val` without checking the
direction constraint. This meant `dp(n-1, true, i)` and `dp(n-1, false, i)` explored
the **same** search spaces, doubling the count. Fix: filter `prevVal` based on
`isUp`.

**Bug 2 — Slice backtracking**:
For debugging, we tracked the constructed array `z` via a pointer. Elements were
appended but never removed after recursive calls returned, causing the slice to
accumulate stale values across branches. Fix: add `*z = (*z)[:len(*z)-1]` after
each recursive call (standard backtracking undo).

### Step 7: Memoization

**State space size**: `n × 2 × (r - l + 1)` — the Cartesian product of all state
dimensions.

We used two 2D memo arrays:
- `memoUp[pos][val-l]` — caches `dp(pos, true, val)`
- `memoDown[pos][val-l]` — caches `dp(pos, false, val)`

We initialized with sentinel value `-1` (not `0`) because `0` is a valid result
(it means no valid arrays exist from that state). Using `0` as "not computed"
would cause repeated recomputation of zero-result states.

**Known limitation (TLE)**: Even with memoization, the time complexity is
**O(n × range²)** because each state iterates over up to `range` values. For large
inputs (e.g., n=43, range=365), this hits LeetCode's time limit. The bottom-up
iterative approach with prefix/suffix sums can reduce this to **O(n × range)**.

### The Code

```go
func zigZagArrays(n int, l int, r int) int {
	const MOD = 1000000007

	memoUp := make([][]int, n)
	memoDown := make([][]int, n)
	for i := range n {
		rangeLength := (r - l + 1)
		for j := 0; j <= rangeLength; j++ {
			memoUp[i] = append(memoUp[i], -1)
			memoDown[i] = append(memoDown[i], -1)
		}
	}

	var dp func(pos int, isUp bool, val int) int
	dp = func(pos int, isUp bool, val int) int {
		if pos == 0 {
			return 1
		}

		count := 0

		rangeStart, rangeEnd := l, val-1
		if !isUp {
			rangeStart, rangeEnd = val+1, r
		}
		for prevVal := rangeStart; prevVal <= rangeEnd; prevVal++ {
			if prevVal == val {
				continue
			}

			if pos-1 >= 1 {
				res := 0

				if isUp {
					if memoDown[pos-1][prevVal-l] != -1 {
						res = memoDown[pos-1][prevVal-l]
					} else {
						res = dp(pos-1, false, prevVal)
						memoDown[pos-1][prevVal-l] = res
					}
				} else {
					if memoUp[pos-1][prevVal-l] != -1 {
						res = memoUp[pos-1][prevVal-l]
					} else {
						res = dp(pos-1, true, prevVal)
						memoUp[pos-1][prevVal-l] = res
					}
				}
				count = (count + res) % MOD
			} else {
				count = (count + 1) % MOD
			}
		}
		return count
	}

	totalCount := 0
	for i := r; i >= l; i-- {
		totalCount = (totalCount + dp(n-1, true, i)) % MOD
		totalCount = (totalCount + dp(n-1, false, i)) % MOD
	}
	return totalCount
}
```

### Complexity

- **Time**: O(n × range²) where range = r - l + 1
- **Space**: O(n × range) for the memo arrays

## Approach 2: Iterative Bottom-Up with Prefix/Suffix Sums

### Step 1: Why Bottom-Up?

The top-down approach with memoization has time complexity **O(n × range²)**
because each state iterates over up to `range` values. For large inputs
(e.g., n=43, range=365), this hits LeetCode's TLE.

The bottleneck is the inner loop: for each state, we sum over a range of
previous values. This is exactly the kind of operation that **prefix/suffix
sums** can optimize to O(1).

### Step 2: Re-framing the State for Bottom-Up

Instead of recursing backwards from position n-1, we **build forward** from
position 0. At each position, we track two arrays:

- `up[v]` = number of valid arrays ending at the current position with value `v`,
  where the last step went **up** (v increased from previous)
- `down[v]` = same, but last step went **down**

These arrays have size `range = r - l + 1`, indexed by `v - l`.

### Step 3: Base Case (Position 0)

At position 0, there's no direction yet — just place any value:

```
count[v] = 1   for all v in [l, r]
```

### Step 4: Bootstrapping Position 1

At position 1, a direction first appears. We compute `up` and `down` from
`count`:

```
up[v]   = number of values less than v    = v - l
down[v] = number of values greater than v = r - v
```

In the code, this is handled by the `else` branch (when `i == 1`) which builds
prefix/suffix sums from `count` instead of from `up`/`down`.

### Step 5: State Transition (Position 2+)

For positions 2 and beyond, the zigzag constraint kicks in:

```
newUp[v]   = sum of down[u] for all u < v
             (came from below → previous step was DOWN)

newDown[v] = sum of up[u] for all u > v
             (came from above → previous step was UP)
```

### Step 6: Prefix/Suffix Sum Optimization

Each transition requires a **range sum** over the previous layer. Instead of
looping, we precompute:

- `prefixDown[k]` = `down[0] + down[1] + ... + down[k-1]` (sum of first k elements)
- `suffixUp[k]` = `up[k] + up[k+1] + ... + up[range-1]` (sum from k to end)

With sentinel values (`prefixDown[0] = 0`, `suffixUp[range] = 0`), the lookups
become O(1):

```
newUp[v]   = prefixDown[v-l]       // sum of down for indices < v-l
newDown[v] = suffixUp[v-l+1]       // sum of up for indices > v-l
```

### Step 7: Accumulating the Answer

After processing all positions, the answer is the sum of all values in `up` and
`down`:

```
answer = Σ (up[v] + down[v]) for all v in [l, r]
```

### The Code

```go
func zigZagArrays(n int, l int, r int) int {
	const MOD = 1000000007

	rangeLength := r - l + 1
	count := make([]int, rangeLength)
	for i := l; i <= r; i++ {
		count[i-l] = 1
	}

	var up, down []int

	for i := 1; i <= n-1; i++ {

		prefixDown := make([]int, (rangeLength + 1))
		prefixDown[0] = 0
		suffixUp := make([]int, (rangeLength + 1))
		suffixUp[rangeLength] = 0

		if i > 1 {
			for j := 1; j <= rangeLength; j++ {
				prefixDown[j] = (prefixDown[j-1] + down[j-1]) % MOD
			}

			for j := rangeLength - 1; j >= 0; j-- {
				suffixUp[j] = (suffixUp[j+1] + up[j]) % MOD
			}
		} else {
			for j := 1; j <= rangeLength; j++ {
				prefixDown[j] = prefixDown[j-1] + count[j-1]
			}

			for j := rangeLength - 1; j >= 0; j-- {
				suffixUp[j] = suffixUp[j+1] + count[j]
			}
		}

		newUp := make([]int, rangeLength)
		newDown := make([]int, rangeLength)
		for j := l; j <= r; j++ {
			newUp[j-l] = prefixDown[j-l]
			newDown[j-l] = suffixUp[j-l+1]
		}

		up, down = newUp, newDown
	}

	answer := 0
	for i := 0; i < rangeLength; i++ {
		answer = (answer + up[i] + down[i]) % MOD
	}
	return answer
}
```

### Complexity

- **Time**: O(n × range) — each position does O(range) work for prefix/suffix
  builds and O(range) for the transition
- **Space**: O(range) — only need the current `up`/`down` arrays and temporary
  prefix/suffix arrays

### Comparison with Top-Down

| | Top-Down | Bottom-Up |
|---|---|---|
| Time | O(n × range²) | O(n × range) |
| Space | O(n × range) | O(range) |
| Implementation | Recursive with memo | Iterative with prefix/suffix |
| TLE on large input | Yes | No |

The bottom-up approach is both **faster** and **more memory-efficient** because
we only keep the previous layer's values, not the entire memo table.
