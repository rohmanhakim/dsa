# 9. Floyd's Cycle Detection (Fast-Slow Pointers)

## When to Use
Linked list problems involving:
- Detecting whether a cycle exists (LC 141)
- Finding the start node of a cycle (LC 142)
- Finding the middle of a linked list (LC 876)
- Any scenario where you need O(1) space cycle detection on a linked structure

## How It Works
Use two pointers moving at different speeds:
1. **Slow pointer** advances 1 step per iteration
2. **Fast pointer** advances 2 steps per iteration
3. If a cycle exists, fast will eventually "lap" slow and they'll point to the same node
4. If no cycle, fast will reach the end (nil) first

### Detecting a Cycle (LC 141)
```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
        return true  // cycle detected
    }
}
return false
```

### Finding Cycle Start (LC 142)
After detecting the meeting point, reset one pointer to head. Move both at speed 1. The node where they meet again is the cycle start:
```go
// Phase 1: detect cycle (find meeting point)
// Phase 2: find start
slow = head
for slow != fast {
    slow = slow.Next
    fast = fast.Next
}
return slow  // cycle start
```

## Why It Works (Invariant Proof)

**Phase 1 (detection):** Let the cycle have length `C`, and let the distance from head to cycle start be `F`. When slow enters the cycle (after `F` steps), fast is already `F` steps ahead inside the cycle. Since fast gains 1 step per iteration relative to slow, it will catch up in at most `C` steps.

**Phase 2 (finding start):** At the meeting point, slow has traveled `F + a` steps (where `a` is the distance into the cycle). Fast has traveled `2(F + a)`. The difference `F + a` must be a multiple of `C` (that's why they met). So `F ≡ -a (mod C)`, meaning moving `F` steps from both head and the meeting point lands on the same node — the cycle start.

## Complexity
- **Time:** O(n) — slow pointer traverses at most n nodes
- **Space:** O(1) — only two pointers

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Linked List Cycle (LC 141)** | Basic cycle detection — do they meet? |
| **Linked List Cycle II (LC 142)** | Phase 2: reset one pointer to head, both move at speed 1 to find start |
| **Middle of the Linked List (LC 876)** | When fast reaches end, slow is at middle |
| **Happy Number (LC 202)** | Treat digit-sum sequence as a linked list; cycle = unhappy |
| **Find the Duplicate Number (LC 287)** | Treat array as linked list where nums[i] is next pointer; Floyd's finds the duplicate |