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
