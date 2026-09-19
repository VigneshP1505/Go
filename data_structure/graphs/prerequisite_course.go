package graphs

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		course := prerequisite[0]
		preRequisite := prerequisite[1]
		graph[preRequisite] = append(graph[preRequisite], course)
	}

	state := make([]int, numCourses)

	var dfs func(node int) bool

	dfs = func(course int) bool {
		if state[course] == 1 {
			return false
		}

		if state[course] == 2 {
			return true
		}

		state[course] = 1
		for _, next := range graph[course] {
			if !dfs(next) {
				return false
			}
		}
		state[course] = 2
		return true
	}

	for course := 0; course < numCourses; course++ {
		if state[course] == 0 {
			if !dfs(course) {
				return false
			}
		}
	}

	return true
}

func has_cycle_dfs(node int, parent int, adj [][]int, visited []bool) bool {
	visited[node] = true

	for _, neighbor := range adj[node] {
		if !visited[neighbor] {
			if has_cycle_dfs(neighbor, node, adj, visited) {
				return true
			}
		} else if neighbor != parent {
			return true
		}
	}

	return false
}

func cyclic(numNodes int, adj [][]int) bool {
	visited := make([]bool, numNodes)
	for i := range numNodes {
		if !visited[i] {
			if has_cycle_dfs(i, -1, adj, visited) {
				return true
			}
		}
	}
	return false
}
