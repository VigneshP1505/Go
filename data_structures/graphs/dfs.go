package graphs

import "fmt"

func dfs(n int, graph [][]int, visited []bool) {
	visited[n] = true
	fmt.Println(n)
	for _, neighbor := range graph[n] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}

func depthFirstSearch() {
	graph := make([][]int, 5)
	graph[0] = []int{1, 2}
	graph[1] = []int{3}
	graph[2] = []int{5}
	graph[3] = []int{4}
	graph[4] = []int{}
	visited := make([]bool, 5)
	for i := 0; i < 5; i++ {
		if !visited[i] {
			dfs(i, graph, visited)
		}
	}
}
