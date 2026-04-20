package main

func solveDijkstra(graph Graph, startNode, endNode Node) {
	pq := &PriorityQueue{}

	graph.Distances[startNode.ID] = 0
	pq.Enqueue(startNode, 0)

	for !pq.IsEmpty() {
		current := pq.Dequeue()

		if graph.Visited[current.ID] {
			continue
		}
		graph.Visited[current.ID] = true

		if current == endNode {
			return
		}

		for _, neighbor := range graph.AdjacencyList[current.ID] {
			newDist := graph.Distances[current.ID] + 1

			if newDist < graph.Distances[neighbor.ID] {
				graph.Distances[neighbor.ID] = newDist
				graph.Parents[neighbor.ID] = current
				graph.History = append(graph.History, neighbor)

				pq.Enqueue(neighbor, newDist)
			}
		}
	}
}
