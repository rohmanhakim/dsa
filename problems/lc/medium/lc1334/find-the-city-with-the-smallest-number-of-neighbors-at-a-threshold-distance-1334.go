package lc1334

import (
	"container/heap"
	"math"
	"sort"
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

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	adj := make(map[int][]edge)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{
			to:     v,
			weight: w,
		})
		adj[v] = append(adj[v], edge{
			to:     u,
			weight: w,
		})
	}

	validNeighbors := func(start int) int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt
		}

		dist[start] = 0

		pq := PQ{item{
			node: start,
			cost: 0,
		}}

		heap.Init(&pq)

		for pq.Len() > 0 {
			curr := heap.Pop(&pq).(item)

			if curr.cost > dist[curr.node] {
				continue
			}

			for _, e := range adj[curr.node] {
				nc := curr.cost + e.weight

				if nc <= distanceThreshold && nc < dist[e.to] {
					dist[e.to] = nc
					heap.Push(&pq, item{
						node: e.to,
						cost: nc,
					})
				}
			}
		}

		out := 0
		for i, d := range dist {
			if i != start && d < math.MaxInt {
				out++
			}
		}
		return out
	}
	res := [][2]int{}
	for i := range n {
		out := validNeighbors(i)
		res = append(res, [2]int{i, out})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i][1] != res[j][1] {
			return res[i][1] < res[j][1]
		} else {
			return res[i][0] > res[j][0]
		}
	})

	return res[0][0]
}
