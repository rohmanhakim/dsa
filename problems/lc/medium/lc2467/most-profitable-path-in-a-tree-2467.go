package lc2467

import (
	"math"
	"slices"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {

	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	bPath := []int{}
	path := []int{}
	maxProfit := math.MinInt
	visited := make(map[int]bool)
	bSteps := make(map[int]int)

	var dfs func(node int, isAlice bool, profit int, res *[]int)

	dfs = func(node int, isAlice bool, profit int, res *[]int) {
		if !isAlice && node == 0 {
			*res = slices.Clone(path)
			*res = append(*res, node)
			return
		}

		visited[node] = true
		path = append(path, node)

		if isAlice {

			gain := amount[node]
			currPos := len(path) - 1
			if val, ok := bSteps[node]; ok {
				if val == currPos {
					gain = amount[node] / 2
				} else if val < currPos {
					gain = 0
				}
			}
			profit = profit + gain
		}

		neighbors := []int{}
		for _, nb := range adj[node] {
			if !visited[nb] {
				neighbors = append(neighbors, nb)
			}
		}
		if isAlice && len(neighbors) == 0 {
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		for _, nb := range neighbors {
			dfs(nb, isAlice, profit, res)
			path = path[:len(path)-1]
		}
	}

	dfs(bob, false, 0, &bPath)
	for i, p := range bPath {
		bSteps[p] = i
	}
	visited = make(map[int]bool)
	dfs(0, true, 0, &[]int{})
	return maxProfit
}
