package main

type PQNode struct {
	Node     Node
	Priority int
}

type PriorityQueue struct {
	Nodes []PQNode
}

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

func (pq *PriorityQueue) Enqueue(node Node, priority int) {
	pq.Nodes = append(pq.Nodes, PQNode{node, priority})
}

// lowest priority = best
// dequeues node with best priority, which is lowest distance to end node
func (pq *PriorityQueue) Dequeue() Node {
	bestIndex := 0
	for i := 1; i < len(pq.Nodes); i++ {
		if pq.Nodes[i].Priority < pq.Nodes[bestIndex].Priority {
			bestIndex = i
		}
	}
	node := pq.Nodes[bestIndex].Node
	// remove best node so it doesn't get picked again
	pq.Nodes = append(pq.Nodes[:bestIndex], pq.Nodes[bestIndex+1:]...)
	return node
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Nodes) == 0
}
