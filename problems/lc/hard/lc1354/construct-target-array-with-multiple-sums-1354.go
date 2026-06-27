/*
To decide between forward vs backward operation is a skill we can develop. Here are the concrete signals to look for:

## Signal 1: "Many choices forward, but forced choices backward"

In this problem, forward we must **choose** which index to set at each step — there's no rule telling we which one. But backward, the max element is **forced** — it MUST be the one most recently set. When forward gives we freedom but backward eliminates it, go backward.

**Ask ourselves:** "At each step, do I have multiple options, or is the answer forced?" If forward is branching and backward is deterministic, reverse.

## Signal 2: "The operation has a dominant/critical element"

When the forward operation always produces an element that is ≥ all others (like sum ≥ each element), that element becomes a "signature" of the last operation. we can identify it immediately in the current state. This is our backward anchor.

**Ask ourselves:** "Does the operation guarantee that something is the largest/max after it happens?" If yes, that's our greedy hook going backward.

## Signal 3: "Start state is simple, end state is complex"

we're given the complex end state (target) and asked if we can reach it from a simple start (all 1s). The question "can I build X from simple?" is often harder than "can I reduce X to simple?" because:
- Forward: we're searching a huge space of choices
- Backward: we're collapsing toward a known base case

**Ask ourselves:** "Is the start state trivial to verify (all 1s)?" If yes, try reducing toward it.

## Signal 4: "we keep getting stuck in forward"

When we try forward and find ourselves thinking "but which one do I pick?", that's the universe telling we the greedy choice isn't in that direction. The moment we notice we're making arbitrary decisions with no justification, that's a red flag to try backward.

## A practical checklist

When we encounter a problem with iterative operations:

1. **Can I identify what happened LAST?** (→ backward approach)
2. **Do I have a clear greedy choice going forward?** (→ forward approach)
3. **Is the start or end state trivially checkable?** (→ work from that end)

For this problem:
1. ✅ Last element set = max (since sum ≥ each element)
2. ❌ No clear greedy choice forward (which index?)
3. ✅ Start state (all 1s) is trivially checkable

→ Answer: backward.

## Other problems where this pattern appears

- **"Can we reach state A from state B?"** problems (like this one)
- Problems where the operation is **reversible** and the reverse has a **forced choice**
- **Euclidean algorithm** style problems (GCD reduces via modulo — that's backward thinking!)

The core habit: **before writing code, ask "which direction gives me fewer choices?"** — that's usually the right one.

*/

package lc1354

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func isPossible(target []int) bool {
	n := len(target)
	if n == 1 && target[0] != 1 {
		return false
	}
	arr := make([]int, n)

	h := &maxHeap{}
	heap.Init(h)

	currentSum := 0

	for i := range n {
		arr[i] = 1
		heap.Push(h, target[i])
		currentSum += target[i]
	}

	/*
	   Forward operation (what the problem describes):

	   Say we have an array and its current sum is `P`. We pick index `i` and set `arr[i] = P`.

	   Before the operation:

	   - Sum = `P`
	   - The value at index `i` was some unknown `old_val`
	   - Sum of all OTHER elements = `P - old_val`

	   After the operation:

	   - `arr[i]` is now `P` (replacing `old_val`)
	   - New sum `S` = (sum of other elements) + P = `(P - old_val) + P = 2P - old_val`

	*/
	for {
		maxItem := heap.Pop(h).(int)
		if maxItem == 1 {
			return true
		}

		rest := currentSum - maxItem
		if maxItem <= rest {
			return false
		}

		// Modular optimization: instead of reducing by rest one step at a time,
		// skip ahead using modulo. Each backward step reduces max by rest.
		oldVal := maxItem % rest
		if oldVal == 0 {
			oldVal = rest
		}

		currentSum = currentSum - maxItem + oldVal
		heap.Push(h, oldVal)
	}
}

func ConstructTargetArrayWithMultipleSums() {
	target := []int{2}
	out := isPossible(target)
	fmt.Println(out)
}
