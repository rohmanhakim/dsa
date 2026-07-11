package lc2685

func countCompleteComponents(n int, edges [][]int) int {
	adj := make([][]int, n)

	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	components := [][]int{}
	visited := make([]bool, n)
	for i := range n {
		if visited[i] {
			continue
		}
		queue := []int{i}
		members := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]
			neighbors := adj[front]
			for _, neighbor := range neighbors {
				if visited[neighbor] == false {
					members = append(members, neighbor)
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		components = append(components, members)
	}

	count := 0
	for _, c := range components {
		k := len(c)
		connected := true
		for _, v := range c {
			if len(adj[v]) != k-1 {
				connected = false
				break
			}
		}
		if connected {
			count++
		}
	}
	return count
}
