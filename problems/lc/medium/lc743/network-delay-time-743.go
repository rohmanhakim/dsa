package lc743

import (
	"container/heap"
	"math"
)

type edge struct {
	to     int
	weight int
}

type item struct {
	node int
	cost int
}

type PQ []item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ) Push(x any)        { *h = append(*h, x.(item)) }
func (h *PQ) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func networkDelayTime(times [][]int, n int, k int) int {

	adj := make(map[int][]edge)

	for _, t := range times {
		u, v, w := t[0], t[1], t[2]

		adj[u-1] = append(adj[u-1], edge{
			to:     v - 1,
			weight: w,
		})
	}

	shortestTime := func(dest int) int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt
		}

		pq := PQ{{
			node: k - 1,
			cost: 0,
		}}
		heap.Init(&pq)

		for pq.Len() > 0 {
			curr := heap.Pop(&pq).(item)

			if curr.node == dest {
				return dist[curr.node]
			}

			if curr.cost > dist[curr.node] {
				continue
			}

			for _, e := range adj[curr.node] {
				nc := curr.cost + e.weight

				if nc < dist[e.to] {
					dist[e.to] = nc
					heap.Push(&pq, item{
						node: e.to,
						cost: nc,
					})
				}
			}
		}

		return -1
	}

	longest := -1
	for i := range n {
		if i != k-1 {
			time := shortestTime(i)

			if time > -1 {
				if time > longest {
					longest = time
				}
			} else {
				return -1
			}
		}
	}

	return longest
}
