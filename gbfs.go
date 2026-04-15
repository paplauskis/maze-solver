package main

// greedy best-first search
func solveGBFS(graph Graph, startNode, endNode Node) {
	pq := &PriorityQueue{}
	graph.Visited[startNode.ID] = true
	pq.Enqueue(startNode, heuristic(startNode, endNode))

	for !pq.IsEmpty() {
		node := pq.Dequeue()

		if node == endNode {
			return
		}

		neighbors := graph.AdjacencyList[node.ID]

		for _, x := range neighbors {
			if graph.Visited[x.ID] {
				continue
			}

			graph.Visited[x.ID] = true
			graph.Parents[x.ID] = node
			graph.History = append(graph.History, x)

			priority := heuristic(x, endNode)
			pq.Enqueue(x, priority)
		}
	}
}

// manhattan distance
func heuristic(a, b Node) int {
	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}
	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
