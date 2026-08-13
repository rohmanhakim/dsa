# Segment Tree (Range Queries + Point Updates)

## When to Use

- Need **point updates** and **range queries** on a dynamic array (elements change between queries)
- Constraints: n ≤ 10^5, k queries, O(n·k) is too slow — need O((n+k) log n)
- The range property can be **merged** from children (sum, max, min, prefix/suffix runs, etc.)

Key signals: "point update + range query," "dynamic array," "after each update, what is the [max/sum/longest run] over the whole array?"

## Core Idea

A binary tree where each node covers a contiguous interval of the array. Each node stores a **summary** of its interval. Two operations:

1. **Build:** Construct the tree bottom-up from the leaves
2. **Update:** Change one element, then re-merge on the way back up to the root
3. **Query:** The root's summary is the answer for the entire array (for whole-array queries; for arbitrary range queries, you'd walk the tree)

For problems that only need the **global** answer after each point update (like LC 2213), you just read `tree[1].maxLen` — no separate query function needed.

## Template (Iterative, Heap Layout)

```go
type node struct {
    // fields that summarize the interval
}

type segtree struct {
    n    int      // size (next power of 2)
    tree []node   // 1-indexed heap layout: node i has children 2*i and 2*i+1
}

// merge: combine left child and right child into parent
// This is the ENTIRE algorithm — get this right and everything else is mechanical
func merge(left, right node) node {
    // ... merge logic depends on what you're tracking
}

// build: construct tree from array
func (st *segtree) build(arr []Type) {
    // leaves at indices n to n+len(arr)-1
    for i := 0; i < len(arr); i++ {
        st.tree[n+i] = nodeFromValue(arr[i])
    }
    // mark padding leaves as empty (if len(arr) < n)
    for i := len(arr); i < n; i++ {
        st.tree[n+i] = node{isEmpty: true}
    }
    // internal nodes bottom-up
    for i := n - 1; i > 0; i-- {
        st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
    }
}

// update: point update, then walk up re-merging
func (st *segtree) update(pos int, val Type) {
    i := st.n + pos
    st.tree[i] = nodeFromValue(val)
    for i /= 2; i > 0; i /= 2 {
        st.tree[i] = merge(st.tree[i*2], st.tree[2*i+1])
    }
}

// query: root covers the entire array
func (st *segtree) query() node {
    return st.tree[1]
}
```

### Padding Leaves (Non-Power-of-2 Lengths)

When `len(arr)` is not a power of 2, the tree has extra leaves beyond the array. These must be marked as **empty** so they don't corrupt merge results.

**Critical Go gotcha:** Go's zero value for `bool` is `false`, not `true`. You must explicitly set `isEmpty: true` for padding leaves. If you don't, padding leaves have `leftChar=0, rightChar=0, pre=0, suf=0, maxLen=0, iLen=0, isEmpty=false` — and the merge logic treats them as real data.

**Merge rule for empty nodes:**
```go
if left.isEmpty { return right }
if right.isEmpty { return left }
```

## Example: LC 2213 (Longest Substring of One Repeating Character)

### What each node stores

| Field | Meaning |
|-------|---------|
| `pre` | Length of longest identical-char **prefix** of this interval |
| `suf` | Length of longest identical-char **suffix** of this interval |
| `maxLen` | Length of longest identical-char substring **anywhere** in this interval |
| `leftChar` | Character at the leftmost position |
| `rightChar` | Character at the rightmost position |
| `iLen` | Length of the interval (number of positions covered) |

### Merge rule (the entire algorithm)

Given left child `L` and right child `R`:

```
leftChar  = L.leftChar
rightChar = R.rightChar
iLen      = L.iLen + R.iLen

pre = L.pre
if L.rightChar == R.leftChar AND L.pre covers entire L interval (L.pre == L.iLen):
    pre = L.pre + R.pre

suf = R.suf
if L.rightChar == R.leftChar AND R.suf covers entire R interval (R.suf == R.iLen):
    suf = R.suf + L.suf

maxLen = max(L.maxLen, R.maxLen)
if L.rightChar == R.leftChar:
    maxLen = max(maxLen, L.suf + R.pre)
```

### Why the conditions matter

**`pre` extension:** Can only extend across the boundary if the **entire** left child is one character (not just the prefix). Example: L = "aab", L.pre = 2, L.iLen = 3. Even if L.rightChar matches R.leftChar, `pre` stays 2 because L isn't fully uniform.

**`maxLen` boundary:** Only combine `L.suf + R.pre` when characters match at the boundary. When chars differ, the suffix of L and prefix of R are separate runs — summing them would be wrong.

### Trace: merging "aa" and "ba"

```
L = [a,a]  → pre=2, suf=2, maxLen=2, leftChar=a, rightChar=a, iLen=2
R = [b,a]  → pre=1, suf=1, maxLen=1, leftChar=b, rightChar=a, iLen=2

pre:  L.pre = 2. L.pre == L.iLen? Yes (2==2). L.rightChar==R.leftChar? a==b? No. → pre = 2
suf:  R.suf = 1. R.suf == R.iLen? No (1!=2). → suf = 1
maxLen: max(2, 1) = 2. L.rightChar==R.leftChar? a==b? No. → maxLen = 2

Result: pre=2, suf=1, maxLen=2, leftChar=a, rightChar=a, iLen=4
```

### Trace: merging "aa" and "aa"

```
L = [a,a]  → pre=2, suf=2, maxLen=2, leftChar=a, rightChar=a, iLen=2
R = [a,a]  → pre=2, suf=2, maxLen=2, leftChar=a, rightChar=a, iLen=2

pre:  L.pre = 2. L.pre == L.iLen? Yes. L.rightChar==R.leftChar? a==a? Yes. → pre = 2+2 = 4
suf:  R.suf = 2. R.suf == R.iLen? Yes. L.rightChar==R.leftChar? Yes. → suf = 2+2 = 4
maxLen: max(2, 2) = 2. L.rightChar==R.leftChar? Yes. → maxLen = max(2, 2+2) = 4

Result: pre=4, suf=4, maxLen=4, leftChar=a, rightChar=a, iLen=4
```

## Common Mistakes

1. **Unconditionally combining boundary runs in maxLen** — `L.suf + R.pre` must only be a candidate when `L.rightChar == R.leftChar`. Otherwise, the suffix of L and prefix of R are separate runs.

2. **Extending pre/suf without checking full interval coverage** — `pre` can only extend into R if L is **entirely** one character (`L.pre == L.iLen`), not just if the boundary chars match.

3. **Go zero-value `bool` is `false`** — padding leaves need explicit `isEmpty: true`. Forgetting this causes padding leaves (with `leftChar=0, rightChar=0`) to falsely match each other and corrupt merges.

4. **Not propagating `iLen` in merge** — if the merged node doesn't set `iLen = L.iLen + R.iLen`, parent nodes can't check the full-interval coverage condition.

5. **Infinite recursion in update** — the update function should just set the leaf and walk up. No recursive calls needed.

## Complexity

- **Build:** O(n) — visit each node once
- **Update:** O(log n) — walk from leaf to root
- **Space:** O(n) — array of 2n nodes (power-of-2 sizing)

## Alternative: Interval Merging (Editorial Approach 2)

Instead of a segment tree, maintain a sorted set of intervals (runs). Each query finds the interval containing the position, splits it, and merges with neighbors. Requires an ordered set (C++ `std::set`, Java `TreeMap`) for O(log n) operations.

**Why segment tree is often better in Go:** Go doesn't have a built-in sorted set. The segment tree uses only arrays — no external dependencies.

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 2213** — Longest Substring of One Repeating Character | Segment tree with pre/suf/maxLen/leftChar/rightChar per node |
| **Range Sum Query - Mutable (LC 307)** | Classic segment tree for range sum + point update |
| **Range Maximum Query (RMQ)** | Segment tree with `max` as the merge operation |
| **Count of Smaller Numbers After Self (LC 315)** | Segment tree on value range, update + query |
| **Reverse Pairs (LC 493)** | Segment tree or merge sort variant |