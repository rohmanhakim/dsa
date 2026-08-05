package lc3310

func remainingMethods(n int, k int, invocations [][]int) []int {

	adj := make(map[int][]int)

	for _, node := range invocations {
		adj[node[0]] = append(adj[node[0]], node[1])
	}

	queue := []int{k}
	visited := make([]bool, n)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if !visited[front] {
			visited[front] = true
		}

		neighbors := adj[front]

		for _, ne := range neighbors {
			if !visited[ne] {
				queue = append(queue, ne)
			}
		}
	}

	all := []int{}
	for i := range n {
		all = append(all, i)
	}

	for _, inv := range invocations {
		a, b := inv[0], inv[1]

		if !visited[a] && visited[b] {
			return all
		}
	}

	nonSuspicious := []int{}

	for i, val := range visited {
		if !val {
			nonSuspicious = append(nonSuspicious, i)
		}
	}

	return nonSuspicious
}
