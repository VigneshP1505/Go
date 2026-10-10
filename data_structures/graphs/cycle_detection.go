package graphs

func _hasCycle(node int, graph [][]int, state []int) bool {
	if state[node] == 1 {
		return true
	}
	if state[node] == 2 {
		return false
	}
	state[node] = 1
	for _, neighbor := range graph[node] {
		if _hasCycle(neighbor, graph, state) {
			return true
		}
	}
	state[node] = 2
	return false
}

type Pairs struct {
	parent int
	child  int
}

func _hasCycleBfs(start int, graph [][]int, visited []bool) bool {
	visited[start] = true
	pairs := Pairs{
		parent: -1,
		child:  start,
	}
	queue := []Pairs{pairs}
	for len(queue) > 0 {
		pair := queue[0]
		parent := pair.parent
		child := pair.child
		queue = queue[1:]
		for _, neighbor := range graph[child] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, Pairs{parent: child, child: neighbor})
			} else if parent != neighbor {
				return true
			}
		}
	}
	return false
}

func _hasCycleDfs(pair Pairs, graph [][]int, visited []bool) bool {
	parent := pair.parent
	child := pair.child
	visited[child] = true
	for _, neighbor := range graph[child] {
		if !visited[neighbor] {
			if _hasCycleDfs(Pairs{parent: child, child: neighbor}, graph, visited) {
				return true
			}
		} else if neighbor != parent {
			return true
		}
	}
	return false
}
