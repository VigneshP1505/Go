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
