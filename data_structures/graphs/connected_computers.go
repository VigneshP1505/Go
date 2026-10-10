package graphs

func countComponents(n int, connections [][]int) int {
	graph := make([][]int, n)

	for _, edge := range connections {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	count := 0

	var dfs func(int)

	dfs = func(node int) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for node := 0; node < n; node++ {
		if !visited[node] {
			count++
			dfs(node)
		}
	}

	return count
}
