package graphs

func cheapestFlightsWithKStops(flights [][]int) {
	n := len(flights)
	m := len(flights[0])
	graph := make([][]int, n)

	for i := range n {
		graph[i] = make([]int, m)
	}

}
