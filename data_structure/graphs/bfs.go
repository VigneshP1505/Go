package graphs

func bfs(start int, graph [][]int, visited []bool) {
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func bfsUtil() {
	var graph [][]int
	graph[0] = []int{1, 2}
	graph[1] = []int{3, 4}
	visited := make([]bool, 5)
	bfs(0, graph, visited)
}
