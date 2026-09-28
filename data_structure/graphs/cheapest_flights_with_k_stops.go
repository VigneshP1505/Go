package graphs

type FlightPairs struct {
	dest int
	cost int
}

func dfsCheapFlights(node int, dest int, k int, graph [][]FlightPairs) {
	if node==dest{
		
	}
}

func cheapestFlightsWithKStops(flights [][]int, start int, dest int, k int) {
	n := len(flights)
	graph := make([][]FlightPairs, n)

	for i := 0; i < n; i++ {
		flight := flights[i]
		src := flight[0]
		dest := flight[1]
		cost := flight[2]

		graph[src] = append(graph[src], FlightPairs{dest: dest, cost: cost})
	}

	dfsCheapFlights(start, dest, k, graph)

}
