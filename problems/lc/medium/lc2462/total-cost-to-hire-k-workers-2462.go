package lc2462

/*
You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the ith worker.

You are also given two integers k and candidates. We want to hire exactly k workers according to the following rules:

    You will run k sessions and hire exactly one worker in each session.
    In each hiring session, choose the worker with the lowest cost from either the first candidates workers or the last candidates workers. Break the tie by the smallest index.
        For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session, we will choose the 4th worker because they have the lowest cost [3,2,7,7,1,2].
        In the second hiring session, we will choose 1st worker because they have the same lowest cost as 4th worker but they have the smallest index [3,2,7,7,2]. Please note that the indexing may be changed in the process.
    If there are fewer than candidates workers remaining, choose the worker with the lowest cost among them. Break the tie by the smallest index.
    A worker can only be chosen once.

Return the total cost to hire exactly k workers.



Example 1:

Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
Output: 11
Explanation: We hire 3 workers in total. The total cost is initially 0.
- In the first hiring round we choose the worker from [17,12,10,2,7,2,11,20,8]. The lowest cost is 2, and we break the tie by the smallest index, which is 3. The total cost = 0 + 2 = 2.
- In the second hiring round we choose the worker from [17,12,10,7,2,11,20,8]. The lowest cost is 2 (index 4). The total cost = 2 + 2 = 4.
- In the third hiring round we choose the worker from [17,12,10,7,11,20,8]. The lowest cost is 7 (index 3). The total cost = 4 + 7 = 11. Notice that the worker with index 3 was common in the first and last four workers.
The total hiring cost is 11.

Example 2:

Input: costs = [1,2,4,1], k = 3, candidates = 3
Output: 4
Explanation: We hire 3 workers in total. The total cost is initially 0.
- In the first hiring round we choose the worker from [1,2,4,1]. The lowest cost is 1, and we break the tie by the smallest index, which is 0. The total cost = 0 + 1 = 1. Notice that workers with index 1 and 2 are common in the first and last 3 workers.
- In the second hiring round we choose the worker from [2,4,1]. The lowest cost is 1 (index 2). The total cost = 1 + 1 = 2.
- In the third hiring round there are less than three candidates. We choose the worker from the remaining workers [2,4]. The lowest cost is 2 (index 0). The total cost = 2 + 2 = 4.
The total hiring cost is 4.



Constraints:

    1 <= costs.length <= 10^5
    1 <= costs[i] <= 10^5
    1 <= k, candidates <= costs.length

*/

import (
	"container/heap"
	"fmt"
)

type worker struct {
	idx  int
	cost int
}
type minHeap []worker

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(worker))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap, rightHeap := &minHeap{}, &minHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)
	leftNext, rightNext := 0, len(costs)-1

	for i := 0; i < candidates && leftNext <= rightNext; i++ {
		heap.Push(leftHeap, worker{
			idx:  i,
			cost: costs[i],
		})
		leftNext++
	}

	for i := 0; i < candidates && leftNext <= rightNext; i++ {
		heap.Push(rightHeap, worker{
			idx:  rightNext,
			cost: costs[rightNext],
		})
		rightNext--
	}

	total := int64(0)
	for k > 0 {
		if leftHeap.Len() == 0 {
			front := heap.Pop(rightHeap).(worker)
			total += int64(front.cost)
			if rightNext >= leftNext {
				heap.Push(rightHeap, worker{
					idx:  rightNext,
					cost: costs[rightNext],
				})
				rightNext++
			}
		} else if rightHeap.Len() == 0 {
			front := heap.Pop(leftHeap).(worker)
			total += int64(front.cost)
			if leftNext <= rightNext {
				heap.Push(leftHeap, worker{
					idx:  leftNext,
					cost: costs[leftNext],
				})
				leftNext++
			}
		} else {
			leftFront := (*leftHeap)[0]
			rightFront := (*rightHeap)[0]

			if leftFront.cost <= rightFront.cost {
				heap.Pop(leftHeap)
				total += int64(leftFront.cost)
				if leftNext <= rightNext {
					heap.Push(leftHeap, worker{
						idx:  leftNext,
						cost: costs[leftNext],
					})
					leftNext++
				}
			} else {
				heap.Pop(rightHeap)
				total += int64(rightFront.cost)
				if rightNext >= leftNext {
					heap.Push(rightHeap, worker{
						idx:  rightNext,
						cost: costs[rightNext],
					})
					rightNext--
				}
			}
		}
		k--
	}

	return total
}

func TotalCostToHireKWorkers() {

	/*
				{17, 12, 10, 2}, 7, {2, 11, 20, 8}
				                 L
								 R

				{17, 12, 10, [2], 7}, {2, 11, 20, 8}
		                               L
								  R
				{17, 12, 10, [2], 7}, {[2], 11, 20, 8}
								        L
							      R
				{17, 12, 10, [2], [7]}, {[2], 11, 20, 8}
		                                  L
							       R
	*/

	// cost := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	// k := 3
	// candidates := 4 // 11

	// cost := []int{1, 2, 4, 1}
	// k := 3
	// candidates := 3 // 4

	// cost := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	// k := 11
	// candidates := 2 // 423

	// cost := []int{28, 35, 21, 13, 21, 72, 35, 52, 74, 92, 25, 65, 77, 1, 73, 32, 43, 68, 8, 100, 84, 80, 14, 88, 42, 53, 98, 69, 64, 40, 60, 23, 99, 83, 5, 21, 76, 34}
	// k := 32
	// candidates := 12 // 1407

	cost := []int{18, 64, 12, 21, 21, 78, 36, 58, 88, 58, 99, 26, 92, 91, 53, 10, 24, 25, 20, 92, 73, 63, 51, 65, 87, 6, 17, 32, 14, 42, 46, 65, 43, 9, 75}
	k := 13
	candidates := 23 // 223

	/*
		[31,25, | 72,79,74,65,84,91,18,59,27,9,81,33, | 17,58]                          17

		[31,25, | 72,79,74,65,84,91,18,59,27,9,81, | 33,[17],58]                        25

		[31,[25],72, | 79,74,65,84,91,18,59,27,9,81, | 33,[17],58]                      31

		[[31],[25],72,79, | 74,65,84,91,18,59,27,9,81, | 33,[17],58]                    33

		[[31],[25],72,79, | 74,65,84,91,18,59,27,9, | 81,[33],[17],58]                  58

		[[31],[25],72,79, | 74,65,84,91,18,59,27, | 9,81,[33],[17],[58]]                9

		[[31],[25],72,79, | 74,65,84,91,18,59, | 27,[9],81,[33],[17],[58]]              27

		[[31],[25],72,79, | 74,65,84,91,18, | 59,[27],[9],81,[33],[17],[58]]            59

		[[31],[25],72,79, | 74,65,84,91, | 18,[59],[27],[9],81,[33],[17],[58]]          18

		[[31],[25],72,79, | 74,65,84, | 91,[18],[59],[27],[9],81,[33],[17],[58]]        72

		[[31],[25],[72],79,74, | 65,84, | 91,[18],[59],[27],[9],81,[33],[17],[58]]      74
	*/
	out := totalCost(cost, k, candidates)
	fmt.Println(out)
}
