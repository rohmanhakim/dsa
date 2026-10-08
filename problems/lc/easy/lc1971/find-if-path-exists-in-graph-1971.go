package lc1971

func validPath(n int, edges [][]int, source int, destination int) bool {

	adj := make(map[int][]int)

	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	var dfs func(curr int) bool

	visited := make(map[int]bool)

	dfs = func(curr int) bool {

		if curr == destination {
			return true
		}
		neighbors := adj[curr]
		visited[curr] = true
		for _, ne := range neighbors {
			if visited[ne] == false {

				res := dfs(ne)
				if res {
					return true
				}
			}
		}

		return false
	}

	res := dfs(source)
	return res
}
