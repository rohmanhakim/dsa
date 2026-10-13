package cf1473e

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

const inf int64 = 1e16 // Safe upper bound constraint for 2*10^5 * 10^9

type edge struct {
	to     int
	weight int64
}

type item struct {
	node    int
	cost    int64
	subUsed int
	addUsed int
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

func minimumPath(m, n int, edges [][3]int64) []int64 {

	adj := make(map[int][]edge)
	for _, e := range edges {
		adj[int(e[0])] = append(adj[int(e[0])], edge{
			to:     int(e[1]),
			weight: e[2],
		})
		adj[int(e[1])] = append(adj[int(e[1])], edge{
			to:     int(e[0]),
			weight: e[2],
		})
	}

	dist := make([][2][2]int64, n)

	for i := range dist {
		dist[i] = [2][2]int64{{
			inf,
			inf,
		}, {
			inf,
			inf,
		}}
	}

	dist[0] = [2][2]int64{{
		0,
		math.MaxInt,
	}, {
		math.MaxInt,
		math.MaxInt,
	}}

	pq := PQ{item{
		node:    0,
		cost:    0,
		subUsed: 0,
		addUsed: 0,
	}}

	walk := func(curr item, spendSub, spendAdd, dest int, price int64) {
		ns := curr.subUsed | spendSub
		na := curr.addUsed | spendAdd
		newCost := curr.cost + price

		if newCost < dist[dest][ns][na] {
			dist[dest][ns][na] = newCost
			heap.Push(&pq, item{
				node:    dest,
				cost:    newCost,
				subUsed: ns,
				addUsed: na,
			})
		}
	}

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(item)

		if curr.cost > dist[curr.node][curr.subUsed][curr.addUsed] {
			continue
		}

		for _, e := range adj[curr.node] {
			walk(curr, 0, 0, e.to, e.weight)

			if curr.subUsed == 0 {
				walk(curr, 1, 0, e.to, 0)
				if curr.addUsed == 0 {
					walk(curr, 1, 1, e.to, e.weight)
				}
			}
			if curr.addUsed == 0 {
				walk(curr, 0, 1, e.to, e.weight+e.weight)
			}
		}
	}

	res := []int64{}
	for i := 1; i < len(dist); i++ {
		res = append(res, dist[i][1][1])
	}
	return res
}

func MinimumPath() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][3]int64, m)
	for i := range m {
		var u, v, w int64
		fmt.Fscan(reader, &u, &v, &w)
		edges[i][0], edges[i][1], edges[i][2] = u-1, v-1, w
	}

	res := minimumPath(m, n, edges)

	out := make([]any, len(res))
	for i, v := range res {
		out[i] = v
	}
	fmt.Fprintln(writer, out...)
}
