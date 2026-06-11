package hard

import (
	"container/heap"
	"math/bits"
)

type candidate struct {
	value int
	left  int
	right int
}

type candidateMaxHeap []candidate

func (h candidateMaxHeap) Len() int           { return len(h) }
func (h candidateMaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h candidateMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *candidateMaxHeap) Push(x any) {
	*h = append(*h, x.(candidate))
}
func (h *candidateMaxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	left := n - 1
	right := n - 1
	minItem := nums[right]
	maxitem := nums[right]
	h := &candidateMaxHeap{}
	heap.Init(h)

	// --------------------- using heap to track top largest scores among candidates
	for left >= 0 {
		if nums[left] < minItem {
			minItem = nums[left]
		}
		if nums[left] > maxitem {
			maxitem = nums[left]
		}
		score := maxitem - minItem
		heap.Push(h, candidate{
			value: score,
			left:  left,
			right: right,
		})
		left--
	}

	// ----------------------- Initislize Sparse Tables for Min & Max Range Min/Max Queries
	minst := [][]int{}
	minst = append(minst, nums)
	j := 1
	// bits.Len(uint(n)) gives floor(log2(n)) + 1, which is the number of levels needed
	for j < bits.Len(uint(n)) {
		minst = append(minst, make([]int, len(nums)-(pow2(j)-1)))
		for i := 0; i < len(minst[j]); i++ {
			minst[j][i] = min(minst[j-1][i], minst[j-1][i+pow2(j-1)])
		}
		j++
	}
	maxst := [][]int{}
	maxst = append(maxst, nums)
	j = 1
	for j < bits.Len(uint(n)) {
		maxst = append(maxst, make([]int, len(nums)-(pow2(j)-1)))
		for i := 0; i < len(maxst[j]); i++ {
			maxst[j][i] = max(maxst[j-1][i], maxst[j-1][i+pow2(j-1)])
		}
		j++
	}

	total := int64(0)
	for i := 0; i < k; i++ {
		best := heap.Pop(h).(candidate)
		total += int64(best.value)
		if best.right > best.left {
			rangeLength := (best.right - 1) - best.left + 1
			// bits.Len(uint(x)) - 1 gives floor(log2(x))
			j = bits.Len(uint(rangeLength)) - 1
			minRange := min(minst[j][best.left], minst[j][(best.right-1)-pow2(j)+1])
			maxRange := max(maxst[j][best.left], maxst[j][(best.right-1)-pow2(j)+1])
			heap.Push(h, candidate{
				value: maxRange - minRange,
				left:  best.left,
				right: best.right - 1,
			})
		}
	}

	return int64(total)
}

func pow2(x int) int {
	return 1 << x
}

/*
Key Points:

1. The "Monotonicity" Insight (Why the Heap Works)

For a fixed left endpoint, as you expand right to include more elements:

 • max can only increase (or stay same)
 • min can only decrease (or stay same)
 • Therefore, value = max - min is monotonically non-decreasing

This property is crucial. It means for each starting index l, the subarrays form a sorted sequence by value:

[l,l] ≤ [l,l+1] ≤ [l,l+2] ≤ ... ≤ [l,n-1]

When you have n sorted lists and need the global top k elements, the standard pattern is:

 • Push the last element (largest) of each list into a max-heap
 • Pop the global max, then push the previous element from that same list (the next largest for that specific l)

Similar problems: "Find K Pairs with Smallest Sums", "Kth Smallest Element in a Sorted Matrix"


2. Range Query Preprocessing (The Sparse Table)

The heap approach requires calculating max-min for arbitrary ranges [l, r] dynamically as you shrink windows.
Without preprocessing, this is O(n) per query, killing the complexity.

Key takeaway: When you need frequent min/max queries on static arrays, build a Sparse Table (O(n log n) build,
O(1) query) or a Segment Tree (O(n) build, O(log n) query).

The query formula using overlapping intervals:


j = floor(log2(length))

min = min(st[j][l], st[j][r - 2^j + 1])


This works because min/max are idempotent (overlapping doesn't hurt).


3. State Fingerprinting (Avoiding Duplicates)

Your initial instinct to use sorted values failed because sorting destroys spatial information. The correct
fingerprint for a subarray is always its original indices (left, right).

In the heap approach, you naturally avoid duplicates because:

 • You only generate [l, r-1] immediately after consuming [l, r]
 • Each (l, r) pair is generated exactly once

Pattern: If a problem asks for "top k distinct subarrays/intervals", ensure your state representation uses actual
indices, not derived values.


4. Complexity Analysis

 • Heap operations: k pops and at most k pushes → O(k log n)
 • Initialization: n suffixes → O(n log n)
 • Sparse Table: O(n log n) build
 • Total: O((n + k) log n)

Given constraints n ≤ 5×10⁴ and k ≤ 10⁵, this is ~1.5×10⁶ heap operations—easily fast enough.


5. When to Apply This Pattern

Use Heap + Best-First Search when:

 • You need top k results from a search space of O(n²) or larger (too big to enumerate)
 • The search space can be viewed as n "sorted streams" (monotonic sequences)
 • You can compute the "next" candidate from a "current" candidate efficiently

Use Binary Search + Sliding Window (the alternative approach) when:

 • You only need the k-th value (not the sum of top k)
 • You can count "how many subarrays have value ≤ X" efficiently


6. Implementation Detail: Bit Hacks

Using bits.Len(uint(x)) - 1 for floor(log2(x)) is faster and more precise than math.Log2 with floating-point
conversion. Remember:

 • bits.Len(uint(n)) = floor(log2(n)) + 1 (number of bits needed)
 • Number of sparse table levels = bits.Len(uint(n))

This pattern appears frequently in RMQ, binary lifting, and divide-and-conquer algorithms.

Summary: This problem combines three classic techniques: (1) Monotonic heap exploration for top-k, (2) Sparse
Table for range min/max, and (3) Index-based state tracking. Mastering these allows you to solve many "top k
subarray" variants efficiently.

*/
