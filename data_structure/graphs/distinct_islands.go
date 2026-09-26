package graphs

import (
	"fmt"
	"strings"
)

func _dfsIsland(i int, j int, graph [][]int, visited [][]bool, shape *[]string, row0 int, col0 int) {
	visited[i][j] = true
	key := fmt.Sprintf("%d,%d", i-row0, j-col0)
	*shape = append(*shape, key)

	deltaRow := []int{-1, 0, 1, 0}
	deltaCol := []int{0, 1, 0, -1}
	n := len(graph)
	m := len(graph[0])

	for k := 0; k < 4; k++ {
		nrow := i + deltaRow[k]
		ncol := j + deltaCol[k]
		if nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && !visited[nrow][ncol] && graph[nrow][ncol] == 1 {
			_dfsIsland(nrow, ncol, graph, visited, shape, row0, col0)
		}
	}

}

func _distinctIslands(graph [][]int) int {
	n := len(graph)
	m := len(graph[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}

	unique := make(map[string]struct{})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				shape := []string{}
				_dfsIsland(i, j, graph, visited, &shape, i, j)
				key := strings.Join(shape, ";")
				unique[key] = struct{}{}
			}
		}
	}

	return len(unique)

}
