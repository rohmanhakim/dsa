# Binary Search

## When to Use

Binary search applies when:

- The search space is **ordered** (sorted array), or
- There is a **monotonic predicate** (e.g. `false false false true true true`).

Goal:

- Reduce the search space by half every iteration.
- Time complexity: **O(log n)**.

---

# Step 1: Define the Search Space

Before writing any binary search, answer these questions:

1. **What are all possible answers?**
2. **How will I represent the search space?**
   - Closed interval: `[left, right]`
   - Half-open interval: `[left, right)`
3. **When is the search space empty?**
4. **After examining `mid`, can I eliminate it or must I preserve it?**

Everything else (initialization, loop condition, updates, return value) follows from these answers.

---

# Template A — Exact Search (Closed Interval)

Search space:

```
[left, right]
```

Invariant:

> If the answer exists, it is somewhere in `[left, right]`.

Initialization:

```go
left, right := 0, len(nums)-1
```

Empty search space:

```
left > right
```

Loop:

```go
for left <= right {
    mid := left + (right-left)/2

    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}

return -1
```

### Why `left <= right`?

When

```
left == right
```

there is still **one candidate** left:

```
[left]
```

so we must continue searching.

Only when

```
left > right
```

is the search space empty.

### Why `mid ± 1`?

We've already examined `mid`.

If `nums[mid] != target`, then `mid` cannot be the answer.

Therefore it is safe to discard it.

### Use when

- Exact search
- Return found / not found

Examples:

- LC 704 — Binary Search

---

# Template B — Boundary Search (Half-Open Interval)

Search space:

```
[left, right)
```

Invariant:

> The boundary is always inside `[left, right)`.

Initialization:

```go
left, right := 0, len(nums)
```

Empty search space:

```
left == right
```

Loop:

```go
for left < right {
    mid := left + (right-left)/2

    if predicate(mid) {
        right = mid
    } else {
        left = mid + 1
    }
}

return left
```

### Why `left < right`?

When

```
left == right
```

the interval

```
[left, right)
```

contains **no candidates**.

The search space is already empty.

By the invariant, `left` is exactly the desired boundary.

### Why `right = mid`?

`mid` might still be the answer.

We cannot discard it.

Instead, we shrink the search space while keeping `mid` inside.

### Use when

- Lower bound
- Upper bound
- First occurrence
- Last occurrence
- Search insert position
- Binary search on answer

Examples:

- LC 35 — Search Insert Position
- LC 34 — First and Last Position
- LC 278 — First Bad Version

---

# Special Case — Peak Finding (LC 162)

```go
left, right := 0, len(nums)-1

for left < right {
    mid := left + (right-left)/2

    if nums[mid] < nums[mid+1] {
        left = mid + 1
    } else {
        right = mid
    }
}

return left
```

Although this uses

```
right = mid
```

it does **not** use a half-open interval.

Invariant:

> A peak always exists inside `[left, right]`.

We stop when

```
left == right
```

because there is only one possible peak remaining.

### Why is `mid+1` always valid?

Since

```
left < right
```

we always have

```
mid < right
```

Therefore

```
mid + 1 <= right
```

which guarantees `nums[mid+1]` is in bounds.

---

# Choosing the Correct Template

Ask these questions:

### 1. What is my search space?

```
[left,right]
```

or

```
[left,right)
```

---

### 2. When is it empty?

Closed:

```
left > right
```

Half-open:

```
left == right
```

---

### 3. Can I discard `mid`?

If you've proven

```
mid
```

cannot be the answer:

```go
left = mid + 1
```

or

```go
right = mid - 1
```

If `mid` might still be the answer:

```go
right = mid
```

or

```go
left = mid
```

(usually with an upper-mid calculation).

---

# Common Mistakes

### Mixing interval conventions

Wrong:

```go
left < right
right = mid - 1
```

Wrong:

```go
left <= right
right = mid
```

Initialization, loop condition, and updates must all agree with the chosen invariant.

---

### Forgetting why the loop ends

Closed interval:

```
left > right
```

Half-open interval:

```
left == right
```

Never memorize these—derive them from when the search space becomes empty.

---

### Overflow

Instead of

```go
mid := (left + right)/2
```

use

```go
mid := left + (right-left)/2
```

---

# Complexity

Time:

```
O(log n)
```

Space:

```
O(1)
```

---

# Problem Mapping

| Problem | Template | Key Idea |
|----------|----------|----------|
| LC 704 | Exact Search | Find target |
| LC 35 | Boundary Search | Lower bound |
| LC 34 | Boundary Search ×2 | Lower + Upper bound |
| LC 278 | Boundary Search | First true |
| LC 153 | Boundary Search | Minimum in rotated array |
| LC 162 | Peak Search | Preserve candidate peak |
| LC 4 | Binary Search on Partition | Binary search on answer |