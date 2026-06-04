package temp

import (
	"fmt"
	"sort"
)

/*
func Push(queue []int, item int) []int {
	queue = append(queue, item)
	sort.Slice(queue, func(i, j int) bool {
		return queue[i] > queue[j]
	})
	return queue
}
*/

func Push(queue []int, item int) []int {
	fmt.Print("===> ")
	fmt.Println(queue)
	idx := sort.Search(len(queue), func(i int) bool {
		return queue[i] <= item
	})

	queue = append(queue, 0)
	copy(queue[idx+1:], queue[idx:])
	queue[idx] = item

	return queue
}

func Pop(queue []int) (int, []int) {
	front := queue[0]
	queue = queue[1:]
	return front, queue
}

func PriorityQueue() {
	queue := []int{}

	queue = Push(queue, 5)
	queue = Push(queue, 7)
	queue = Push(queue, 3)
	queue = Push(queue, 6)
	queue = Push(queue, 8)

	fmt.Println(queue)

	front, queue := Pop(queue)

	fmt.Println(front)
}
