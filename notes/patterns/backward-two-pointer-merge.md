# 6. Backward Two-Pointer Merge

## When to Use
You have two sorted sequences and need to merge them into a buffer that has **extra space at the end**. Forward insertion causes O(n) shifting per element; filling from the back eliminates shifting entirely.

## How It Works
1. Start pointers at the **end** of both sorted sequences (`p1 = m-1`, `p2 = n-1`)
2. Start the write pointer at the **end** of the result buffer (`p = m+n-1`)
3. Compare `nums1[p1]` vs `nums2[p2]`, place the **larger** value at `nums1[p]`, decrement the winning pointer and `p`
4. When one sequence is exhausted, copy remaining elements from the other (they're already in place if it's nums1)

## Key Insight
The empty slots at the end of nums1 are **unused space** — writing there never overwrites valid data. By always placing the largest remaining element at the current back position, you guarantee sorted order with zero shifting.

## Example: LC 88
```
nums1 = [1,2,3,0,0,0]  nums2 = [2,5,6]

p1=2(3), p2=2(6), p=5  → 6>3 → nums1[5]=6
p1=2(3), p2=1(5), p=4  → 5>3 → nums1[4]=5
p1=2(3), p2=0(2), p=3  → 3>2 → nums1[3]=3
p1=1(2), p2=0(2), p=2  → 2≥2 → nums1[2]=2
p1=0(1), p2=-1,   p=1  → nums2 done, copy nums1[0:1] → already in place
Result: [1,2,2,3,5,6] ✓
```

## Complexity
- **Time:** O(m+n) — single pass, each element placed exactly once
- **Space:** O(1) — in-place

## Classic Problems to Practice
| Problem | Key Idea |
|---|---|
| **Merge Sorted Array (LC 88)** | Classic backward merge into buffer with trailing zeros |
| **Squares of a Sorted Array (LC 977)** | Largest squares at the ends → fill result from back |
| **Sort Transformed Array (LC 360)** | Quadratic transform preserves sorted property at extremes |
