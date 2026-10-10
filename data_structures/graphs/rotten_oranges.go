package graphs

type Pair struct {
	row       int
	col       int
	timestamp int
}

func _rottenOranges(graph [][]int) int {
	var n int = len(graph)
	var m int = len(graph[0])
	queue := []Pair{}

	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			if graph[i][j] == 2 {
				queue = append(queue, Pair{row: i, col: j, timestamp: 0})
				visited[i][j] = 1
			}
		}
	}

	time := 0
	deltaRow := []int{-1, 0, 1, 0}
	deltaCol := []int{0, 1, 0, -1}
	for len(queue) > 0 {
		pair := queue[0]
		row := pair.row
		col := pair.col
		t := pair.timestamp
		queue = queue[1:]

		time = max(time, t)

		for i := range 4 {
			nrow := row + deltaRow[i]
			ncol := col + deltaCol[i]
			if (nrow) >= 0 && (ncol) >= 0 && graph[nrow][ncol] == 1 && visited[nrow][ncol] != 1 {
				visited[nrow][ncol] = 1
				graph[nrow][ncol] = 2
				queue = append(queue, Pair{row: nrow, col: ncol, timestamp: t + 1})
			}
		}
	}

	return time

}
