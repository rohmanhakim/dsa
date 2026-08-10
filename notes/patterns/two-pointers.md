# Two Pointers

## When to Use

- **Opposite-end convergence:** Two pointers at start/end, moving inward based on a condition
- **Same-direction traversal:** Two pointers advancing through two sorted sequences simultaneously
- **Parallel greedy states:** Two independent match/scan states maintained in a single pass
- **Partitioning:** Separating elements by a predicate (e.g., Dutch national flag)

Key signals: "sorted array," "pair with sum/difference," "merge two sorted," "subsequence with one exception"

## Variant 1: Opposite-End Scan

Two pointers at opposite ends of a sorted array, moving inward:

```go
left, right := 0, len(nums)-1
for left < right {
    sum := nums[left] + nums[right]
    if sum == target {
        return []int{left, right}
    } else if sum < target {
        left++
    } else {
        right--
    }
}
```

**Why it works:** In a sorted array, moving the smaller pointer inward can only increase the sum (replacing a small value with a larger one). Moving the larger pointer inward can only decrease the sum. This eliminates half the search space per step.

**Used in:** LC 392 (Is Subsequence — one pointer per string), LC 167 (Two Sum II — sorted array)

**Key insight for LC 392:** When one string is a subsequence of another, a single forward pass with two pointers (one per string) checks the property in O(n+m). The "sorted" property here is the implicit ordering of characters in the target.

## Variant 2: Same-Direction Traversal

Both pointers advance in the same direction, one per sequence:

```go
dummy := &ListNode{}
tail := dummy
for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
        tail.Next = l1
        l1 = l1.Next
    } else {
        tail.Next = l2
        l2 = l2.Next
    }
    tail = tail.Next
}
tail.Next = l1 // or l2 (whichever remains)
return dummy.Next
```

**Key template:** dummy → tail → attach smaller → advance chosen → return dummy.Next

**Used in:** LC 21 (Merge Two Sorted Lists)

## Variant 3: Parallel Greedy States

Maintain two independent match/scan pointers in a single pass:

```go
j0, j1 := 0, 0  // j0 = without exception, j1 = with exception
for _, ch := range t {
    if s[j1] == byte(ch) { j1++ }
    j1 = max(j1, j0+1)  // "use exception now" = one ahead of j0
    if s[j0] == byte(ch) { j0++ }
    if j1 == n { return true }
}
return false
```

**The core insight:** When the problem asks "can I do X with one exception," maintain two parallel states. The `max(j1, j0+1)` line captures the decision to use the exception at the current position — it puts the "with exception" state exactly one step ahead of the "without exception" state.

**Why `j0+1`?** It represents: "I've just consumed one character in `t` as the exception (without matching it to `s`), so my match position is at least one ahead of the no-exception path."

**Used in:** LC 3983 (Subsequence After One Replacement)

**Complexity:** O(m) time, O(1) space — single pass through `t`.

**Pattern recognition signal:** When you can rephrase the problem as "subsequence with one allowed mismatch" or "can I do X if one step is free," think parallel greedy states.

## Variant 4: Read/Write Pointers

Separate read and write positions for in-place array modification:

```go
write := 0
for read := 0; read < len(nums); read++ {
    if shouldKeep(nums[read]) {
        nums[write] = nums[read]
        write++
    }
}
return write  // new length
```

**Used in:** LC 26 (Remove Duplicates), LC 283 (Move Zeroes)

**Key insight:** The write pointer only advances when an element is kept. Everything before `write` is the processed result. Everything after `read` is unprocessed.

## Decision Framework

| Situation | Variant | Why |
|---|---|---|
| Sorted array, pair with property | Opposite-end | Moving the "wrong" pointer eliminates candidates |
| Merge two sorted sequences | Same-direction | Always take the smaller element |
| "X with one exception" | Parallel states | Two independent match states in one pass |
| In-place compaction/filtering | Read/Write | Write pointer tracks result, read pointer scans input |

## Common Mistakes

1. **Not recognizing sorted property enables two pointers** — when the array is sorted, pair-finding reduces from O(n²) to O(n) with opposite-end scan.
2. **Using hash map when sorted makes it unnecessary** — LC 26: sorted array means duplicates are adjacent. Two pointers suffice.
3. **Forgetting the dummy node pattern** — for same-direction merge on linked lists, always start with `dummy := &ListNode{}; tail := dummy`.
4. **Missing parallel-state pattern** — when the problem involves "one allowed mismatch" or "one free operation," the parallel-state approach is O(n) vs O(n²) for explicit enumeration.

## Complexity

- **All variants:** O(n) or O(n+m) time, O(1) space (except linked list merge which uses O(1) extra space)

## Classic Problems to Practice

| Problem | Variant | Key Idea |
|---|---|---|
| **LC 392** — Is Subsequence | Opposite-end (one per string) | Forward scan: advance s-pointer on match |
| **LC 167** — Two Sum II | Opposite-end | Sorted → move smaller left, larger right |
| **LC 21** — Merge Two Sorted Lists | Same-direction | Dummy node + attach smaller |
| **LC 26** — Remove Duplicates | Read/Write | Write pointer tracks unique count |
| **LC 283** — Move Zeroes | Read/Write | Write non-zeros first, fill rest with 0 |
| **LC 3983** — Subsequence After One Replacement | Parallel states | `j1 = max(j1, j0+1)` for exception handling |
| **LC 88** — Merge Sorted Array | Backward merge | See [[backward-two-pointer-merge]] |
| **LC 141** — Linked List Cycle | Fast/Slow | See [[floyd-cycle-detection]] |